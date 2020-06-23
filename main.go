package main

import (
	"log"

	"github.com/knight0zh/demo_config/config"
	"github.com/knight0zh/demo_server/routers"
)

func main() {

	r := routers.InitRouter()

	port := config.Get("app.port").(string)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
