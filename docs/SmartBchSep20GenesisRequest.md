# SmartBchSep20GenesisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | serialized wallet | 
**Name** | **string** | Token name | 
**Ticker** | **string** | Token ticker | 
**InitialAmount** | [**AnyOfnumberstring**](anyOf&lt;number,string&gt;.md) | Value is represented as number or string to avoid precision loss | 
**Decimals** | **float32** | Indicates that 1 token is divisible into 10^decimals base units | 
**EndBaton** | **bool** | Indicates if token should not be &#39;mintable&#39;, e.g. total circulation amount increased | [optional] 
**TokenReceiverAddress** | **string** |  | [optional] 
**BatonReceiverAddress** | **string** |  | [optional] 
**Overrides** | [**SmartBchOverrides**](SmartBchOverrides.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


