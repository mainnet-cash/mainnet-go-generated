/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.7
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TokenBurnRequest struct for TokenBurnRequest
type TokenBurnRequest struct {
	// The walletId to make a request to.
	WalletId string `json:"walletId"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId"`
	// Capability of the new NFT
	Capability string `json:"capability,omitempty"`
	// Token commitment message, hexadecimal encoded, 40 bytes max length
	Commitment string `json:"commitment,omitempty"`
	// amount of fungible tokens to burn
	Amount float32 `json:"amount,omitempty"`
	// address to return token and satoshi change to, default to the sender's cashaddr
	Cashaddr string `json:"cashaddr,omitempty"`
	// optional message to include in OP_RETURN
	Message string `json:"message,omitempty"`
}