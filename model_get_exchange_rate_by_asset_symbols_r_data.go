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

// GetExchangeRateByAssetSymbolsRData struct for GetExchangeRateByAssetSymbolsRData
type GetExchangeRateByAssetSymbolsRData struct {
	Item GetExchangeRateByAssetSymbolsRI `json:"item"`
}

// NewGetExchangeRateByAssetSymbolsRData instantiates a new GetExchangeRateByAssetSymbolsRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeRateByAssetSymbolsRData(item GetExchangeRateByAssetSymbolsRI) *GetExchangeRateByAssetSymbolsRData {
	this := GetExchangeRateByAssetSymbolsRData{}
	this.Item = item
	return &this
}

// NewGetExchangeRateByAssetSymbolsRDataWithDefaults instantiates a new GetExchangeRateByAssetSymbolsRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeRateByAssetSymbolsRDataWithDefaults() *GetExchangeRateByAssetSymbolsRData {
	this := GetExchangeRateByAssetSymbolsRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetExchangeRateByAssetSymbolsRData) GetItem() GetExchangeRateByAssetSymbolsRI {
	if o == nil {
		var ret GetExchangeRateByAssetSymbolsRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetSymbolsRData) GetItemOk() (*GetExchangeRateByAssetSymbolsRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetExchangeRateByAssetSymbolsRData) SetItem(v GetExchangeRateByAssetSymbolsRI) {
	o.Item = v
}

func (o GetExchangeRateByAssetSymbolsRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetExchangeRateByAssetSymbolsRData struct {
	value *GetExchangeRateByAssetSymbolsRData
	isSet bool
}

func (v NullableGetExchangeRateByAssetSymbolsRData) Get() *GetExchangeRateByAssetSymbolsRData {
	return v.value
}

func (v *NullableGetExchangeRateByAssetSymbolsRData) Set(val *GetExchangeRateByAssetSymbolsRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeRateByAssetSymbolsRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeRateByAssetSymbolsRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeRateByAssetSymbolsRData(val *GetExchangeRateByAssetSymbolsRData) *NullableGetExchangeRateByAssetSymbolsRData {
	return &NullableGetExchangeRateByAssetSymbolsRData{value: val, isSet: true}
}

func (v NullableGetExchangeRateByAssetSymbolsRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeRateByAssetSymbolsRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


