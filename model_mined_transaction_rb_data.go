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

// MinedTransactionRBData struct for MinedTransactionRBData
type MinedTransactionRBData struct {
	Item MinedTransactionRBDataItem `json:"item"`
}

// NewMinedTransactionRBData instantiates a new MinedTransactionRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinedTransactionRBData(item MinedTransactionRBDataItem) *MinedTransactionRBData {
	this := MinedTransactionRBData{}
	this.Item = item
	return &this
}

// NewMinedTransactionRBDataWithDefaults instantiates a new MinedTransactionRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinedTransactionRBDataWithDefaults() *MinedTransactionRBData {
	this := MinedTransactionRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *MinedTransactionRBData) GetItem() MinedTransactionRBDataItem {
	if o == nil {
		var ret MinedTransactionRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionRBData) GetItemOk() (*MinedTransactionRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *MinedTransactionRBData) SetItem(v MinedTransactionRBDataItem) {
	o.Item = v
}

func (o MinedTransactionRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableMinedTransactionRBData struct {
	value *MinedTransactionRBData
	isSet bool
}

func (v NullableMinedTransactionRBData) Get() *MinedTransactionRBData {
	return v.value
}

func (v *NullableMinedTransactionRBData) Set(val *MinedTransactionRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableMinedTransactionRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableMinedTransactionRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinedTransactionRBData(val *MinedTransactionRBData) *NullableMinedTransactionRBData {
	return &NullableMinedTransactionRBData{value: val, isSet: true}
}

func (v NullableMinedTransactionRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinedTransactionRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


