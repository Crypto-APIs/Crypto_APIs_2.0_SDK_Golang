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

// GenerateDepositAddressRBData struct for GenerateDepositAddressRBData
type GenerateDepositAddressRBData struct {
	Item GenerateDepositAddressRBDataItem `json:"item"`
}

// NewGenerateDepositAddressRBData instantiates a new GenerateDepositAddressRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDepositAddressRBData(item GenerateDepositAddressRBDataItem) *GenerateDepositAddressRBData {
	this := GenerateDepositAddressRBData{}
	this.Item = item
	return &this
}

// NewGenerateDepositAddressRBDataWithDefaults instantiates a new GenerateDepositAddressRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDepositAddressRBDataWithDefaults() *GenerateDepositAddressRBData {
	this := GenerateDepositAddressRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *GenerateDepositAddressRBData) GetItem() GenerateDepositAddressRBDataItem {
	if o == nil {
		var ret GenerateDepositAddressRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GenerateDepositAddressRBData) GetItemOk() (*GenerateDepositAddressRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GenerateDepositAddressRBData) SetItem(v GenerateDepositAddressRBDataItem) {
	o.Item = v
}

func (o GenerateDepositAddressRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateDepositAddressRBData struct {
	value *GenerateDepositAddressRBData
	isSet bool
}

func (v NullableGenerateDepositAddressRBData) Get() *GenerateDepositAddressRBData {
	return v.value
}

func (v *NullableGenerateDepositAddressRBData) Set(val *GenerateDepositAddressRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDepositAddressRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDepositAddressRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDepositAddressRBData(val *GenerateDepositAddressRBData) *NullableGenerateDepositAddressRBData {
	return &NullableGenerateDepositAddressRBData{value: val, isSet: true}
}

func (v NullableGenerateDepositAddressRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateDepositAddressRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


