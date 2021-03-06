# Fizz-Buzz REST Server

Calculates the Fizz-Buzz sequence from number `1` to `limit`. Multiples of `int1` will be replaced with `str1`,
multiples of `int2` with `str2`, and multiples of `int1 and int2` with `str1str2`.

The arguments will be cached for calculating statistics about the most frequent requests.

## API

- POST /fizzbuzz

  ```shell
  curl localhost:8080/fizzbuzz -X POST -d '{ "int1": 3, "int2": 5, "limit": 15, "str1": "Fizz", "str2": "Buzz" }'
  ```

  ```json
  {
    "result": "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz"
  }
  ```

- GET /statistics

  ```shell
  curl localhost:8080/statistics
  ```

  ```json
  {
    "highest_score": "1",
    "requests": [
      {
        "int1": "3",
        "int2": "5",
        "limit": "15",
        "str1": "Fizz",
        "str2": "Buzz"
      }
    ]
  }
  ```

### Health probes

- GET /healthy

  ```shell
  curl localhost:8080/healthy
  ```

  ```json
  {
    "result": "ok"
  }
  ```

- GET /ready

  ```shell
  curl localhost:8080/ready
  ```

  ```json
  {
    "result": "ok"
  }
  ```

## Storage

The statistics will be cached with [Redis](https://redis.io/) and saved to the volume `fizz-buzz-rest-server_data`.

## Run the Fizz-Buzz server from source

1. Start the Redis service, which needs to be accessible on localhost (the default endpoint).

   ```shell
   docker run -d --name redis --rm -p 6379:6379 -v fizz-buzz-rest-server_data:/data redis:6.2.5-alpine --appendonly yes
   ```
2. Run `go run .` to start the server. You can set the Redis endpoint with `-redisEndpoint localhost:6379` and wait for
   it with `-waitForRedis=true`.

## Run the Fizz-Buzz server as a multi-container application

Run `docker-compose up -d` to start the server together with a Redis instance.

Run `docker-compose down` to stop the server and Redis.

## Run the Fizz-Buzz server on Kubernetes (docker-desktop)

Run `kubectl apply -f deployment.yaml` to start the server together with a Redis instance.

Run `kubectl port-forward services/fizz-buzz -n fizz-buzz 8080:8080` to access the server via `localhost:8080`.

Run `kubectl delete -f deployment.yaml` to stop the server and Redis.

## Requirements

- Golang: https://go.dev

- Protocol buffers v3: https://developers.google.com/protocol-buffers/docs/downloads

- Protocol buffers plugin for Go: https://developers.google.com/protocol-buffers/docs/gotutorial

  ```shell
  go install google.golang.org/protobuf/cmd/protoc-gen-go
  ```

- gRPC: https://grpc.io/blog/installation

  ```shell
  go get google.golang.org/grpc
  ```

- gRPC Gateway: https://github.com/grpc-ecosystem/grpc-gateway
    - v1: https://github.com/grpc-ecosystem/grpc-gateway/tree/v1

  ```shell
  go install \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        github.com/golang/protobuf/protoc-gen-go
  ```

## Build

Run `make all` to compile the protobuf files, generate the gRPC client/server code, and build the Go code.

Run `make docker` to build and push the docker image, with the tag set in .env file, to the docker hub.

Run `make docker FIZZ_BUZZ_TAG=x.y.z` to build and push the docker image with tag `x.y.z` to the docker hub.
