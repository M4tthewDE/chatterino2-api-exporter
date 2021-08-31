package main

import (
	"time"

	"github.com/m4tthewde/chatterino2-api-exporter/config"
	"github.com/m4tthewde/chatterino2-api-exporter/internal/chatterino"
	"github.com/m4tthewde/chatterino2-api-exporter/internal/metrics"
	"github.com/m4tthewde/chatterino2-api-exporter/internal/server"
)

func main() {
	cfg := config.GetConfig()

	metrics.Init()

	initChatterinoClient(cfg)

	s := server.NewServer(cfg.ServerPort)
	go s.ListenAndServe()

	for {
		select {}
	}
}

func initChatterinoClient(cfg *config.Config) {
	client := chatterino.NewClient(time.Second*time.Duration(cfg.Interval), cfg.Protocol, cfg.Hostname)
	go client.Scrape()
}
