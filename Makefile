PROTOC=protoc
PROTO_DIR=internal/models/pb
IMPORTS=-I/usr/local/include -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$(PROTO_DIR)
SOURCES=$(PROTO_DIR)/*.proto
GW_SOURCES=$(PROTO_DIR)/*service.proto

DOCKER_USER=spelmezan
IMAGE_NAME=fizz-buzz-rest-server
TAG_NAME=0.1.0

DIR_BIN=bin
DIR_GOLIB=golib
DIR_TEMP=tmp
DIRS=$(DIR_BIN)/ $(DIR_GOLIB)/ $(DIR_TEMP)/

DIR_MODULE_PATH=github.com/danhspe/fizz-buzz-rest-server

.PHONY: all clean build docker models

all: clean models test build

clean:
	rm -rf $(DIRS)
	mkdir -p $(DIRS)

build:
	go mod tidy
	go mod download
	go build -o $(DIR_BIN)/main

docker:
	docker build -t $(DOCKER_USER)/$(IMAGE_NAME):$(TAG_NAME) .
	docker login -u $(DOCKER_USER)
	docker push $(DOCKER_USER)/$(IMAGE_NAME):$(TAG_NAME)

models:
	# generate go grpc server & client
	$(PROTOC) $(IMPORTS) --go_out=plugins=grpc:$(DIR_TEMP) $(SOURCES)
	# generate grpc gateway
	$(PROTOC) $(IMPORTS) --grpc-gateway_out=logtostderr=true:$(DIR_TEMP) $(GW_SOURCES)
	# move generated files
	mv $(DIR_TEMP)/$(DIR_MODULE_PATH)/$(DIR_GOLIB)/* $(DIR_GOLIB)
	rm -rf $(DIR_TEMP)

test:
	go test -cover -race ./...
