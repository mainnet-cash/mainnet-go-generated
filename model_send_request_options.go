/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.12
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SendRequestOptions struct for SendRequestOptions
type SendRequestOptions struct {
	UtxoIds []string `json:"utxoIds,omitempty"`
	// Cash address to receive change to 
	ChangeAddress string `json:"changeAddress,omitempty"`
	// If you have any SLP tokens on the address, you should set `slpAware` to `true` to prevent accidental token burning So far this option is not switched on by default
	SlpAware bool `json:"slpAware,omitempty"`
}
