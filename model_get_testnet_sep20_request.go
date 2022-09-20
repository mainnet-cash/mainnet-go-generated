/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.5.8
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetTestnetSep20Request struct for GetTestnetSep20Request
type GetTestnetSep20Request struct {
	// Receiver's SmartBch testnet address
	Address string `json:"address"`
	// Token unique hexadecimal identifier
	TokenId string `json:"tokenId"`
}