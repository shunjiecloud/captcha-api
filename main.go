package main

import (
	"log"

	"github.com/micro/go-micro/v2/web"
	"github.com/shunjiecloud/captcha_api/modules"
	"github.com/shunjiecloud/captcha_api/routers"
)

func main() {
	//  modules init
	modules.Setup()

	//  Create web
	webSrv := web.NewService(
		web.Name("go.micro.api.captcha"),
	)

	//  register web handler
	webSrv.Handle("/", routers.InitRouter())

	//  init
	if err := webSrv.Init(); err != nil {
		log.Fatal(err)
	}

	if err := webSrv.Run(); err != nil {
		log.Fatal(err)
	}
}
