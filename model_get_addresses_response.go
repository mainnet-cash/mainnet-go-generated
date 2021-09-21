/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.3
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetAddressesResponse struct for GetAddressesResponse
type GetAddressesResponse struct {
	// Cashaddr to return testnet BCH
	Bchtest string `json:"bchtest,omitempty"`
	// Slpaddr to return testnet SLP tokens
	Slptest string `json:"slptest,omitempty"`
}
