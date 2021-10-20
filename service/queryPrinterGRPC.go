package service

import (
	"context"
	"log"

	"github.com/ronen25/task2/service/protos"
)

type QueryPrinterGRPCService struct {
	protos.UnimplementedQueryPrinterGRPCServer
}

func NewQueryPrinterGRPCService() *QueryPrinterGRPCService {
	return &QueryPrinterGRPCService{}
}

func (svc *QueryPrinterGRPCService) PrintParameters(ctx context.Context, r *protos.Request) (*protos.Response, error) {
	// TODO
	log.Printf("%v", r)
	return &protos.Response{
		Param1: "",
		Param2: "",
		Param3: "",
	}, nil
}
