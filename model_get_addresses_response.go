/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.34-alpha.1
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetAddressesResponse struct for GetAddressesResponse
type GetAddressesResponse struct {
	// Cashaddr to return testnet BCH
	Bchtest string `json:"bchtest,omitempty"`
	// Slpaddr to return testnet SLP tokens
	Slptest string `json:"slptest,omitempty"`
	// SmartBch testnet faucet contract owner address
	Sbchtest string `json:"sbchtest,omitempty"`
	// SmartBch testnet contract address to return BCH and SEP20 tokens
	Sbchcontract string `json:"sbchcontract,omitempty"`
	// SmartBch testnet sample SEP20 token address
	Sbchtoken string `json:"sbchtoken,omitempty"`
}
