/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.16
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20GenesisRequest struct for SmartBchSep20GenesisRequest
type SmartBchSep20GenesisRequest struct {
	// serialized wallet
	WalletId string `json:"walletId"`
	// Token name
	Name string `json:"name"`
	// Token ticker
	Ticker string `json:"ticker"`
	// Value is represented as number or string to avoid precision loss
	InitialAmount AnyOfnumberstring `json:"initialAmount"`
	// Indicates that 1 token is divisible into 10^decimals base units
	Decimals float32 `json:"decimals"`
	// Indicates if token should not be 'mintable', e.g. total circulation amount increased
	EndBaton bool `json:"endBaton,omitempty"`
	TokenReceiverAddress string `json:"tokenReceiverAddress,omitempty"`
	BatonReceiverAddress string `json:"batonReceiverAddress,omitempty"`
	Overrides SmartBchOverrides `json:"overrides,omitempty"`
}
