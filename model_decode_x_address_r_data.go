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

// DecodeXAddressRData struct for DecodeXAddressRData
type DecodeXAddressRData struct {
	Item DecodeXAddressRI `json:"item"`
}

// NewDecodeXAddressRData instantiates a new DecodeXAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeXAddressRData(item DecodeXAddressRI) *DecodeXAddressRData {
	this := DecodeXAddressRData{}
	this.Item = item
	return &this
}

// NewDecodeXAddressRDataWithDefaults instantiates a new DecodeXAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeXAddressRDataWithDefaults() *DecodeXAddressRData {
	this := DecodeXAddressRData{}
	return &this
}

// GetItem returns the Item field value
func (o *DecodeXAddressRData) GetItem() DecodeXAddressRI {
	if o == nil {
		var ret DecodeXAddressRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *DecodeXAddressRData) GetItemOk() (*DecodeXAddressRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *DecodeXAddressRData) SetItem(v DecodeXAddressRI) {
	o.Item = v
}

func (o DecodeXAddressRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeXAddressRData struct {
	value *DecodeXAddressRData
	isSet bool
}

func (v NullableDecodeXAddressRData) Get() *DecodeXAddressRData {
	return v.value
}

func (v *NullableDecodeXAddressRData) Set(val *DecodeXAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeXAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeXAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeXAddressRData(val *DecodeXAddressRData) *NullableDecodeXAddressRData {
	return &NullableDecodeXAddressRData{value: val, isSet: true}
}

func (v NullableDecodeXAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeXAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


