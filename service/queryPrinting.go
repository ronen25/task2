package service

import (
	"fmt"
	"net/http"
	"strings"
)

type QueryPrintingService struct {
}

func NewService() QueryPrintingService {
	return QueryPrintingService{}
}

func (service *QueryPrintingService) QueryPrintingHandler(w http.ResponseWriter, r *http.Request) {
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

		return
	}

	// At this point we are certain there are 3 params so just format them.
	fmt.Fprintf(w, "%s-%s-%s", params[0], params[1], params[2])
}
