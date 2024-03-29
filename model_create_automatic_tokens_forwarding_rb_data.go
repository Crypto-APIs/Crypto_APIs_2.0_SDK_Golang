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

// CreateAutomaticTokensForwardingRBData struct for CreateAutomaticTokensForwardingRBData
type CreateAutomaticTokensForwardingRBData struct {
	Item CreateAutomaticTokensForwardingRBDataItem `json:"item"`
}

// NewCreateAutomaticTokensForwardingRBData instantiates a new CreateAutomaticTokensForwardingRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticTokensForwardingRBData(item CreateAutomaticTokensForwardingRBDataItem) *CreateAutomaticTokensForwardingRBData {
	this := CreateAutomaticTokensForwardingRBData{}
	this.Item = item
	return &this
}

// NewCreateAutomaticTokensForwardingRBDataWithDefaults instantiates a new CreateAutomaticTokensForwardingRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticTokensForwardingRBDataWithDefaults() *CreateAutomaticTokensForwardingRBData {
	this := CreateAutomaticTokensForwardingRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *CreateAutomaticTokensForwardingRBData) GetItem() CreateAutomaticTokensForwardingRBDataItem {
	if o == nil {
		var ret CreateAutomaticTokensForwardingRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticTokensForwardingRBData) GetItemOk() (*CreateAutomaticTokensForwardingRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *CreateAutomaticTokensForwardingRBData) SetItem(v CreateAutomaticTokensForwardingRBDataItem) {
	o.Item = v
}

func (o CreateAutomaticTokensForwardingRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticTokensForwardingRBData struct {
	value *CreateAutomaticTokensForwardingRBData
	isSet bool
}

func (v NullableCreateAutomaticTokensForwardingRBData) Get() *CreateAutomaticTokensForwardingRBData {
	return v.value
}

func (v *NullableCreateAutomaticTokensForwardingRBData) Set(val *CreateAutomaticTokensForwardingRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticTokensForwardingRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticTokensForwardingRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticTokensForwardingRBData(val *CreateAutomaticTokensForwardingRBData) *NullableCreateAutomaticTokensForwardingRBData {
	return &NullableCreateAutomaticTokensForwardingRBData{value: val, isSet: true}
}

func (v NullableCreateAutomaticTokensForwardingRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticTokensForwardingRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


