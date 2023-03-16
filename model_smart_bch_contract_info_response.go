/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.17
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchContractInfoResponse struct for SmartBchContractInfoResponse
type SmartBchContractInfoResponse struct {
	// serialized contract
	ContractId string `json:"contractId,omitempty"`
	// Address of an already deployed contract
	Address string `json:"address,omitempty"`
	// Contract ABI (Application Binary Interface), which describes the contract interaction
	Abi []string `json:"abi,omitempty"`
	// solidity source code used to create this contract instance
	Script string `json:"script,omitempty"`
	// constructor parameters used to create this contract instance
	Parameters []interface{} `json:"parameters,omitempty"`
}
