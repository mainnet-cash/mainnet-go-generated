/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.34-alpha.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateSignedMessageRequest struct for CreateSignedMessageRequest
type CreateSignedMessageRequest struct {
	WalletId string `json:"walletId,omitempty"`
	Message *interface{} `json:"message,omitempty"`
}
