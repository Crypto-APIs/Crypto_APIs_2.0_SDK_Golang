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

// ListConfirmedTransactionsByAddressRIBSD2 Dash
type ListConfirmedTransactionsByAddressRIBSD2 struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the transaction's version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListConfirmedTransactionsByAddressRIBSD2VinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListConfirmedTransactionsByAddressRIBSD2VoutInner `json:"vout"`
}

// NewListConfirmedTransactionsByAddressRIBSD2 instantiates a new ListConfirmedTransactionsByAddressRIBSD2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSD2(locktime int64, size int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSD2VinInner, vout []ListConfirmedTransactionsByAddressRIBSD2VoutInner) *ListConfirmedTransactionsByAddressRIBSD2 {
	this := ListConfirmedTransactionsByAddressRIBSD2{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSD2WithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSD2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSD2WithDefaults() *ListConfirmedTransactionsByAddressRIBSD2 {
	this := ListConfirmedTransactionsByAddressRIBSD2{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetVin() []ListConfirmedTransactionsByAddressRIBSD2VinInner {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSD2VinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetVinOk() ([]ListConfirmedTransactionsByAddressRIBSD2VinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) SetVin(v []ListConfirmedTransactionsByAddressRIBSD2VinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetVout() []ListConfirmedTransactionsByAddressRIBSD2VoutInner {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSD2VoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2) GetVoutOk() ([]ListConfirmedTransactionsByAddressRIBSD2VoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2) SetVout(v []ListConfirmedTransactionsByAddressRIBSD2VoutInner) {
	o.Vout = v
}

func (o ListConfirmedTransactionsByAddressRIBSD2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locktime"] = o.Locktime
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["vin"] = o.Vin
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSD2 struct {
	value *ListConfirmedTransactionsByAddressRIBSD2
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2) Get() *ListConfirmedTransactionsByAddressRIBSD2 {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2) Set(val *ListConfirmedTransactionsByAddressRIBSD2) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSD2(val *ListConfirmedTransactionsByAddressRIBSD2) *NullableListConfirmedTransactionsByAddressRIBSD2 {
	return &NullableListConfirmedTransactionsByAddressRIBSD2{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


