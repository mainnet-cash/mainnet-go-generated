/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.13
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SmartBchSendRequestOptions struct for SmartBchSendRequestOptions
type SmartBchSendRequestOptions struct {
	// A boolean flag to include the wallet balance after the successful `send` operation to the returned result If set to false, the balance will not be queried and returned, making the `send` call faster
	QueryBalance bool `json:"queryBalance,omitempty"`
	// A boolean flag to wait for transaction to propagate through the network. If set to false, the `send` call will be faster, but other subsequent operations might fail due to nonce number mismatch.
	AwaitTransactionPropagation bool `json:"awaitTransactionPropagation,omitempty"`
}
