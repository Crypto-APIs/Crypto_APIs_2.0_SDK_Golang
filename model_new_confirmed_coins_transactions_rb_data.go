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

// NewConfirmedCoinsTransactionsRBData struct for NewConfirmedCoinsTransactionsRBData
type NewConfirmedCoinsTransactionsRBData struct {
	Item NewConfirmedCoinsTransactionsRBDataItem `json:"item"`
}

// NewNewConfirmedCoinsTransactionsRBData instantiates a new NewConfirmedCoinsTransactionsRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedCoinsTransactionsRBData(item NewConfirmedCoinsTransactionsRBDataItem) *NewConfirmedCoinsTransactionsRBData {
	this := NewConfirmedCoinsTransactionsRBData{}
	this.Item = item
	return &this
}

// NewNewConfirmedCoinsTransactionsRBDataWithDefaults instantiates a new NewConfirmedCoinsTransactionsRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedCoinsTransactionsRBDataWithDefaults() *NewConfirmedCoinsTransactionsRBData {
	this := NewConfirmedCoinsTransactionsRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewConfirmedCoinsTransactionsRBData) GetItem() NewConfirmedCoinsTransactionsRBDataItem {
	if o == nil {
		var ret NewConfirmedCoinsTransactionsRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsRBData) GetItemOk() (*NewConfirmedCoinsTransactionsRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewConfirmedCoinsTransactionsRBData) SetItem(v NewConfirmedCoinsTransactionsRBDataItem) {
	o.Item = v
}

func (o NewConfirmedCoinsTransactionsRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedCoinsTransactionsRBData struct {
	value *NewConfirmedCoinsTransactionsRBData
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsRBData) Get() *NewConfirmedCoinsTransactionsRBData {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsRBData) Set(val *NewConfirmedCoinsTransactionsRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsRBData(val *NewConfirmedCoinsTransactionsRBData) *NullableNewConfirmedCoinsTransactionsRBData {
	return &NullableNewConfirmedCoinsTransactionsRBData{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


