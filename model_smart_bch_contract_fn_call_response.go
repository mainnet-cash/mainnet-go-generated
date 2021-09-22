/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.9
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchContractFnCallResponse Contract function evaluation response. A constant function (non state changing, also does not cost gas) evaluates to a `result` field A state changing function will generate a transaction and `txId` and `receipt` fields will be filled 
type SmartBchContractFnCallResponse struct {
	// serialized contract
	ContractId string `json:"contractId,omitempty"`
	// contract function evaluation result. Can be of any type.
	Result *interface{} `json:"result,omitempty"`
	// The hash of a transaction
	TxId string `json:"txId,omitempty"`
	Receipt SmartBchTransactionReceipt `json:"receipt,omitempty"`
}
