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

// EncodeXAddressRI struct for EncodeXAddressRI
type EncodeXAddressRI struct {
	// Represents the encoded classic address with its destination tag.
	XAddress string `json:"xAddress"`
}

// NewEncodeXAddressRI instantiates a new EncodeXAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncodeXAddressRI(xAddress string) *EncodeXAddressRI {
	this := EncodeXAddressRI{}
	this.XAddress = xAddress
	return &this
}

// NewEncodeXAddressRIWithDefaults instantiates a new EncodeXAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncodeXAddressRIWithDefaults() *EncodeXAddressRI {
	this := EncodeXAddressRI{}
	return &this
}

// GetXAddress returns the XAddress field value
func (o *EncodeXAddressRI) GetXAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XAddress
}

// GetXAddressOk returns a tuple with the XAddress field value
// and a boolean to check if the value has been set.
func (o *EncodeXAddressRI) GetXAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XAddress, true
}

// SetXAddress sets field value
func (o *EncodeXAddressRI) SetXAddress(v string) {
	o.XAddress = v
}

func (o EncodeXAddressRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["xAddress"] = o.XAddress
	}
	return json.Marshal(toSerialize)
}

type NullableEncodeXAddressRI struct {
	value *EncodeXAddressRI
	isSet bool
}

func (v NullableEncodeXAddressRI) Get() *EncodeXAddressRI {
	return v.value
}

func (v *NullableEncodeXAddressRI) Set(val *EncodeXAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodeXAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodeXAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodeXAddressRI(val *EncodeXAddressRI) *NullableEncodeXAddressRI {
	return &NullableEncodeXAddressRI{value: val, isSet: true}
}

func (v NullableEncodeXAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodeXAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

