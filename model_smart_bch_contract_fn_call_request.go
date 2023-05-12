/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.15
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchContractFnCallRequest Contract function evaluation request. A constant function (non state changing, also does not cost gas) does not require a signer (walletId) A state changing function will generate a transaction and does require a signer (walletId) 
type SmartBchContractFnCallRequest struct {
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
