/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// GetBlockDetailsByBlockHashFromCallbackRData struct for GetBlockDetailsByBlockHashFromCallbackRData
type GetBlockDetailsByBlockHashFromCallbackRData struct {
	Item GetBlockDetailsByBlockHashFromCallbackRI `json:"item"`
}

// NewGetBlockDetailsByBlockHashFromCallbackRData instantiates a new GetBlockDetailsByBlockHashFromCallbackRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashFromCallbackRData(item GetBlockDetailsByBlockHashFromCallbackRI) *GetBlockDetailsByBlockHashFromCallbackRData {
	this := GetBlockDetailsByBlockHashFromCallbackRData{}
	this.Item = item
	return &this
}

// NewGetBlockDetailsByBlockHashFromCallbackRDataWithDefaults instantiates a new GetBlockDetailsByBlockHashFromCallbackRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashFromCallbackRDataWithDefaults() *GetBlockDetailsByBlockHashFromCallbackRData {
	this := GetBlockDetailsByBlockHashFromCallbackRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetBlockDetailsByBlockHashFromCallbackRData) GetItem() GetBlockDetailsByBlockHashFromCallbackRI {
	if o == nil {
		var ret GetBlockDetailsByBlockHashFromCallbackRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRData) GetItemOk() (*GetBlockDetailsByBlockHashFromCallbackRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRData) SetItem(v GetBlockDetailsByBlockHashFromCallbackRI) {
	o.Item = v
}

func (o GetBlockDetailsByBlockHashFromCallbackRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHashFromCallbackRData struct {
	value *GetBlockDetailsByBlockHashFromCallbackRData
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRData) Get() *GetBlockDetailsByBlockHashFromCallbackRData {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRData) Set(val *GetBlockDetailsByBlockHashFromCallbackRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashFromCallbackRData(val *GetBlockDetailsByBlockHashFromCallbackRData) *NullableGetBlockDetailsByBlockHashFromCallbackRData {
	return &NullableGetBlockDetailsByBlockHashFromCallbackRData{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


