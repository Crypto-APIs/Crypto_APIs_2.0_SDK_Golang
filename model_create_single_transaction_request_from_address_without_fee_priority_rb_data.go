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

// CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData struct for CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData
type CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData struct {
	Item CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem `json:"item"`
}

// NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData(item CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData {
	this := CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData{}
	this.Item = item
	return &this
}

// NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataWithDefaults instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataWithDefaults() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData {
	this := CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) GetItem() CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem {
	if o == nil {
		var ret CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) GetItemOk() (*CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) SetItem(v CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) {
	o.Item = v
}

func (o CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData struct {
	value *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData
	isSet bool
}

func (v NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) Get() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData {
	return v.value
}

func (v *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) Set(val *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData(val *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData {
	return &NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData{value: val, isSet: true}
}

func (v NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

