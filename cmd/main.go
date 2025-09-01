package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/faujiahmat/zentra-user-service/src/infrastructure/database"
	"github.com/faujiahmat/zentra-user-service/src/repository"
)

func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func handleCloseApp(closeCH chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	postgresDB := database.NewPostgres()
	defer database.ClosePostgres(postgresDB)

	userRepository := repository.NewUser(postgresDB)

	go func() {
		<-c
		close(closeCH)
	}()
}

func main() {
	closeCH := make(chan struct{})
	handleCloseApp(closeCH)

	<-closeCH
}
