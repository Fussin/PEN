package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ScansProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "reconnaissance_scans_processed_total",
		Help: "The total number of processed scans",
	})
)
