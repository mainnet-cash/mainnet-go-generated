# VerifySignedMessageResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignatureType** | **string** | The type of signature passed | [optional] 
**MessageHash** | **string** | The double sha256 hash of the string encoded as Bitcoin Message binary. | [optional] 
**SignatureValid** | **bool** | A boolean indicating whether the signature is valid for the message | [optional] 
**PublicKeyHashMatch** | **bool** | A boolean indicating whether the publicKeyHash of a recoverable signature matches the provided cashaddr, false otherwise | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


