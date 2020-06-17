package main

import (
	"log"

	. "github.com/knight0zh/demo_server/base"
	"github.com/knight0zh/demo_server/routers"
)

func main() {

	r := routers.InitRouter()

	port := Config.Get("app.port").(string)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
