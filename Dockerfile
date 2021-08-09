FROM golang:1.16-alpine AS builder

RUN apk update

WORKDIR $GOPATH/src/fizz-buzz-rest-server/
COPY go.mod ./
COPY go.sum ./
COPY main.go main.go
COPY golib/ golib/
COPY internal/ internal/

RUN go mod tidy
RUN go mod download
RUN go mod verify

## small binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/fizz-buzz-rest-server main.go

FROM alpine
RUN mkdir /fizz-buzz-rest-server
COPY --from=builder /go/bin/fizz-buzz-rest-server /go/bin/fizz-buzz-rest-server
ENTRYPOINT ["/go/bin/fizz-buzz-rest-server"]
