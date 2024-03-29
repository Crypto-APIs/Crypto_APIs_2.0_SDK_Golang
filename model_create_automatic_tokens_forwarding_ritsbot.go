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

// CreateAutomaticTokensForwardingRITSBOT Bitcoin Omni Token
type CreateAutomaticTokensForwardingRITSBOT struct {
	// Defines the `propertyId` of the Omni Layer token.
	PropertyId int32 `json:"propertyId"`
}

// NewCreateAutomaticTokensForwardingRITSBOT instantiates a new CreateAutomaticTokensForwardingRITSBOT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticTokensForwardingRITSBOT(propertyId int32) *CreateAutomaticTokensForwardingRITSBOT {
	this := CreateAutomaticTokensForwardingRITSBOT{}
	this.PropertyId = propertyId
	return &this
}

// NewCreateAutomaticTokensForwardingRITSBOTWithDefaults instantiates a new CreateAutomaticTokensForwardingRITSBOT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticTokensForwardingRITSBOTWithDefaults() *CreateAutomaticTokensForwardingRITSBOT {
	this := CreateAutomaticTokensForwardingRITSBOT{}
	return &this
}

// GetPropertyId returns the PropertyId field value
func (o *CreateAutomaticTokensForwardingRITSBOT) GetPropertyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticTokensForwardingRITSBOT) GetPropertyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *CreateAutomaticTokensForwardingRITSBOT) SetPropertyId(v int32) {
	o.PropertyId = v
}

func (o CreateAutomaticTokensForwardingRITSBOT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["propertyId"] = o.PropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticTokensForwardingRITSBOT struct {
	value *CreateAutomaticTokensForwardingRITSBOT
	isSet bool
}

func (v NullableCreateAutomaticTokensForwardingRITSBOT) Get() *CreateAutomaticTokensForwardingRITSBOT {
	return v.value
}

func (v *NullableCreateAutomaticTokensForwardingRITSBOT) Set(val *CreateAutomaticTokensForwardingRITSBOT) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticTokensForwardingRITSBOT) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticTokensForwardingRITSBOT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticTokensForwardingRITSBOT(val *CreateAutomaticTokensForwardingRITSBOT) *NullableCreateAutomaticTokensForwardingRITSBOT {
	return &NullableCreateAutomaticTokensForwardingRITSBOT{value: val, isSet: true}
}

func (v NullableCreateAutomaticTokensForwardingRITSBOT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticTokensForwardingRITSBOT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


