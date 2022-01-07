/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.31
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EscrowResponse struct for EscrowResponse
type EscrowResponse struct {
	// serialized escrow contract 
	EscrowContractId string `json:"escrowContractId,omitempty"`
	// The funding address for the escrow contract
	Cashaddr *interface{} `json:"cashaddr,omitempty"`
}
