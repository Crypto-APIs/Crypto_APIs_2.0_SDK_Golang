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

// ValidateAddressRequestBodyData struct for ValidateAddressRequestBodyData
type ValidateAddressRequestBodyData struct {
	Item ValidateAddressRequestBodyDataItem `json:"item"`
}

// NewValidateAddressRequestBodyData instantiates a new ValidateAddressRequestBodyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateAddressRequestBodyData(item ValidateAddressRequestBodyDataItem) *ValidateAddressRequestBodyData {
	this := ValidateAddressRequestBodyData{}
	this.Item = item
	return &this
}

// NewValidateAddressRequestBodyDataWithDefaults instantiates a new ValidateAddressRequestBodyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateAddressRequestBodyDataWithDefaults() *ValidateAddressRequestBodyData {
	this := ValidateAddressRequestBodyData{}
	return &this
}

// GetItem returns the Item field value
func (o *ValidateAddressRequestBodyData) GetItem() ValidateAddressRequestBodyDataItem {
	if o == nil {
		var ret ValidateAddressRequestBodyDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *ValidateAddressRequestBodyData) GetItemOk() (*ValidateAddressRequestBodyDataItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *ValidateAddressRequestBodyData) SetItem(v ValidateAddressRequestBodyDataItem) {
	o.Item = v
}

func (o ValidateAddressRequestBodyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableValidateAddressRequestBodyData struct {
	value *ValidateAddressRequestBodyData
	isSet bool
}

func (v NullableValidateAddressRequestBodyData) Get() *ValidateAddressRequestBodyData {
	return v.value
}

func (v *NullableValidateAddressRequestBodyData) Set(val *ValidateAddressRequestBodyData) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateAddressRequestBodyData) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateAddressRequestBodyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateAddressRequestBodyData(val *ValidateAddressRequestBodyData) *NullableValidateAddressRequestBodyData {
	return &NullableValidateAddressRequestBodyData{value: val, isSet: true}
}

func (v NullableValidateAddressRequestBodyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateAddressRequestBodyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


