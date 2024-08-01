# HistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The walletId to get a history for.  | 
**Unit** | **string** | Unit of account for running balance. | [optional] [default to UNIT_SAT]
**FromHeight** | **float32** | optional, if set, history will be limited. Default 0 | [optional] [default to 0]
**ToHeight** | **float32** | optional, if set, history will be limited. Default -1, meaning that all history items will be returned, including mempool | [optional] [default to -1]
**Start** | **float32** | optional, if set, the result set will be paginated with offset &#x60;start&#x60;. Default 0 | [optional] [default to 0]
**Count** | **float32** | optional, if set, the result set will be paginated with &#x60;count&#x60;. Default -1, meaning that all history items will be returned | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


