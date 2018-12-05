package metricdata

import (
	"github.com/gophercloud/gophercloud"
)

// BatchQueryOptsBuilder allows extensions to add additional parameters to the
// BatchQuery request.
type BatchQueryOptsBuilder interface {
	ToBatchQueryOptsMap() (map[string]interface{}, error)
}

type Metric struct {
	// Specifies the namespace in service.
	Namespace string `json:"namespace" required:"true"`

	// The value can be a string of 1 to 64 characters
	// and must start with a letter and contain only uppercase
	// letters, lowercase letters, digits, and underscores.
	MetricName string `json:"metric_name" required:"true"`

	// Specifies the list of the metric dimensions.
	Dimensions []map[string]string `json:"dimensions" required:"true"`
}

// BatchQueryOpts represents options for batch query metric data.
type BatchQueryOpts struct {
	// Specifies the metric data.
	Metrics []Metric `json:"metrics" required:"true"`

	// Specifies the start time of the query.
	From int64 `json:"from" required:"true"`

	// Specifies the end time of the query.
	To int64 `json:"to" required:"true"`

	// Specifies the data monitoring granularity.
	Period string `json:"period" required:"true"`

	// Specifies the data rollup method.
	Filter string `json:"filter" required:"true"`
}

// ToBatchQueryOptsMap builds a request body from BatchQueryOpts.
func (opts BatchQueryOpts) ToBatchQueryOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Querying Monitoring Data in Batches.
func BatchQuery(client *gophercloud.ServiceClient, opts BatchQueryOptsBuilder) (r MetricDatasResult) {
	b, err := opts.ToBatchQueryOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchQueryMetricDataURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},})
	return
}
