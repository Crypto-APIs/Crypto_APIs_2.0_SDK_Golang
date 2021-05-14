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

// GetExchangeRateByAssetSymbolsResponseData struct for GetExchangeRateByAssetSymbolsResponseData
type GetExchangeRateByAssetSymbolsResponseData struct {
	Item GetExchangeRateByAssetSymbolsResponseItem `json:"item"`
}

// NewGetExchangeRateByAssetSymbolsResponseData instantiates a new GetExchangeRateByAssetSymbolsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeRateByAssetSymbolsResponseData(item GetExchangeRateByAssetSymbolsResponseItem) *GetExchangeRateByAssetSymbolsResponseData {
	this := GetExchangeRateByAssetSymbolsResponseData{}
	this.Item = item
	return &this
}

// NewGetExchangeRateByAssetSymbolsResponseDataWithDefaults instantiates a new GetExchangeRateByAssetSymbolsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeRateByAssetSymbolsResponseDataWithDefaults() *GetExchangeRateByAssetSymbolsResponseData {
	this := GetExchangeRateByAssetSymbolsResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetExchangeRateByAssetSymbolsResponseData) GetItem() GetExchangeRateByAssetSymbolsResponseItem {
	if o == nil {
		var ret GetExchangeRateByAssetSymbolsResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetSymbolsResponseData) GetItemOk() (*GetExchangeRateByAssetSymbolsResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetExchangeRateByAssetSymbolsResponseData) SetItem(v GetExchangeRateByAssetSymbolsResponseItem) {
	o.Item = v
}

func (o GetExchangeRateByAssetSymbolsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetExchangeRateByAssetSymbolsResponseData struct {
	value *GetExchangeRateByAssetSymbolsResponseData
	isSet bool
}

func (v NullableGetExchangeRateByAssetSymbolsResponseData) Get() *GetExchangeRateByAssetSymbolsResponseData {
	return v.value
}

func (v *NullableGetExchangeRateByAssetSymbolsResponseData) Set(val *GetExchangeRateByAssetSymbolsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeRateByAssetSymbolsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeRateByAssetSymbolsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeRateByAssetSymbolsResponseData(val *GetExchangeRateByAssetSymbolsResponseData) *NullableGetExchangeRateByAssetSymbolsResponseData {
	return &NullableGetExchangeRateByAssetSymbolsResponseData{value: val, isSet: true}
}

func (v NullableGetExchangeRateByAssetSymbolsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeRateByAssetSymbolsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

