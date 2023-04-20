/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.4
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ContractFnResponse struct for ContractFnResponse
type ContractFnResponse struct {
	// serialized contract 
	ContractId string `json:"contractId,omitempty"`
	// The hash of a transaction
	TxId string `json:"txId,omitempty"`
	// The transaction as bitcoin encoded hex.
	Hex string `json:"hex,omitempty"`
	// if a debugger action is passed in the request, the result will return here.
	Debug string `json:"debug,omitempty"`
}
