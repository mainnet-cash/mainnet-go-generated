/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.3.14
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetAddrsByXpubKeyRequest struct for GetAddrsByXpubKeyRequest
type GetAddrsByXpubKeyRequest struct {
	Xpubkey string `json:"xpubkey,omitempty"`
	Path string `json:"path,omitempty"`
	Count float32 `json:"count,omitempty"`
}
