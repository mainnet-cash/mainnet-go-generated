/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.6.2
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Token struct for Token
type Token struct {
	// Fungible token amount
	Amount float32 `json:"amount,omitempty"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId,omitempty"`
	// Capability of the NFT
	Capability *string `json:"capability,omitempty"`
	// Token commitment message, hexadecimal encoded, 40 bytes max length
	Commitment *string `json:"commitment,omitempty"`
}
