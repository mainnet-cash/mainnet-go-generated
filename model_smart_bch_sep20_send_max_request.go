/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.14
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20SendMaxRequest struct for SmartBchSep20SendMaxRequest
type SmartBchSep20SendMaxRequest struct {
	// serialized wallet
	WalletId string `json:"walletId"`
	Address string `json:"address"`
	// Token unique hexadecimal identifier
	TokenId string `json:"tokenId"`
}
