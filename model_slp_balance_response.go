/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.34-alpha.1
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpBalanceResponse struct for SlpBalanceResponse
type SlpBalanceResponse struct {
	// A big number represented as a string to avoid precision loss
	Value float32 `json:"value,omitempty"`
	// Token ticker
	Ticker string `json:"ticker,omitempty"`
	// Full name of the SLP token
	Name string `json:"name,omitempty"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId,omitempty"`
	// Token type. 0x01 Type1, 0x81 NFT Parent, 0x41 NFT Child
	Type float32 `json:"type,omitempty"`
}
