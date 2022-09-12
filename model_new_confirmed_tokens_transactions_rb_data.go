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

// NewConfirmedTokensTransactionsRBData struct for NewConfirmedTokensTransactionsRBData
type NewConfirmedTokensTransactionsRBData struct {
	Item NewConfirmedTokensTransactionsRBDataItem `json:"item"`
}

// NewNewConfirmedTokensTransactionsRBData instantiates a new NewConfirmedTokensTransactionsRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedTokensTransactionsRBData(item NewConfirmedTokensTransactionsRBDataItem) *NewConfirmedTokensTransactionsRBData {
	this := NewConfirmedTokensTransactionsRBData{}
	this.Item = item
	return &this
}

// NewNewConfirmedTokensTransactionsRBDataWithDefaults instantiates a new NewConfirmedTokensTransactionsRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedTokensTransactionsRBDataWithDefaults() *NewConfirmedTokensTransactionsRBData {
	this := NewConfirmedTokensTransactionsRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewConfirmedTokensTransactionsRBData) GetItem() NewConfirmedTokensTransactionsRBDataItem {
	if o == nil {
		var ret NewConfirmedTokensTransactionsRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokensTransactionsRBData) GetItemOk() (*NewConfirmedTokensTransactionsRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewConfirmedTokensTransactionsRBData) SetItem(v NewConfirmedTokensTransactionsRBDataItem) {
	o.Item = v
}

func (o NewConfirmedTokensTransactionsRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedTokensTransactionsRBData struct {
	value *NewConfirmedTokensTransactionsRBData
	isSet bool
}

func (v NullableNewConfirmedTokensTransactionsRBData) Get() *NewConfirmedTokensTransactionsRBData {
	return v.value
}

func (v *NullableNewConfirmedTokensTransactionsRBData) Set(val *NewConfirmedTokensTransactionsRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedTokensTransactionsRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedTokensTransactionsRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedTokensTransactionsRBData(val *NewConfirmedTokensTransactionsRBData) *NullableNewConfirmedTokensTransactionsRBData {
	return &NullableNewConfirmedTokensTransactionsRBData{value: val, isSet: true}
}

func (v NullableNewConfirmedTokensTransactionsRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedTokensTransactionsRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


