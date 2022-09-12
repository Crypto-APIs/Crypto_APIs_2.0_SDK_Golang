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

// ListTransactionsByBlockHashRIBSB Bitcoin
type ListTransactionsByBlockHashRIBSB struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the virtual size of this transaction.
	VSize int32 `json:"vSize"`
	// Represents the transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByBlockHashRIBSBVinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByBlockHashRIBSBVoutInner `json:"vout"`
}

// NewListTransactionsByBlockHashRIBSB instantiates a new ListTransactionsByBlockHashRIBSB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSB(locktime int64, size int32, vSize int32, version int32, vin []ListTransactionsByBlockHashRIBSBVinInner, vout []ListTransactionsByBlockHashRIBSBVoutInner) *ListTransactionsByBlockHashRIBSB {
	this := ListTransactionsByBlockHashRIBSB{}
	this.Locktime = locktime
	this.Size = size
	this.VSize = vSize
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHashRIBSBWithDefaults instantiates a new ListTransactionsByBlockHashRIBSB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSBWithDefaults() *ListTransactionsByBlockHashRIBSB {
	this := ListTransactionsByBlockHashRIBSB{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByBlockHashRIBSB) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSB) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByBlockHashRIBSB) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByBlockHashRIBSB) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSB) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByBlockHashRIBSB) SetSize(v int32) {
	o.Size = v
}

// GetVSize returns the VSize field value
func (o *ListTransactionsByBlockHashRIBSB) GetVSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VSize
}

// GetVSizeOk returns a tuple with the VSize field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSB) GetVSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VSize, true
}

// SetVSize sets field value
func (o *ListTransactionsByBlockHashRIBSB) SetVSize(v int32) {
	o.VSize = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByBlockHashRIBSB) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSB) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByBlockHashRIBSB) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByBlockHashRIBSB) GetVin() []ListTransactionsByBlockHashRIBSBVinInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSBVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSB) GetVinOk() ([]ListTransactionsByBlockHashRIBSBVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByBlockHashRIBSB) SetVin(v []ListTransactionsByBlockHashRIBSBVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHashRIBSB) GetVout() []ListTransactionsByBlockHashRIBSBVoutInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSBVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSB) GetVoutOk() ([]ListTransactionsByBlockHashRIBSBVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHashRIBSB) SetVout(v []ListTransactionsByBlockHashRIBSBVoutInner) {
	o.Vout = v
}

func (o ListTransactionsByBlockHashRIBSB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locktime"] = o.Locktime
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["vSize"] = o.VSize
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

type NullableListTransactionsByBlockHashRIBSB struct {
	value *ListTransactionsByBlockHashRIBSB
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSB) Get() *ListTransactionsByBlockHashRIBSB {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSB) Set(val *ListTransactionsByBlockHashRIBSB) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSB) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSB(val *ListTransactionsByBlockHashRIBSB) *NullableListTransactionsByBlockHashRIBSB {
	return &NullableListTransactionsByBlockHashRIBSB{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


