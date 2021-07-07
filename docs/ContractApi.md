# \ContractApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractFn**](ContractApi.md#ContractFn) | **Post** /contract/call | Call a method on a contract
[**ContractUtxos**](ContractApi.md#ContractUtxos) | **Post** /contract/utxos | List specific utxos on any contract
[**CreateContract**](ContractApi.md#CreateContract) | **Post** /contract/create | Create a cashscript contract



## ContractFn

> ContractFnResponse ContractFn(ctx, contractFnRequest)

Call a method on a contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractFnRequest** | [**ContractFnRequest**](ContractFnRequest.md)| Request a new cashscript contract | 

### Return type

[**ContractFnResponse**](ContractFnResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractUtxos

> UtxoResponse ContractUtxos(ctx, contract)

List specific utxos on any contract

Returns all UTXOs that can be spent by the contract. Both confirmed and unconfirmed UTXOs are included. The endpoint works with contracts generated from templates (i.e. escrow). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | [**Contract**](Contract.md)|  | 

### Return type

[**UtxoResponse**](UtxoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContract

> ContractResponse CreateContract(ctx, contractRequest)

Create a cashscript contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractRequest** | [**ContractRequest**](ContractRequest.md)| Request a new cashscript contract | 

### Return type

[**ContractResponse**](ContractResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

