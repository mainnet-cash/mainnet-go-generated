/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.0-rc.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchContractDeployResponse struct for SmartBchContractDeployResponse
type SmartBchContractDeployResponse struct {
	// serialized contract
	ContractId string `json:"contractId,omitempty"`
	// contract address
	Address string `json:"address,omitempty"`
	// The hash of a transaction
	TxId string `json:"txId,omitempty"`
	Receipt SmartBchTransactionReceipt `json:"receipt,omitempty"`
}
