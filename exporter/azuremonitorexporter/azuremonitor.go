package azuremonitorexporter

import (
	"github.com/spf13/viper"

	//"contrib.go.opencensus.io/exporter/jaeger"
	// TODO: once this repository has been transferred to the
	// official census-ecosystem location, update this import path.
	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor"
	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor/common"
	"log"
	"fmt"
	//"go.opencensus.io/trace"

	"github.com/census-instrumentation/opencensus-service/consumer"
	"github.com/census-instrumentation/opencensus-service/exporter/exporterwrapper"
)

type azuerMonitorConfig struct {
	InstrumentationKey string `mapstructure:"instrumentationKey"`
}

// JaegerExportersFromViper unmarshals the viper and returns exporter.TraceExporters targeting
// Jaeger according to the configuration settings.
func AzureMonitorExportersFromViper(v *viper.Viper) (tps []consumer.TraceConsumer, mps []consumer.MetricsConsumer, doneFns []func() error, err error) {
	var cfg struct {
		AzureMonitor *azuerMonitorConfig `mapstructure:"azuremonitor"`
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, nil, nil, err
	}
	amc := cfg.AzureMonitor
	if amc == nil {
		return nil, nil, nil, nil
	}

	// TODO: i am updating the constructor but first figure out import
	// // jaeger.NewExporter performs configurqtion validation
	fmt.Println("LAUGH Y")
	fmt.Println("This is ikey")
	fmt.Println(amc.InstrumentationKey)
	azureexporter, err := azuremonitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: amc.InstrumentationKey, // add your InstrumentationKey
	})
	if err != nil {
		log.Fatal(err)
	}
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

	amte, err := exporterwrapper.NewExporterWrapper("azuremonitor_trace", "ocservice.exporter.AzureMonitor.ConsumeTraceData", azureexporter)
	if err != nil {
		return nil, nil, nil, err
	}

	tps = append(tps, amte)
	return
}

