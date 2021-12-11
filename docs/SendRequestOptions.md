# SendRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UtxoIds** | **[]string** |  | [optional] 
**ChangeAddress** | **string** | Cash address to receive change to  | [optional] 
**SlpAware** | **bool** | If you have any SLP tokens on the address, you should set &#x60;slpAware&#x60; to &#x60;true&#x60; to prevent accidental token burning This flag activates utxo checking against an external slp indexer So far this option is not switched on by default | [optional] [default to false]
**SlpSemiAware** | **bool** | This flag requires an utxo to have more than 546 sats to be spendable and counted in the balance This option is not switched on by default | [optional] [default to false]
**QueryBalance** | **bool** | A boolean flag to include the wallet balance after the successful &#x60;send&#x60; operation to the returned result If set to false, the balance will not be queried and returned, making the &#x60;send&#x60; call faster | [optional] [default to true]
**AwaitTransactionPropagation** | **bool** | A boolean flag to wait for transaction to propagate through the network and be registered in the bitcoind and indexer. If set to false, the &#x60;send&#x60; call will be very fast, but the wallet UTXO state might be invalid for some 500ms. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


