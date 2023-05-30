/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.26
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WalletReplaceNamedRequest struct for WalletReplaceNamedRequest
type WalletReplaceNamedRequest struct {
	// network type
	Network string `json:"network,omitempty"`
	// Wallet type, either a mnemonic seed single address wallet, a simple private key (wif) or a *Hierarchical Deterministic wallet determined from a seed (not yet implemented)* .
	Type string `json:"type,omitempty"`
}
