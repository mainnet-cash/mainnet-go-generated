/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.3.6
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SendMaxResponse struct for SendMaxResponse
type SendMaxResponse struct {
	// The hash of a transaction
	TxId string `json:"txId,omitempty"`
	Balance ZeroBalanceResponse `json:"balance,omitempty"`
}
