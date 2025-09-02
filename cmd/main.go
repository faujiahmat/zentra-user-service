package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/faujiahmat/zentra-user-service/src/core/grpc"
	"github.com/faujiahmat/zentra-user-service/src/infrastructure/database"
	"github.com/faujiahmat/zentra-user-service/src/repository"
	"github.com/faujiahmat/zentra-user-service/src/service"
)

func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func handleCloseApp(closeCH chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		close(closeCH)
	}()
}

func main() {
	closeCH := make(chan struct{})
	handleCloseApp(closeCH)

	postgresDB := database.NewPostgres()
	defer database.ClosePostgres(postgresDB)

	userRepository := repository.NewUser(postgresDB)

	grpcClient := grpc.InitClient()
	defer grpcClient.Close()

	userService := service.NewUser(grpcClient, userRepository)

	grpcServer := grpc.InitServer(userService)
	defer grpcServer.Stop()

	go grpcServer.Run()

	<-closeCH
}
