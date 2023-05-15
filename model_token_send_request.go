/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.18
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TokenSendRequest struct for TokenSendRequest
type TokenSendRequest struct {
	// Cashaddress to send tokens to
	Cashaddr string `json:"cashaddr"`
	// Satoshi value to send alongside with tokens
	Value float32 `json:"value,omitempty"`
	// Fungible token amount to send
	Amount float32 `json:"amount,omitempty"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId"`
	// Capability of the NFT
	Capability string `json:"capability,omitempty"`
	// Token commitment message, hexadecimal encoded, 40 bytes max length
	Commitment string `json:"commitment,omitempty"`
}
