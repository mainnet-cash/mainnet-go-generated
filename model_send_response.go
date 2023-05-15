/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.20
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SendResponse struct for SendResponse
type SendResponse struct {
	// The hash of a transaction
	TxId string `json:"txId,omitempty"`
	Balance BalanceResponse `json:"balance,omitempty"`
	// Web url to a transaction on a block explorer
	ExplorerUrl string `json:"explorerUrl,omitempty"`
	TokenIds []string `json:"tokenIds,omitempty"`
}
