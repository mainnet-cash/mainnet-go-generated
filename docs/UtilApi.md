# \UtilApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Convert**](UtilApi.md#Convert) | **Post** /util/convert | convert between units



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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

