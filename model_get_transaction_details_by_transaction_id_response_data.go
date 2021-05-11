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

// GetTransactionDetailsByTransactionIDResponseData struct for GetTransactionDetailsByTransactionIDResponseData
type GetTransactionDetailsByTransactionIDResponseData struct {
	Item GetTransactionDetailsByTransactionIDResponseItem `json:"item"`
}

// NewGetTransactionDetailsByTransactionIDResponseData instantiates a new GetTransactionDetailsByTransactionIDResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDResponseData(item GetTransactionDetailsByTransactionIDResponseItem) *GetTransactionDetailsByTransactionIDResponseData {
	this := GetTransactionDetailsByTransactionIDResponseData{}
	this.Item = item
	return &this
}

// NewGetTransactionDetailsByTransactionIDResponseDataWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDResponseDataWithDefaults() *GetTransactionDetailsByTransactionIDResponseData {
	this := GetTransactionDetailsByTransactionIDResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetTransactionDetailsByTransactionIDResponseData) GetItem() GetTransactionDetailsByTransactionIDResponseItem {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseData) GetItemOk() (*GetTransactionDetailsByTransactionIDResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetTransactionDetailsByTransactionIDResponseData) SetItem(v GetTransactionDetailsByTransactionIDResponseItem) {
	o.Item = v
}

func (o GetTransactionDetailsByTransactionIDResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDResponseData struct {
	value *GetTransactionDetailsByTransactionIDResponseData
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDResponseData) Get() *GetTransactionDetailsByTransactionIDResponseData {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseData) Set(val *GetTransactionDetailsByTransactionIDResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDResponseData(val *GetTransactionDetailsByTransactionIDResponseData) *NullableGetTransactionDetailsByTransactionIDResponseData {
	return &NullableGetTransactionDetailsByTransactionIDResponseData{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


