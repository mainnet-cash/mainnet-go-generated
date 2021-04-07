/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.3.12
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SignedMessageResponseDetails struct for SignedMessageResponseDetails
type SignedMessageResponseDetails struct {
	// A tag used for recovering the public key from a compact signature.
	RecoveryId float32 `json:"recoveryId,omitempty"`
	Compressed bool `json:"compressed,omitempty"`
	// The double sha256 hash of the string encoded as Bitcoin Message binary.
	MessageHash string `json:"messageHash,omitempty"`
}