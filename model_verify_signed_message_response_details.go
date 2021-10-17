/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 0.4.22
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VerifySignedMessageResponseDetails struct for VerifySignedMessageResponseDetails
type VerifySignedMessageResponseDetails struct {
	// The type of signature passed
	SignatureType string `json:"signatureType,omitempty"`
	// The double sha256 hash of the string encoded as Bitcoin Message binary.
	MessageHash string `json:"messageHash,omitempty"`
	// A boolean indicating whether the signature is valid for the message
	SignatureValid bool `json:"signatureValid,omitempty"`
	// A boolean indicating whether the publicKeyHash of a recoverable signature matches the provided cashaddr, false otherwise
	PublicKeyHashMatch bool `json:"publicKeyHashMatch,omitempty"`
}
