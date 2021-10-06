/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.21
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchContractEstimateGasRequest Estimate the gas for a contract function interaction. For a state-changing methods `walletId` is required.
type SmartBchContractEstimateGasRequest struct {
	// Serialized wallet
	WalletId string `json:"walletId,omitempty"`
	// serialized contract
	ContractId string `json:"contractId"`
	// Function to call on the SmartBch contract.
	Function string `json:"function"`
	// Arguments for the contract function call. Binary and BigNumber data should be passed as hexadecimal. 
	Arguments []interface{} `json:"arguments,omitempty"`
	Overrides SmartBchOverrides `json:"overrides,omitempty"`
}
