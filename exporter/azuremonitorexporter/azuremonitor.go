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

// AzureMonitorExportersFromViper unmarshals the viper and returns exporter.TraceExporters targeting
// Azure Monitor according to the configuration settings.
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
	azureexporter, err := azuremonitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: amc.InstrumentationKey, // add InstrumentationKey
	})

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

	tps = append(tps, amte)
	return
}
