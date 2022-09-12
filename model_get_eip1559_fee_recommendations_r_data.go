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

// GetEIP1559FeeRecommendationsRData struct for GetEIP1559FeeRecommendationsRData
type GetEIP1559FeeRecommendationsRData struct {
	Item GetEIP1559FeeRecommendationsRI `json:"item"`
}

// NewGetEIP1559FeeRecommendationsRData instantiates a new GetEIP1559FeeRecommendationsRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEIP1559FeeRecommendationsRData(item GetEIP1559FeeRecommendationsRI) *GetEIP1559FeeRecommendationsRData {
	this := GetEIP1559FeeRecommendationsRData{}
	this.Item = item
	return &this
}

// NewGetEIP1559FeeRecommendationsRDataWithDefaults instantiates a new GetEIP1559FeeRecommendationsRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEIP1559FeeRecommendationsRDataWithDefaults() *GetEIP1559FeeRecommendationsRData {
	this := GetEIP1559FeeRecommendationsRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetEIP1559FeeRecommendationsRData) GetItem() GetEIP1559FeeRecommendationsRI {
	if o == nil {
		var ret GetEIP1559FeeRecommendationsRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsRData) GetItemOk() (*GetEIP1559FeeRecommendationsRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetEIP1559FeeRecommendationsRData) SetItem(v GetEIP1559FeeRecommendationsRI) {
	o.Item = v
}

func (o GetEIP1559FeeRecommendationsRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetEIP1559FeeRecommendationsRData struct {
	value *GetEIP1559FeeRecommendationsRData
	isSet bool
}

func (v NullableGetEIP1559FeeRecommendationsRData) Get() *GetEIP1559FeeRecommendationsRData {
	return v.value
}

func (v *NullableGetEIP1559FeeRecommendationsRData) Set(val *GetEIP1559FeeRecommendationsRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEIP1559FeeRecommendationsRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEIP1559FeeRecommendationsRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEIP1559FeeRecommendationsRData(val *GetEIP1559FeeRecommendationsRData) *NullableGetEIP1559FeeRecommendationsRData {
	return &NullableGetEIP1559FeeRecommendationsRData{value: val, isSet: true}
}

func (v NullableGetEIP1559FeeRecommendationsRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEIP1559FeeRecommendationsRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


