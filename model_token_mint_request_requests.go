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
// TokenMintRequestRequests struct for TokenMintRequestRequests
type TokenMintRequestRequests struct {
	// Capability of the new NFT
	Capability string `json:"capability,omitempty"`
	// Token commitment message, hexadecimal encoded, 40 bytes max length
	Commitment string `json:"commitment,omitempty"`
	// Cashaddress to send tokens to
	Cashaddr string `json:"cashaddr,omitempty"`
	// Satoshi value to send alongside with tokens
	Value float32 `json:"value,omitempty"`
}
