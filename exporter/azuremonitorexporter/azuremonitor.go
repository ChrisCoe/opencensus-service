package azuremonitorexporter

import (
	"github.com/spf13/viper"

	// TODO: once this repository has been transferred to the
	// official census-ecosystem location, update this import path.
	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor"
	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor/common"

	"github.com/census-instrumentation/opencensus-service/consumer"
	"github.com/census-instrumentation/opencensus-service/exporter/exporterwrapper"
)

type azuermonitorconfig struct {
	InstrumentationKey string `mapstructure:"instrumentationKey"`
}

// JaegerExportersFromViper unmarshals the viper and returns exporter.TraceExporters targeting
// Jaeger according to the configuration settings.
func AzureMonitorExportersFromViper(v *viper.Viper) (tps []consumer.TraceConsumer, mps []consumer.MetricsConsumer, doneFns []func() error, err error) {
	var cfg struct { // cfg stands for config. I am following the naming convention 
					 // used for all the exporters in this package
		AzureMonitor *azuermonitorconfig `mapstructure:"azuremonitor"`
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, nil, nil, err
	}
	amc := cfg.AzureMonitor
	if amc == nil {
		return nil, nil, nil, nil
	}

	return
}
