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

// GetTokenDetailsByContractAddressRData struct for GetTokenDetailsByContractAddressRData
type GetTokenDetailsByContractAddressRData struct {
	Item GetTokenDetailsByContractAddressRI `json:"item"`
}

// NewGetTokenDetailsByContractAddressRData instantiates a new GetTokenDetailsByContractAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTokenDetailsByContractAddressRData(item GetTokenDetailsByContractAddressRI) *GetTokenDetailsByContractAddressRData {
	this := GetTokenDetailsByContractAddressRData{}
	this.Item = item
	return &this
}

// NewGetTokenDetailsByContractAddressRDataWithDefaults instantiates a new GetTokenDetailsByContractAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTokenDetailsByContractAddressRDataWithDefaults() *GetTokenDetailsByContractAddressRData {
	this := GetTokenDetailsByContractAddressRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetTokenDetailsByContractAddressRData) GetItem() GetTokenDetailsByContractAddressRI {
	if o == nil {
		var ret GetTokenDetailsByContractAddressRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetTokenDetailsByContractAddressRData) GetItemOk() (*GetTokenDetailsByContractAddressRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetTokenDetailsByContractAddressRData) SetItem(v GetTokenDetailsByContractAddressRI) {
	o.Item = v
}

func (o GetTokenDetailsByContractAddressRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetTokenDetailsByContractAddressRData struct {
	value *GetTokenDetailsByContractAddressRData
	isSet bool
}

func (v NullableGetTokenDetailsByContractAddressRData) Get() *GetTokenDetailsByContractAddressRData {
	return v.value
}

func (v *NullableGetTokenDetailsByContractAddressRData) Set(val *GetTokenDetailsByContractAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTokenDetailsByContractAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTokenDetailsByContractAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTokenDetailsByContractAddressRData(val *GetTokenDetailsByContractAddressRData) *NullableGetTokenDetailsByContractAddressRData {
	return &NullableGetTokenDetailsByContractAddressRData{value: val, isSet: true}
}

func (v NullableGetTokenDetailsByContractAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTokenDetailsByContractAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


