/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.15
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20TokenInfoResponse struct for SmartBchSep20TokenInfoResponse
type SmartBchSep20TokenInfoResponse struct {
	// Token name
	Name string `json:"name,omitempty"`
	// Token ticker
	Ticker string `json:"ticker,omitempty"`
	TokenId string `json:"tokenId,omitempty"`
	// Indicates that 1 token is divisible into 10^decimals base units
	Decimals float32 `json:"decimals,omitempty"`
	// Total amount of tokens in circulation
	TotalSupply string `json:"totalSupply,omitempty"`
}
