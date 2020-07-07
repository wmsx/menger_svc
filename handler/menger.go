package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"
	menger "github.com/wmsx/menger_svc/proto/menger"
)

type Menger struct{}

func (e *Menger) Login(ctx context.Context, req *menger.LoginRequest, rsp *menger.LoginResponse) error {
	log.Info("Received Login request")
	rsp.Msg = "Hello"
	return nil
}

