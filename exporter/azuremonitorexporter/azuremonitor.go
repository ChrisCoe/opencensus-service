package azuremonitorexporter

import (
	"github.com/spf13/viper"

	//"contrib.go.opencensus.io/exporter/jaeger"
	// TODO: once this repository has been transferred to the
	// official census-ecosystem location, update this import path.
	"fmt"
	"log"

	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor"
	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor/common"

	"go.opencensus.io/trace"
	//"context"

	"github.com/census-instrumentation/opencensus-service/consumer"
	"github.com/census-instrumentation/opencensus-service/exporter/exporterwrapper"
)

type azuermonitorconfig struct {
	InstrumentationKey string `mapstructure:"instrumentationKey"`
}

// JaegerExportersFromViper unmarshals the viper and returns exporter.TraceExporters targeting
// Jaeger according to the configuration settings.
func AzureMonitorExportersFromViper(v *viper.Viper) (tps []consumer.TraceConsumer, mps []consumer.MetricsConsumer, doneFns []func() error, err error) {
	var cfg struct {
		AzureMonitor *azuermonitorconfig `mapstructure:"azuremonitor"`
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, nil, nil, err
	}
	amc := cfg.AzureMonitor
	if amc == nil {
		return nil, nil, nil, nil
	}

	// TODO: i am updating the constructor but first figure out import
	fmt.Println("LAUGH Y")
	fmt.Println("This is ikey")
	fmt.Println(amc.InstrumentationKey)
	azureexporter, err := azuremonitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: amc.InstrumentationKey, // add your InstrumentationKey
	})
	if err != nil {
		log.Fatal(err)
	}

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(azureexporter)

	fmt.Println("My azure trace exporter:")
	fmt.Println(azureexporter)
	fmt.Println("END STUFF")

	// ame, err := jaeger.NewExporter(jaeger.Options{
	// 	CollectorEndpoint: jc.CollectorEndpoint,
	// 	Username:          jc.Username,
	// 	Password:          jc.Password,
	// 	Process: jaeger.Process{
	// 		ServiceName: jc.ServiceName,
	// 	},
	// })

	if err != nil {
		return nil, nil, nil, err
	}

	doneFns = append(doneFns, func() error {
		return nil
	})

	amte, err := exporterwrapper.NewExporterWrapper("azuremonitor", "ocservice.exporter.AzureMonitor.ConsumeTraceData", azureexporter)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Println("My wrapper azure trace exporter:")
	fmt.Println(amte)

	doneFns = append(doneFns, func() error {
		return nil
	})
	tps = append(tps, amte)
	return
}
