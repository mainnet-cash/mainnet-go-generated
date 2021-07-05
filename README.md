# Go API client for openapi

A developer friendly bitcoin cash wallet api

This API is currently in *active* development, breaking changes may
be made prior to official release of version 1.0.0.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.3.18
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://rest-unstable.mainnet.cash*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContractApi* | [**ContractFn**](docs/ContractApi.md#contractfn) | **Post** /contract/call | Call a method on a contract
*ContractApi* | [**ContractUtxos**](docs/ContractApi.md#contractutxos) | **Post** /contract/utxos | List specific utxos on any contract
*ContractApi* | [**CreateContract**](docs/ContractApi.md#createcontract) | **Post** /contract/create | Create a cashscript contract
*ContractEscrowApi* | [**CreateEscrow**](docs/ContractEscrowApi.md#createescrow) | **Post** /contract/escrow/create | Create an escrow contract
*ContractEscrowApi* | [**EscrowFn**](docs/ContractEscrowApi.md#escrowfn) | **Post** /contract/escrow/call | Finalize an escrow contract
*FaucetApi* | [**GetAddresses**](docs/FaucetApi.md#getaddresses) | **Post** /faucet/get_addresses | Get addresses to return back or donate the testnet bch and tokens 
*FaucetApi* | [**GetTestnetBch**](docs/FaucetApi.md#gettestnetbch) | **Post** /faucet/get_testnet_bch | Get testnet bch 
*FaucetApi* | [**GetTestnetSlp**](docs/FaucetApi.md#gettestnetslp) | **Post** /faucet/get_testnet_slp | Get testnet slp tokens 
*MineApi* | [**Mine**](docs/MineApi.md#mine) | **Post** /mine | Mine regtest coins to a specified address
*UtilApi* | [**Convert**](docs/UtilApi.md#convert) | **Post** /util/convert | convert between units
*WalletApi* | [**Balance**](docs/WalletApi.md#balance) | **Post** /wallet/balance | Get total balance for wallet
*WalletApi* | [**CreateWallet**](docs/WalletApi.md#createwallet) | **Post** /wallet/create | create a new wallet
*WalletApi* | [**DepositAddress**](docs/WalletApi.md#depositaddress) | **Post** /wallet/deposit_address | Get a deposit address in cash address format
*WalletApi* | [**DepositQr**](docs/WalletApi.md#depositqr) | **Post** /wallet/deposit_qr | Get receiving cash address as a qrcode
*WalletApi* | [**Info**](docs/WalletApi.md#info) | **Post** /wallet/info | Get information about a wallet
*WalletApi* | [**MaxAmountToSend**](docs/WalletApi.md#maxamounttosend) | **Post** /wallet/max_amount_to_send | Get maximum spendable amount
*WalletApi* | [**NamedExists**](docs/WalletApi.md#namedexists) | **Post** /wallet/named_exists | Check if a named wallet already exists
*WalletApi* | [**ReplaceNamed**](docs/WalletApi.md#replacenamed) | **Post** /wallet/replace_named | Replace (recover) named wallet with a new walletId. If wallet with a provided name does not exist yet, it will be creted with a &#x60;walletId&#x60; supplied If wallet exists it will be overwritten without exception 
*WalletApi* | [**Send**](docs/WalletApi.md#send) | **Post** /wallet/send | Send some amount to a given address
*WalletApi* | [**SendMax**](docs/WalletApi.md#sendmax) | **Post** /wallet/send_max | Send all available funds to a given address
*WalletApi* | [**SignedMessageSign**](docs/WalletApi.md#signedmessagesign) | **Post** /wallet/signed/sign | Returns the message signature
*WalletApi* | [**SignedMessageVerify**](docs/WalletApi.md#signedmessageverify) | **Post** /wallet/signed/verify | Returns a boolean indicating whether message was valid for a given address
*WalletApi* | [**Utxos**](docs/WalletApi.md#utxos) | **Post** /wallet/utxo | Get detailed information about unspent outputs (utxos)
*WalletSlpApi* | [**NftChildGenesis**](docs/WalletSlpApi.md#nftchildgenesis) | **Post** /wallet/slp/nft_child_genesis | Get created tokenId back and new NFT child token balance of the wallet
*WalletSlpApi* | [**NftParentGenesis**](docs/WalletSlpApi.md#nftparentgenesis) | **Post** /wallet/slp/nft_parent_genesis | Get created tokenId back and new NFT token balance of the wallet
*WalletSlpApi* | [**SlpAllBalances**](docs/WalletSlpApi.md#slpallbalances) | **Post** /wallet/slp/all_balances | Get all slp balances of the wallet
*WalletSlpApi* | [**SlpBalance**](docs/WalletSlpApi.md#slpbalance) | **Post** /wallet/slp/balance | Get total slp token balance of the wallet
*WalletSlpApi* | [**SlpCreateWallet**](docs/WalletSlpApi.md#slpcreatewallet) | **Post** /wallet/slp/create | create a new SLP wallet
*WalletSlpApi* | [**SlpDepositAddress**](docs/WalletSlpApi.md#slpdepositaddress) | **Post** /wallet/slp/deposit_address | Get an SLP deposit address in cash address format
*WalletSlpApi* | [**SlpDepositQr**](docs/WalletSlpApi.md#slpdepositqr) | **Post** /wallet/slp/deposit_qr | Get an SLP receiving cash address as a qrcode
*WalletSlpApi* | [**SlpGenesis**](docs/WalletSlpApi.md#slpgenesis) | **Post** /wallet/slp/genesis | Get created tokenId back and new slp token balance of the wallet
*WalletSlpApi* | [**SlpMint**](docs/WalletSlpApi.md#slpmint) | **Post** /wallet/slp/mint | Get created tokenId back and new slp token balance of the wallet
*WalletSlpApi* | [**SlpOutpoints**](docs/WalletSlpApi.md#slpoutpoints) | **Post** /wallet/slp/outpoints | Get list of unspent SLP outpoints.
*WalletSlpApi* | [**SlpSend**](docs/WalletSlpApi.md#slpsend) | **Post** /wallet/slp/send | Send some SLP token amount to a given cash address
*WalletSlpApi* | [**SlpSendMax**](docs/WalletSlpApi.md#slpsendmax) | **Post** /wallet/slp/send_max | Send all available SLP funds to a given address
*WalletSlpApi* | [**SlpTokenInfo**](docs/WalletSlpApi.md#slptokeninfo) | **Post** /wallet/slp/token_info | Get information about the token
*WalletSlpApi* | [**SlpUtxos**](docs/WalletSlpApi.md#slputxos) | **Post** /wallet/slp/utxo | Get detailed information about unspent SLP outputs (utxos)
*WebhookApi* | [**WatchAddress**](docs/WebhookApi.md#watchaddress) | **Post** /webhook/watch_address | Create a webhook to watch cashaddress balance and transactions. 


## Documentation For Models

 - [BalanceRequest](docs/BalanceRequest.md)
 - [BalanceResponse](docs/BalanceResponse.md)
 - [CashscriptReceipt](docs/CashscriptReceipt.md)
 - [Contract](docs/Contract.md)
 - [ContractFnRequest](docs/ContractFnRequest.md)
 - [ContractFnResponse](docs/ContractFnResponse.md)
 - [ContractRequest](docs/ContractRequest.md)
 - [ContractResponse](docs/ContractResponse.md)
 - [ConvertRequest](docs/ConvertRequest.md)
 - [CreateSignedMessageRequest](docs/CreateSignedMessageRequest.md)
 - [DepositAddressResponse](docs/DepositAddressResponse.md)
 - [Error](docs/Error.md)
 - [EscrowFnRequest](docs/EscrowFnRequest.md)
 - [EscrowRequest](docs/EscrowRequest.md)
 - [GetAddressesResponse](docs/GetAddressesResponse.md)
 - [GetTestnetBchRequest](docs/GetTestnetBchRequest.md)
 - [GetTestnetBchResponse](docs/GetTestnetBchResponse.md)
 - [GetTestnetSlpRequest](docs/GetTestnetSlpRequest.md)
 - [GetTestnetSlpResponse](docs/GetTestnetSlpResponse.md)
 - [MaxAmountToSendRequest](docs/MaxAmountToSendRequest.md)
 - [MineRequest](docs/MineRequest.md)
 - [NetworkEnum](docs/NetworkEnum.md)
 - [ScalableVectorGraphic](docs/ScalableVectorGraphic.md)
 - [SendMaxRequest](docs/SendMaxRequest.md)
 - [SendMaxResponse](docs/SendMaxResponse.md)
 - [SendRequest](docs/SendRequest.md)
 - [SendRequestItem](docs/SendRequestItem.md)
 - [SendRequestOptions](docs/SendRequestOptions.md)
 - [SendResponse](docs/SendResponse.md)
 - [SerializedWallet](docs/SerializedWallet.md)
 - [SignedMessageResponse](docs/SignedMessageResponse.md)
 - [SignedMessageResponseDetails](docs/SignedMessageResponseDetails.md)
 - [SignedMessageResponseRaw](docs/SignedMessageResponseRaw.md)
 - [SlpBalanceRequest](docs/SlpBalanceRequest.md)
 - [SlpBalanceResponse](docs/SlpBalanceResponse.md)
 - [SlpDepositAddressResponse](docs/SlpDepositAddressResponse.md)
 - [SlpGenesisRequest](docs/SlpGenesisRequest.md)
 - [SlpGenesisResponse](docs/SlpGenesisResponse.md)
 - [SlpMintRequest](docs/SlpMintRequest.md)
 - [SlpMintResponse](docs/SlpMintResponse.md)
 - [SlpOutpointsResponse](docs/SlpOutpointsResponse.md)
 - [SlpSendMaxRequest](docs/SlpSendMaxRequest.md)
 - [SlpSendRequest](docs/SlpSendRequest.md)
 - [SlpSendRequestItem](docs/SlpSendRequestItem.md)
 - [SlpSendRequestOptions](docs/SlpSendRequestOptions.md)
 - [SlpSendResponse](docs/SlpSendResponse.md)
 - [SlpTokenInfoRequest](docs/SlpTokenInfoRequest.md)
 - [SlpTokenInfoResponseItem](docs/SlpTokenInfoResponseItem.md)
 - [SlpUtxo](docs/SlpUtxo.md)
 - [SlpUtxoResponse](docs/SlpUtxoResponse.md)
 - [UnitType](docs/UnitType.md)
 - [Utxo](docs/Utxo.md)
 - [UtxoResponse](docs/UtxoResponse.md)
 - [VerifySignedMessageRequest](docs/VerifySignedMessageRequest.md)
 - [VerifySignedMessageResponse](docs/VerifySignedMessageResponse.md)
 - [VerifySignedMessageResponseDetails](docs/VerifySignedMessageResponseDetails.md)
 - [WalletInfo](docs/WalletInfo.md)
 - [WalletMnemonic](docs/WalletMnemonic.md)
 - [WalletNamedExistsRequest](docs/WalletNamedExistsRequest.md)
 - [WalletReplaceNamedRequest](docs/WalletReplaceNamedRequest.md)
 - [WalletRequest](docs/WalletRequest.md)
 - [WalletResponse](docs/WalletResponse.md)
 - [WalletType](docs/WalletType.md)
 - [WatchAddressRequest](docs/WatchAddressRequest.md)
 - [WatchAddressResponse](docs/WatchAddressResponse.md)
 - [Wif](docs/Wif.md)
 - [ZeroBalanceResponse](docs/ZeroBalanceResponse.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author

hello@mainnet.cash

