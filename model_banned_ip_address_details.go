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

// BannedIpAddressDetails struct for BannedIpAddressDetails
type BannedIpAddressDetails struct {
	// Specifies an attribute of the error by name.
	Attribute string `json:"attribute"`
	// Specifies the details of an attribute as part from the error.
	Message string `json:"message"`
}

// NewBannedIpAddressDetails instantiates a new BannedIpAddressDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBannedIpAddressDetails(attribute string, message string) *BannedIpAddressDetails {
	this := BannedIpAddressDetails{}
	this.Attribute = attribute
	this.Message = message
	return &this
}

// NewBannedIpAddressDetailsWithDefaults instantiates a new BannedIpAddressDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBannedIpAddressDetailsWithDefaults() *BannedIpAddressDetails {
	this := BannedIpAddressDetails{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *BannedIpAddressDetails) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *BannedIpAddressDetails) GetAttributeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *BannedIpAddressDetails) SetAttribute(v string) {
	o.Attribute = v
}

// GetMessage returns the Message field value
func (o *BannedIpAddressDetails) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BannedIpAddressDetails) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BannedIpAddressDetails) SetMessage(v string) {
	o.Message = v
}

func (o BannedIpAddressDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attribute"] = o.Attribute
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableBannedIpAddressDetails struct {
	value *BannedIpAddressDetails
	isSet bool
}

func (v NullableBannedIpAddressDetails) Get() *BannedIpAddressDetails {
	return v.value
}

func (v *NullableBannedIpAddressDetails) Set(val *BannedIpAddressDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBannedIpAddressDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBannedIpAddressDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBannedIpAddressDetails(val *BannedIpAddressDetails) *NullableBannedIpAddressDetails {
	return &NullableBannedIpAddressDetails{value: val, isSet: true}
}

func (v NullableBannedIpAddressDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBannedIpAddressDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

