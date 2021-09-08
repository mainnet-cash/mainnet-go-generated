# \SmartbchWalletApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartbchBalance**](SmartbchWalletApi.md#SmartbchBalance) | **Post** /smartbch/wallet/balance | Get total balance for wallet
[**SmartbchCreateWallet**](SmartbchWalletApi.md#SmartbchCreateWallet) | **Post** /smartbch/wallet/create | create a new wallet
[**SmartbchDepositAddress**](SmartbchWalletApi.md#SmartbchDepositAddress) | **Post** /smartbch/wallet/deposit_address | Get a deposit address
[**SmartbchDepositQr**](SmartbchWalletApi.md#SmartbchDepositQr) | **Post** /smartbch/wallet/deposit_qr | Get receiving cash address as a qrcode
[**SmartbchMaxAmountToSend**](SmartbchWalletApi.md#SmartbchMaxAmountToSend) | **Post** /smartbch/wallet/max_amount_to_send | Get maximum spendable amount
[**SmartbchSend**](SmartbchWalletApi.md#SmartbchSend) | **Post** /smartbch/wallet/send | Send some amount to a given address
[**SmartbchSendMax**](SmartbchWalletApi.md#SmartbchSendMax) | **Post** /smartbch/wallet/send_max | Send all available funds to a given address
[**SmartbchSignedMessageSign**](SmartbchWalletApi.md#SmartbchSignedMessageSign) | **Post** /smartbch/wallet/signed/sign | Returns the message signature
[**SmartbchSignedMessageVerify**](SmartbchWalletApi.md#SmartbchSignedMessageVerify) | **Post** /smartbch/wallet/signed/verify | Returns a boolean indicating whether message was valid for a given address



## SmartbchBalance

> BalanceResponse SmartbchBalance(ctx, balanceRequest)

Get total balance for wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balanceRequest** | [**BalanceRequest**](BalanceRequest.md)| Request for a wallet balance  | 

### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartbchCreateWallet

> WalletResponse SmartbchCreateWallet(ctx, walletRequest)

create a new wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletRequest** | [**WalletRequest**](WalletRequest.md)| Request a new random wallet | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartbchDepositAddress

> SmartBchDepositAddressResponse SmartbchDepositAddress(ctx, serializedWallet)

Get a deposit address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a deposit address given a wallet  | 

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


## SmartbchDepositQr

> ScalableVectorGraphic SmartbchDepositQr(ctx, serializedWallet)

Get receiving cash address as a qrcode

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a deposit cash address as a Quick Response code (qrcode)  | 

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


## SmartbchMaxAmountToSend

> BalanceResponse SmartbchMaxAmountToSend(ctx, smartBchMaxAmountToSendRequest)

Get maximum spendable amount

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchMaxAmountToSendRequest** | [**SmartBchMaxAmountToSendRequest**](SmartBchMaxAmountToSendRequest.md)| get amount that will be spend with a spend max request. If a unit type is specified, a numeric value will be returned. | 

### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartbchSend

> []SmartBchSendResponseItem SmartbchSend(ctx, smartBchSendRequest)

Send some amount to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSendRequest** | [**SmartBchSendRequest**](SmartBchSendRequest.md)| place a send request | 

### Return type

[**[]SmartBchSendResponseItem**](SmartBchSendResponseItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartbchSendMax

> SmartBchSendResponseItem SmartbchSendMax(ctx, smartBchSendMaxRequest)

Send all available funds to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartBchSendMaxRequest** | [**SmartBchSendMaxRequest**](SmartBchSendMaxRequest.md)| Request to send all available funds to a given address | 

### Return type

[**SmartBchSendResponseItem**](SmartBchSendResponseItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartbchSignedMessageSign

> SignedMessageResponse SmartbchSignedMessageSign(ctx, optional)

Returns the message signature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmartbchSignedMessageSignOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SmartbchSignedMessageSignOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSignedMessageRequest** | [**optional.Interface of CreateSignedMessageRequest**](CreateSignedMessageRequest.md)| Sign a message  | 

### Return type

[**SignedMessageResponse**](SignedMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartbchSignedMessageVerify

> VerifySignedMessageResponse SmartbchSignedMessageVerify(ctx, optional)

Returns a boolean indicating whether message was valid for a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmartbchSignedMessageVerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SmartbchSignedMessageVerifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifySignedMessageRequest** | [**optional.Interface of VerifySignedMessageRequest**](VerifySignedMessageRequest.md)| Sign a message  | 

### Return type

[**VerifySignedMessageResponse**](VerifySignedMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

