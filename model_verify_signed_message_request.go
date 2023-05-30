/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 1.1.27
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VerifySignedMessageRequest struct for VerifySignedMessageRequest
type VerifySignedMessageRequest struct {
	WalletId string `json:"walletId,omitempty"`
	Message string `json:"message,omitempty"`
	// The base64 signature of the double sha265 hash of a bitcoin message formatted string signed using the private key associated with the related cashaddr
	Signature *interface{} `json:"signature,omitempty"`
	// If the publicKey is not recoverable from the signature, the base64 encoded public key may be passed as instead of the walletId
	PublicKey string `json:"publicKey,omitempty"`
}
