package cmd

import (
	"bufio"
	"fmt"
	"os"
)

type SgridCmd struct{
    registryMap map[string]func()
}


func (s *SgridCmd) Registry(name string,cb func()) {
    s.registryMap[name] = cb
}

func (s *SgridCmd) Invoke(name string) {
	fn,ok := s.registryMap[name]
    if !ok{
		return
	}
	fn()
}

func (s *SgridCmd) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for {
			// 扫描是否有新的输入行
			if scanner.Scan() {
				input := scanner.Text()
				fmt.Printf("从管道接收到的内容是：%s\n", input)
                s.Invoke(input)
            } else {
				if err := scanner.Err(); err != nil {
					fmt.Printf("读取管道输入时出错：%s\n", err)
				}
				break
			}
		}
	}()
}


func NewSgridCmd() *SgridCmd{
    s := &SgridCmd{}
    s.registryMap = make(map[string]func())
	s.Run()
    return s
}