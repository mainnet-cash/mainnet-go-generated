# \WalletSlpApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NftChildGenesis**](WalletSlpApi.md#NftChildGenesis) | **Post** /wallet/slp/nft_child_genesis | Get created tokenId back and new NFT child token balance of the wallet
[**NftParentGenesis**](WalletSlpApi.md#NftParentGenesis) | **Post** /wallet/slp/nft_parent_genesis | Get created tokenId back and new NFT token balance of the wallet
[**SlpAllBalances**](WalletSlpApi.md#SlpAllBalances) | **Post** /wallet/slp/all_balances | Get all slp balances of the wallet
[**SlpBalance**](WalletSlpApi.md#SlpBalance) | **Post** /wallet/slp/balance | Get total slp token balance of the wallet
[**SlpCreateWallet**](WalletSlpApi.md#SlpCreateWallet) | **Post** /wallet/slp/create | create a new SLP wallet
[**SlpDepositAddress**](WalletSlpApi.md#SlpDepositAddress) | **Post** /wallet/slp/deposit_address | Get an SLP deposit address in cash address format
[**SlpDepositQr**](WalletSlpApi.md#SlpDepositQr) | **Post** /wallet/slp/deposit_qr | Get an SLP receiving cash address as a qrcode
[**SlpGenesis**](WalletSlpApi.md#SlpGenesis) | **Post** /wallet/slp/genesis | Get created tokenId back and new slp token balance of the wallet
[**SlpMint**](WalletSlpApi.md#SlpMint) | **Post** /wallet/slp/mint | Get created tokenId back and new slp token balance of the wallet
[**SlpOutpoints**](WalletSlpApi.md#SlpOutpoints) | **Post** /wallet/slp/outpoints | Get list of unspent SLP outpoints.
[**SlpSend**](WalletSlpApi.md#SlpSend) | **Post** /wallet/slp/send | Send some SLP token amount to a given cash address
[**SlpSendMax**](WalletSlpApi.md#SlpSendMax) | **Post** /wallet/slp/send_max | Send all available SLP funds to a given address
[**SlpTokenInfo**](WalletSlpApi.md#SlpTokenInfo) | **Post** /wallet/slp/token_info | Get information about the token
[**SlpUtxos**](WalletSlpApi.md#SlpUtxos) | **Post** /wallet/slp/utxo | Get detailed information about unspent SLP outputs (utxos)



## NftChildGenesis

> SlpGenesisResponse NftChildGenesis(ctx, slpGenesisRequest)

Get created tokenId back and new NFT child token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpGenesisRequest** | [**SlpGenesisRequest**](SlpGenesisRequest.md)| Request to create a new NFT child token (genesis) by consuming a parent token  | 

### Return type

[**SlpGenesisResponse**](SlpGenesisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NftParentGenesis

> SlpGenesisResponse NftParentGenesis(ctx, slpGenesisRequest)

Get created tokenId back and new NFT token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpGenesisRequest** | [**SlpGenesisRequest**](SlpGenesisRequest.md)| Request to create a new NFT parent token (genesis)  | 

### Return type

[**SlpGenesisResponse**](SlpGenesisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpAllBalances

> []SlpBalanceResponse SlpAllBalances(ctx, serializedWallet)

Get all slp balances of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a wallet slp balances  | 

### Return type

[**[]SlpBalanceResponse**](SlpBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpBalance

> SlpBalanceResponse SlpBalance(ctx, slpBalanceRequest)

Get total slp token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpBalanceRequest** | [**SlpBalanceRequest**](SlpBalanceRequest.md)| Request for a wallet slp token balance  | 

### Return type

[**SlpBalanceResponse**](SlpBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpCreateWallet

> WalletResponse SlpCreateWallet(ctx, walletRequest)

create a new SLP wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletRequest** | [**WalletRequest**](WalletRequest.md)| Request a new SLP wallet | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpDepositAddress

> SlpDepositAddressResponse SlpDepositAddress(ctx, serializedWallet)

Get an SLP deposit address in cash address format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for an SLP deposit address given a wallet  | 

### Return type

[**SlpDepositAddressResponse**](SlpDepositAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpDepositQr

> ScalableVectorGraphic SlpDepositQr(ctx, serializedWallet)

Get an SLP receiving cash address as a qrcode

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for an SLP deposit cash address as a Quick Response code (qrcode)  | 

### Return type

[**ScalableVectorGraphic**](ScalableVectorGraphic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpGenesis

> SlpGenesisResponse SlpGenesis(ctx, slpGenesisRequest)

Get created tokenId back and new slp token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpGenesisRequest** | [**SlpGenesisRequest**](SlpGenesisRequest.md)| Request to create a new SLP token (genesis)  | 

### Return type

[**SlpGenesisResponse**](SlpGenesisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpMint

> SlpMintResponse SlpMint(ctx, slpMintRequest)

Get created tokenId back and new slp token balance of the wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpMintRequest** | [**SlpMintRequest**](SlpMintRequest.md)| Request to mint more of SLP tokens  | 

### Return type

[**SlpMintResponse**](SlpMintResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpOutpoints

> SlpOutpointsResponse SlpOutpoints(ctx, serializedWallet)

Get list of unspent SLP outpoints.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request of unspent SLP outpoints  | 

### Return type

[**SlpOutpointsResponse**](SlpOutpointsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpSend

> SlpSendResponse SlpSend(ctx, slpSendRequest)

Send some SLP token amount to a given cash address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpSendRequest** | [**SlpSendRequest**](SlpSendRequest.md)| place an SLP send request | 

### Return type

[**SlpSendResponse**](SlpSendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpSendMax

> SlpSendResponse SlpSendMax(ctx, slpSendMaxRequest)

Send all available SLP funds to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpSendMaxRequest** | [**SlpSendMaxRequest**](SlpSendMaxRequest.md)| Request to send all available SLP funds to a given address | 

### Return type

[**SlpSendResponse**](SlpSendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpTokenInfo

> SlpTokenInfoResponseItem SlpTokenInfo(ctx, slpTokenInfoRequest)

Get information about the token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slpTokenInfoRequest** | [**SlpTokenInfoRequest**](SlpTokenInfoRequest.md)| Request to get information about the token  | 

### Return type

[**SlpTokenInfoResponseItem**](SlpTokenInfoResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlpUtxos

> SlpUtxoResponse SlpUtxos(ctx, serializedWallet)

Get detailed information about unspent SLP outputs (utxos)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request detailed list of unspent SLP transaction outputs  | 

### Return type

[**SlpUtxoResponse**](SlpUtxoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

