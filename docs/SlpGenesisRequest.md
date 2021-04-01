# SlpGenesisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | ID that is returned in &#x60;wallet&#x60; field of /wallet call  | 
**Name** | **string** |  | 
**Ticker** | **string** | Token ticker | 
**InitialAmount** | **string** | Value is represented as a string to avoid precision loss | 
**Decimals** | **float32** | Indicates that 1 token is divisible into 10^decimals base units | 
**DocumentUrl** | **string** |  | [optional] 
**DocumentHash** | **string** | Document hash of the token. Empty or 64 character long hex string. | [optional] 
**EndBaton** | **bool** | Indicates if token should not be &#39;mintable&#39;, e.g. total circulation amount increased | [optional] 
**Type** | **float32** | Token type. 0x01 Type1, 0x81 NFT Parent, 0x41 NFT Child | [optional] 
**TokenReceiverSlpAddr** | **string** |  | [optional] 
**BatonReceiverSlpAddr** | **string** |  | [optional] 
**ParentTokenId** | **string** | Identifier of the NFT parent token | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


