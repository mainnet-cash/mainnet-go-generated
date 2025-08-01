/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.7.14
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ElectrumRawTransactionScriptPubKey struct for ElectrumRawTransactionScriptPubKey
type ElectrumRawTransactionScriptPubKey struct {
	Addresses []string `json:"addresses,omitempty"`
	Asm string `json:"asm,omitempty"`
	Hex string `json:"hex,omitempty"`
	ReqSigs float32 `json:"reqSigs,omitempty"`
	Type string `json:"type,omitempty"`
}
