/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.14
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SendRequestOptions struct for SendRequestOptions
type SendRequestOptions struct {
	UtxoIds []string `json:"utxoIds,omitempty"`
	// Cash address to receive change to 
	ChangeAddress string `json:"changeAddress,omitempty"`
	// If you have any SLP tokens on the address, you should set `slpAware` to `true` to prevent accidental token burning This flag activates utxo checking against an external slp indexer So far this option is not switched on by default
	SlpAware bool `json:"slpAware,omitempty"`
	// This flag requires an utxo to have more than 546 sats to be spendable and counted in the balance This option is not switched on by default
	SlpSemiAware bool `json:"slpSemiAware,omitempty"`
	// A boolean flag to include the wallet balance after the successful `send` operation to the returned result If set to false, the balance will not be queried and returned, making the `send` call faster
	QueryBalance bool `json:"queryBalance,omitempty"`
	// A boolean flag to wait for transaction to propagate through the network and be registered in the bitcoind and indexer. If set to false, the `send` call will be very fast, but the wallet UTXO state might be invalid for some 500ms.
	AwaitTransactionPropagation bool `json:"awaitTransactionPropagation,omitempty"`
	// Fee allocation stragety. Convenience option to subtract fees from outputs if change is not sufficent to cover transaction costs.   * `change` - Change, pay the fees from change or error   * `firstOutput` - First Output, pay the fee from the first output or error   * `lastOutput` - Last Output, pay the fee from the last output or error   * `anyOutput` - Any Output, pay the fee from dust outputs or divide across all remaining non-dust outputs.   * `changeThenFirst` - Use change then first output or error.   * `changeThenLast` - Use change then last output or error.   * `changeThenAny` - Use change then any output stragety or error. 
	FeePaidBy string `json:"feePaidBy,omitempty"`
	// Check that sum of input amounts and output amount for each token category matches. Prevents accidental burns. 
	CheckTokenQuantities bool `json:"checkTokenQuantities,omitempty"`
}
