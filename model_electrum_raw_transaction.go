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
// ElectrumRawTransaction struct for ElectrumRawTransaction
type ElectrumRawTransaction struct {
	Blockhash string `json:"blockhash,omitempty"`
	Blocktime float32 `json:"blocktime,omitempty"`
	Confirmations float32 `json:"confirmations,omitempty"`
	Hash string `json:"hash,omitempty"`
	Hex string `json:"hex,omitempty"`
	Locktime float32 `json:"locktime,omitempty"`
	Size float32 `json:"size,omitempty"`
	Time float32 `json:"time,omitempty"`
	Txid string `json:"txid,omitempty"`
	Version float32 `json:"version,omitempty"`
	Vin []ElectrumRawTransactionVin `json:"vin,omitempty"`
	Vout []ElectrumRawTransactionVout `json:"vout,omitempty"`
}
