/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.32
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpSendRequestItem struct for SlpSendRequestItem
type SlpSendRequestItem struct {
	Slpaddr string `json:"slpaddr"`
	Value float32 `json:"value"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId"`
}
