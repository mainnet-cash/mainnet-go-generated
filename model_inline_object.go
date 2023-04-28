/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.8
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject struct for InlineObject
type InlineObject struct {
	// The walletId to make a request to.
	WalletId string `json:"walletId"`
	// tokenId (category) to filter utxos by, if not set will return utxos from all tokens
	TokenId string `json:"tokenId,omitempty"`
}
