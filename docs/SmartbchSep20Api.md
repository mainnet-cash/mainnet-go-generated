# \SmartbchSep20Api

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartBchSep20Balance**](SmartbchSep20Api.md#SmartBchSep20Balance) | **Post** /smartbch/sep20/balance | Get total SmartBch SEP20 token balance of the wallet
[**SmartBchSep20DepositAddress**](SmartbchSep20Api.md#SmartBchSep20DepositAddress) | **Post** /smartbch/sep20/deposit_address | Get an SmartBch SEP20 deposit address
[**SmartBchSep20DepositQr**](SmartbchSep20Api.md#SmartBchSep20DepositQr) | **Post** /smartbch/sep20/deposit_qr | Get an SmartBch SEP20 receiving address as a qrcode
[**SmartBchSep20Genesis**](SmartbchSep20Api.md#SmartBchSep20Genesis) | **Post** /smartbch/sep20/genesis | Get created tokenId back and new SmartBch SEP20 token balance of the wallet
[**SmartBchSep20Mint**](SmartbchSep20Api.md#SmartBchSep20Mint) | **Post** /smartbch/sep20/mint | Get created tokenId back and new SmartBch SEP20 token balance of the wallet
[**SmartBchSep20Send**](SmartbchSep20Api.md#SmartBchSep20Send) | **Post** /smartbch/sep20/send | Send some SmartBch SEP20 token amount to a given address
[**SmartBchSep20SendMax**](SmartbchSep20Api.md#SmartBchSep20SendMax) | **Post** /smartbch/sep20/send_max | Send all available SmartBch SEP20 token funds to a given address
[**SmartBchSep20TokenInfo**](SmartbchSep20Api.md#SmartBchSep20TokenInfo) | **Post** /smartbch/sep20/token_info | Get information about the SmartBch SEP20 token



## SmartBchSep20Balance

> SmartBchSep20BalanceResponse SmartBchSep20Balance(ctx, smartBchSep20BalanceRequest)

Get total SmartBch SEP20 token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSep20BalanceRequest** | [**SmartBchSep20BalanceRequest**](SmartBchSep20BalanceRequest.md)| Request for a wallet SmartBch SEP20 token balance  | 

### Return type

[**SmartBchSep20BalanceResponse**](SmartBchSep20BalanceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchSep20DepositAddress

> SmartBchDepositAddressResponse SmartBchSep20DepositAddress(ctx, serializedWallet)

Get an SmartBch SEP20 deposit address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for an SmartBch SEP20 deposit address given a wallet  | 

### Return type

[**SmartBchDepositAddressResponse**](SmartBchDepositAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchSep20DepositQr

> ScalableVectorGraphic SmartBchSep20DepositQr(ctx, serializedWallet)

Get an SmartBch SEP20 receiving address as a qrcode

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a SmartBch SEP20 deposit address as a Quick Response code (qrcode)  | 

### Return type

[**ScalableVectorGraphic**](ScalableVectorGraphic.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchSep20Genesis

> SmartBchSep20GenesisResponse SmartBchSep20Genesis(ctx, smartBchSep20GenesisRequest)

Get created tokenId back and new SmartBch SEP20 token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSep20GenesisRequest** | [**SmartBchSep20GenesisRequest**](SmartBchSep20GenesisRequest.md)| Request to create a new SmartBch SEP20 token (genesis)  | 

### Return type

[**SmartBchSep20GenesisResponse**](SmartBchSep20GenesisResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchSep20Mint

> SmartBchSep20MintResponse SmartBchSep20Mint(ctx, smartBchSep20MintRequest)

Get created tokenId back and new SmartBch SEP20 token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSep20MintRequest** | [**SmartBchSep20MintRequest**](SmartBchSep20MintRequest.md)| Request to mint more of SmartBch SEP20 tokens  | 

### Return type

[**SmartBchSep20MintResponse**](SmartBchSep20MintResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchSep20Send

> []map[string]interface{} SmartBchSep20Send(ctx, smartBchSep20SendRequest)

Send some SmartBch SEP20 token amount to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSep20SendRequest** | [**SmartBchSep20SendRequest**](SmartBchSep20SendRequest.md)| place a SmartBch SEP20 token send request | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchSep20SendMax

> []map[string]interface{} SmartBchSep20SendMax(ctx, smartBchSep20SendMaxRequest)

Send all available SmartBch SEP20 token funds to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSep20SendMaxRequest** | [**SmartBchSep20SendMaxRequest**](SmartBchSep20SendMaxRequest.md)| Request to send all available SmartBch SEP20 token funds to a given address | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchSep20TokenInfo

> SmartBchSep20TokenInfoResponse SmartBchSep20TokenInfo(ctx, smartBchSep20TokenInfoRequest)

Get information about the SmartBch SEP20 token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSep20TokenInfoRequest** | [**SmartBchSep20TokenInfoRequest**](SmartBchSep20TokenInfoRequest.md)| Request to get information about the SmartBch SEP20 token  | 

### Return type

[**SmartBchSep20TokenInfoResponse**](SmartBchSep20TokenInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

