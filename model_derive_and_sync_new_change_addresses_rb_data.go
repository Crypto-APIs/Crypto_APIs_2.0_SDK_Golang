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

// DeriveAndSyncNewChangeAddressesRBData struct for DeriveAndSyncNewChangeAddressesRBData
type DeriveAndSyncNewChangeAddressesRBData struct {
	Item DeriveAndSyncNewChangeAddressesRBDataItem `json:"item"`
}

// NewDeriveAndSyncNewChangeAddressesRBData instantiates a new DeriveAndSyncNewChangeAddressesRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeriveAndSyncNewChangeAddressesRBData(item DeriveAndSyncNewChangeAddressesRBDataItem) *DeriveAndSyncNewChangeAddressesRBData {
	this := DeriveAndSyncNewChangeAddressesRBData{}
	this.Item = item
	return &this
}

// NewDeriveAndSyncNewChangeAddressesRBDataWithDefaults instantiates a new DeriveAndSyncNewChangeAddressesRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeriveAndSyncNewChangeAddressesRBDataWithDefaults() *DeriveAndSyncNewChangeAddressesRBData {
	this := DeriveAndSyncNewChangeAddressesRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *DeriveAndSyncNewChangeAddressesRBData) GetItem() DeriveAndSyncNewChangeAddressesRBDataItem {
	if o == nil {
		var ret DeriveAndSyncNewChangeAddressesRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *DeriveAndSyncNewChangeAddressesRBData) GetItemOk() (*DeriveAndSyncNewChangeAddressesRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *DeriveAndSyncNewChangeAddressesRBData) SetItem(v DeriveAndSyncNewChangeAddressesRBDataItem) {
	o.Item = v
}

func (o DeriveAndSyncNewChangeAddressesRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableDeriveAndSyncNewChangeAddressesRBData struct {
	value *DeriveAndSyncNewChangeAddressesRBData
	isSet bool
}

func (v NullableDeriveAndSyncNewChangeAddressesRBData) Get() *DeriveAndSyncNewChangeAddressesRBData {
	return v.value
}

func (v *NullableDeriveAndSyncNewChangeAddressesRBData) Set(val *DeriveAndSyncNewChangeAddressesRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeriveAndSyncNewChangeAddressesRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeriveAndSyncNewChangeAddressesRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeriveAndSyncNewChangeAddressesRBData(val *DeriveAndSyncNewChangeAddressesRBData) *NullableDeriveAndSyncNewChangeAddressesRBData {
	return &NullableDeriveAndSyncNewChangeAddressesRBData{value: val, isSet: true}
}

func (v NullableDeriveAndSyncNewChangeAddressesRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeriveAndSyncNewChangeAddressesRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

