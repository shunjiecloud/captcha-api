package main

import (
	"log"

	"github.com/dchest/captcha"
	"github.com/micro/go-micro/v2/web"
)

func main() {

	//  Create web
	webSrv := web.NewService(
		web.Name("go.micro.api.captcha"),
	)

	//  register web handler
	webSrv.Handle("/", captcha.Server(captcha.StdWidth, captcha.StdHeight))

	//  init
	webSrv.Init()

	if err := webSrv.Run(); err != nil {
		log.Fatal(err)
	}
}
