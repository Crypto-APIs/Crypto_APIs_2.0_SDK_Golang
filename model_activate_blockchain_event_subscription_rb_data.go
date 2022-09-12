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

// ActivateBlockchainEventSubscriptionRBData struct for ActivateBlockchainEventSubscriptionRBData
type ActivateBlockchainEventSubscriptionRBData struct {
	Item map[string]interface{} `json:"item"`
}

// NewActivateBlockchainEventSubscriptionRBData instantiates a new ActivateBlockchainEventSubscriptionRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivateBlockchainEventSubscriptionRBData(item map[string]interface{}) *ActivateBlockchainEventSubscriptionRBData {
	this := ActivateBlockchainEventSubscriptionRBData{}
	this.Item = item
	return &this
}

// NewActivateBlockchainEventSubscriptionRBDataWithDefaults instantiates a new ActivateBlockchainEventSubscriptionRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivateBlockchainEventSubscriptionRBDataWithDefaults() *ActivateBlockchainEventSubscriptionRBData {
	this := ActivateBlockchainEventSubscriptionRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *ActivateBlockchainEventSubscriptionRBData) GetItem() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *ActivateBlockchainEventSubscriptionRBData) GetItemOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item, true
}

// SetItem sets field value
func (o *ActivateBlockchainEventSubscriptionRBData) SetItem(v map[string]interface{}) {
	o.Item = v
}

func (o ActivateBlockchainEventSubscriptionRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableActivateBlockchainEventSubscriptionRBData struct {
	value *ActivateBlockchainEventSubscriptionRBData
	isSet bool
}

func (v NullableActivateBlockchainEventSubscriptionRBData) Get() *ActivateBlockchainEventSubscriptionRBData {
	return v.value
}

func (v *NullableActivateBlockchainEventSubscriptionRBData) Set(val *ActivateBlockchainEventSubscriptionRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateBlockchainEventSubscriptionRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateBlockchainEventSubscriptionRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateBlockchainEventSubscriptionRBData(val *ActivateBlockchainEventSubscriptionRBData) *NullableActivateBlockchainEventSubscriptionRBData {
	return &NullableActivateBlockchainEventSubscriptionRBData{value: val, isSet: true}
}

func (v NullableActivateBlockchainEventSubscriptionRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateBlockchainEventSubscriptionRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


