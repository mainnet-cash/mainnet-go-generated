# \WalletSignedApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SignedMessageSign**](WalletSignedApi.md#SignedMessageSign) | **Post** /wallet/signed/sign | Returns the message signature
[**SignedMessageVerify**](WalletSignedApi.md#SignedMessageVerify) | **Post** /wallet/signed/verify | Returns a boolean indicating whether message was valid for a given address



## SignedMessageSign

> SignedMessageResponse SignedMessageSign(ctx, optional)

Returns the message signature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SignedMessageSignOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SignedMessageSignOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSignedMessageRequest** | [**optional.Interface of CreateSignedMessageRequest**](CreateSignedMessageRequest.md)| Sign a message  | 

### Return type

[**SignedMessageResponse**](SignedMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignedMessageVerify

> VerifySignedMessageResponse SignedMessageVerify(ctx, optional)

Returns a boolean indicating whether message was valid for a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SignedMessageVerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SignedMessageVerifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifySignedMessageRequest** | [**optional.Interface of VerifySignedMessageRequest**](VerifySignedMessageRequest.md)| Sign a message  | 

### Return type

[**VerifySignedMessageResponse**](VerifySignedMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

