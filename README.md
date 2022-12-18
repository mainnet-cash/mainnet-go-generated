# Go API client for openapi

A developer friendly bitcoin cash wallet api

This API is currently in *active* development, breaking changes may
be made prior to official release of version 1.0.0.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.1
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
*FaucetApi* | [**GetTestnetSbch**](docs/FaucetApi.md#gettestnetsbch) | **Post** /faucet/get_testnet_sbch | Request testnet SmartBCH funds. The request is enqueued and served within 1-3 block confirmations. If the target address holds more that 0.1 BCH, the request will fail. Otherwise the address will be topped up to 0.1 BCH. 
*FaucetApi* | [**GetTestnetSep20**](docs/FaucetApi.md#gettestnetsep20) | **Post** /faucet/get_testnet_sep20 | Request testnet SmartBch SEP20 tokens. The request is enqueued and served within 1-3 block confirmations. If the target address holds more that 10 tokens of requested kind, the request will fail. Otherwise the address will be topped up to 10 tokens. 
*FaucetApi* | [**GetTestnetSlp**](docs/FaucetApi.md#gettestnetslp) | **Post** /faucet/get_testnet_slp | Get testnet slp tokens 
*MineApi* | [**Mine**](docs/MineApi.md#mine) | **Post** /mine | Mine regtest coins to a specified address
*SmartbchContractApi* | [**SmartBchContractCall**](docs/SmartbchContractApi.md#smartbchcontractcall) | **Post** /smartbch/contract/call | Call a SmartBch contract function
*SmartbchContractApi* | [**SmartBchContractCreate**](docs/SmartbchContractApi.md#smartbchcontractcreate) | **Post** /smartbch/contract/create | Create a SmartBch contractId
*SmartbchContractApi* | [**SmartBchContractDeploy**](docs/SmartbchContractApi.md#smartbchcontractdeploy) | **Post** /smartbch/contract/deploy | Request to deploy a SmartBch contract
*SmartbchContractApi* | [**SmartBchContractEstimateGas**](docs/SmartbchContractApi.md#smartbchcontractestimategas) | **Post** /smartbch/contract/estimate_gas | Estimate the gas for a contract interaction function given the arguments
*SmartbchContractApi* | [**SmartBchContractInfo**](docs/SmartbchContractApi.md#smartbchcontractinfo) | **Post** /smartbch/contract/info | Get information about a SmartBch contract from the contractId
*SmartbchSep20Api* | [**SmartBchSep20AllBalances**](docs/SmartbchSep20Api.md#smartbchsep20allbalances) | **Post** /smartbch/sep20/all_balances | Get all SmartBch SEP20 balances of the wallet
*SmartbchSep20Api* | [**SmartBchSep20Balance**](docs/SmartbchSep20Api.md#smartbchsep20balance) | **Post** /smartbch/sep20/balance | Get total SmartBch SEP20 token balance of the wallet
*SmartbchSep20Api* | [**SmartBchSep20DepositAddress**](docs/SmartbchSep20Api.md#smartbchsep20depositaddress) | **Post** /smartbch/sep20/deposit_address | Get an SmartBch SEP20 deposit address
*SmartbchSep20Api* | [**SmartBchSep20DepositQr**](docs/SmartbchSep20Api.md#smartbchsep20depositqr) | **Post** /smartbch/sep20/deposit_qr | Get an SmartBch SEP20 receiving address as a qrcode
*SmartbchSep20Api* | [**SmartBchSep20Genesis**](docs/SmartbchSep20Api.md#smartbchsep20genesis) | **Post** /smartbch/sep20/genesis | Get created tokenId back and new SmartBch SEP20 token balance of the wallet
*SmartbchSep20Api* | [**SmartBchSep20Mint**](docs/SmartbchSep20Api.md#smartbchsep20mint) | **Post** /smartbch/sep20/mint | Get created tokenId back and new SmartBch SEP20 token balance of the wallet
*SmartbchSep20Api* | [**SmartBchSep20Send**](docs/SmartbchSep20Api.md#smartbchsep20send) | **Post** /smartbch/sep20/send | Send some SmartBch SEP20 token amount to a given address
*SmartbchSep20Api* | [**SmartBchSep20SendMax**](docs/SmartbchSep20Api.md#smartbchsep20sendmax) | **Post** /smartbch/sep20/send_max | Send all available SmartBch SEP20 token funds to a given address
*SmartbchSep20Api* | [**SmartBchSep20TokenInfo**](docs/SmartbchSep20Api.md#smartbchsep20tokeninfo) | **Post** /smartbch/sep20/token_info | Get information about the SmartBch SEP20 token
*SmartbchWalletApi* | [**SmartbchBalance**](docs/SmartbchWalletApi.md#smartbchbalance) | **Post** /smartbch/wallet/balance | Get total balance for wallet
*SmartbchWalletApi* | [**SmartbchCreateWallet**](docs/SmartbchWalletApi.md#smartbchcreatewallet) | **Post** /smartbch/wallet/create | create a new wallet
*SmartbchWalletApi* | [**SmartbchDepositAddress**](docs/SmartbchWalletApi.md#smartbchdepositaddress) | **Post** /smartbch/wallet/deposit_address | Get a deposit address
*SmartbchWalletApi* | [**SmartbchDepositQr**](docs/SmartbchWalletApi.md#smartbchdepositqr) | **Post** /smartbch/wallet/deposit_qr | Get receiving cash address as a qrcode
*SmartbchWalletApi* | [**SmartbchMaxAmountToSend**](docs/SmartbchWalletApi.md#smartbchmaxamounttosend) | **Post** /smartbch/wallet/max_amount_to_send | Get maximum spendable amount
*SmartbchWalletApi* | [**SmartbchSend**](docs/SmartbchWalletApi.md#smartbchsend) | **Post** /smartbch/wallet/send | Send some amount to a given address
*SmartbchWalletApi* | [**SmartbchSendMax**](docs/SmartbchWalletApi.md#smartbchsendmax) | **Post** /smartbch/wallet/send_max | Send all available funds to a given address
*SmartbchWalletApi* | [**SmartbchSignedMessageSign**](docs/SmartbchWalletApi.md#smartbchsignedmessagesign) | **Post** /smartbch/wallet/signed/sign | Returns the message signature
*SmartbchWalletApi* | [**SmartbchSignedMessageVerify**](docs/SmartbchWalletApi.md#smartbchsignedmessageverify) | **Post** /smartbch/wallet/signed/verify | Returns a boolean indicating whether message was valid for a given address
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
 - [GetTestnetSbchRequest](docs/GetTestnetSbchRequest.md)
 - [GetTestnetSbchResponse](docs/GetTestnetSbchResponse.md)
 - [GetTestnetSep20Request](docs/GetTestnetSep20Request.md)
 - [GetTestnetSep20Response](docs/GetTestnetSep20Response.md)
 - [GetTestnetSlpRequest](docs/GetTestnetSlpRequest.md)
 - [GetTestnetSlpResponse](docs/GetTestnetSlpResponse.md)
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
 - [SlpTokenInfoResponse](docs/SlpTokenInfoResponse.md)
 - [SlpUtxo](docs/SlpUtxo.md)
 - [SlpUtxoResponse](docs/SlpUtxoResponse.md)
 - [SmartBchContractDeployRequest](docs/SmartBchContractDeployRequest.md)
 - [SmartBchContractDeployResponse](docs/SmartBchContractDeployResponse.md)
 - [SmartBchContractEstimateGasRequest](docs/SmartBchContractEstimateGasRequest.md)
 - [SmartBchContractEstimateGasResponse](docs/SmartBchContractEstimateGasResponse.md)
 - [SmartBchContractFnCallRequest](docs/SmartBchContractFnCallRequest.md)
 - [SmartBchContractFnCallResponse](docs/SmartBchContractFnCallResponse.md)
 - [SmartBchContractInfoRequest](docs/SmartBchContractInfoRequest.md)
 - [SmartBchContractInfoResponse](docs/SmartBchContractInfoResponse.md)
 - [SmartBchContractRequest](docs/SmartBchContractRequest.md)
 - [SmartBchContractResponse](docs/SmartBchContractResponse.md)
 - [SmartBchDepositAddressResponse](docs/SmartBchDepositAddressResponse.md)
 - [SmartBchMaxAmountToSendRequest](docs/SmartBchMaxAmountToSendRequest.md)
 - [SmartBchOverrides](docs/SmartBchOverrides.md)
 - [SmartBchSendMaxRequest](docs/SmartBchSendMaxRequest.md)
 - [SmartBchSendRequest](docs/SmartBchSendRequest.md)
 - [SmartBchSendRequestItem](docs/SmartBchSendRequestItem.md)
 - [SmartBchSendRequestItemAnyOf](docs/SmartBchSendRequestItemAnyOf.md)
 - [SmartBchSendRequestOptions](docs/SmartBchSendRequestOptions.md)
 - [SmartBchSendResponseItem](docs/SmartBchSendResponseItem.md)
 - [SmartBchSep20AllBalancesRequest](docs/SmartBchSep20AllBalancesRequest.md)
 - [SmartBchSep20BalanceRequest](docs/SmartBchSep20BalanceRequest.md)
 - [SmartBchSep20BalanceResponse](docs/SmartBchSep20BalanceResponse.md)
 - [SmartBchSep20GenesisRequest](docs/SmartBchSep20GenesisRequest.md)
 - [SmartBchSep20GenesisResponse](docs/SmartBchSep20GenesisResponse.md)
 - [SmartBchSep20MintRequest](docs/SmartBchSep20MintRequest.md)
 - [SmartBchSep20MintResponse](docs/SmartBchSep20MintResponse.md)
 - [SmartBchSep20SendMaxRequest](docs/SmartBchSep20SendMaxRequest.md)
 - [SmartBchSep20SendRequest](docs/SmartBchSep20SendRequest.md)
 - [SmartBchSep20SendRequestItem](docs/SmartBchSep20SendRequestItem.md)
 - [SmartBchSep20TokenInfoRequest](docs/SmartBchSep20TokenInfoRequest.md)
 - [SmartBchSep20TokenInfoResponse](docs/SmartBchSep20TokenInfoResponse.md)
 - [SmartBchTransactionReceipt](docs/SmartBchTransactionReceipt.md)
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

