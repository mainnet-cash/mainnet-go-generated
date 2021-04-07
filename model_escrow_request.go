/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.12
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EscrowRequest struct for EscrowRequest
type EscrowRequest struct {
	// The cashaddress of the buyer
	BuyerAddr *interface{} `json:"buyerAddr"`
	// The cashaddress of the arbiter
	ArbiterAddr *interface{} `json:"arbiterAddr"`
	// The cashaddress of the seller
	SellerAddr *interface{} `json:"sellerAddr"`
	// Numeric amount to be transferred by the contract in satoshi, amount must be less than 21 BCH.
	Amount float32 `json:"amount"`
}