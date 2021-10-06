/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.21
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ElectrumRawTransactionVout struct for ElectrumRawTransactionVout
type ElectrumRawTransactionVout struct {
	N float32 `json:"n,omitempty"`
	ScriptPubKey ElectrumRawTransactionScriptPubKey `json:"scriptPubKey,omitempty"`
	Value float32 `json:"value,omitempty"`
}
