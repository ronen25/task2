ALL: build

protos:
	@echo "--- Regenerating protobufs"
	(protoc \
	    -I . \
	    -I ${GOPATH}/bin \
	    --gogofaster_out=plugins=grpc,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:. \
	    --govalidators_out=gogoimport=true,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:. \
	    service/grpc/*.proto)

build: protos
	@echo "--- Building application"
	go build -o task2 ./cmd/task2

clean:
	@echo "--- Cleaning binaries and protobuf outputs"
	@rm task2
	@rm service/grpc/*.pb.go