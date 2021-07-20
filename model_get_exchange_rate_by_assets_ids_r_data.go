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

// GetExchangeRateByAssetsIDsRData struct for GetExchangeRateByAssetsIDsRData
type GetExchangeRateByAssetsIDsRData struct {
	Item GetExchangeRateByAssetsIDsRI `json:"item"`
}

// NewGetExchangeRateByAssetsIDsRData instantiates a new GetExchangeRateByAssetsIDsRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeRateByAssetsIDsRData(item GetExchangeRateByAssetsIDsRI) *GetExchangeRateByAssetsIDsRData {
	this := GetExchangeRateByAssetsIDsRData{}
	this.Item = item
	return &this
}

// NewGetExchangeRateByAssetsIDsRDataWithDefaults instantiates a new GetExchangeRateByAssetsIDsRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeRateByAssetsIDsRDataWithDefaults() *GetExchangeRateByAssetsIDsRData {
	this := GetExchangeRateByAssetsIDsRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetExchangeRateByAssetsIDsRData) GetItem() GetExchangeRateByAssetsIDsRI {
	if o == nil {
		var ret GetExchangeRateByAssetsIDsRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetsIDsRData) GetItemOk() (*GetExchangeRateByAssetsIDsRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetExchangeRateByAssetsIDsRData) SetItem(v GetExchangeRateByAssetsIDsRI) {
	o.Item = v
}

func (o GetExchangeRateByAssetsIDsRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetExchangeRateByAssetsIDsRData struct {
	value *GetExchangeRateByAssetsIDsRData
	isSet bool
}

func (v NullableGetExchangeRateByAssetsIDsRData) Get() *GetExchangeRateByAssetsIDsRData {
	return v.value
}

func (v *NullableGetExchangeRateByAssetsIDsRData) Set(val *GetExchangeRateByAssetsIDsRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeRateByAssetsIDsRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeRateByAssetsIDsRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeRateByAssetsIDsRData(val *GetExchangeRateByAssetsIDsRData) *NullableGetExchangeRateByAssetsIDsRData {
	return &NullableGetExchangeRateByAssetsIDsRData{value: val, isSet: true}
}

func (v NullableGetExchangeRateByAssetsIDsRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeRateByAssetsIDsRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

