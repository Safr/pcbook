LOCAL_BIN:=$(CURDIR)/bin/

PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc
PROTO_PATH:=$(CURDIR)/proto

install-grpc:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
gen:
	$(PROTOC) --proto_path=$(PROTO_PATH) --go_out=$(CURDIR) $(PROTO_PATH)/*.proto
clean:
	rm pb/*.go
run:
	go run cmd/main.go
install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml
test:
	go test -cover -race ./...
