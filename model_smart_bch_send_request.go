/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.30
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSendRequest struct for SmartBchSendRequest
type SmartBchSendRequest struct {
	// The walletId of the source of funds to spend from. 
	WalletId string `json:"walletId"`
	To AnyOfSmartBchSendRequestItemarray `json:"to"`
	Options SmartBchSendRequestOptions `json:"options,omitempty"`
	Overrides SmartBchOverrides `json:"overrides,omitempty"`
}
