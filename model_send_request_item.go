/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.4.1
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SendRequestItem struct for SendRequestItem
type SendRequestItem struct {
	Cashaddr string `json:"cashaddr"`
	Value float32 `json:"value"`
}
