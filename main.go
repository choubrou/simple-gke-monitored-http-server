package main

import "fmt"
import "net/http"
import "time"
import "github.com/google/uuid"
import "github.com/prometheus/client_golang/prometheus"
import "github.com/prometheus/client_golang/prometheus/promauto"
import "github.com/prometheus/client_golang/prometheus/promhttp"

var coolCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "cool_calls_total",
	Help: "The total number of calls to the cool handler.",
})

func main() {
	newID := uuid.New()

	hello := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "hello")
	}

	coolHandler := func(w http.ResponseWriter, req *http.Request) {
		coolCounter.Inc()
		fmt.Fprintln(w, "my ID is", newID, "and the current time is", time.Now())
	}

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/cool", coolHandler)
	http.HandleFunc("/", hello)
	fmt.Println("launching the http server...")
	http.ListenAndServe(":8080", nil)
}
