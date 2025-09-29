package util

import (
	"github.com/faujiahmat/zentra-user-service/src/cache"
	"github.com/faujiahmat/zentra-user-service/src/core/grpc/client"
	"github.com/faujiahmat/zentra-user-service/src/interface/service"
	"github.com/faujiahmat/zentra-user-service/src/repository"
	serviceimpl "github.com/faujiahmat/zentra-user-service/src/service"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitUserServiceTest(gc *client.Grpc, postgresDB *gorm.DB, r *redis.ClusterClient) service.User {
	userCache := cache.NewUser(r)
	userRepository := repository.NewUser(postgresDB, userCache)

	userService := serviceimpl.NewUser(gc, userRepository, userCache)
	return userService
}
