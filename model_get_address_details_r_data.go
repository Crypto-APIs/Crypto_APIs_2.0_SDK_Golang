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

// GetAddressDetailsRData struct for GetAddressDetailsRData
type GetAddressDetailsRData struct {
	Item GetAddressDetailsRI `json:"item"`
}

// NewGetAddressDetailsRData instantiates a new GetAddressDetailsRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAddressDetailsRData(item GetAddressDetailsRI) *GetAddressDetailsRData {
	this := GetAddressDetailsRData{}
	this.Item = item
	return &this
}

// NewGetAddressDetailsRDataWithDefaults instantiates a new GetAddressDetailsRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAddressDetailsRDataWithDefaults() *GetAddressDetailsRData {
	this := GetAddressDetailsRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetAddressDetailsRData) GetItem() GetAddressDetailsRI {
	if o == nil {
		var ret GetAddressDetailsRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsRData) GetItemOk() (*GetAddressDetailsRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetAddressDetailsRData) SetItem(v GetAddressDetailsRI) {
	o.Item = v
}

func (o GetAddressDetailsRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetAddressDetailsRData struct {
	value *GetAddressDetailsRData
	isSet bool
}

func (v NullableGetAddressDetailsRData) Get() *GetAddressDetailsRData {
	return v.value
}

func (v *NullableGetAddressDetailsRData) Set(val *GetAddressDetailsRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddressDetailsRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddressDetailsRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddressDetailsRData(val *GetAddressDetailsRData) *NullableGetAddressDetailsRData {
	return &NullableGetAddressDetailsRData{value: val, isSet: true}
}

func (v NullableGetAddressDetailsRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddressDetailsRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

