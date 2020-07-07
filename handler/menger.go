package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

)

type Menger struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Menger) Login(ctx context.Context, req *menger.LoginRequest, rsp *menger.LoginResponse) error {
	log.Info("Received Login request")
	rsp.Msg = "Hello "
	return nil
}

