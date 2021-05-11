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

// DeleteAutomaticCoinsForwardingResponseData struct for DeleteAutomaticCoinsForwardingResponseData
type DeleteAutomaticCoinsForwardingResponseData struct {
	Item DeleteAutomaticCoinsForwardingResponseItem `json:"item"`
}

// NewDeleteAutomaticCoinsForwardingResponseData instantiates a new DeleteAutomaticCoinsForwardingResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAutomaticCoinsForwardingResponseData(item DeleteAutomaticCoinsForwardingResponseItem) *DeleteAutomaticCoinsForwardingResponseData {
	this := DeleteAutomaticCoinsForwardingResponseData{}
	this.Item = item
	return &this
}

// NewDeleteAutomaticCoinsForwardingResponseDataWithDefaults instantiates a new DeleteAutomaticCoinsForwardingResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAutomaticCoinsForwardingResponseDataWithDefaults() *DeleteAutomaticCoinsForwardingResponseData {
	this := DeleteAutomaticCoinsForwardingResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *DeleteAutomaticCoinsForwardingResponseData) GetItem() DeleteAutomaticCoinsForwardingResponseItem {
	if o == nil {
		var ret DeleteAutomaticCoinsForwardingResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *DeleteAutomaticCoinsForwardingResponseData) GetItemOk() (*DeleteAutomaticCoinsForwardingResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *DeleteAutomaticCoinsForwardingResponseData) SetItem(v DeleteAutomaticCoinsForwardingResponseItem) {
	o.Item = v
}

func (o DeleteAutomaticCoinsForwardingResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAutomaticCoinsForwardingResponseData struct {
	value *DeleteAutomaticCoinsForwardingResponseData
	isSet bool
}

func (v NullableDeleteAutomaticCoinsForwardingResponseData) Get() *DeleteAutomaticCoinsForwardingResponseData {
	return v.value
}

func (v *NullableDeleteAutomaticCoinsForwardingResponseData) Set(val *DeleteAutomaticCoinsForwardingResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAutomaticCoinsForwardingResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAutomaticCoinsForwardingResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAutomaticCoinsForwardingResponseData(val *DeleteAutomaticCoinsForwardingResponseData) *NullableDeleteAutomaticCoinsForwardingResponseData {
	return &NullableDeleteAutomaticCoinsForwardingResponseData{value: val, isSet: true}
}

func (v NullableDeleteAutomaticCoinsForwardingResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAutomaticCoinsForwardingResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


