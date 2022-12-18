# \WalletBcmrApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BcmrAddMetadataRegistryAuthChain**](WalletBcmrApi.md#BcmrAddMetadataRegistryAuthChain) | **Post** /wallet/bcmr/add_registry_authchain | Add BCMR metadata registry from autchain, returning the built chain
[**BcmrAddRegistry**](WalletBcmrApi.md#BcmrAddRegistry) | **Post** /wallet/bcmr/add_registry | Add BCMR registry from parameter
[**BcmrAddRegistryFromUri**](WalletBcmrApi.md#BcmrAddRegistryFromUri) | **Post** /wallet/bcmr/add_registry_from_uri | Reset tracked BCMR registries
[**BcmrBuildAuthChain**](WalletBcmrApi.md#BcmrBuildAuthChain) | **Post** /wallet/bcmr/build_authchain | Build a BCMR authchain
[**BcmrGetRegistries**](WalletBcmrApi.md#BcmrGetRegistries) | **Post** /wallet/bcmr/get_registries | Get tracked BCMR registries
[**BcmrGetTokenInfo**](WalletBcmrApi.md#BcmrGetTokenInfo) | **Post** /wallet/bcmr/get_token_info | Get token info
[**BcmrResetRegistries**](WalletBcmrApi.md#BcmrResetRegistries) | **Post** /wallet/bcmr/reset_registries | Reset tracked BCMR registries



## BcmrAddMetadataRegistryAuthChain

> []AuthChainElement BcmrAddMetadataRegistryAuthChain(ctx, uNKNOWNBASETYPE)

Add BCMR metadata registry from autchain, returning the built chain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Add BCMR metadata registry by resolving an authchain, allowing for token info lookup  | 

### Return type

[**[]AuthChainElement**](AuthChainElement.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BcmrAddRegistry

> map[string]interface{} BcmrAddRegistry(ctx, requestBody)

Add BCMR registry from parameter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**map[string]map[string]interface{}**](map[string]interface{}.md)| Add a BCMR registry to the list of tracked, allowing for token info lookup  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BcmrAddRegistryFromUri

> map[string]interface{} BcmrAddRegistryFromUri(ctx, uNKNOWNBASETYPE)

Reset tracked BCMR registries

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Add a BCMR registry from remote URI to the list of tracked, allowing for token info lookup  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BcmrBuildAuthChain

> []AuthChainElement BcmrBuildAuthChain(ctx, uNKNOWNBASETYPE)

Build a BCMR authchain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Build an authchain - Zeroth-Descendant Transaction Chain, refer to https://github.com/bitjson/chip-bcmr#zeroth-descendant-transaction-chains The authchain in this implementation is specific to resolve to a valid metadata registy  | 

### Return type

[**[]AuthChainElement**](AuthChainElement.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BcmrGetRegistries

> []map[string]map[string]interface{} BcmrGetRegistries(ctx, optional)

Get tracked BCMR registries

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BcmrGetRegistriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BcmrGetRegistriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**optional.Interface of map[string]map[string]interface{}**](map[string]interface{}.md)| Get tracked BCMR registries.  | 

### Return type

[**[]map[string]map[string]interface{}**](map.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BcmrGetTokenInfo

> map[string]interface{} BcmrGetTokenInfo(ctx, uNKNOWNBASETYPE)

Get token info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Return the token info (the identity snapshot as per BCMR spec)  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BcmrResetRegistries

> map[string]interface{} BcmrResetRegistries(ctx, optional)

Reset tracked BCMR registries

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BcmrResetRegistriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BcmrResetRegistriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **optional.Map[string]interface{}**| Reset tracked BCMR registries, dropping all token info.  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

