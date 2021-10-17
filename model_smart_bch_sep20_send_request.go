/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.22
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20SendRequest struct for SmartBchSep20SendRequest
type SmartBchSep20SendRequest struct {
	// serialized wallet
	WalletId string `json:"walletId"`
	To []SmartBchSep20SendRequestItem `json:"to"`
	Overrides SmartBchOverrides `json:"overrides,omitempty"`
}
