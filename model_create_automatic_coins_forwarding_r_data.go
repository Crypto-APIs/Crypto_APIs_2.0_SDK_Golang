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

// CreateAutomaticCoinsForwardingRData struct for CreateAutomaticCoinsForwardingRData
type CreateAutomaticCoinsForwardingRData struct {
	Item CreateAutomaticCoinsForwardingRI `json:"item"`
}

// NewCreateAutomaticCoinsForwardingRData instantiates a new CreateAutomaticCoinsForwardingRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticCoinsForwardingRData(item CreateAutomaticCoinsForwardingRI) *CreateAutomaticCoinsForwardingRData {
	this := CreateAutomaticCoinsForwardingRData{}
	this.Item = item
	return &this
}

// NewCreateAutomaticCoinsForwardingRDataWithDefaults instantiates a new CreateAutomaticCoinsForwardingRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticCoinsForwardingRDataWithDefaults() *CreateAutomaticCoinsForwardingRData {
	this := CreateAutomaticCoinsForwardingRData{}
	return &this
}

// GetItem returns the Item field value
func (o *CreateAutomaticCoinsForwardingRData) GetItem() CreateAutomaticCoinsForwardingRI {
	if o == nil {
		var ret CreateAutomaticCoinsForwardingRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRData) GetItemOk() (*CreateAutomaticCoinsForwardingRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *CreateAutomaticCoinsForwardingRData) SetItem(v CreateAutomaticCoinsForwardingRI) {
	o.Item = v
}

func (o CreateAutomaticCoinsForwardingRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticCoinsForwardingRData struct {
	value *CreateAutomaticCoinsForwardingRData
	isSet bool
}

func (v NullableCreateAutomaticCoinsForwardingRData) Get() *CreateAutomaticCoinsForwardingRData {
	return v.value
}

func (v *NullableCreateAutomaticCoinsForwardingRData) Set(val *CreateAutomaticCoinsForwardingRData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticCoinsForwardingRData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticCoinsForwardingRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticCoinsForwardingRData(val *CreateAutomaticCoinsForwardingRData) *NullableCreateAutomaticCoinsForwardingRData {
	return &NullableCreateAutomaticCoinsForwardingRData{value: val, isSet: true}
}

func (v NullableCreateAutomaticCoinsForwardingRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticCoinsForwardingRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


