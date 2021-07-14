// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: service.proto

/*
Package fizzbuzz is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package fizzbuzz

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_FizzBuzzService_GetFizzBuzz_0(ctx context.Context, marshaler runtime.Marshaler, client FizzBuzzServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FizzBuzzRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFizzBuzz(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FizzBuzzService_GetFizzBuzz_0(ctx context.Context, marshaler runtime.Marshaler, server FizzBuzzServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FizzBuzzRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFizzBuzz(ctx, &protoReq)
	return msg, metadata, err

}

func request_FizzBuzzService_GetStatistics_0(ctx context.Context, marshaler runtime.Marshaler, client FizzBuzzServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetStatistics(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FizzBuzzService_GetStatistics_0(ctx context.Context, marshaler runtime.Marshaler, server FizzBuzzServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetStatistics(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterFizzBuzzServiceHandlerServer registers the http handlers for service FizzBuzzService to "mux".
// UnaryRPC     :call FizzBuzzServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterFizzBuzzServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server FizzBuzzServiceServer) error {

	mux.Handle("POST", pattern_FizzBuzzService_GetFizzBuzz_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FizzBuzzService_GetFizzBuzz_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FizzBuzzService_GetFizzBuzz_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FizzBuzzService_GetStatistics_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FizzBuzzService_GetStatistics_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FizzBuzzService_GetStatistics_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterFizzBuzzServiceHandlerFromEndpoint is same as RegisterFizzBuzzServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFizzBuzzServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterFizzBuzzServiceHandler(ctx, mux, conn)
}

// RegisterFizzBuzzServiceHandler registers the http handlers for service FizzBuzzService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFizzBuzzServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFizzBuzzServiceHandlerClient(ctx, mux, NewFizzBuzzServiceClient(conn))
}

// RegisterFizzBuzzServiceHandlerClient registers the http handlers for service FizzBuzzService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "FizzBuzzServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FizzBuzzServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FizzBuzzServiceClient" to call the correct interceptors.
func RegisterFizzBuzzServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FizzBuzzServiceClient) error {

	mux.Handle("POST", pattern_FizzBuzzService_GetFizzBuzz_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FizzBuzzService_GetFizzBuzz_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FizzBuzzService_GetFizzBuzz_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_FizzBuzzService_GetStatistics_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FizzBuzzService_GetStatistics_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FizzBuzzService_GetStatistics_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_FizzBuzzService_GetFizzBuzz_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"fizzbuzz"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_FizzBuzzService_GetStatistics_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"statistics"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_FizzBuzzService_GetFizzBuzz_0 = runtime.ForwardResponseMessage

	forward_FizzBuzzService_GetStatistics_0 = runtime.ForwardResponseMessage
)