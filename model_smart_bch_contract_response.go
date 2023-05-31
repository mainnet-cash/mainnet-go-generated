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
// SmartBchContractResponse struct for SmartBchContractResponse
type SmartBchContractResponse struct {
	// serialized contract
	ContractId string `json:"contractId,omitempty"`
	// Contract address
	Address *interface{} `json:"address,omitempty"`
}
