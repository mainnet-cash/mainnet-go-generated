/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.41
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SignedMessageResponse struct for SignedMessageResponse
type SignedMessageResponse struct {
	Signature string `json:"signature,omitempty"`
	Raw SignedMessageResponseRaw `json:"raw,omitempty"`
	Details SignedMessageResponseDetails `json:"details,omitempty"`
}
