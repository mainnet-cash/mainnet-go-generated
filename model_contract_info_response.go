/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.27
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ContractInfoResponse struct for ContractInfoResponse
type ContractInfoResponse struct {
	// serialized contract 
	ContractId string `json:"contractId,omitempty"`
	// The smart contract in cashscript syntax 
	Script string `json:"script,omitempty"`
	// Parameters passed when the contract was created
	Parameters []string `json:"parameters,omitempty"`
	Cashaddr string `json:"cashaddr,omitempty"`
	Tokenaddr string `json:"tokenaddr,omitempty"`
}
