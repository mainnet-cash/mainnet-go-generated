# \UtilApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Convert**](UtilApi.md#Convert) | **Post** /util/convert | convert between units
[**GetAddrsByXpubKey**](UtilApi.md#GetAddrsByXpubKey) | **Post** /util/get_addrs_by_xpubkey | Derive heristic determined addresses from an extended public key, per BIP32
[**GetXpubKeyInfo**](UtilApi.md#GetXpubKeyInfo) | **Post** /util/get_xpubkey_info | Decode information about an extended public key, per BIP32



## Convert

> float32 Convert(ctx, optional)

convert between units

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConvertOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConvertOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convertRequest** | [**optional.Interface of ConvertRequest**](ConvertRequest.md)|  | 

### Return type

**float32**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddrsByXpubKey

> []AnyOfstring GetAddrsByXpubKey(ctx, optional)

Derive heristic determined addresses from an extended public key, per BIP32

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAddrsByXpubKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAddrsByXpubKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAddrsByXpubKeyRequest** | [**optional.Interface of GetAddrsByXpubKeyRequest**](GetAddrsByXpubKeyRequest.md)|  | 

### Return type

[**[]AnyOfstring**](anyOf&lt;string&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXpubKeyInfo

> GetXpubKeyInfoResponse GetXpubKeyInfo(ctx, getXpubKeyInfoRequest)

Decode information about an extended public key, per BIP32

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getXpubKeyInfoRequest** | [**GetXpubKeyInfoRequest**](GetXpubKeyInfoRequest.md)| Decode information about an extended public key, per BIP32 | 

### Return type

[**GetXpubKeyInfoResponse**](getXpubKeyInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

