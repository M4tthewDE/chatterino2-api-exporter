package main

import (
	"time"

	"github.com/m4tthewde/chatterino2-api-exporter/internal/chatterino"
	"github.com/m4tthewde/chatterino2-api-exporter/internal/metrics"
	"github.com/m4tthewde/chatterino2-api-exporter/internal/server"
)

func main() {
	metrics.Init()

	initChatterinoClient()
	s := server.NewServer("9500")
	go s.ListenAndServe()

	for {
		select {}
	}
}

func initChatterinoClient() {
	client := chatterino.NewClient(time.Second*10, "localhost", "1234")
	go client.Scrape()
}
