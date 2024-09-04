package public

import (
	"com_sgrid_gotrade/src/framework/config"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	ENV_PRODUCTION         = "SGRID_PRODUCTION"
	ENV_TARGET_PORT        = "SGRID_TARGET_PORT"
	ENV_PROCESS_INDEX      = "SGRID_PROCESS_INDEX"
	ENV_SGRID_SERVANT_NAME = "SGRID_SERVANT_NAME"
	SGRID_CONFIG           = "SGRID_CONFIG"
	DEV_CONF_NAME          = "sgrid.yml"
	PROD_CONF_NAME         = "sgridProd.yml"
)

func GetWd() string {
	dir, _ := os.Getwd()
	s := os.Getenv(ENV_PRODUCTION)
	if len(s) == 0 {
		return dir
	} else {
		return s
	}
}

func Join(args ...string) string {
	s := GetWd()
	arr := []string{}
	arr = append(arr, s)
	arr = append(arr, args...)
	return filepath.Join(arr...)
}

type ConfOpt func(*withConf)

func WithTargetPath(targetPath string) ConfOpt {
	return func(conf *withConf) {
		conf.targetPath = targetPath
	}
}

func WithNameSpace(nameSpace string) ConfOpt {
	return func(conf *withConf) {
		conf.nameSpace = nameSpace
	}
}

type withConf struct {
	targetPath string
	nameSpace  string
}

func SgridProduction() bool {
	s := os.Getenv(ENV_PRODUCTION)
	fmt.Println("s", s)
	if len(s) == 0 {
		return false
	} else {
		return true
	}
}

func NewConfig(opts ...ConfOpt) (conf *config.SgridConf, err error) {
	prod := os.Getenv(SGRID_CONFIG)
	wc := &withConf{}
	if len(prod) > 0 {
		err = yaml.Unmarshal([]byte(prod), &conf)
		if err != nil {
			fmt.Println("err", err.Error())
		}
		fmt.Printf("SGRID_PROD_CONFIG %+v\n", conf)
	} else {
		for _, opt := range opts {
			opt(wc)
		}
		var path string
		if len(wc.targetPath) != 0 {
			path = wc.targetPath
		} else if len(wc.nameSpace) != 0 {
			path = Join(DEV_CONF_NAME)
		} else if SgridProduction() {
			path = Join(PROD_CONF_NAME)
		} else {
			path = Join(DEV_CONF_NAME)
		}
		// 读取 YAML 文件
		yamlFile, err := os.ReadFile(path)
		fmt.Println("Get FilePath from ", path)
		if err != nil {
			fmt.Println("Error reading YAML file:", err)
			return conf, err
		}
		// 解析 YAML 数据
		err = yaml.Unmarshal(yamlFile, &conf)
		if err != nil {
			fmt.Println("Error unmarshalling YAML:", err)
			return conf, err
		}
		fmt.Printf("SGRID_DEV_CONFIG %+v\n", conf)

	}
	// 打印解析后的配置
	return conf, nil
}
