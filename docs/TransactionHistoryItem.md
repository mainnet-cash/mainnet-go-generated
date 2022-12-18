# TransactionHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | cashaddr of the receiving cash address. | [optional] 
**From** | **string** | cashaddr of the sending cash address. | [optional] 
**Unit** | [**UnitType**](UnitType.md) |  | [optional] 
**Index** | **float32** | the index of the input or output in the transaction | [optional] 
**Blockheight** | **float32** | the blockheight of transaction | [optional] 
**Txn** | **string** | The hash of the transaction | [optional] 
**TxId** | **string** | a unique identifier for the sub-transaction | [optional] 
**Value** | **float32** | The amount of the transaction, in the unit provided, with respect to the wallet provided. | [optional] 
**Fee** | **float32** | The amount paid, if any, by the wallet for the transaction (incoming tranactions are \&quot;free\&quot;) | [optional] 
**Balance** | **float32** | A running balance | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


