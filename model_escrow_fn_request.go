/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.10
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EscrowFnRequest struct for EscrowFnRequest
type EscrowFnRequest struct {
	// serialized contract 
	EscrowContractId string `json:"escrowContractId"`
	// The walletId of the transaction signer. 
	WalletId string `json:"walletId"`
	// Function to call on the escrow contract.
	Function string `json:"function"`
	// Output address for the transaction
	To string `json:"to,omitempty"`
	// In addition to `send`ing the built transaction, the built transaction hex may be returned (without broadcasting) with `build` action, or the [`meep 🔗`](https://github.com/gcash/meep) debugger command
	Action string `json:"action,omitempty"`
	UtxoIds []string `json:"utxoIds,omitempty"`
}
