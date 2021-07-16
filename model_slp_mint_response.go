/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.28
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpMintResponse struct for SlpMintResponse
type SlpMintResponse struct {
	// Transaction id 
	TxId string `json:"txId,omitempty"`
	Balance SlpBalanceResponse `json:"balance,omitempty"`
}
