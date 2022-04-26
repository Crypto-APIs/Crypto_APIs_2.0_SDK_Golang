/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ListAssetsDetailsRIAssetLogo Defines the logo of the asset.
type ListAssetsDetailsRIAssetLogo struct {
	// Defines the encoding of the image data which is usually base64.
	Encoding string `json:"encoding"`
	// Defines the encoded image data as a string.
	ImageData string `json:"imageData"`
	// Defines the image type of the logo - jpg, png, svg, etc.
	MimeType string `json:"mimeType"`
}

// NewListAssetsDetailsRIAssetLogo instantiates a new ListAssetsDetailsRIAssetLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAssetsDetailsRIAssetLogo(encoding string, imageData string, mimeType string) *ListAssetsDetailsRIAssetLogo {
	this := ListAssetsDetailsRIAssetLogo{}
	this.Encoding = encoding
	this.ImageData = imageData
	this.MimeType = mimeType
	return &this
}

// NewListAssetsDetailsRIAssetLogoWithDefaults instantiates a new ListAssetsDetailsRIAssetLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAssetsDetailsRIAssetLogoWithDefaults() *ListAssetsDetailsRIAssetLogo {
	this := ListAssetsDetailsRIAssetLogo{}
	return &this
}

// GetEncoding returns the Encoding field value
func (o *ListAssetsDetailsRIAssetLogo) GetEncoding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRIAssetLogo) GetEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value
func (o *ListAssetsDetailsRIAssetLogo) SetEncoding(v string) {
	o.Encoding = v
}

// GetImageData returns the ImageData field value
func (o *ListAssetsDetailsRIAssetLogo) GetImageData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageData
}

// GetImageDataOk returns a tuple with the ImageData field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRIAssetLogo) GetImageDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageData, true
}

// SetImageData sets field value
func (o *ListAssetsDetailsRIAssetLogo) SetImageData(v string) {
	o.ImageData = v
}

// GetMimeType returns the MimeType field value
func (o *ListAssetsDetailsRIAssetLogo) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRIAssetLogo) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *ListAssetsDetailsRIAssetLogo) SetMimeType(v string) {
	o.MimeType = v
}

func (o ListAssetsDetailsRIAssetLogo) MarshalJSON() ([]byte, error) {
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

type NullableListAssetsDetailsRIAssetLogo struct {
	value *ListAssetsDetailsRIAssetLogo
	isSet bool
}

func (v NullableListAssetsDetailsRIAssetLogo) Get() *ListAssetsDetailsRIAssetLogo {
	return v.value
}

func (v *NullableListAssetsDetailsRIAssetLogo) Set(val *ListAssetsDetailsRIAssetLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableListAssetsDetailsRIAssetLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableListAssetsDetailsRIAssetLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAssetsDetailsRIAssetLogo(val *ListAssetsDetailsRIAssetLogo) *NullableListAssetsDetailsRIAssetLogo {
	return &NullableListAssetsDetailsRIAssetLogo{value: val, isSet: true}
}

func (v NullableListAssetsDetailsRIAssetLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAssetsDetailsRIAssetLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


