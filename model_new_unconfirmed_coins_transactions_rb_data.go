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

// NewUnconfirmedCoinsTransactionsRBData struct for NewUnconfirmedCoinsTransactionsRBData
type NewUnconfirmedCoinsTransactionsRBData struct {
	Item NewUnconfirmedCoinsTransactionsRBDataItem `json:"item"`
}

// NewNewUnconfirmedCoinsTransactionsRBData instantiates a new NewUnconfirmedCoinsTransactionsRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUnconfirmedCoinsTransactionsRBData(item NewUnconfirmedCoinsTransactionsRBDataItem) *NewUnconfirmedCoinsTransactionsRBData {
	this := NewUnconfirmedCoinsTransactionsRBData{}
	this.Item = item
	return &this
}

// NewNewUnconfirmedCoinsTransactionsRBDataWithDefaults instantiates a new NewUnconfirmedCoinsTransactionsRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUnconfirmedCoinsTransactionsRBDataWithDefaults() *NewUnconfirmedCoinsTransactionsRBData {
	this := NewUnconfirmedCoinsTransactionsRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewUnconfirmedCoinsTransactionsRBData) GetItem() NewUnconfirmedCoinsTransactionsRBDataItem {
	if o == nil {
		var ret NewUnconfirmedCoinsTransactionsRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedCoinsTransactionsRBData) GetItemOk() (*NewUnconfirmedCoinsTransactionsRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewUnconfirmedCoinsTransactionsRBData) SetItem(v NewUnconfirmedCoinsTransactionsRBDataItem) {
	o.Item = v
}

func (o NewUnconfirmedCoinsTransactionsRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewUnconfirmedCoinsTransactionsRBData struct {
	value *NewUnconfirmedCoinsTransactionsRBData
	isSet bool
}

func (v NullableNewUnconfirmedCoinsTransactionsRBData) Get() *NewUnconfirmedCoinsTransactionsRBData {
	return v.value
}

func (v *NullableNewUnconfirmedCoinsTransactionsRBData) Set(val *NewUnconfirmedCoinsTransactionsRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUnconfirmedCoinsTransactionsRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUnconfirmedCoinsTransactionsRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUnconfirmedCoinsTransactionsRBData(val *NewUnconfirmedCoinsTransactionsRBData) *NullableNewUnconfirmedCoinsTransactionsRBData {
	return &NullableNewUnconfirmedCoinsTransactionsRBData{value: val, isSet: true}
}

func (v NullableNewUnconfirmedCoinsTransactionsRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUnconfirmedCoinsTransactionsRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


