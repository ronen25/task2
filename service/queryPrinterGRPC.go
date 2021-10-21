package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/ronen25/task2/instrumentation"
	"github.com/ronen25/task2/service/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueryPrinterGRPCService struct {
	protos.UnimplementedQueryPrinterGRPCServer
}

func NewQueryPrinterGRPCService() *QueryPrinterGRPCService {
	return &QueryPrinterGRPCService{}
}

func (service *QueryPrinterGRPCService) updateTotalCounters() {
	instrumentation.Instrumentation.TotalRequests.Add(1)
	instrumentation.Instrumentation.TotalRequestsGRPC.Add(1)
}

func (service *QueryPrinterGRPCService) updateCountersBad() {
	service.updateTotalCounters()

	instrumentation.Instrumentation.TotalBadRequests.Add(1)
}

func (service *QueryPrinterGRPCService) updateCountersGood() {
	service.updateTotalCounters()

	instrumentation.Instrumentation.TotalGoodRequests.Add(1)
}

func (svc *QueryPrinterGRPCService) PrintParameters(ctx context.Context, r *protos.Request) (*protos.Response, error) {
	svc.updateTotalCounters()

	// Check for 3 parameters. We expect the query to be similar
	// to the HTTP format.
	params := make(map[string]string)
	values := strings.Split(r.Query, "&")

	if len(values) != 3 {
		err := status.Error(codes.InvalidArgument,
			fmt.Sprintf("Expected 3 arguments, got %d", len(values)))

		svc.updateCountersBad()
		return nil, err
	}

	// Parse parameters
	for _, value := range values {
		parts := strings.Split(value, "=")
		params[parts[0]] = parts[1]
	}

	var valueParams []string
	for _, v := range params {
		valueParams = append(valueParams, v)
	}

	svc.updateCountersGood()

	return &protos.Response{
		Param1: valueParams[0],
		Param2: valueParams[1],
		Param3: valueParams[2],
	}, nil
}
