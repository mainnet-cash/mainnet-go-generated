# \WalletUtilApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UtilDecodeTransaction**](WalletUtilApi.md#UtilDecodeTransaction) | **Post** /wallet/util/decode_transaction | Decode a bitcoin transaction. Accepts both transaction hash or raw transaction in hex format.



## UtilDecodeTransaction

> ElectrumRawTransaction UtilDecodeTransaction(ctx, utilDecodeTransactionRequest)

Decode a bitcoin transaction. Accepts both transaction hash or raw transaction in hex format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**utilDecodeTransactionRequest** | [**UtilDecodeTransactionRequest**](UtilDecodeTransactionRequest.md)| Request to decode a transaction  | 

### Return type

[**ElectrumRawTransaction**](ElectrumRawTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

