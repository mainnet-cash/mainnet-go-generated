# \FaucetApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddresses**](FaucetApi.md#GetAddresses) | **Post** /faucet/get_addresses | Get addresses to return back or donate the testnet bch and tokens 
[**GetTestnetBch**](FaucetApi.md#GetTestnetBch) | **Post** /faucet/get_testnet_bch | Get testnet bch 
[**GetTestnetSbch**](FaucetApi.md#GetTestnetSbch) | **Post** /faucet/get_testnet_sbch | Request testnet SmartBCH funds. The request is enqueued and served within 1-3 block confirmations. If the target address holds more that 0.1 BCH, the request will fail. Otherwise the address will be topped up to 0.1 BCH. 
[**GetTestnetSep20**](FaucetApi.md#GetTestnetSep20) | **Post** /faucet/get_testnet_sep20 | Request testnet SmartBch SEP20 tokens. The request is enqueued and served within 1-3 block confirmations. If the target address holds more that 10 tokens of requested kind, the request will fail. Otherwise the address will be topped up to 10 tokens. 
[**GetTestnetSlp**](FaucetApi.md#GetTestnetSlp) | **Post** /faucet/get_testnet_slp | Get testnet slp tokens 



## GetAddresses

> GetAddressesResponse GetAddresses(ctx, )

Get addresses to return back or donate the testnet bch and tokens 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetAddressesResponse**](GetAddressesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestnetBch

> GetTestnetBchResponse GetTestnetBch(ctx, getTestnetBchRequest)

Get testnet bch 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getTestnetBchRequest** | [**GetTestnetBchRequest**](GetTestnetBchRequest.md)| Request to bch faucet  | 

### Return type

[**GetTestnetBchResponse**](GetTestnetBchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestnetSbch

> GetTestnetSbchResponse GetTestnetSbch(ctx, getTestnetSbchRequest)

Request testnet SmartBCH funds. The request is enqueued and served within 1-3 block confirmations. If the target address holds more that 0.1 BCH, the request will fail. Otherwise the address will be topped up to 0.1 BCH. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getTestnetSbchRequest** | [**GetTestnetSbchRequest**](GetTestnetSbchRequest.md)| Request to bch faucet  | 

### Return type

[**GetTestnetSbchResponse**](GetTestnetSbchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestnetSep20

> GetTestnetSep20Response GetTestnetSep20(ctx, getTestnetSep20Request)

Request testnet SmartBch SEP20 tokens. The request is enqueued and served within 1-3 block confirmations. If the target address holds more that 10 tokens of requested kind, the request will fail. Otherwise the address will be topped up to 10 tokens. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getTestnetSep20Request** | [**GetTestnetSep20Request**](GetTestnetSep20Request.md)| Request to SEP20 faucet  | 

### Return type

[**GetTestnetSep20Response**](GetTestnetSep20Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestnetSlp

> GetTestnetSlpResponse GetTestnetSlp(ctx, getTestnetSlpRequest)

Get testnet slp tokens 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getTestnetSlpRequest** | [**GetTestnetSlpRequest**](GetTestnetSlpRequest.md)| Request to slp faucet  | 

### Return type

[**GetTestnetSlpResponse**](GetTestnetSlpResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

