/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.7.14
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EncodeTransactionRequest struct for EncodeTransactionRequest
type EncodeTransactionRequest struct {
	// The walletId of the source of funds to spend from. 
	WalletId string `json:"walletId,omitempty"`
	DiscardChange bool `json:"discardChange,omitempty"`
	To AnyOfSendRequestItemarrayOpReturnData `json:"to,omitempty"`
	Options SendRequestOptions `json:"options,omitempty"`
}
