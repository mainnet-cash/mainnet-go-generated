# TokenGenesisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The walletId to make a request to. | 
**Amount** | **float32** | amount of *fungible* tokens to create | [optional] 
**Capability** | **string** | Capability of the new NFT | [optional] 
**Commitment** | **string** | Token commitment message, hexadecimal encoded, 40 bytes max length | [optional] 
**Cashaddr** | **string** | Cashaddress to send tokens to | [optional] 
**Value** | **float32** | Satoshi value to send alongside with tokens | [optional] [default to 1000]
**SendRequests** | [**AnyOfSendRequestItemarrayTokenSendRequestOpReturnData**](anyOf&lt;SendRequestItem,array,TokenSendRequest,OpReturnData&gt;.md) | Single or an array of extra send requests (OP_RETURN, value transfer, etc.) to include in genesis transaction. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


