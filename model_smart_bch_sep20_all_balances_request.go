/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.14
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20AllBalancesRequest struct for SmartBchSep20AllBalancesRequest
type SmartBchSep20AllBalancesRequest struct {
	// serialized wallet
	WalletId string `json:"walletId"`
	Options map[string]interface{} `json:"options,omitempty"`
}
