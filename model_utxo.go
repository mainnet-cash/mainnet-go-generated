/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.22
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Utxo struct for Utxo
type Utxo struct {
	Vout float32 `json:"vout"`
	// The hash of a transaction
	Txid string `json:"txid"`
	Satoshis float32 `json:"satoshis"`
	Token *UtxoToken `json:"token,omitempty"`
}
