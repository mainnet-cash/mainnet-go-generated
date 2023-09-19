/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.2.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ContractFnRequest struct for ContractFnRequest
type ContractFnRequest struct {
	// serialized contract 
	ContractId string `json:"contractId"`
	// In addition to `send`ing the built transaction, the built transaction hex may be returned (without broadcasting) with `build` action, or the [`meep 🔗`](https://github.com/gcash/meep) debugger command.
	Action string `json:"action,omitempty"`
	// Function to call on the cashscript contract.
	Function string `json:"function"`
	// Arguments for the contract function call as strings.  Binary data should be passed as hexidecimal.  Signatures may be passed as wallet import format (WIF) or wallet strings (walletId). Cashscript expects `pubkey`s to be compressed 35 byte values. 
	Arguments []string `json:"arguments,omitempty"`
	// The output destination, as a SendRequest, cashscript style output(s), array of either.
	To AnyOfSendRequestItemarrayCashscriptReceiptarray `json:"to"`
	// Serialized utxoId(s) to spend from
	UtxoIds []string `json:"utxoIds,omitempty"`
	// Add OP_RETURN outputs to the transaction. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withopreturn) 
	OpReturn []string `json:"opReturn,omitempty"`
	// The withFeePerByte() function allows you to specify the fee per per bytes for the transaction. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withfeeperbyte) 
	FeePerByte float32 `json:"feePerByte,omitempty"`
	// Specify a hardcoded fee to the transaction. By default the transaction fee is automatically calculated by the CashScript SDK. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withhardcodedfee) 
	HardcodedFee float32 `json:"hardcodedFee,omitempty"`
	// Set a threshold for including a change output. Any remaining amount under this threshold will be added to the transaction fee instead. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withminchange) 
	MinChange float32 `json:"minChange,omitempty"`
	// Disable the change output. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withoutchange) 
	WithoutChange bool `json:"withoutChange,omitempty"`
	// Specify the minimum age of the transaction inputs. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withage) 
	Age float32 `json:"age,omitempty"`
	// Specify the minimum block number that the transaction can be included in. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withtime) 
	Time float32 `json:"time,omitempty"`
}
