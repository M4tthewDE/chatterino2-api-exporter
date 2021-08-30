package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	Alloc = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "alloc_memory",
			Namespace: "chatterino",
			Help:      "Currently allocated memory",
		},
		[]string{"hostname"},
	)

	TotalAlloc = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "total_alloc_memory",
			Namespace: "chatterino",
			Help:      "Total allocated memory",
		},
		[]string{"hostname"},
	)

	SystemAlloc = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "system_alloc_memory",
			Namespace: "chatterino",
			Help:      "Currently allocated memory by the system",
		},
		[]string{"hostname"},
	)

	NumGC = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "num_garbage_collector",
			Namespace: "chatterino",
			Help:      "Number of completed GC cycles",
		},
		[]string{"hostname"},
	)
)

func Init() {
	initMetric("alloc_memory", Alloc)
	initMetric("total_alloc_memory", TotalAlloc)
	initMetric("system_alloc_memory", SystemAlloc)
	initMetric("num_garbage_collector", NumGC)
}

func initMetric(name string, metric *prometheus.GaugeVec) {
	prometheus.MustRegister(metric)
	log.Printf("New Prometheus metric registered: %s", name)
}
