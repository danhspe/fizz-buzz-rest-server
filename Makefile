PROTOC=protoc
PROTO_DIR=models/pb
IMPORTS=-I/usr/local/include -I$(GOPATH)/src -I$(DIR_VENDOR)/github.com/googleapis/googleapis -I$(PROTO_DIR)
#-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
SOURCES=$(PROTO_DIR)/*.proto
GW_SOURCES=$(PROTO_DIR)/*service.proto

DIR_BIN=bin
DIR_GOLIB=golib
DIR_TEMP=tmp
DIR_VENDOR=vendor
DIRS=$(DIR_BIN)/ $(DIR_GOLIB)/ $(DIR_TEMP)/ $(DIR_VENDOR)/

DIR_MODULE_PATH=github.com/danhspe/fizz-buzz-rest-server

.PHONY: all clean build models vendor

all: clean vendor models build

clean:
	rm -rf $(DIRS)
	mkdir -p $(DIRS)

build:
	go mod tidy
	go mod download
	go build -o $(DIR_BIN)/main -mod=readonly

models:
	# generate go grpc server & client
	$(PROTOC) $(IMPORTS) --go_out=plugins=grpc:$(DIR_TEMP) $(SOURCES)
	# generate grpc gateway
	$(PROTOC) $(IMPORTS) --grpc-gateway_out=logtostderr=true:$(DIR_TEMP) $(GW_SOURCES)
	# move generated files
	mv $(DIR_TEMP)/$(DIR_MODULE_PATH)/$(DIR_GOLIB)/* $(DIR_GOLIB)
	rm -rf $(DIR_TEMP)

test:
	go test -race -mod=readonly ./...

vendor:
	# third-party proto files
	git clone --depth 1 https://github.com/googleapis/googleapis $(DIR_VENDOR)/github.com/googleapis/googleapis
