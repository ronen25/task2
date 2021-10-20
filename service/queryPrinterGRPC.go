package service

import (
	"context"
	"log"

	"github.com/ronen25/task2/service/protos"
)

type QueryPrinterGRPCService struct {
}

func NewQueryPrinterGRPCService() *QueryPrinterGRPCService {
	return &QueryPrinterGRPCService{}
}

func (svc *QueryPrinterGRPCService) PrintParameters(ctx context.Context, r *protos.Request) (*protos.Response, error) {
	// TODO
	log.Printf("%v", r)
	return nil, nil
}
