/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.15
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchOverrides struct for SmartBchOverrides
type SmartBchOverrides struct {
	GasLimit AnyOfnumberstring `json:"gasLimit,omitempty"`
	GasPrice AnyOfnumberstring `json:"gasPrice,omitempty"`
	MaxFeePerGas AnyOfnumberstring `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas AnyOfnumberstring `json:"maxPriorityFeePerGas,omitempty"`
	Nonce AnyOfnumberstring `json:"nonce,omitempty"`
	Value AnyOfnumberstring `json:"value,omitempty"`
	BlockTag AnyOfnumberstring `json:"blockTag,omitempty"`
	From string `json:"from,omitempty"`
}
