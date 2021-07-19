/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.30
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ConvertRequest struct for ConvertRequest
type ConvertRequest struct {
	Value float32 `json:"value"`
	// Provided value unit of account.
	From string `json:"from"`
	// Unit of account to be returned
	To string `json:"to"`
}
