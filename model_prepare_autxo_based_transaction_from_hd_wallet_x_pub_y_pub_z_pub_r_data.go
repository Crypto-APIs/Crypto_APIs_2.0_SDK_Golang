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

// PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData struct for PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData
type PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData struct {
	Item PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI `json:"item"`
}

// NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData(item PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData {
	this := PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData{}
	this.Item = item
	return &this
}

// NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRDataWithDefaults instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRDataWithDefaults() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData {
	this := PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData{}
	return &this
}

// GetItem returns the Item field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) GetItem() PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI {
	if o == nil {
		var ret PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) GetItemOk() (*PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) SetItem(v PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) {
	o.Item = v
}

func (o PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData struct {
	value *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData
	isSet bool
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) Get() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData {
	return v.value
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) Set(val *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData(val *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData {
	return &NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData{value: val, isSet: true}
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


