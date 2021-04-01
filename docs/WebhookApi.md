# \WebhookApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchAddress**](WebhookApi.md#WatchAddress) | **Post** /webhook/watch_address | Create a webhook to watch cashaddress balance and transactions. 



## WatchAddress

> WatchAddressResponse WatchAddress(ctx, watchAddressRequest)

Create a webhook to watch cashaddress balance and transactions. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchAddressRequest** | [**WatchAddressRequest**](WatchAddressRequest.md)| Based on the &#39;type&#39; parameter the webhook will be triggered to either post balance or raw transactions to the &#39;url&#39; - &#39;transaction:in&#39; for incoming only, &#39;transaction:out&#39; for outgoing only and &#39;transaction:in,out&#39; both for incoming and outgoing transactions. &#39;balance&#39; will post the object according to &#39;BalanceResponse&#39; schema  | 

### Return type

[**WatchAddressResponse**](WatchAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

