/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.3
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VerifySignedMessageResponse struct for VerifySignedMessageResponse
type VerifySignedMessageResponse struct {
	// Whether the message was signed by a given address
	Valid bool `json:"valid,omitempty"`
	Details VerifySignedMessageResponseDetails `json:"details,omitempty"`
}
