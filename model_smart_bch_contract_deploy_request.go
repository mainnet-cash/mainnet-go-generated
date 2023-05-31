/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.28
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchContractDeployRequest struct for SmartBchContractDeployRequest
type SmartBchContractDeployRequest struct {
	// Serialized wallet
	WalletId string `json:"walletId"`
	// solidity source code used to create this contract instance
	Script string `json:"script"`
	// constructor parameters used to create this contract instance
	Parameters []interface{} `json:"parameters,omitempty"`
	Overrides SmartBchOverrides `json:"overrides,omitempty"`
}
