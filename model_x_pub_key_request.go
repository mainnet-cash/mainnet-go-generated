/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.31
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// XPubKeyRequest struct for XPubKeyRequest
type XPubKeyRequest struct {
	// walletId for a seed wallet 
	WalletId string `json:"walletId"`
	Paths []string `json:"paths,omitempty"`
}
