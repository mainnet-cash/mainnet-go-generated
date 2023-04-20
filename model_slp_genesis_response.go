/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.4
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SlpGenesisResponse struct for SlpGenesisResponse
type SlpGenesisResponse struct {
	// Full token identifier freshly created
	TokenId string `json:"tokenId,omitempty"`
	Balance SlpBalanceResponse `json:"balance,omitempty"`
}
