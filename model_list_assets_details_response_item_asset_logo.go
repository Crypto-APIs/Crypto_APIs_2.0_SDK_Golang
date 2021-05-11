/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ListAssetsDetailsResponseItemAssetLogo Defines the logo of the asset.
type ListAssetsDetailsResponseItemAssetLogo struct {
	// Defines the encoding of the image data which is usually base64.
	Encoding string `json:"encoding"`
	// Defines the encoded image data as a string.
	ImageData string `json:"imageData"`
	// Defines the image type of the logo - jpg, png, svg, etc.
	MimeType string `json:"mimeType"`
}

// NewListAssetsDetailsResponseItemAssetLogo instantiates a new ListAssetsDetailsResponseItemAssetLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAssetsDetailsResponseItemAssetLogo(encoding string, imageData string, mimeType string) *ListAssetsDetailsResponseItemAssetLogo {
	this := ListAssetsDetailsResponseItemAssetLogo{}
	this.Encoding = encoding
	this.ImageData = imageData
	this.MimeType = mimeType
	return &this
}

// NewListAssetsDetailsResponseItemAssetLogoWithDefaults instantiates a new ListAssetsDetailsResponseItemAssetLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAssetsDetailsResponseItemAssetLogoWithDefaults() *ListAssetsDetailsResponseItemAssetLogo {
	this := ListAssetsDetailsResponseItemAssetLogo{}
	return &this
}

// GetEncoding returns the Encoding field value
func (o *ListAssetsDetailsResponseItemAssetLogo) GetEncoding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsResponseItemAssetLogo) GetEncodingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value
func (o *ListAssetsDetailsResponseItemAssetLogo) SetEncoding(v string) {
	o.Encoding = v
}

// GetImageData returns the ImageData field value
func (o *ListAssetsDetailsResponseItemAssetLogo) GetImageData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageData
}

// GetImageDataOk returns a tuple with the ImageData field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsResponseItemAssetLogo) GetImageDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageData, true
}

// SetImageData sets field value
func (o *ListAssetsDetailsResponseItemAssetLogo) SetImageData(v string) {
	o.ImageData = v
}

// GetMimeType returns the MimeType field value
func (o *ListAssetsDetailsResponseItemAssetLogo) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsResponseItemAssetLogo) GetMimeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *ListAssetsDetailsResponseItemAssetLogo) SetMimeType(v string) {
	o.MimeType = v
}

func (o ListAssetsDetailsResponseItemAssetLogo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["encoding"] = o.Encoding
	}
	if true {
		toSerialize["imageData"] = o.ImageData
	}
	if true {
		toSerialize["mimeType"] = o.MimeType
	}
	return json.Marshal(toSerialize)
}

type NullableListAssetsDetailsResponseItemAssetLogo struct {
	value *ListAssetsDetailsResponseItemAssetLogo
	isSet bool
}

func (v NullableListAssetsDetailsResponseItemAssetLogo) Get() *ListAssetsDetailsResponseItemAssetLogo {
	return v.value
}

func (v *NullableListAssetsDetailsResponseItemAssetLogo) Set(val *ListAssetsDetailsResponseItemAssetLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableListAssetsDetailsResponseItemAssetLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableListAssetsDetailsResponseItemAssetLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAssetsDetailsResponseItemAssetLogo(val *ListAssetsDetailsResponseItemAssetLogo) *NullableListAssetsDetailsResponseItemAssetLogo {
	return &NullableListAssetsDetailsResponseItemAssetLogo{value: val, isSet: true}
}

func (v NullableListAssetsDetailsResponseItemAssetLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAssetsDetailsResponseItemAssetLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


