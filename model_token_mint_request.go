/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.34
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TokenMintRequest struct for TokenMintRequest
type TokenMintRequest struct {
	// The walletId to make a request to.
	WalletId string `json:"walletId"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId"`
	Requests []TokenMintRequestRequests `json:"requests,omitempty"`
	// if minting token contains fungible amount, deduct from it by amount of minted tokens
	DeductTokenAmount bool `json:"deductTokenAmount,omitempty"`
}
