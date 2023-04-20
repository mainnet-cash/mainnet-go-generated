/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.3
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AuthChainElement struct for AuthChainElement
type AuthChainElement struct {
	// Hex encoded transaction hash
	TxHash string `json:"txHash,omitempty"`
	// Content hash of the remote registry, either a SHA256 hash of an HTTPS Publication output or an IPFS CID
	ContentHash string `json:"contentHash,omitempty"`
	// URI to fetch the registry from
	Uri string `json:"uri,omitempty"`
}
