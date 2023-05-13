/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.16
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpTokenInfoResponse struct for SlpTokenInfoResponse
type SlpTokenInfoResponse struct {
	// Token name
	Name string `json:"name,omitempty"`
	// Token ticker
	Ticker string `json:"ticker,omitempty"`
	TokenId string `json:"tokenId,omitempty"`
	// Value is represented as a string to avoid precision loss
	InitialAmount string `json:"initialAmount,omitempty"`
	// Indicates that 1 token is divisible into 10^decimals base units
	Decimals float32 `json:"decimals,omitempty"`
	DocumentUrl string `json:"documentUrl,omitempty"`
	// Document hash of the token. Empty or 64 character long hex string.
	DocumentHash string `json:"documentHash,omitempty"`
	// Token type. 0x01 Type1, 0x81 NFT Parent, 0x41 NFT Child
	Type float32 `json:"type,omitempty"`
}
