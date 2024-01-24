/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.3.3
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScalableVectorGraphic struct for ScalableVectorGraphic
type ScalableVectorGraphic struct {
	// A Qr code image data in svg format as utf-8 encoded string. Suitable for inclusion in html using:     - \\<img src\\=\\\"**data:image/svg+xml;base64,PD94bWwgdm... ==**\"\\> 
	Src string `json:"src,omitempty"`
	// hover text for graphic
	Title string `json:"title,omitempty"`
	// assistive text
	Alt string `json:"alt,omitempty"`
}
