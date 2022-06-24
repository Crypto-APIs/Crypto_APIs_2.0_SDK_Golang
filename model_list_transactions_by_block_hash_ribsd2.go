/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ListTransactionsByBlockHashRIBSD2 Dash
type ListTransactionsByBlockHashRIBSD2 struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByBlockHashRIBSD2VinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByBlockHashRIBSD2VoutInner `json:"vout"`
}

// NewListTransactionsByBlockHashRIBSD2 instantiates a new ListTransactionsByBlockHashRIBSD2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSD2(locktime int64, size int32, version int32, vin []ListTransactionsByBlockHashRIBSD2VinInner, vout []ListTransactionsByBlockHashRIBSD2VoutInner) *ListTransactionsByBlockHashRIBSD2 {
	this := ListTransactionsByBlockHashRIBSD2{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHashRIBSD2WithDefaults instantiates a new ListTransactionsByBlockHashRIBSD2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSD2WithDefaults() *ListTransactionsByBlockHashRIBSD2 {
	this := ListTransactionsByBlockHashRIBSD2{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByBlockHashRIBSD2) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByBlockHashRIBSD2) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByBlockHashRIBSD2) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByBlockHashRIBSD2) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByBlockHashRIBSD2) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByBlockHashRIBSD2) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByBlockHashRIBSD2) GetVin() []ListTransactionsByBlockHashRIBSD2VinInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSD2VinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2) GetVinOk() ([]ListTransactionsByBlockHashRIBSD2VinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByBlockHashRIBSD2) SetVin(v []ListTransactionsByBlockHashRIBSD2VinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHashRIBSD2) GetVout() []ListTransactionsByBlockHashRIBSD2VoutInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSD2VoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2) GetVoutOk() ([]ListTransactionsByBlockHashRIBSD2VoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHashRIBSD2) SetVout(v []ListTransactionsByBlockHashRIBSD2VoutInner) {
	o.Vout = v
}

func (o ListTransactionsByBlockHashRIBSD2) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashRIBSD2 struct {
	value *ListTransactionsByBlockHashRIBSD2
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSD2) Get() *ListTransactionsByBlockHashRIBSD2 {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSD2) Set(val *ListTransactionsByBlockHashRIBSD2) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSD2) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSD2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSD2(val *ListTransactionsByBlockHashRIBSD2) *NullableListTransactionsByBlockHashRIBSD2 {
	return &NullableListTransactionsByBlockHashRIBSD2{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSD2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSD2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


