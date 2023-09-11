/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.34-alpha.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TransactionHistoryItem struct for TransactionHistoryItem
type TransactionHistoryItem struct {
	// cashaddr of the receiving cash address.
	To string `json:"to,omitempty"`
	// cashaddr of the sending cash address.
	From string `json:"from,omitempty"`
	Unit UnitType `json:"unit,omitempty"`
	// the index of the input or output in the transaction
	Index float32 `json:"index,omitempty"`
	// the blockheight of transaction
	Blockheight float32 `json:"blockheight,omitempty"`
	// The hash of the transaction
	Txn string `json:"txn,omitempty"`
	// a unique identifier for the sub-transaction
	TxId string `json:"txId,omitempty"`
	// The amount of the transaction, in the unit provided, with respect to the wallet provided.
	Value float32 `json:"value,omitempty"`
	// The amount paid, if any, by the wallet for the transaction (incoming tranactions are \"free\")
	Fee float32 `json:"fee,omitempty"`
	// A running balance
	Balance float32 `json:"balance,omitempty"`
}
