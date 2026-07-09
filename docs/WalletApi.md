# \WalletApi

All URIs are relative to *https://rest-unstable.mainnet.cash*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Balance**](WalletApi.md#Balance) | **Post** /wallet/balance | Get total balance for wallet
[**CreateWallet**](WalletApi.md#CreateWallet) | **Post** /wallet/create | create a new wallet
[**DepositAddress**](WalletApi.md#DepositAddress) | **Post** /wallet/deposit_address | Get a deposit address in cash address format
[**EncodeTransaction**](WalletApi.md#EncodeTransaction) | **Post** /wallet/encode_transaction | Encode and sign a transaction given a list of sendRequests, options and estimate fees
[**GetAllNftTokenBalances**](WalletApi.md#GetAllNftTokenBalances) | **Post** /wallet/get_all_nft_token_balances | Get non-fungible token balance
[**GetAllTokenBalances**](WalletApi.md#GetAllTokenBalances) | **Post** /wallet/get_all_token_balances | Get non-fungible token balance
[**GetHistory**](WalletApi.md#GetHistory) | **Post** /wallet/get_history | Get a list of transactions related to a wallet
[**GetNftTokenBalance**](WalletApi.md#GetNftTokenBalance) | **Post** /wallet/get_nft_token_balance | Get non-fungible token balance
[**GetTokenBalance**](WalletApi.md#GetTokenBalance) | **Post** /wallet/get_token_balance | Get fungible token balance
[**GetTokenUtxos**](WalletApi.md#GetTokenUtxos) | **Post** /wallet/get_token_utxos | Get token utxos
[**Info**](WalletApi.md#Info) | **Post** /wallet/info | Get information about a wallet
[**MaxAmountToSend**](WalletApi.md#MaxAmountToSend) | **Post** /wallet/max_amount_to_send | Get maximum spendable amount
[**NamedExists**](WalletApi.md#NamedExists) | **Post** /wallet/named_exists | Check if a named wallet already exists
[**ReplaceNamed**](WalletApi.md#ReplaceNamed) | **Post** /wallet/replace_named | Replace (recover) named wallet with a new walletId. If wallet with a provided name does not exist yet, it will be creted with a &#x60;walletId&#x60; supplied If wallet exists it will be overwritten without exception 
[**Send**](WalletApi.md#Send) | **Post** /wallet/send | Send some amount to a given address
[**SendMax**](WalletApi.md#SendMax) | **Post** /wallet/send_max | Send all available funds to a given address
[**SubmitTransaction**](WalletApi.md#SubmitTransaction) | **Post** /wallet/submit_transaction | submit an encoded and signed transaction to the network
[**TokenBurn**](WalletApi.md#TokenBurn) | **Post** /wallet/token_burn | Perform an explicit token burn
[**TokenDepositAddress**](WalletApi.md#TokenDepositAddress) | **Post** /wallet/token_deposit_address | Get a token aware deposit address in cash address format
[**TokenGenesis**](WalletApi.md#TokenGenesis) | **Post** /wallet/token_genesis | Create new token category
[**TokenMint**](WalletApi.md#TokenMint) | **Post** /wallet/token_mint | Mint new non-fungible tokens
[**Utxos**](WalletApi.md#Utxos) | **Post** /wallet/utxo | Get detailed information about unspent outputs (utxos)
[**Xpubkeys**](WalletApi.md#Xpubkeys) | **Post** /wallet/xpubkeys | A set of xpubkeys and paths for the wallet.



## Balance

> BalanceResponse Balance(ctx, balanceRequest)

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWallet

> WalletResponse CreateWallet(ctx, walletRequest)

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


## DepositAddress

> DepositAddressResponse DepositAddress(ctx, serializedWallet)

Get a deposit address in cash address format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a deposit address given a wallet  | 

### Return type

[**DepositAddressResponse**](DepositAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncodeTransaction

> EncodeTransactionResponse EncodeTransaction(ctx, encodeTransactionRequest)

Encode and sign a transaction given a list of sendRequests, options and estimate fees

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodeTransactionRequest** | [**EncodeTransactionRequest**](EncodeTransactionRequest.md)| encode a transaction | 

### Return type

[**EncodeTransactionResponse**](EncodeTransactionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNftTokenBalances

> map[string]float32 GetAllNftTokenBalances(ctx, inlineObject5)

Get non-fungible token balance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject5** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

**map[string]float32**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTokenBalances

> map[string]float32 GetAllTokenBalances(ctx, inlineObject4)

Get non-fungible token balance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject4** | [**InlineObject4**](InlineObject4.md)|  | 

### Return type

**map[string]float32**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistory

> []TransactionHistoryItem GetHistory(ctx, historyRequest)

Get a list of transactions related to a wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**historyRequest** | [**HistoryRequest**](HistoryRequest.md)| Gets transaction history of this wallet with most data decoded and ready to present to user   Note: balance calculations are valid only if querying to the blockchain tip (&#x60;toHeight&#x60; &#x3D;&#x3D;&#x3D; -1, &#x60;count&#x60; &#x3D;&#x3D;&#x3D; -1)   Note: this method is heavy on network calls, if invoked in browser use of cache is advised, @see &#x60;Config.UseLocalStorageCache&#x60;   Note: this method tries to recreate the history tab view of Electron Cash wallet, however, it may not be 100% accurate if the tnransaction value changes are the same in the same block (ordering)  | 

### Return type

[**[]TransactionHistoryItem**](TransactionHistoryItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftTokenBalance

> map[string]interface{} GetNftTokenBalance(ctx, inlineObject3)

Get non-fungible token balance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | 

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


## GetTokenBalance

> map[string]interface{} GetTokenBalance(ctx, inlineObject2)

Get fungible token balance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | 

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


## GetTokenUtxos

> []Utxo GetTokenUtxos(ctx, inlineObject1)

Get token utxos

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**[]Utxo**](Utxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Info

> WalletInfo Info(ctx, serializedWallet)

Get information about a wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| The wallet to request information about, in serialized form.  | 

### Return type

[**WalletInfo**](WalletInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaxAmountToSend

> BalanceResponse MaxAmountToSend(ctx, maxAmountToSendRequest)

Get maximum spendable amount

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maxAmountToSendRequest** | [**MaxAmountToSendRequest**](MaxAmountToSendRequest.md)| get amount that will be spend with a spend max request. If a unit type is specified, a numeric value will be returned. | 

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


## NamedExists

> WalletNamedExistsResponse NamedExists(ctx, walletNamedExistsRequest)

Check if a named wallet already exists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletNamedExistsRequest** | [**WalletNamedExistsRequest**](WalletNamedExistsRequest.md)| Request parameters | 

### Return type

[**WalletNamedExistsResponse**](WalletNamedExistsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamed

> WalletReplaceNamedResponse ReplaceNamed(ctx, walletReplaceNamedRequest)

Replace (recover) named wallet with a new walletId. If wallet with a provided name does not exist yet, it will be creted with a `walletId` supplied If wallet exists it will be overwritten without exception 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletReplaceNamedRequest** | [**WalletReplaceNamedRequest**](WalletReplaceNamedRequest.md)| Request parameters | 

### Return type

[**WalletReplaceNamedResponse**](WalletReplaceNamedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> SendResponse Send(ctx, sendRequest)

Send some amount to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sendRequest** | [**SendRequest**](SendRequest.md)| place a send request | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMax

> SendMaxResponse SendMax(ctx, sendMaxRequest)

Send all available funds to a given address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sendMaxRequest** | [**SendMaxRequest**](SendMaxRequest.md)| Request to send all available funds to a given address | 

### Return type

[**SendMaxResponse**](SendMaxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitTransaction

> SubmitTransactionResponse SubmitTransaction(ctx, submitTransactionRequest)

submit an encoded and signed transaction to the network

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submitTransactionRequest** | [**SubmitTransactionRequest**](SubmitTransactionRequest.md)| submit an encoded and signed transaction to the network | 

### Return type

[**SubmitTransactionResponse**](SubmitTransactionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenBurn

> SendResponse TokenBurn(ctx, tokenBurnRequest)

Perform an explicit token burn

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenBurnRequest** | [**TokenBurnRequest**](TokenBurnRequest.md)| Perform an explicit token burning by spending a token utxo to an OP_RETURN Behaves differently for fungible and non-fungible tokens:  * NFTs are always \&quot;destroyed\&quot;  * FTs&#39; amount is reduced by the amount specified, if 0 FT amount is left and no NFT present, the token is \&quot;destroyed\&quot; Refer to spec https://github.com/bitjson/cashtokens  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenDepositAddress

> DepositAddressResponse TokenDepositAddress(ctx, serializedWallet)

Get a token aware deposit address in cash address format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request for a token aware deposit address given a wallet  | 

### Return type

[**DepositAddressResponse**](DepositAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenGenesis

> SendResponse TokenGenesis(ctx, tokenGenesisRequest)

Create new token category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenGenesisRequest** | [**TokenGenesisRequest**](TokenGenesisRequest.md)| Create new cashtoken, both funglible and/or non-fungible (NFT) Refer to spec https://github.com/bitjson/cashtokens Newly created token identifier can be found in &#x60;categories&#x60; field.  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenMint

> SendResponse TokenMint(ctx, tokenMintRequest)

Mint new non-fungible tokens

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenMintRequest** | [**TokenMintRequest**](TokenMintRequest.md)| Mint new NFT cashtokens using an existing minting token Refer to spec https://github.com/bitjson/cashtokens Newly minted tokens will retain the parent&#39;s category.  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Utxos

> []Utxo Utxos(ctx, serializedWallet)

Get detailed information about unspent outputs (utxos)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serializedWallet** | [**SerializedWallet**](SerializedWallet.md)| Request detailed list of unspent transaction outputs  | 

### Return type

[**[]Utxo**](Utxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Xpubkeys

> XPubKeyResponse Xpubkeys(ctx, xPubKeyRequest)

A set of xpubkeys and paths for the wallet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xPubKeyRequest** | [**XPubKeyRequest**](XPubKeyRequest.md)| x  | 

### Return type

[**XPubKeyResponse**](XPubKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

