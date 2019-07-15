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

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	// Now register it as a trace exporter.
	trace.RegisterExporter(exp)

	fmt.Println("START SPAN")
	// Then use the OpenCensus tracing library, like we normally would.
	ctx, span := trace.StartSpan(ctx, "AgentExporter-Example")
	defer span.End()

	_, span2 := trace.StartSpan(context.Background(), "/boi") // This calls the function ExportSpan written in azuremonitor.go 
	span2.End()

	for i := 0; i < 5; i++ {
		_, iSpan := trace.StartSpan(ctx, fmt.Sprintf("Sample-%d", i))
		<-time.After(6 * time.Millisecond)
		iSpan.End()
	}
	fmt.Println("END MAIN")
}
