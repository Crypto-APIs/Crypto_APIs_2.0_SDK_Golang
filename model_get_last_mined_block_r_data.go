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

// GetLastMinedBlockRData struct for GetLastMinedBlockRData
type GetLastMinedBlockRData struct {
	Item GetLastMinedBlockRI `json:"item"`
}

// NewGetLastMinedBlockRData instantiates a new GetLastMinedBlockRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLastMinedBlockRData(item GetLastMinedBlockRI) *GetLastMinedBlockRData {
	this := GetLastMinedBlockRData{}
	this.Item = item
	return &this
}

// NewGetLastMinedBlockRDataWithDefaults instantiates a new GetLastMinedBlockRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLastMinedBlockRDataWithDefaults() *GetLastMinedBlockRData {
	this := GetLastMinedBlockRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetLastMinedBlockRData) GetItem() GetLastMinedBlockRI {
	if o == nil {
		var ret GetLastMinedBlockRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetLastMinedBlockRData) GetItemOk() (*GetLastMinedBlockRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetLastMinedBlockRData) SetItem(v GetLastMinedBlockRI) {
	o.Item = v
}

func (o GetLastMinedBlockRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetLastMinedBlockRData struct {
	value *GetLastMinedBlockRData
	isSet bool
}

func (v NullableGetLastMinedBlockRData) Get() *GetLastMinedBlockRData {
	return v.value
}

func (v *NullableGetLastMinedBlockRData) Set(val *GetLastMinedBlockRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLastMinedBlockRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLastMinedBlockRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLastMinedBlockRData(val *GetLastMinedBlockRData) *NullableGetLastMinedBlockRData {
	return &NullableGetLastMinedBlockRData{value: val, isSet: true}
}

func (v NullableGetLastMinedBlockRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLastMinedBlockRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


