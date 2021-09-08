/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WalletMnemonic struct for WalletMnemonic
type WalletMnemonic struct {
	Seed string `json:"seed,omitempty"`
	DerivationPath string `json:"derivationPath,omitempty"`
}
