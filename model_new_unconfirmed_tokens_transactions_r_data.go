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

// NewUnconfirmedTokensTransactionsRData struct for NewUnconfirmedTokensTransactionsRData
type NewUnconfirmedTokensTransactionsRData struct {
	Item NewUnconfirmedTokensTransactionsRI `json:"item"`
}

// NewNewUnconfirmedTokensTransactionsRData instantiates a new NewUnconfirmedTokensTransactionsRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUnconfirmedTokensTransactionsRData(item NewUnconfirmedTokensTransactionsRI) *NewUnconfirmedTokensTransactionsRData {
	this := NewUnconfirmedTokensTransactionsRData{}
	this.Item = item
	return &this
}

// NewNewUnconfirmedTokensTransactionsRDataWithDefaults instantiates a new NewUnconfirmedTokensTransactionsRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUnconfirmedTokensTransactionsRDataWithDefaults() *NewUnconfirmedTokensTransactionsRData {
	this := NewUnconfirmedTokensTransactionsRData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewUnconfirmedTokensTransactionsRData) GetItem() NewUnconfirmedTokensTransactionsRI {
	if o == nil {
		var ret NewUnconfirmedTokensTransactionsRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRData) GetItemOk() (*NewUnconfirmedTokensTransactionsRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewUnconfirmedTokensTransactionsRData) SetItem(v NewUnconfirmedTokensTransactionsRI) {
	o.Item = v
}

func (o NewUnconfirmedTokensTransactionsRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewUnconfirmedTokensTransactionsRData struct {
	value *NewUnconfirmedTokensTransactionsRData
	isSet bool
}

func (v NullableNewUnconfirmedTokensTransactionsRData) Get() *NewUnconfirmedTokensTransactionsRData {
	return v.value
}

func (v *NullableNewUnconfirmedTokensTransactionsRData) Set(val *NewUnconfirmedTokensTransactionsRData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUnconfirmedTokensTransactionsRData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUnconfirmedTokensTransactionsRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUnconfirmedTokensTransactionsRData(val *NewUnconfirmedTokensTransactionsRData) *NullableNewUnconfirmedTokensTransactionsRData {
	return &NullableNewUnconfirmedTokensTransactionsRData{value: val, isSet: true}
}

func (v NullableNewUnconfirmedTokensTransactionsRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUnconfirmedTokensTransactionsRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


