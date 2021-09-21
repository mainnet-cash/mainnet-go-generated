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
// WalletType struct for WalletType
type WalletType struct {
	// Wallet type, either a mnemonic seed single address wallet, a simple private key (wif) or a *Hierarchical Deterministic wallet determined from a seed (not yet implemented)* .
	Type string `json:"type,omitempty"`
}
