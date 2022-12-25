/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.4
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpSendMaxRequest struct for SlpSendMaxRequest
type SlpSendMaxRequest struct {
	// The walletId of the sending wallet 
	WalletId string `json:"walletId"`
	Slpaddr string `json:"slpaddr"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId"`
	Options SlpSendRequestOptions `json:"options,omitempty"`
}
