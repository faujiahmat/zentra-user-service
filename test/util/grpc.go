package util

import (
	pb "github.com/faujiahmat/zentra-proto/protogen/user"
	"github.com/faujiahmat/zentra-user-service/src/common/log"
	"github.com/faujiahmat/zentra-user-service/src/core/grpc/client"
	"github.com/faujiahmat/zentra-user-service/src/core/grpc/handler"
	"github.com/faujiahmat/zentra-user-service/src/core/grpc/interceptor"
	"github.com/faujiahmat/zentra-user-service/src/core/grpc/server"
	"github.com/faujiahmat/zentra-user-service/src/infrastructure/config"
	"github.com/faujiahmat/zentra-user-service/src/interface/service"
	"github.com/faujiahmat/zentra-user-service/src/mock/delivery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpcServerTest(us service.User) *server.Grpc {
	userGrpcHandler := handler.NewUserGrpc(us)
	unaryResponseInterceptor := interceptor.NewUnaryResponse()

	grpcServer := server.NewGrpc(userGrpcHandler, unaryResponseInterceptor)
	return grpcServer
}

func InitGrpcClientTest(ogdm *delivery.OtpGrpcMock) *client.Grpc {
	otpGrpcConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(ogdm, otpGrpcConn)
	return grpcClient
}

func InitUserGrpcDelivery() (pb.UserServiceClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(config.Conf.ApiGateway.BaseUrl, opts...)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.InitUserGrpcDelivery", "section": "grpc.NewClient"}).Fatal(err)
	}

	UserGrpcDeliver := pb.NewUserServiceClient(conn)
	return UserGrpcDeliver, conn
}
