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

// BlockHeightReachedRBData struct for BlockHeightReachedRBData
type BlockHeightReachedRBData struct {
	Item BlockHeightReachedRBDataItem `json:"item"`
}

// NewBlockHeightReachedRBData instantiates a new BlockHeightReachedRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockHeightReachedRBData(item BlockHeightReachedRBDataItem) *BlockHeightReachedRBData {
	this := BlockHeightReachedRBData{}
	this.Item = item
	return &this
}

// NewBlockHeightReachedRBDataWithDefaults instantiates a new BlockHeightReachedRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockHeightReachedRBDataWithDefaults() *BlockHeightReachedRBData {
	this := BlockHeightReachedRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *BlockHeightReachedRBData) GetItem() BlockHeightReachedRBDataItem {
	if o == nil {
		var ret BlockHeightReachedRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *BlockHeightReachedRBData) GetItemOk() (*BlockHeightReachedRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *BlockHeightReachedRBData) SetItem(v BlockHeightReachedRBDataItem) {
	o.Item = v
}

func (o BlockHeightReachedRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableBlockHeightReachedRBData struct {
	value *BlockHeightReachedRBData
	isSet bool
}

func (v NullableBlockHeightReachedRBData) Get() *BlockHeightReachedRBData {
	return v.value
}

func (v *NullableBlockHeightReachedRBData) Set(val *BlockHeightReachedRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeightReachedRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeightReachedRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeightReachedRBData(val *BlockHeightReachedRBData) *NullableBlockHeightReachedRBData {
	return &NullableBlockHeightReachedRBData{value: val, isSet: true}
}

func (v NullableBlockHeightReachedRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeightReachedRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

