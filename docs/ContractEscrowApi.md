# \ContractEscrowApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEscrow**](ContractEscrowApi.md#CreateEscrow) | **Post** /contract/escrow/create | Create an escrow contract
[**EscrowFn**](ContractEscrowApi.md#EscrowFn) | **Post** /contract/escrow/call | Finalize an escrow contract



## CreateEscrow

> ContractResponse CreateEscrow(ctx, escrowRequest)

Create an escrow contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**escrowRequest** | [**EscrowRequest**](EscrowRequest.md)| Request a new escrow contract from a template | 

### Return type

[**ContractResponse**](ContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EscrowFn

> ContractFnResponse EscrowFn(ctx, escrowFnRequest)

Finalize an escrow contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**escrowFnRequest** | [**EscrowFnRequest**](EscrowFnRequest.md)|  | 

### Return type

[**ContractFnResponse**](ContractFnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

