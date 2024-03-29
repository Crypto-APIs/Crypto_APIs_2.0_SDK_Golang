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

// GenerateDepositAddressRBDataItem struct for GenerateDepositAddressRBDataItem
type GenerateDepositAddressRBDataItem struct {
	// Represents a custom tag that customers can set up for their Wallets and addresses. E.g. custom label named \"Special addresses\".
	Label string `json:"label"`
}

// NewGenerateDepositAddressRBDataItem instantiates a new GenerateDepositAddressRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDepositAddressRBDataItem(label string) *GenerateDepositAddressRBDataItem {
	this := GenerateDepositAddressRBDataItem{}
	this.Label = label
	return &this
}

// NewGenerateDepositAddressRBDataItemWithDefaults instantiates a new GenerateDepositAddressRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDepositAddressRBDataItemWithDefaults() *GenerateDepositAddressRBDataItem {
	this := GenerateDepositAddressRBDataItem{}
	return &this
}

// GetLabel returns the Label field value
func (o *GenerateDepositAddressRBDataItem) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *GenerateDepositAddressRBDataItem) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *GenerateDepositAddressRBDataItem) SetLabel(v string) {
	o.Label = v
}

func (o GenerateDepositAddressRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateDepositAddressRBDataItem struct {
	value *GenerateDepositAddressRBDataItem
	isSet bool
}

func (v NullableGenerateDepositAddressRBDataItem) Get() *GenerateDepositAddressRBDataItem {
	return v.value
}

func (v *NullableGenerateDepositAddressRBDataItem) Set(val *GenerateDepositAddressRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDepositAddressRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDepositAddressRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDepositAddressRBDataItem(val *GenerateDepositAddressRBDataItem) *NullableGenerateDepositAddressRBDataItem {
	return &NullableGenerateDepositAddressRBDataItem{value: val, isSet: true}
}

func (v NullableGenerateDepositAddressRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateDepositAddressRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


