package main

import (
	"fmt"
	"github.com/autonomouspen/reconnaissance/pkg/db"
	"github.com/autonomouspen/reconnaissance/pkg/metrics"
	"github.com/autonomouspen/reconnaissance/pkg/plugin"
	"github.com/autonomouspen/reconnaissance/pkg/queue"
	"github.com/autonomouspen/reconnaissance/pkg/ratelimiter"
	"github.com/autonomouspen/reconnaissance/pkg/scanner"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

type ExamplePlugin struct{}

func (p *ExamplePlugin) Name() string {
	return "Example Plugin"
}

func (p *ExamplePlugin) Run(target string) (string, error) {
	return fmt.Sprintf("Example plugin ran on %s", target), nil
}

func main() {
	pluginManager := plugin.NewManager()
	pluginManager.Register(&ExamplePlugin{})

	rateLimiter := ratelimiter.NewRateLimiter(100, time.Second)
	jobQueue := queue.NewRedisQueue("localhost:6379", "scan_jobs")

	connStr := "user=postgres dbname=autonomouspen password=password sslmode=disable"
	database, err := db.NewPostgres(connStr)
	if err != nil {
		log.Fatal(err)
	}

	s := scanner.NewScanner(10000, pluginManager, rateLimiter, jobQueue)
	s.Start()

	go func() {
		for i := 0; i < 100000; i++ {
			jobQueue.Enqueue(queue.Job{ID: fmt.Sprintf("job-%d", i), Target: fmt.Sprintf("target-%d", i)})
		}
	}()

	go func() {
		for result := range s.Results() {
			fmt.Println(result)
			database.SaveResult(result)
			metrics.ScansProcessed.Inc()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Reconnaissance Service")
	})

	http.ListenAndServe(":8080", nil)
}
