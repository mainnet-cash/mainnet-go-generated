# TransactionHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**[]InOutput**](InOutput.md) |  | [optional] 
**Outputs** | [**[]InOutput**](InOutput.md) |  | [optional] 
**Blockheight** | **float32** | The blockheight of transaction | [optional] 
**Timestamp** | Pointer to **string** | The timestamp of transaction, undefined if unconfirmed | [optional] 
**Hash** | **string** | The hash of the transaction | [optional] 
**Size** | **float32** | The size of the transaction | [optional] 
**Fee** | **float32** | Transaction fee | [optional] 
**Balance** | **float32** | A running balance, in units | [optional] 
**ValueChange** | **float32** | Change of value in the transaction, in units | [optional] 
**TokenAmountChanges** | [**[]TokenAmountChange**](TokenAmountChange.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


