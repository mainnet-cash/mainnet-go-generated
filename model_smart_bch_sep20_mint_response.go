/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.0.8
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSep20MintResponse struct for SmartBchSep20MintResponse
type SmartBchSep20MintResponse struct {
	// The hash of a transaction
	TxId string `json:"txId,omitempty"`
	Balance SmartBchSep20BalanceResponse `json:"balance,omitempty"`
}
