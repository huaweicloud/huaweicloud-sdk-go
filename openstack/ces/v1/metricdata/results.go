package metricdata

import (
	"github.com/gophercloud/gophercloud"
)

type MetricData struct {
	// Specifies the namespace in service.
	Namespace string `json:"namespace"`

	// The value can be a string of 1 to 64 characters
	// and must start with a letter and contain only uppercase
	// letters, lowercase letters, digits, and underscores.
	MetricName string `json:"metric_name"`

	//Specifies the list of the metric dimensions.
	Dimensions []map[string]interface{} `json:"dimensions"`

	// Specifies the metric data list.
	Datapoints []map[string]interface{} `json:"datapoints"`

	// Specifies the metric unit.
	Unit string `json:"unit"`
}

type MetricDatasResult struct {
	gophercloud.Result
}

// ExtractMetricDatas is a function that accepts a result and extracts metric datas.
func (r MetricDatasResult) ExtractMetricDatas() ([]MetricData, error) {
	var s struct {
		// Specifies the metric data.
		MetricDatas []MetricData `json:"metrics"`
	}
	err := r.ExtractInto(&s)
	return s.MetricDatas, err
}
