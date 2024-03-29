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

// PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner struct for PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner
type PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner struct {
	// Representation of the address
	Address string `json:"address"`
	// Representation of the satoshis value
	Satoshis int64 `json:"satoshis"`
	// Representation of the script
	Script string `json:"script"`
}

// NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner(address string, satoshis int64, script string) *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner {
	this := PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner{}
	this.Address = address
	this.Satoshis = satoshis
	this.Script = script
	return &this
}

// NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInnerWithDefaults instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInnerWithDefaults() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner {
	this := PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) SetAddress(v string) {
	o.Address = v
}

// GetSatoshis returns the Satoshis field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) GetSatoshis() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Satoshis
}

// GetSatoshisOk returns a tuple with the Satoshis field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) GetSatoshisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Satoshis, true
}

// SetSatoshis sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) SetSatoshis(v int64) {
	o.Satoshis = v
}

// GetScript returns the Script field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) SetScript(v string) {
	o.Script = v
}

func (o PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["satoshis"] = o.Satoshis
	}
	if true {
		toSerialize["script"] = o.Script
	}
	return json.Marshal(toSerialize)
}

type NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner struct {
	value *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner
	isSet bool
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) Get() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner {
	return v.value
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) Set(val *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner(val *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner {
	return &NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner{value: val, isSet: true}
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


