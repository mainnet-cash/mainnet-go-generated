/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.5.1
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpGenesisRequest struct for SlpGenesisRequest
type SlpGenesisRequest struct {
	// ID that is returned in `wallet` field of /wallet call 
	WalletId string `json:"walletId"`
	// Token name
	Name string `json:"name"`
	// Token ticker
	Ticker string `json:"ticker"`
	// Value is represented as a string to avoid precision loss
	InitialAmount string `json:"initialAmount"`
	// Indicates that 1 token is divisible into 10^decimals base units
	Decimals float32 `json:"decimals"`
	DocumentUrl string `json:"documentUrl,omitempty"`
	// Document hash of the token. Empty or 64 character long hex string.
	DocumentHash string `json:"documentHash,omitempty"`
	// Indicates if token should not be 'mintable', e.g. total circulation amount increased
	EndBaton bool `json:"endBaton,omitempty"`
	// Token type. 0x01 Type1, 0x81 NFT Parent, 0x41 NFT Child
	Type float32 `json:"type,omitempty"`
	TokenReceiverSlpAddr string `json:"tokenReceiverSlpAddr,omitempty"`
	BatonReceiverSlpAddr string `json:"batonReceiverSlpAddr,omitempty"`
	// Identifier of the NFT parent token
	ParentTokenId string `json:"parentTokenId,omitempty"`
}
