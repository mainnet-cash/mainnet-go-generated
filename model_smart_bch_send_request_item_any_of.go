/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.4
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSendRequestItemAnyOf Send request in object notation
type SmartBchSendRequestItemAnyOf struct {
	Address string `json:"address"`
	Value float32 `json:"value"`
}
