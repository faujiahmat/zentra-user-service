package restful

import (
	"github.com/faujiahmat/zentra-user-service/src/core/restful/client"
	"github.com/faujiahmat/zentra-user-service/src/core/restful/delivery"
	"github.com/faujiahmat/zentra-user-service/src/core/restful/handler"
	"github.com/faujiahmat/zentra-user-service/src/core/restful/middleware"
	"github.com/faujiahmat/zentra-user-service/src/core/restful/server"
	"github.com/faujiahmat/zentra-user-service/src/interface/service"
)

func InitServer(rc *client.Restful, us service.User) *server.Restful {

	userHandler := handler.NewUser(us, rc)
	middleware := middleware.New(rc)

	restfulServer := server.NewRestful(userHandler, middleware)
	return restfulServer
}

func InitClient() *client.Restful {
	imageKitDelivery := delivery.NewImageKit()

	restfulClient := client.NewRestful(imageKitDelivery)
	return restfulClient
}
