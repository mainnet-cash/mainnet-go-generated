/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.40
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SendRequestOptions struct for SendRequestOptions
type SendRequestOptions struct {
	UtxoIds []string `json:"utxoIds,omitempty"`
	// Cash address to receive change to 
	ChangeAddress string `json:"changeAddress,omitempty"`
	// If you have any SLP tokens on the address, you should set `slpAware` to `true` to prevent accidental token burning So far this option is not switched on by default
	SlpAware bool `json:"slpAware,omitempty"`
	// A boolean flag to include the wallet balance after the successful `send` operation to the returned result If set to false, the balance will not be queried and returned, making the `send` call faster
	QueryBalance bool `json:"queryBalance,omitempty"`
	// A boolean flag to wait for transaction to propagate through the network and be registered in the bitcoind and indexer. If set to false, the `send` call will be very fast, but the wallet UTXO state might be invalid for some 500ms.
	AwaitTransactionPropagation bool `json:"awaitTransactionPropagation,omitempty"`
}
