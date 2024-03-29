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

// EstimateTransactionSmartFeeRData struct for EstimateTransactionSmartFeeRData
type EstimateTransactionSmartFeeRData struct {
	Item EstimateTransactionSmartFeeRI `json:"item"`
}

// NewEstimateTransactionSmartFeeRData instantiates a new EstimateTransactionSmartFeeRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateTransactionSmartFeeRData(item EstimateTransactionSmartFeeRI) *EstimateTransactionSmartFeeRData {
	this := EstimateTransactionSmartFeeRData{}
	this.Item = item
	return &this
}

// NewEstimateTransactionSmartFeeRDataWithDefaults instantiates a new EstimateTransactionSmartFeeRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateTransactionSmartFeeRDataWithDefaults() *EstimateTransactionSmartFeeRData {
	this := EstimateTransactionSmartFeeRData{}
	return &this
}

// GetItem returns the Item field value
func (o *EstimateTransactionSmartFeeRData) GetItem() EstimateTransactionSmartFeeRI {
	if o == nil {
		var ret EstimateTransactionSmartFeeRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *EstimateTransactionSmartFeeRData) GetItemOk() (*EstimateTransactionSmartFeeRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *EstimateTransactionSmartFeeRData) SetItem(v EstimateTransactionSmartFeeRI) {
	o.Item = v
}

func (o EstimateTransactionSmartFeeRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableEstimateTransactionSmartFeeRData struct {
	value *EstimateTransactionSmartFeeRData
	isSet bool
}

func (v NullableEstimateTransactionSmartFeeRData) Get() *EstimateTransactionSmartFeeRData {
	return v.value
}

func (v *NullableEstimateTransactionSmartFeeRData) Set(val *EstimateTransactionSmartFeeRData) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateTransactionSmartFeeRData) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateTransactionSmartFeeRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateTransactionSmartFeeRData(val *EstimateTransactionSmartFeeRData) *NullableEstimateTransactionSmartFeeRData {
	return &NullableEstimateTransactionSmartFeeRData{value: val, isSet: true}
}

func (v NullableEstimateTransactionSmartFeeRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateTransactionSmartFeeRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


