/*
 * Detectors API
 *
 * **Detectors** define rules for identifying conditions of interest to the customer, and the notifications to send when the conditions occur or stop occurring.
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package detector

type SearchResults struct {
	// Number of objects that match the search query. If you use paging, `count` is either the value of the `limit` request parameter or the number of objects still undelivered after the request reached the `offset` request parameter.<br> *Note:* If you use paging, count is not the number of objects returned in the response.
	Count int32 `json:"count,omitempty"`
	// An array of detector definitions that match the request criteria. Each definition is defined as a `Detector`.
	Results []Detector `json:"results,omitempty"`
}