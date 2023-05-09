/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.13
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSendMaxRequest struct for SmartBchSendMaxRequest
type SmartBchSendMaxRequest struct {
	// The walletId of the sender 
	WalletId string `json:"walletId"`
	Address string `json:"address"`
	Options SmartBchSendRequestOptions `json:"options,omitempty"`
	Overrides SmartBchOverrides `json:"overrides,omitempty"`
}
