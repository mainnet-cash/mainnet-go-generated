# EscrowFnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | serialized contract  | 
**WalletId** | **string** | The walletId of the transaction signer.  | 
**Function** | **string** | Function to call on the escrow contract. | [optional] 
**To** | **string** | Output address for the transaction | [optional] 
**Action** | **string** | In addition to &#x60;send&#x60;ing the built transaction, the built transaction hex may be returned (without broadcasting) with &#x60;build&#x60; action, or the [&#x60;meep 🔗&#x60;](https://github.com/gcash/meep) debugger command | [optional] 
**UtxoIds** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


