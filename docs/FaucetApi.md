# \FaucetApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddresses**](FaucetApi.md#GetAddresses) | **Post** /faucet/get_addresses | Get addresses to return back or donate the testnet bch and tokens 
[**GetTestnetBch**](FaucetApi.md#GetTestnetBch) | **Post** /faucet/get_testnet_bch | Get testnet bch 
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

