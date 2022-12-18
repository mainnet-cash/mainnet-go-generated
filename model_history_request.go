/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.5.9
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// HistoryRequest struct for HistoryRequest
type HistoryRequest struct {
	// The walletId to get a history for. 
	WalletId string `json:"walletId"`
	// Unit of account for running balance.
	Unit string `json:"unit,omitempty"`
	// The first record to return starting from zero
	Start float32 `json:"start,omitempty"`
	// The number of records to return
	Count float32 `json:"count,omitempty"`
	// Exclude transactions returned to the wallet as change in response.
	CollapseChange bool `json:"collapseChange,omitempty"`
}
