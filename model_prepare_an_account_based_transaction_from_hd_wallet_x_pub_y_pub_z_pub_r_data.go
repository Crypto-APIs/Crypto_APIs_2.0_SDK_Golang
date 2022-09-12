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

// PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData struct for PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData
type PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData struct {
	Item PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI `json:"item"`
}

// NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData(item PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData {
	this := PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData{}
	this.Item = item
	return &this
}

// NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRDataWithDefaults instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRDataWithDefaults() *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData {
	this := PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData{}
	return &this
}

// GetItem returns the Item field value
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) GetItem() PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI {
	if o == nil {
		var ret PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) GetItemOk() (*PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) SetItem(v PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) {
	o.Item = v
}

func (o PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData struct {
	value *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData
	isSet bool
}

func (v NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) Get() *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData {
	return v.value
}

func (v *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) Set(val *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData(val *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData {
	return &NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData{value: val, isSet: true}
}

func (v NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


