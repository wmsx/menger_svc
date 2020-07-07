package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wmsx/menger_svc/handler"
	menger "github.com/wmsx/menger_svc/proto/menger"
)

func main() {
	service := micro.NewService(
		micro.Name("wm.sx.svc.menger"),
		micro.Version("latest"),
	)

	service.Init()

	_ = menger.RegisterMengerHandler(service.Server(), new(handler.Menger))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
