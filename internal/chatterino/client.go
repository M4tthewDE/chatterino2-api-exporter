package chatterino

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/m4tthewde/chatterino2-api-exporter/internal/metrics"
)

type Client struct {
	httpClient http.Client
	interval   time.Duration
	hostname   string
	port       string
}

func NewClient(interval time.Duration, hostname string, port string) *Client {
	return &Client{
		httpClient: http.Client{},
		interval:   interval,
		hostname:   hostname,
		port:       port,
	}
}

func (c *Client) Scrape() {
	for range time.Tick(c.interval) {
		allStats := c.getStatistics()
		c.setMetrics(allStats)
	}
}

func (c *Client) getStatistics() *AllStats {
	var allStats AllStats
	var memory Memory

	memoryUrl := "http://" + c.hostname + ":" + c.port + "/health/memory"
	body := c.MakeRequest(memoryUrl)
	memoryString := string(body)

	parts := strings.Split(memoryString, ",")

	for i, part := range parts {
		keyAndValue := strings.Split(part, "=")
		allocString := strings.Split(keyAndValue[1], " ")[0]
		alloc, err := strconv.Atoi(allocString)
		if err != nil {
			log.Fatal(err)
		}
		byteAlloc := int64(alloc) * int64(1048576)

		switch i {
		case 0:
			memory.alloc = byteAlloc
		case 1:
			memory.totalAlloc = byteAlloc
		case 2:
			memory.systemAlloc = byteAlloc
		}

		if i == 3 {
			memory.numGC = alloc
		}
	}

	allStats.memory = &memory
	return &allStats
}

func (c *Client) setMetrics(allStats *AllStats) {
	metrics.Alloc.WithLabelValues(c.hostname).Set(float64(allStats.memory.alloc))
	metrics.TotalAlloc.WithLabelValues(c.hostname).Set(float64(allStats.memory.totalAlloc))
	metrics.SystemAlloc.WithLabelValues(c.hostname).Set(float64(allStats.memory.systemAlloc))
	metrics.NumGC.WithLabelValues(c.hostname).Set(float64(allStats.memory.numGC))
}

func (c *Client) MakeRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
