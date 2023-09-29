/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.2.2
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WatchAddressRequest struct for WatchAddressRequest
type WatchAddressRequest struct {
	// Cash address to watch 
	Cashaddr string `json:"cashaddr"`
	// Url to be called when configured action is triggered
	Url string `json:"url"`
	// Type of watch operation
	Type string `json:"type"`
	// Action recurrence. Indicates if webhook should be triggered recurrently until expire or only once
	Recurrence string `json:"recurrence,omitempty"`
	// Duration of the webhook lifetime in seconds before it will expire.
	DurationSec float32 `json:"duration_sec,omitempty"`
}
