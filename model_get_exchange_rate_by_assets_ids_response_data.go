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

// GetExchangeRateByAssetsIDsResponseData struct for GetExchangeRateByAssetsIDsResponseData
type GetExchangeRateByAssetsIDsResponseData struct {
	Item GetExchangeRateByAssetsIDsResponseItem `json:"item"`
}

// NewGetExchangeRateByAssetsIDsResponseData instantiates a new GetExchangeRateByAssetsIDsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeRateByAssetsIDsResponseData(item GetExchangeRateByAssetsIDsResponseItem) *GetExchangeRateByAssetsIDsResponseData {
	this := GetExchangeRateByAssetsIDsResponseData{}
	this.Item = item
	return &this
}

// NewGetExchangeRateByAssetsIDsResponseDataWithDefaults instantiates a new GetExchangeRateByAssetsIDsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeRateByAssetsIDsResponseDataWithDefaults() *GetExchangeRateByAssetsIDsResponseData {
	this := GetExchangeRateByAssetsIDsResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetExchangeRateByAssetsIDsResponseData) GetItem() GetExchangeRateByAssetsIDsResponseItem {
	if o == nil {
		var ret GetExchangeRateByAssetsIDsResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetsIDsResponseData) GetItemOk() (*GetExchangeRateByAssetsIDsResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetExchangeRateByAssetsIDsResponseData) SetItem(v GetExchangeRateByAssetsIDsResponseItem) {
	o.Item = v
}

func (o GetExchangeRateByAssetsIDsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetExchangeRateByAssetsIDsResponseData struct {
	value *GetExchangeRateByAssetsIDsResponseData
	isSet bool
}

func (v NullableGetExchangeRateByAssetsIDsResponseData) Get() *GetExchangeRateByAssetsIDsResponseData {
	return v.value
}

func (v *NullableGetExchangeRateByAssetsIDsResponseData) Set(val *GetExchangeRateByAssetsIDsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeRateByAssetsIDsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeRateByAssetsIDsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeRateByAssetsIDsResponseData(val *GetExchangeRateByAssetsIDsResponseData) *NullableGetExchangeRateByAssetsIDsResponseData {
	return &NullableGetExchangeRateByAssetsIDsResponseData{value: val, isSet: true}
}

func (v NullableGetExchangeRateByAssetsIDsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeRateByAssetsIDsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

