/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.5.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TokenAmountChange struct for TokenAmountChange
type TokenAmountChange struct {
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId,omitempty"`
	// Fungible token amount
	Amount float32 `json:"amount,omitempty"`
	// Non-fungible token amount
	NftAmount float32 `json:"nftAmount,omitempty"`
}
