package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exp, err := ocagent.NewExporter(ocagent.WithInsecure(), ocagent.WithServiceName("your-service-name"))
	if err != nil {
		log.Fatalf("Failed to create the agent exporter: %v", err)
	}
	defer exp.Stop()

	// Now register it as a trace exporter.
	trace.RegisterExporter(exp)

	fmt.Println("START SPAN")
	// Then use the OpenCensus tracing library, like we normally would.
	ctx, span := trace.StartSpan(context.Background(), "AgentExporter-Example")
	defer span.End()

	for i := 0; i < 10; i++ {
		_, iSpan := trace.StartSpan(ctx, fmt.Sprintf("Sample-%d", i))
		<-time.After(6 * time.Millisecond)
		iSpan.End()
	}
	fmt.Println("END MAIN")
}
