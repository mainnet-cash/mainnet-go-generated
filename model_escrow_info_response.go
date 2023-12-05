/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.3.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EscrowInfoResponse struct for EscrowInfoResponse
type EscrowInfoResponse struct {
	// serialized escrow contract 
	EscrowContractId string `json:"escrowContractId,omitempty"`
	// The escrow contract in cashscript syntax 
	Script string `json:"script,omitempty"`
	// Parameters passed when the contract was created
	Parameters []string `json:"parameters,omitempty"`
	Cashaddr string `json:"cashaddr,omitempty"`
	// The cashaddress of the buyer
	BuyerAddr *interface{} `json:"buyerAddr,omitempty"`
	// The cashaddress of the arbiter
	ArbiterAddr *interface{} `json:"arbiterAddr,omitempty"`
	// The cashaddress of the seller
	SellerAddr *interface{} `json:"sellerAddr,omitempty"`
	// Numeric amount to be transferred by the contract in satoshi, amount must be less than 21 BCH.
	Amount float32 `json:"amount,omitempty"`
	// A unique number used once for each instance of a contract.
	Nonce float32 `json:"nonce,omitempty"`
}
