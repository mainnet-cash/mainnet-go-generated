/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.34
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MaxAmountToSendRequest struct for MaxAmountToSendRequest
type MaxAmountToSendRequest struct {
	// ID that is returned in `wallet` field of /wallet call 
	WalletId string `json:"walletId"`
	// (optional) if sending all funds to multiple addresses, the count of the number of address funds will be sent to may be included.             
	OutputCount int32 `json:"output_count,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}
