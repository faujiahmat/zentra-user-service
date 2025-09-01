package server

import (
	"fmt"
	"net"

	"github.com/faujiahmat/zentra-proto/protogen/user"
	"github.com/faujiahmat/zentra-user-service/src/common/log"
	"github.com/faujiahmat/zentra-user-service/src/infrastructure/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Grpc struct {
	port            string
	server          *grpc.Server
	userGrpcHandler user.UserServiceServer
}

func NewGrpc(userGrpcHandler user.UserServiceServer) *Grpc {
	port := config.Conf.CurrentApp.GrpcPort

	return &Grpc{
		port:            port,
		userGrpcHandler: userGrpcHandler,
	}
}

func (g *Grpc) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "server.Grpc/Run", "section": "net.Listen"}).Fatal(err)
	}

	log.Logger.Infof("grpc run in port: %s", g.port)

	grpcServer := grpc.NewServer()

	g.server = grpcServer

	user.RegisterUserServiceServer(grpcServer, g.userGrpcHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "server.Grpc/Run", "section": "grpcServer.Serve"}).Fatal(err)
	}
}

func (g *Grpc) Stop() {
	g.server.Stop()
}
