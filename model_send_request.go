/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.3.15
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SendRequest struct for SendRequest
type SendRequest struct {
	// The walletId of the source of funds to spend from. 
	WalletId string `json:"walletId,omitempty"`
	To AnyOfSendRequestItemarrayTokenSendRequestOpReturnData `json:"to,omitempty"`
	Options SendRequestOptions `json:"options,omitempty"`
}
