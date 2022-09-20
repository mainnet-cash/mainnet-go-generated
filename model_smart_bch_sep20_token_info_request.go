/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.5.8
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20TokenInfoRequest struct for SmartBchSep20TokenInfoRequest
type SmartBchSep20TokenInfoRequest struct {
	// serialized wallet
	WalletId string `json:"walletId"`
	// Token unique hexadecimal identifier
	TokenId string `json:"tokenId"`
}