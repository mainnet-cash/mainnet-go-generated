/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.5.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ElectrumRawTransactionVin struct for ElectrumRawTransactionVin
type ElectrumRawTransactionVin struct {
	ScriptSig ElectrumRawTransactionScriptSig `json:"scriptSig,omitempty"`
	Sequence float32 `json:"sequence,omitempty"`
	Txid string `json:"txid,omitempty"`
	Vout float32 `json:"vout,omitempty"`
	// optional extention by mainnet.cash
	Value float32 `json:"value,omitempty"`
	// optional extention by mainnet.cash
	Address string `json:"address,omitempty"`
}
