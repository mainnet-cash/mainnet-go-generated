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
// SignedMessageResponseRaw struct for SignedMessageResponseRaw
type SignedMessageResponseRaw struct {
	// Elliptic Curve Digital messageHash encoded as a base64 string
	Ecdsa string `json:"ecdsa,omitempty"`
	// Schnorr signature of the messageHash encoded as a base64 string, Note from libauth: Nonces are generated using RFC6979, where the Section 3.6, 16-byte ASCII \"additional data\" is set to Schnorr+SHA256, see https://libauth.org/interfaces/secp256k1.html#signmessagehashschnorr 
	Schnorr string `json:"schnorr,omitempty"`
	// Signature of messageHash using DER (Distinguished Encoding Rules) 
	Der string `json:"der,omitempty"`
}