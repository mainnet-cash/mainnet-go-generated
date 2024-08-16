/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.4.3
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TransactionHistoryItem struct for TransactionHistoryItem
type TransactionHistoryItem struct {
	Inputs []InOutput `json:"inputs,omitempty"`
	Outputs []InOutput `json:"outputs,omitempty"`
	// The blockheight of transaction
	Blockheight float32 `json:"blockheight,omitempty"`
	// The timestamp of transaction, undefined if unconfirmed
	Timestamp *string `json:"timestamp,omitempty"`
	// The hash of the transaction
	Hash string `json:"hash,omitempty"`
	// The size of the transaction
	Size float32 `json:"size,omitempty"`
	// Transaction fee
	Fee float32 `json:"fee,omitempty"`
	// A running balance, in units
	Balance float32 `json:"balance,omitempty"`
	// Change of value in the transaction, in units
	ValueChange float32 `json:"valueChange,omitempty"`
	TokenAmountChanges []TokenAmountChange `json:"tokenAmountChanges,omitempty"`
}
