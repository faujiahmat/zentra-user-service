package grpc

import (
	"github.com/faujiahmat/zentra-user-service/src/core/grpc/handler"
	"github.com/faujiahmat/zentra-user-service/src/core/grpc/server"
	"github.com/faujiahmat/zentra-user-service/src/interface/service"
)

func InitServer(us service.User) *server.Grpc {
	userHandler := handler.NewUserGrpc(us)

	grpcServer := server.NewGrpc(userHandler)

	return grpcServer
}
