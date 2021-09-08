# \SmartbchContractApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartBchContractCall**](SmartbchContractApi.md#SmartBchContractCall) | **Post** /smartbch/contract/call | Call a SmartBch contract function
[**SmartBchContractCreate**](SmartbchContractApi.md#SmartBchContractCreate) | **Post** /smartbch/contract/create | Create a SmartBch contractId
[**SmartBchContractDeploy**](SmartbchContractApi.md#SmartBchContractDeploy) | **Post** /smartbch/contract/deploy | Request to deploy a SmartBch contract
[**SmartBchContractEstimateGas**](SmartbchContractApi.md#SmartBchContractEstimateGas) | **Post** /smartbch/contract/estimate_gas | Estimate the gas for a contract interaction function given the arguments
[**SmartBchContractInfo**](SmartbchContractApi.md#SmartBchContractInfo) | **Post** /smartbch/contract/info | Get information about a SmartBch contract from the contractId



## SmartBchContractCall

> SmartBchContractFnCallResponse SmartBchContractCall(ctx, smartBchContractFnCallRequest)

Call a SmartBch contract function

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchContractFnCallRequest** | [**SmartBchContractFnCallRequest**](SmartBchContractFnCallRequest.md)|  | 

### Return type

[**SmartBchContractFnCallResponse**](SmartBchContractFnCallResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchContractCreate

> SmartBchContractResponse SmartBchContractCreate(ctx, smartBchContractRequest)

Create a SmartBch contractId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchContractRequest** | [**SmartBchContractRequest**](SmartBchContractRequest.md)| Create a SmartBch contractId | 

### Return type

[**SmartBchContractResponse**](SmartBchContractResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchContractDeploy

> SmartBchContractDeployResponse SmartBchContractDeploy(ctx, smartBchContractDeployRequest)

Request to deploy a SmartBch contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchContractDeployRequest** | [**SmartBchContractDeployRequest**](SmartBchContractDeployRequest.md)| Request to deploy a SmartBch contract | 

### Return type

[**SmartBchContractDeployResponse**](SmartBchContractDeployResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchContractEstimateGas

> SmartBchContractEstimateGasResponse SmartBchContractEstimateGas(ctx, smartBchContractEstimateGasRequest)

Estimate the gas for a contract interaction function given the arguments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchContractEstimateGasRequest** | [**SmartBchContractEstimateGasRequest**](SmartBchContractEstimateGasRequest.md)|  | 

### Return type

[**SmartBchContractEstimateGasResponse**](SmartBchContractEstimateGasResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartBchContractInfo

> SmartBchContractInfoResponse SmartBchContractInfo(ctx, smartBchContractInfoRequest)

Get information about a SmartBch contract from the contractId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchContractInfoRequest** | [**SmartBchContractInfoRequest**](SmartBchContractInfoRequest.md)| Request parsed information regarding a SmartBch contract from the smartBchContractId | 

### Return type

[**SmartBchContractInfoResponse**](SmartBchContractInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

