package main
// Package: Runs code 

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	agenttracepb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/trace/v1"
)

func main() {
	var deferFuncs []func() error
	lis, err := net.Listen("tcp", ":55678")
	if err != nil {
		log.Fatalf("Failed to get an address: %v", err)
	}

	// create a trace server instance
	s := agenttracepb.TraceServiceServer{} //look up how to do this properly and ask Odeke

	deferFuncs = append(deferFuncs, lis.Close)

	// create a gRPC server object
	srv := grpc.NewServer()

	//ma := makeMockAgent(t)

	// attach the trace service to the server
	agenttracepb.RegisterTraceServiceServer(srv, s) // yes? s was ma

	// start the server
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve pollo: %s", err)
	}

	fmt.Println("We got herea")
	// go func() {
	// 	_ = srv.Serve(lis)
	// }()

	// deferFunc := func() error {
	// 	srv.Stop()
	// 	return lis.Close()
	// }
}
