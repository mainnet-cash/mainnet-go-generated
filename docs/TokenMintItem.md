# TokenMintItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | Token unique hexadecimal identifier, also the id of the token creation transaction. The minting NFT spent for this mint must be of this category. | 
**Nft** | Pointer to [**TokenNft**](Token_nft.md) |  | [optional] 
**Cashaddr** | **string** | Cashaddress to send tokens to | [optional] 
**Value** | **float32** | Satoshi value to send alongside with tokens | [optional] [default to 1000]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


