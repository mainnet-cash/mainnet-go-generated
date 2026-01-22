# ContractFnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | serialized contract  | 
**Action** | **string** | In addition to &#x60;send&#x60;ing the built transaction, the built transaction hex may be returned (without broadcasting) with &#x60;build&#x60; action. | [optional] [default to ACTION_SEND]
**Function** | **string** | Function to call on the cashscript contract. | 
**Arguments** | **[]string** | Arguments for the contract function call as strings.  Binary data should be passed as hexidecimal.  Signatures may be passed as wallet import format (WIF) or wallet strings (walletId). Cashscript expects &#x60;pubkey&#x60;s to be compressed 35 byte values.  | [optional] 
**To** | [**AnyOfSendRequestItemarrayCashscriptReceiptarray**](anyOf&lt;SendRequestItem,array,CashscriptReceipt,array&gt;.md) | The output destination, as a SendRequest, cashscript style output(s), array of either. | 
**UtxoIds** | **[]string** | Serialized utxoId(s) to spend from | [optional] 
**OpReturn** | **[]string** | Add OP_RETURN outputs to the transaction. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withopreturn)  | [optional] 
**FeePerByte** | **float32** | The withFeePerByte() function allows you to specify the fee per per bytes for the transaction. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withfeeperbyte)  | [optional] 
**HardcodedFee** | **float32** | Specify a hardcoded fee to the transaction. By default the transaction fee is automatically calculated by the CashScript SDK. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withhardcodedfee)  | [optional] 
**MinChange** | **float32** | Set a threshold for including a change output. Any remaining amount under this threshold will be added to the transaction fee instead. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withminchange)  | [optional] 
**WithoutChange** | **bool** | Disable the change output. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withoutchange)  | [optional] [default to false]
**Age** | **float32** | Specify the minimum age of the transaction inputs. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withage)  | [optional] 
**Time** | **float32** | Specify the minimum block number that the transaction can be included in. See [cashscript docs](https://cashscript.org/docs/sdk/transactions#withtime)  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


