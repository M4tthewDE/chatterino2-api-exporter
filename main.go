package main

import (
	"time"

	"github.com/m4tthewde/chatterino2-api-exporter/internal/chatterino"
	"github.com/m4tthewde/chatterino2-api-exporter/internal/server"
)

func main() {
	s := server.NewServer("9500")
	go s.ListenAndServe()

	initChatterinoClient()
	for {
		select {}
	}
}

func initChatterinoClient() {
	client := chatterino.NewClient(time.Second*10, "")
	go client.Scrape()
}
