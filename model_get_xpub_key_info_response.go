/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.2.8
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetXpubKeyInfoResponse struct for GetXpubKeyInfoResponse
type GetXpubKeyInfoResponse struct {
	Version string `json:"version,omitempty"`
	Depth float32 `json:"depth,omitempty"`
	// the first four bytes of the public key as hex
	ParentFingerprint string `json:"parentFingerprint,omitempty"`
	// the public key as hex
	Data string `json:"data,omitempty"`
	// The chain code, 32 bytes of entrophy extended for the public key
	Chain *interface{} `json:"chain,omitempty"`
}
