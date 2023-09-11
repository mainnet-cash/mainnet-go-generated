/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.34-alpha.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TokenGenesisRequest struct for TokenGenesisRequest
type TokenGenesisRequest struct {
	// The walletId to make a request to.
	WalletId string `json:"walletId"`
	// amount of *fungible* tokens to create
	Amount float32 `json:"amount,omitempty"`
	// Capability of the new NFT
	Capability string `json:"capability,omitempty"`
	// Token commitment message, hexadecimal encoded, 40 bytes max length
	Commitment string `json:"commitment,omitempty"`
	// Cashaddress to send tokens to
	Cashaddr string `json:"cashaddr,omitempty"`
	// Satoshi value to send alongside with tokens
	Value float32 `json:"value,omitempty"`
	// Single or an array of extra send requests (OP_RETURN, value transfer, etc.) to include in genesis transaction.
	SendRequests AnyOfSendRequestItemarrayTokenSendRequestOpReturnData `json:"sendRequests,omitempty"`
}
