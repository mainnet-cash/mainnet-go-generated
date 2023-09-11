# WatchAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cashaddr** | **string** | Cash address to watch  | 
**Url** | **string** | Url to be called when configured action is triggered | 
**Type** | **string** | Type of watch operation | [default to TYPE_TRANSACTIONINOUT]
**Recurrence** | **string** | Action recurrence. Indicates if webhook should be triggered recurrently until expire or only once | [optional] [default to RECURRENCE_ONCE]
**DurationSec** | **float32** | Duration of the webhook lifetime in seconds before it will expire. | [optional] [default to 86400]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


