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

// BroadcastLocallySignedTransactionRData struct for BroadcastLocallySignedTransactionRData
type BroadcastLocallySignedTransactionRData struct {
	Item BroadcastLocallySignedTransactionRI `json:"item"`
}

// NewBroadcastLocallySignedTransactionRData instantiates a new BroadcastLocallySignedTransactionRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastLocallySignedTransactionRData(item BroadcastLocallySignedTransactionRI) *BroadcastLocallySignedTransactionRData {
	this := BroadcastLocallySignedTransactionRData{}
	this.Item = item
	return &this
}

// NewBroadcastLocallySignedTransactionRDataWithDefaults instantiates a new BroadcastLocallySignedTransactionRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastLocallySignedTransactionRDataWithDefaults() *BroadcastLocallySignedTransactionRData {
	this := BroadcastLocallySignedTransactionRData{}
	return &this
}

// GetItem returns the Item field value
func (o *BroadcastLocallySignedTransactionRData) GetItem() BroadcastLocallySignedTransactionRI {
	if o == nil {
		var ret BroadcastLocallySignedTransactionRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *BroadcastLocallySignedTransactionRData) GetItemOk() (*BroadcastLocallySignedTransactionRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *BroadcastLocallySignedTransactionRData) SetItem(v BroadcastLocallySignedTransactionRI) {
	o.Item = v
}

func (o BroadcastLocallySignedTransactionRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableBroadcastLocallySignedTransactionRData struct {
	value *BroadcastLocallySignedTransactionRData
	isSet bool
}

func (v NullableBroadcastLocallySignedTransactionRData) Get() *BroadcastLocallySignedTransactionRData {
	return v.value
}

func (v *NullableBroadcastLocallySignedTransactionRData) Set(val *BroadcastLocallySignedTransactionRData) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastLocallySignedTransactionRData) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastLocallySignedTransactionRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastLocallySignedTransactionRData(val *BroadcastLocallySignedTransactionRData) *NullableBroadcastLocallySignedTransactionRData {
	return &NullableBroadcastLocallySignedTransactionRData{value: val, isSet: true}
}

func (v NullableBroadcastLocallySignedTransactionRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastLocallySignedTransactionRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

