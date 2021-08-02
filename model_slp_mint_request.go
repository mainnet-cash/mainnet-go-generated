/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.40
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpMintRequest struct for SlpMintRequest
type SlpMintRequest struct {
	// ID that is returned in `wallet` field of /wallet call 
	WalletId string `json:"walletId"`
	// Value is represented as a string to avoid precision loss
	Value string `json:"value"`
	TokenId string `json:"tokenId"`
	// Indicates if token should not be 'mintable' anymore, e.g. total circulation amount increased
	EndBaton bool `json:"endBaton,omitempty"`
	TokenReceiverSlpAddr string `json:"tokenReceiverSlpAddr,omitempty"`
	BatonReceiverSlpAddr string `json:"batonReceiverSlpAddr,omitempty"`
}
