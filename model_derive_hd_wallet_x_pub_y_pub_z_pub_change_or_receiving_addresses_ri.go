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

// DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI struct for DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI
type DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI struct {
	// Represents the address details.
	Addresses []DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIAddressesInner `json:"addresses"`
}

// NewDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI instantiates a new DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI(addresses []DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIAddressesInner) *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI {
	this := DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI{}
	this.Addresses = addresses
	return &this
}

// NewDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIWithDefaults instantiates a new DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIWithDefaults() *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI {
	this := DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) GetAddresses() []DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIAddressesInner {
	if o == nil {
		var ret []DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIAddressesInner
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) GetAddressesOk() ([]DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIAddressesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) SetAddresses(v []DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRIAddressesInner) {
	o.Addresses = v
}

func (o DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI struct {
	value *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI
	isSet bool
}

func (v NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) Get() *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI {
	return v.value
}

func (v *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) Set(val *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) {
	v.value = val
	v.isSet = true
}

func (v NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) IsSet() bool {
	return v.isSet
}

func (v *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI(val *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI {
	return &NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI{value: val, isSet: true}
}

func (v NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


