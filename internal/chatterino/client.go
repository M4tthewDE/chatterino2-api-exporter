package chatterino

import (
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	interval   time.Duration
	hostname   string
}

func NewClient(interval time.Duration, hostname string) *Client {
	return &Client{
		httpClient: http.Client{},
		interval:   interval,
		hostname:   hostname,
	}
}

func (c *Client) Scrape() {
	for range time.Tick(c.interval) {
		log.Println("test")
	}
}
