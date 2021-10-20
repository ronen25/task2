ALL: build

protos:
	@echo "--- Regenerating protobufs"
	(protoc \
	    -I . \
	    -I ${GOPATH}/bin \
	    --go_out=service\
		--go-grpc_out=service\
	    service/protos/*.proto)

build: protos
	@echo "--- Building application"
	go build -o task2 ./cmd/task2

clean:
	@echo "--- Cleaning binaries and protobuf outputs"
	@rm task2
	@rm service/grpc/*.pb.go