# EscrowInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EscrowContractId** | **string** | serialized escrow contract  | [optional] 
**Script** | **string** | The escrow contract in cashscript syntax  | [optional] 
**Parameters** | **[]string** | Parameters passed when the contract was created | [optional] 
**Cashaddr** | **string** |  | [optional] 
**BuyerAddr** | Pointer to **interface{}** | The cashaddress of the buyer | [optional] 
**ArbiterAddr** | Pointer to **interface{}** | The cashaddress of the arbiter | [optional] 
**SellerAddr** | Pointer to **interface{}** | The cashaddress of the seller | [optional] 
**Amount** | **float32** | Numeric amount to be transferred by the contract in satoshi, amount must be less than 21 BCH. | [optional] 
**Nonce** | **float32** | A unique number used once for each instance of a contract. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


