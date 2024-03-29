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

// ConvertBitcoinCashAddressRBData struct for ConvertBitcoinCashAddressRBData
type ConvertBitcoinCashAddressRBData struct {
	Item ConvertBitcoinCashAddressRBDataItem `json:"item"`
}

// NewConvertBitcoinCashAddressRBData instantiates a new ConvertBitcoinCashAddressRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertBitcoinCashAddressRBData(item ConvertBitcoinCashAddressRBDataItem) *ConvertBitcoinCashAddressRBData {
	this := ConvertBitcoinCashAddressRBData{}
	this.Item = item
	return &this
}

// NewConvertBitcoinCashAddressRBDataWithDefaults instantiates a new ConvertBitcoinCashAddressRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertBitcoinCashAddressRBDataWithDefaults() *ConvertBitcoinCashAddressRBData {
	this := ConvertBitcoinCashAddressRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *ConvertBitcoinCashAddressRBData) GetItem() ConvertBitcoinCashAddressRBDataItem {
	if o == nil {
		var ret ConvertBitcoinCashAddressRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *ConvertBitcoinCashAddressRBData) GetItemOk() (*ConvertBitcoinCashAddressRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *ConvertBitcoinCashAddressRBData) SetItem(v ConvertBitcoinCashAddressRBDataItem) {
	o.Item = v
}

func (o ConvertBitcoinCashAddressRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableConvertBitcoinCashAddressRBData struct {
	value *ConvertBitcoinCashAddressRBData
	isSet bool
}

func (v NullableConvertBitcoinCashAddressRBData) Get() *ConvertBitcoinCashAddressRBData {
	return v.value
}

func (v *NullableConvertBitcoinCashAddressRBData) Set(val *ConvertBitcoinCashAddressRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertBitcoinCashAddressRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertBitcoinCashAddressRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertBitcoinCashAddressRBData(val *ConvertBitcoinCashAddressRBData) *NullableConvertBitcoinCashAddressRBData {
	return &NullableConvertBitcoinCashAddressRBData{value: val, isSet: true}
}

func (v NullableConvertBitcoinCashAddressRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertBitcoinCashAddressRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


