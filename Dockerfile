FROM golang:1.17-alpine

# Container listens on HTTP port and a GRPC port
ENV HTTP_PORT=6001
ENV GRPC_PORT=9002

# Copy all sources and build
COPY . /task2

WORKDIR /task2
RUN apk add --no-cache protobuf git make \
  && go install github.com/golang/protobuf/protoc-gen-go \
  && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  && cp /go/bin/protoc-gen-go /usr/local/bin/
RUN make

EXPOSE ${HTTP_PORT}
EXPOSE ${GRPC_PORT}

CMD ["/task2/task2"]
