# \ContractEscrowApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEscrow**](ContractEscrowApi.md#CreateEscrow) | **Post** /contract/escrow/create | Create an escrow contract
[**EscrowFn**](ContractEscrowApi.md#EscrowFn) | **Post** /contract/escrow/call | Finalize an escrow contract
[**EscrowInfo**](ContractEscrowApi.md#EscrowInfo) | **Post** /contract/escrow/info | Get information about an escrow contract from the escrowContractId
[**EscrowUtxos**](ContractEscrowApi.md#EscrowUtxos) | **Post** /contract/escrow/utxos | List specific utxos on any escrow contract



## CreateEscrow

> EscrowResponse CreateEscrow(ctx, escrowRequest)

Create an escrow contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**escrowRequest** | [**EscrowRequest**](EscrowRequest.md)| Request a new escrow contract from a template | 

### Return type

[**EscrowResponse**](EscrowResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EscrowInfo

> EscrowInfoResponse EscrowInfo(ctx, escrowInfoRequest)

Get information about an escrow contract from the escrowContractId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**escrowInfoRequest** | [**EscrowInfoRequest**](EscrowInfoRequest.md)| Request parsed information regarding a contract from the escrowContractId | 

### Return type

[**EscrowInfoResponse**](EscrowInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EscrowUtxos

> UtxoResponse EscrowUtxos(ctx, escrowContract)

List specific utxos on any escrow contract

Returns all UTXOs that can be spent by the contract. Both confirmed and unconfirmed UTXOs are included. The endpoint works with contracts generated from templates (i.e. escrow). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**escrowContract** | [**EscrowContract**](EscrowContract.md)|  | 

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

