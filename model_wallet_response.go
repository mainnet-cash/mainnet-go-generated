/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in active development, breaking changes may be made prior to official release of version 1.  **Important:** This library is in active development 
 *
 * API version: 0.0.1-rc
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WalletResponse struct for WalletResponse
type WalletResponse struct {
	// ID that is returned in `wallet` field of /wallet call 
	WalletId string `json:"walletId,omitempty"`
	// The wallet in Wallet Import Format (WIF) 
	Wif string `json:"wif,omitempty"`
	// User defined string for wallet
	Name string `json:"name,omitempty"`
	// The address in cashaddr format. 
	Cashaddr string `json:"cashaddr,omitempty"`
	// network type
	Network string `json:"network,omitempty"`
}
