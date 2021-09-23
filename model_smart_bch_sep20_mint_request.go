/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.11
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20MintRequest struct for SmartBchSep20MintRequest
type SmartBchSep20MintRequest struct {
	// serialized wallet
	WalletId string `json:"walletId"`
	// Value is represented as number or string to avoid precision loss
	Value AnyOfnumberstring `json:"value"`
	TokenId string `json:"tokenId"`
	TokenReceiverAddress string `json:"tokenReceiverAddress,omitempty"`
	Overrides SmartBchOverrides `json:"overrides,omitempty"`
}
