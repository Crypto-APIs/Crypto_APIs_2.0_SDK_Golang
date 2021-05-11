/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ValidateAddressResponseData struct for ValidateAddressResponseData
type ValidateAddressResponseData struct {
	Item ValidateAddressResponseItem `json:"item"`
}

// NewValidateAddressResponseData instantiates a new ValidateAddressResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateAddressResponseData(item ValidateAddressResponseItem) *ValidateAddressResponseData {
	this := ValidateAddressResponseData{}
	this.Item = item
	return &this
}

// NewValidateAddressResponseDataWithDefaults instantiates a new ValidateAddressResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateAddressResponseDataWithDefaults() *ValidateAddressResponseData {
	this := ValidateAddressResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *ValidateAddressResponseData) GetItem() ValidateAddressResponseItem {
	if o == nil {
		var ret ValidateAddressResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *ValidateAddressResponseData) GetItemOk() (*ValidateAddressResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *ValidateAddressResponseData) SetItem(v ValidateAddressResponseItem) {
	o.Item = v
}

func (o ValidateAddressResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableValidateAddressResponseData struct {
	value *ValidateAddressResponseData
	isSet bool
}

func (v NullableValidateAddressResponseData) Get() *ValidateAddressResponseData {
	return v.value
}

func (v *NullableValidateAddressResponseData) Set(val *ValidateAddressResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateAddressResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateAddressResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateAddressResponseData(val *ValidateAddressResponseData) *NullableValidateAddressResponseData {
	return &NullableValidateAddressResponseData{value: val, isSet: true}
}

func (v NullableValidateAddressResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateAddressResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


