# HistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The walletId to get a history for.  | 
**Unit** | **string** | Unit of account for running balance. | [optional] [default to UNIT_SAT]
**Start** | **float32** | The first record to return starting from zero | [optional] 
**Count** | **float32** | The number of records to return | [optional] 
**CollapseChange** | **bool** | Exclude transactions returned to the wallet as change in response. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


