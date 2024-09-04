package main

import (
	"com_sgrid_gotrade/src/components"
	"com_sgrid_gotrade/src/router"
	"fmt"
	"os"
)

func init() {
	components.LoadComponents()
	router.LoadRouter(components.Gin_Server)
}

func main() {
	var port = fmt.Sprintf(":%s", os.Getenv("SGRID_TARGET_PORT"))
	if port == ":" {
		port = fmt.Sprintf(":%v", components.Sgrid_Conf.Server.Port)
	}
	fmt.Println("Starting server on port " + port)
	components.Gin_Server.Run(port)
}
