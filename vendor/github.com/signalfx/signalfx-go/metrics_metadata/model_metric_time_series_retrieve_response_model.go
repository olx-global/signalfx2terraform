/*
 * Metrics Metadata API
 *
 * API for creating, retrieving, updating, and deleting metric names and MTS metadata.<br> **NOTE:*() Although you can't set custom properties or tags for a metric, you *can* retrieve them for metrics and metric time series (**MTS**).
 *
 * API version: 3.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metrics_metadata

// JSON returned by the GET method of the `/metrictimeseries` endpoint
type MetricTimeSeriesRetrieveResponseModel struct {
	// An array of JSON objects (dictionaries). Each object contains the metadata for one MTS that matched the query.
	Results []MetricTimeSeries `json:"results,omitempty"`
	// The number of objects the API returned. This is the same as the size of the `results` property array.
	Count int32 `json:"count,omitempty"`
}
