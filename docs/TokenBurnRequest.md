# TokenBurnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The walletId to make a request to. | 
**Category** | **string** | Token unique hexadecimal identifier, also the id of the token creation transaction | 
**Nft** | Pointer to [**TokenNft**](Token_nft.md) |  | [optional] 
**Amount** | **float32** | amount of fungible tokens to burn | [optional] 
**Cashaddr** | **string** | address to return token and satoshi change to, default to the sender&#39;s cashaddr | [optional] 
**Message** | **string** | optional message to include in OP_RETURN | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


