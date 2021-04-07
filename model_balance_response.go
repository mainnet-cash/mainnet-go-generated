/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.12
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BalanceResponse struct for BalanceResponse
type BalanceResponse struct {
	// Amount in whole Bitcoin Cash
	Bch float32 `json:"bch,omitempty"`
	// Amount in satoshis
	Sat int32 `json:"sat,omitempty"`
	// Amount in United States Dollars
	Usd float32 `json:"usd,omitempty"`
}