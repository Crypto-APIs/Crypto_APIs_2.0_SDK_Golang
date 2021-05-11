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

// DeleteAutomaticTokensForwardingResponseData struct for DeleteAutomaticTokensForwardingResponseData
type DeleteAutomaticTokensForwardingResponseData struct {
	Item DeleteAutomaticTokensForwardingResponseItem `json:"item"`
}

// NewDeleteAutomaticTokensForwardingResponseData instantiates a new DeleteAutomaticTokensForwardingResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAutomaticTokensForwardingResponseData(item DeleteAutomaticTokensForwardingResponseItem) *DeleteAutomaticTokensForwardingResponseData {
	this := DeleteAutomaticTokensForwardingResponseData{}
	this.Item = item
	return &this
}

// NewDeleteAutomaticTokensForwardingResponseDataWithDefaults instantiates a new DeleteAutomaticTokensForwardingResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAutomaticTokensForwardingResponseDataWithDefaults() *DeleteAutomaticTokensForwardingResponseData {
	this := DeleteAutomaticTokensForwardingResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *DeleteAutomaticTokensForwardingResponseData) GetItem() DeleteAutomaticTokensForwardingResponseItem {
	if o == nil {
		var ret DeleteAutomaticTokensForwardingResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *DeleteAutomaticTokensForwardingResponseData) GetItemOk() (*DeleteAutomaticTokensForwardingResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *DeleteAutomaticTokensForwardingResponseData) SetItem(v DeleteAutomaticTokensForwardingResponseItem) {
	o.Item = v
}

func (o DeleteAutomaticTokensForwardingResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAutomaticTokensForwardingResponseData struct {
	value *DeleteAutomaticTokensForwardingResponseData
	isSet bool
}

func (v NullableDeleteAutomaticTokensForwardingResponseData) Get() *DeleteAutomaticTokensForwardingResponseData {
	return v.value
}

func (v *NullableDeleteAutomaticTokensForwardingResponseData) Set(val *DeleteAutomaticTokensForwardingResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAutomaticTokensForwardingResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAutomaticTokensForwardingResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAutomaticTokensForwardingResponseData(val *DeleteAutomaticTokensForwardingResponseData) *NullableDeleteAutomaticTokensForwardingResponseData {
	return &NullableDeleteAutomaticTokensForwardingResponseData{value: val, isSet: true}
}

func (v NullableDeleteAutomaticTokensForwardingResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAutomaticTokensForwardingResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


