/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.26
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20BalanceResponse struct for SmartBchSep20BalanceResponse
type SmartBchSep20BalanceResponse struct {
	// A big number represented as a string to avoid precision loss
	Value float32 `json:"value,omitempty"`
	// Token ticker
	Ticker string `json:"ticker,omitempty"`
	// Token name
	Name string `json:"name,omitempty"`
	// Token unique hexadecimal identifier
	TokenId string `json:"tokenId,omitempty"`
}
