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

// GetRawTransactionDataRData struct for GetRawTransactionDataRData
type GetRawTransactionDataRData struct {
	Item GetRawTransactionDataRI `json:"item"`
}

// NewGetRawTransactionDataRData instantiates a new GetRawTransactionDataRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawTransactionDataRData(item GetRawTransactionDataRI) *GetRawTransactionDataRData {
	this := GetRawTransactionDataRData{}
	this.Item = item
	return &this
}

// NewGetRawTransactionDataRDataWithDefaults instantiates a new GetRawTransactionDataRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawTransactionDataRDataWithDefaults() *GetRawTransactionDataRData {
	this := GetRawTransactionDataRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetRawTransactionDataRData) GetItem() GetRawTransactionDataRI {
	if o == nil {
		var ret GetRawTransactionDataRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetRawTransactionDataRData) GetItemOk() (*GetRawTransactionDataRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetRawTransactionDataRData) SetItem(v GetRawTransactionDataRI) {
	o.Item = v
}

func (o GetRawTransactionDataRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetRawTransactionDataRData struct {
	value *GetRawTransactionDataRData
	isSet bool
}

func (v NullableGetRawTransactionDataRData) Get() *GetRawTransactionDataRData {
	return v.value
}

func (v *NullableGetRawTransactionDataRData) Set(val *GetRawTransactionDataRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawTransactionDataRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawTransactionDataRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawTransactionDataRData(val *GetRawTransactionDataRData) *NullableGetRawTransactionDataRData {
	return &NullableGetRawTransactionDataRData{value: val, isSet: true}
}

func (v NullableGetRawTransactionDataRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawTransactionDataRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


