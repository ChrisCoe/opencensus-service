// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azuremonitor

import (
	"github.com/spf13/viper"
	//"contrib.go.opencensus.io/exporter/jaeger"
	// (Use this for public use) "github.com/ChrisCoe/opencensus-go/tree/master/exporter/azure_monitor"
	// (Only developer should use import below while in progress)
	"go.opencensus.io/exporter/azure_monitor"
	//TODO: Once finished, transfer Azure Monitor Exporter to it's own repo like jaeger does

	"github.com/census-instrumentation/opencensus-service/consumer"
	"github.com/census-instrumentation/opencensus-service/exporter/exporterwrapper"
)

type azureMonitorConfig struct {
	InstrumentationKey string `mapstructure:"instrumentation_key"`
}

// AzureMonitorTraceExportersFromViper unmarshals the viper and returns exporter.TraceExporters targeting
// Azure Monitor according to the configuration settings. 
// For now, there is only a trace exporter
func AzureMonitorTraceExportersFromViper(v *viper.Viper) (tps []consumer.TraceConsumer, mps []consumer.MetricsConsumer, doneFns []func() error, err error) {
	var configStruct struct {
		AzureMonitor *azureMonitorConfig `mapstructure:"azuremonitor"`
	}
	if err := v.Unmarshal(&configStruct); err != nil {
		return nil, nil, nil, err
	}
	config := configStruct.AzureMonitor
	if config == nil {
		return nil, nil, nil, nil
	}

	rawExporter, err := azure_monitor.NewAzureTraceExporter("111a0d2f-ab53-4b62-a54f-4722f09fd136")
	if err != nil {
		log.Fatal(err)
	}

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("Cannot configure Azure Monitor Trace exporter: %v", err)
	}

	exporterWrapper, err := exporterwrapper.NewExporterWrapper("azuremonitor", "ocservice.exporter.AzureMonitor.ConsumeTraceData", rawExporter)
	if err != nil {
		return nil, nil, nil, err
	}
	return
}
