package azuremonitorexporter

import (
	"github.com/spf13/viper"

	//"contrib.go.opencensus.io/exporter/jaeger"
	// TODO: once this repository has been transferred to the
	// official census-ecosystem location, update this import path.
	"github.com/ChrisCoe/opencensus-go/tree/master/exporter/azure_monitor"

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
		AzureMonitor *azuerMonitorConfig `mapstructure:"azureMonitor"`
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

	amte, err := exporterwrapper.NewExporterWrapper("azureMonitor", "name", ame)
	if err != nil {
		return nil, nil, nil, err
	}

	tps = append(tps, amte)
	return
}

