/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.30
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpSendRequest struct for SlpSendRequest
type SlpSendRequest struct {
	// ID that is returned in `wallet` field of /wallet call 
	WalletId string `json:"walletId"`
	To []SlpSendRequestItem `json:"to"`
}
