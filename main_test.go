//go:generate mockgen -destination mocks/mock_trace_service_server.go -package mocks go.opentelemetry.io/proto/otlp/collector/trace/v1 TraceServiceServer
package main

import (
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/miparnisari/mockGo/mocks"
	otlpcollector "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/grpc"
)

func Test(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	var mockTraceServiceServer otlpcollector.TraceServiceServer = mocks.NewMockTraceServiceServer(mockController)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5132))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	otlpcollector.RegisterTraceServiceServer(server, mockTraceServiceServer)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
