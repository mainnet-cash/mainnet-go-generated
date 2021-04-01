/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.8
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpUtxo struct for SlpUtxo
type SlpUtxo struct {
	Index float32 `json:"index,omitempty"`
	// The hash of a transaction
	TxId string `json:"txId,omitempty"`
	// Locked satoshi
	Satoshis float32 `json:"satoshis,omitempty"`
	// serialized outpoint
	UtxoId string `json:"utxoId,omitempty"`
	// Token amount represented as a string to avoid precision loss
	Value string `json:"value,omitempty"`
	// Indicates that 1 token is divisible into 10^decimals base units
	Decimals float32 `json:"decimals,omitempty"`
	// Token ticker
	Ticker string `json:"ticker,omitempty"`
	// Token unique hexadecimal identifier, also the id of the token creation transaction
	TokenId string `json:"tokenId,omitempty"`
	// Token type. 0x01 Type1, 0x81 NFT Parent, 0x41 NFT Child
	Type float32 `json:"type,omitempty"`
}
