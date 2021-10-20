package service

import "github.com/ronen25/task2/service/grpc"

type QueryPrinterGRPCService struct {
}

func NewQueryPrinterGRPCService() *QueryPrinterGRPCService {
	return &QueryPrinterGRPCService{}
}

func (svc *QueryPrinterGRPCService) PrintParameters(r *grpc.Request) (*grpc.Response, error) {
	// TODO
	return nil, nil
}
