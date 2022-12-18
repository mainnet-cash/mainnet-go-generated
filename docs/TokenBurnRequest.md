# TokenBurnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The walletId to make a request to. | 
**TokenId** | **string** | Token unique hexadecimal identifier, also the id of the token creation transaction | 
**Capability** | **string** | Capability of the new NFT | [optional] 
**Commitment** | **string** | Token commitment message, hexadecimal encoded, 40 bytes max length | [optional] 
**Amount** | **float32** | amount of fungible tokens to burn | [optional] 
**Cashaddr** | **string** | address to return token and satoshi change to, default to the sender&#39;s cashaddr | [optional] 
**Message** | **string** | optional message to include in OP_RETURN | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


