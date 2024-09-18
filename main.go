package main

import (
	"com_sgrid_gotrade/src/components"
	component_cache "com_sgrid_gotrade/src/components/cache"
	component_db "com_sgrid_gotrade/src/components/db"
	"com_sgrid_gotrade/src/router"
	"com_sgrid_gotrade/src/schedule"
	"fmt"
	"os"
)

func init() {
	components.LoadComponents()
	router.LoadRouter(components.Gin_Server)
	component_db.LoadDB(components.Sgrid_Conf)
	component_cache.LoadCache(components.Sgrid_Conf)
	defer schedule.InitSchedule()
}

func main() {
	var port = fmt.Sprintf(":%s", os.Getenv("SGRID_TARGET_PORT"))
	if port == ":" {
		port = fmt.Sprintf(":%v", components.Sgrid_Conf.Server.Port)
	}
	fmt.Println("Starting server on port " + port)
	components.Gin_Server.Run(port)
}
