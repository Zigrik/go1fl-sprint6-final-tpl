package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)
	srv := server.StartServer(logger)
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatal("Error while server start: ", err)
	}
}
