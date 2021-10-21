package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ronen25/task2/instrumentation"
)

type QueryPrinterHTTPService struct {
}

func NewHTTPService() QueryPrinterHTTPService {
	return QueryPrinterHTTPService{}
}

func (service *QueryPrinterHTTPService) updateTotalCounters() {
	instrumentation.Instrumentation.TotalRequests.Add(1)
	instrumentation.Instrumentation.TotalRequestsHTTP.Add(1)
}

func (service *QueryPrinterHTTPService) updateCountersBad() {
	service.updateTotalCounters()

	instrumentation.Instrumentation.TotalBadRequests.Add(1)
}

func (service *QueryPrinterHTTPService) updateCountersGood() {
	service.updateTotalCounters()

	instrumentation.Instrumentation.TotalGoodRequests.Add(1)
}

func (service *QueryPrinterHTTPService) QueryPrintingHandler(w http.ResponseWriter, r *http.Request) {
	var params []string
	query := r.URL.Query()

	for _, value := range query {
		combined := strings.Join(value, "")
		params = append(params, combined)
	}

	// We only expect 3 parameters, anything else is an error
	if len(params) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected 3 query parameters, got %d", len(params))

		service.updateCountersBad()

		return
	}

	// At this point we are certain there are 3 params so just format them.
	fmt.Fprintf(w, "%s-%s-%s", params[0], params[1], params[2])

	service.updateCountersGood()
}
