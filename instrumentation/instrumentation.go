package instrumentation

import (
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type instrumentation struct {
	TotalRequests     prometheus.Counter
	TotalRequestsHTTP prometheus.Counter
	TotalRequestsGRPC prometheus.Counter
	TotalGoodRequests prometheus.Counter
	TotalBadRequests  prometheus.Counter

	HTTPHandler http.Handler
}

var Instrumentation instrumentation
var onceInstrumentation sync.Once

func NewInstrumentation() {
	onceInstrumentation.Do(func() {
		Instrumentation = instrumentation{
			TotalRequests: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace: "task2",
				Name:      "total_requests",
				Help:      "Total amount of requests handled by the service",
			}),
			TotalRequestsHTTP: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace: "task2",
				Name:      "total_requests_http",
				Help:      "Total amount of requests handled by the service on the HTTP endpoint only",
			}),
			TotalRequestsGRPC: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace: "task2",
				Name:      "total_requests_grpc",
				Help:      "Total amount of requests handled by the service on the GRPC endpoint only",
			}),
			TotalGoodRequests: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace: "task2",
				Name:      "total_good_requests",
				Help:      "Total amount of good requests handled by the service",
			}),
			TotalBadRequests: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace: "task2",
				Name:      "total_bad_requests",
				Help:      "Total amount of bad requests handled by the service",
			}),
			HTTPHandler: promhttp.Handler(),
		}

		Instrumentation.registerInstrumentation()
	})
}

func (ins *instrumentation) registerInstrumentation() {
	prometheus.MustRegister(ins.TotalRequests)
	prometheus.MustRegister(ins.TotalRequestsHTTP)
	prometheus.MustRegister(ins.TotalRequestsGRPC)
	prometheus.MustRegister(ins.TotalGoodRequests)
	prometheus.MustRegister(ins.TotalBadRequests)
}
