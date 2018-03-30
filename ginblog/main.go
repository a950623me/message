package main

import (
	// "fmt"
	// "log"
	"ginblog/router"
	// "net/http"
)

func main() {

	router := routers.InitRouter()

	// openHttpListen()

	// router.Use(cors.Default())
	router.Run(":8010")
}

