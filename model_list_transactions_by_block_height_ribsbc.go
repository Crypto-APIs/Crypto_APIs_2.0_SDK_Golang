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

// ListTransactionsByBlockHeightRIBSBC Bitcoin Cash
type ListTransactionsByBlockHeightRIBSBC struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the total size of this transaction.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByBlockHashRIBSBCVinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByBlockHashRIBSBCVoutInner `json:"vout"`
}

// NewListTransactionsByBlockHeightRIBSBC instantiates a new ListTransactionsByBlockHeightRIBSBC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHeightRIBSBC(locktime int64, size int32, version int32, vin []ListTransactionsByBlockHashRIBSBCVinInner, vout []ListTransactionsByBlockHashRIBSBCVoutInner) *ListTransactionsByBlockHeightRIBSBC {
	this := ListTransactionsByBlockHeightRIBSBC{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHeightRIBSBCWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSBC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHeightRIBSBCWithDefaults() *ListTransactionsByBlockHeightRIBSBC {
	this := ListTransactionsByBlockHeightRIBSBC{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByBlockHeightRIBSBC) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSBC) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByBlockHeightRIBSBC) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByBlockHeightRIBSBC) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSBC) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByBlockHeightRIBSBC) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByBlockHeightRIBSBC) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSBC) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByBlockHeightRIBSBC) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByBlockHeightRIBSBC) GetVin() []ListTransactionsByBlockHashRIBSBCVinInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSBCVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSBC) GetVinOk() ([]ListTransactionsByBlockHashRIBSBCVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByBlockHeightRIBSBC) SetVin(v []ListTransactionsByBlockHashRIBSBCVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHeightRIBSBC) GetVout() []ListTransactionsByBlockHashRIBSBCVoutInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSBCVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSBC) GetVoutOk() ([]ListTransactionsByBlockHashRIBSBCVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHeightRIBSBC) SetVout(v []ListTransactionsByBlockHashRIBSBCVoutInner) {
	o.Vout = v
}

func (o ListTransactionsByBlockHeightRIBSBC) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHeightRIBSBC struct {
	value *ListTransactionsByBlockHeightRIBSBC
	isSet bool
}

func (v NullableListTransactionsByBlockHeightRIBSBC) Get() *ListTransactionsByBlockHeightRIBSBC {
	return v.value
}

func (v *NullableListTransactionsByBlockHeightRIBSBC) Set(val *ListTransactionsByBlockHeightRIBSBC) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHeightRIBSBC) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHeightRIBSBC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHeightRIBSBC(val *ListTransactionsByBlockHeightRIBSBC) *NullableListTransactionsByBlockHeightRIBSBC {
	return &NullableListTransactionsByBlockHeightRIBSBC{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHeightRIBSBC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHeightRIBSBC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


