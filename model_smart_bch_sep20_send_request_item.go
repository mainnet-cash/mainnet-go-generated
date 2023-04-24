/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.7
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20SendRequestItem struct for SmartBchSep20SendRequestItem
type SmartBchSep20SendRequestItem struct {
	Address string `json:"address"`
	Value AnyOfnumberstring `json:"value"`
	// Token unique hexadecimal identifier
	TokenId string `json:"tokenId"`
}
