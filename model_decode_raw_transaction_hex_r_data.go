/*
CryptoAPIs

Crypto APIs is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2021-03-20
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// DecodeRawTransactionHexRData struct for DecodeRawTransactionHexRData
type DecodeRawTransactionHexRData struct {
	Item DecodeRawTransactionHexRI `json:"item"`
}

// NewDecodeRawTransactionHexRData instantiates a new DecodeRawTransactionHexRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRData(item DecodeRawTransactionHexRI) *DecodeRawTransactionHexRData {
	this := DecodeRawTransactionHexRData{}
	this.Item = item
	return &this
}

// NewDecodeRawTransactionHexRDataWithDefaults instantiates a new DecodeRawTransactionHexRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRDataWithDefaults() *DecodeRawTransactionHexRData {
	this := DecodeRawTransactionHexRData{}
	return &this
}

// GetItem returns the Item field value
func (o *DecodeRawTransactionHexRData) GetItem() DecodeRawTransactionHexRI {
	if o == nil {
		var ret DecodeRawTransactionHexRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRData) GetItemOk() (*DecodeRawTransactionHexRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *DecodeRawTransactionHexRData) SetItem(v DecodeRawTransactionHexRI) {
	o.Item = v
}

func (o DecodeRawTransactionHexRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRData struct {
	value *DecodeRawTransactionHexRData
	isSet bool
}

func (v NullableDecodeRawTransactionHexRData) Get() *DecodeRawTransactionHexRData {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRData) Set(val *DecodeRawTransactionHexRData) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRData) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRData(val *DecodeRawTransactionHexRData) *NullableDecodeRawTransactionHexRData {
	return &NullableDecodeRawTransactionHexRData{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


