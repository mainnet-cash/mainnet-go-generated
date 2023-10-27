# Go API client for openapi

A developer friendly bitcoin cash wallet api

This API is currently in *active* development, breaking changes may
be made prior to official release of version 1.0.0.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.2.6
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
*ContractApi* | [**ContractInfo**](docs/ContractApi.md#contractinfo) | **Post** /contract/info | Get information about a contract from the contractId
*ContractApi* | [**ContractUtxos**](docs/ContractApi.md#contractutxos) | **Post** /contract/utxos | List specific utxos on any contract
*ContractApi* | [**CreateContract**](docs/ContractApi.md#createcontract) | **Post** /contract/create | Create a cashscript contract
*ContractEscrowApi* | [**CreateEscrow**](docs/ContractEscrowApi.md#createescrow) | **Post** /contract/escrow/create | Create an escrow contract
*ContractEscrowApi* | [**EscrowFn**](docs/ContractEscrowApi.md#escrowfn) | **Post** /contract/escrow/call | Finalize an escrow contract
*ContractEscrowApi* | [**EscrowInfo**](docs/ContractEscrowApi.md#escrowinfo) | **Post** /contract/escrow/info | Get information about an escrow contract from the escrowContractId
*ContractEscrowApi* | [**EscrowUtxos**](docs/ContractEscrowApi.md#escrowutxos) | **Post** /contract/escrow/utxos | List specific utxos on any escrow contract
*FaucetApi* | [**GetAddresses**](docs/FaucetApi.md#getaddresses) | **Post** /faucet/get_addresses | Get addresses to return back or donate the testnet bch and tokens 
*FaucetApi* | [**GetTestnetBch**](docs/FaucetApi.md#gettestnetbch) | **Post** /faucet/get_testnet_bch | Get testnet bch 
*MineApi* | [**Mine**](docs/MineApi.md#mine) | **Post** /mine | Mine regtest coins to a specified address
*UtilApi* | [**Convert**](docs/UtilApi.md#convert) | **Post** /util/convert | convert between units
*UtilApi* | [**GetAddrsByXpubKey**](docs/UtilApi.md#getaddrsbyxpubkey) | **Post** /util/get_addrs_by_xpubkey | Derive heristic determined addresses from an extended public key, per BIP32
*UtilApi* | [**GetXpubKeyInfo**](docs/UtilApi.md#getxpubkeyinfo) | **Post** /util/get_xpubkey_info | Decode information about an extended public key, per BIP32
*WalletApi* | [**Balance**](docs/WalletApi.md#balance) | **Post** /wallet/balance | Get total balance for wallet
*WalletApi* | [**CreateWallet**](docs/WalletApi.md#createwallet) | **Post** /wallet/create | create a new wallet
*WalletApi* | [**DepositAddress**](docs/WalletApi.md#depositaddress) | **Post** /wallet/deposit_address | Get a deposit address in cash address format
*WalletApi* | [**DepositQr**](docs/WalletApi.md#depositqr) | **Post** /wallet/deposit_qr | Get receiving cash address as a qrcode
*WalletApi* | [**EncodeTransaction**](docs/WalletApi.md#encodetransaction) | **Post** /wallet/encode_transaction | Encode and sign a transaction given a list of sendRequests, options and estimate fees
*WalletApi* | [**GetAllNftTokenBalances**](docs/WalletApi.md#getallnfttokenbalances) | **Post** /wallet/get_all_nft_token_balances | Get non-fungible token balance
*WalletApi* | [**GetAllTokenBalances**](docs/WalletApi.md#getalltokenbalances) | **Post** /wallet/get_all_token_balances | Get non-fungible token balance
*WalletApi* | [**GetHistory**](docs/WalletApi.md#gethistory) | **Post** /wallet/get_history | Get a simplified list of transactions related to a wallet
*WalletApi* | [**GetNftTokenBalance**](docs/WalletApi.md#getnfttokenbalance) | **Post** /wallet/get_nft_token_balance | Get non-fungible token balance
*WalletApi* | [**GetTokenBalance**](docs/WalletApi.md#gettokenbalance) | **Post** /wallet/get_token_balance | Get fungible token balance
*WalletApi* | [**GetTokenUtxos**](docs/WalletApi.md#gettokenutxos) | **Post** /wallet/get_token_utxos | Get token utxos
*WalletApi* | [**Info**](docs/WalletApi.md#info) | **Post** /wallet/info | Get information about a wallet
*WalletApi* | [**MaxAmountToSend**](docs/WalletApi.md#maxamounttosend) | **Post** /wallet/max_amount_to_send | Get maximum spendable amount
*WalletApi* | [**NamedExists**](docs/WalletApi.md#namedexists) | **Post** /wallet/named_exists | Check if a named wallet already exists
*WalletApi* | [**ReplaceNamed**](docs/WalletApi.md#replacenamed) | **Post** /wallet/replace_named | Replace (recover) named wallet with a new walletId. If wallet with a provided name does not exist yet, it will be creted with a &#x60;walletId&#x60; supplied If wallet exists it will be overwritten without exception 
*WalletApi* | [**Send**](docs/WalletApi.md#send) | **Post** /wallet/send | Send some amount to a given address
*WalletApi* | [**SendMax**](docs/WalletApi.md#sendmax) | **Post** /wallet/send_max | Send all available funds to a given address
*WalletApi* | [**SubmitTransaction**](docs/WalletApi.md#submittransaction) | **Post** /wallet/submit_transaction | submit an encoded and signed transaction to the network
*WalletApi* | [**TokenBurn**](docs/WalletApi.md#tokenburn) | **Post** /wallet/token_burn | Perform an explicit token burn
*WalletApi* | [**TokenDepositAddress**](docs/WalletApi.md#tokendepositaddress) | **Post** /wallet/token_deposit_address | Get a token aware deposit address in cash address format
*WalletApi* | [**TokenDepositQr**](docs/WalletApi.md#tokendepositqr) | **Post** /wallet/token_deposit_qr | Get receiving token aware cash address as a qrcode
*WalletApi* | [**TokenGenesis**](docs/WalletApi.md#tokengenesis) | **Post** /wallet/token_genesis | Create new token category
*WalletApi* | [**TokenMint**](docs/WalletApi.md#tokenmint) | **Post** /wallet/token_mint | Mint new non-fungible tokens
*WalletApi* | [**Utxos**](docs/WalletApi.md#utxos) | **Post** /wallet/utxo | Get detailed information about unspent outputs (utxos)
*WalletApi* | [**Xpubkeys**](docs/WalletApi.md#xpubkeys) | **Post** /wallet/xpubkeys | A set of xpubkeys and paths for the wallet.
*WalletBcmrApi* | [**BcmrAddMetadataRegistryAuthChain**](docs/WalletBcmrApi.md#bcmraddmetadataregistryauthchain) | **Post** /wallet/bcmr/add_registry_authchain | Add BCMR metadata registry from autchain, returning the built chain
*WalletBcmrApi* | [**BcmrAddRegistry**](docs/WalletBcmrApi.md#bcmraddregistry) | **Post** /wallet/bcmr/add_registry | Add BCMR registry from parameter
*WalletBcmrApi* | [**BcmrAddRegistryFromUri**](docs/WalletBcmrApi.md#bcmraddregistryfromuri) | **Post** /wallet/bcmr/add_registry_from_uri | Reset tracked BCMR registries
*WalletBcmrApi* | [**BcmrBuildAuthChain**](docs/WalletBcmrApi.md#bcmrbuildauthchain) | **Post** /wallet/bcmr/build_authchain | Build a BCMR authchain
*WalletBcmrApi* | [**BcmrGetRegistries**](docs/WalletBcmrApi.md#bcmrgetregistries) | **Post** /wallet/bcmr/get_registries | Get tracked BCMR registries
*WalletBcmrApi* | [**BcmrGetTokenInfo**](docs/WalletBcmrApi.md#bcmrgettokeninfo) | **Post** /wallet/bcmr/get_token_info | Get token info
*WalletBcmrApi* | [**BcmrResetRegistries**](docs/WalletBcmrApi.md#bcmrresetregistries) | **Post** /wallet/bcmr/reset_registries | Reset tracked BCMR registries
*WalletSignedApi* | [**SignedMessageSign**](docs/WalletSignedApi.md#signedmessagesign) | **Post** /wallet/signed/sign | Returns the message signature
*WalletSignedApi* | [**SignedMessageVerify**](docs/WalletSignedApi.md#signedmessageverify) | **Post** /wallet/signed/verify | Returns a boolean indicating whether message was valid for a given address
*WalletUtilApi* | [**UtilDecodeTransaction**](docs/WalletUtilApi.md#utildecodetransaction) | **Post** /wallet/util/decode_transaction | Decode a bitcoin transaction. Accepts both transaction hash or raw transaction in hex format.
*WalletUtilApi* | [**UtilGetRawTransaction**](docs/WalletUtilApi.md#utilgetrawtransaction) | **Post** /wallet/util/get_raw_transaction | Get bitcoin raw transaction
*WebhookApi* | [**WatchAddress**](docs/WebhookApi.md#watchaddress) | **Post** /webhook/watch_address | Create a webhook to watch cashaddress balance and transactions. 


## Documentation For Models

 - [AuthChainElement](docs/AuthChainElement.md)
 - [BalanceRequest](docs/BalanceRequest.md)
 - [BalanceResponse](docs/BalanceResponse.md)
 - [CashscriptReceipt](docs/CashscriptReceipt.md)
 - [Contract](docs/Contract.md)
 - [ContractFnRequest](docs/ContractFnRequest.md)
 - [ContractFnResponse](docs/ContractFnResponse.md)
 - [ContractInfoRequest](docs/ContractInfoRequest.md)
 - [ContractInfoResponse](docs/ContractInfoResponse.md)
 - [ContractRequest](docs/ContractRequest.md)
 - [ContractResponse](docs/ContractResponse.md)
 - [ConvertRequest](docs/ConvertRequest.md)
 - [CreateSignedMessageRequest](docs/CreateSignedMessageRequest.md)
 - [DepositAddressResponse](docs/DepositAddressResponse.md)
 - [ElectrumRawTransaction](docs/ElectrumRawTransaction.md)
 - [ElectrumRawTransactionScriptPubKey](docs/ElectrumRawTransactionScriptPubKey.md)
 - [ElectrumRawTransactionScriptSig](docs/ElectrumRawTransactionScriptSig.md)
 - [ElectrumRawTransactionVin](docs/ElectrumRawTransactionVin.md)
 - [ElectrumRawTransactionVout](docs/ElectrumRawTransactionVout.md)
 - [EncodeTransactionRequest](docs/EncodeTransactionRequest.md)
 - [EncodeTransactionResponse](docs/EncodeTransactionResponse.md)
 - [Error](docs/Error.md)
 - [EscrowContract](docs/EscrowContract.md)
 - [EscrowFnRequest](docs/EscrowFnRequest.md)
 - [EscrowInfoRequest](docs/EscrowInfoRequest.md)
 - [EscrowInfoResponse](docs/EscrowInfoResponse.md)
 - [EscrowRequest](docs/EscrowRequest.md)
 - [EscrowResponse](docs/EscrowResponse.md)
 - [GetAddressesResponse](docs/GetAddressesResponse.md)
 - [GetAddrsByXpubKeyRequest](docs/GetAddrsByXpubKeyRequest.md)
 - [GetTestnetBchRequest](docs/GetTestnetBchRequest.md)
 - [GetTestnetBchResponse](docs/GetTestnetBchResponse.md)
 - [GetXpubKeyInfoRequest](docs/GetXpubKeyInfoRequest.md)
 - [GetXpubKeyInfoResponse](docs/GetXpubKeyInfoResponse.md)
 - [HistoryRequest](docs/HistoryRequest.md)
 - [HistoryResponse](docs/HistoryResponse.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [MaxAmountToSendRequest](docs/MaxAmountToSendRequest.md)
 - [MineRequest](docs/MineRequest.md)
 - [NetworkEnum](docs/NetworkEnum.md)
 - [OpReturnData](docs/OpReturnData.md)
 - [ScalableVectorGraphic](docs/ScalableVectorGraphic.md)
 - [SendMaxRequest](docs/SendMaxRequest.md)
 - [SendMaxResponse](docs/SendMaxResponse.md)
 - [SendRequest](docs/SendRequest.md)
 - [SendRequestItem](docs/SendRequestItem.md)
 - [SendRequestItemAnyOf](docs/SendRequestItemAnyOf.md)
 - [SendRequestOptions](docs/SendRequestOptions.md)
 - [SendResponse](docs/SendResponse.md)
 - [SerializedWallet](docs/SerializedWallet.md)
 - [SignedMessageResponse](docs/SignedMessageResponse.md)
 - [SignedMessageResponseDetails](docs/SignedMessageResponseDetails.md)
 - [SignedMessageResponseRaw](docs/SignedMessageResponseRaw.md)
 - [SubmitTransactionRequest](docs/SubmitTransactionRequest.md)
 - [SubmitTransactionResponse](docs/SubmitTransactionResponse.md)
 - [TokenBurnRequest](docs/TokenBurnRequest.md)
 - [TokenGenesisRequest](docs/TokenGenesisRequest.md)
 - [TokenMintRequest](docs/TokenMintRequest.md)
 - [TokenMintRequestRequests](docs/TokenMintRequestRequests.md)
 - [TokenSendRequest](docs/TokenSendRequest.md)
 - [TransactionHistoryItem](docs/TransactionHistoryItem.md)
 - [UnitType](docs/UnitType.md)
 - [UtilDecodeTransactionRequest](docs/UtilDecodeTransactionRequest.md)
 - [Utxo](docs/Utxo.md)
 - [UtxoToken](docs/UtxoToken.md)
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
 - [XPubKey](docs/XPubKey.md)
 - [XPubKeyRequest](docs/XPubKeyRequest.md)
 - [XPubKeyResponse](docs/XPubKeyResponse.md)
 - [ZeroBalanceResponse](docs/ZeroBalanceResponse.md)


## Documentation For Authorization



## bearerAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author

hello@mainnet.cash

