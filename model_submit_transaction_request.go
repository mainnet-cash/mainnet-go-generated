/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.3.13
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SubmitTransactionRequest struct for SubmitTransactionRequest
type SubmitTransactionRequest struct {
	WalletId string `json:"walletId"`
	// encoded transaction in hex format
	TransactionHex string `json:"transactionHex"`
	// A boolean flag to wait for transaction to propagate through the network and be registered in the bitcoind and indexer. If set to false, the `send` call will be very fast, but the wallet UTXO state might be invalid for some 500ms.
	AwaitPropagation bool `json:"awaitPropagation,omitempty"`
}
