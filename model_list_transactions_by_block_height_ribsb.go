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

// ListTransactionsByBlockHeightRIBSB Bitcoin
type ListTransactionsByBlockHeightRIBSB struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the virtual size of this transaction.
	VSize int32 `json:"vSize"`
	// Represents the transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByBlockHashRIBSBVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByBlockHeightRIBSBVout `json:"vout"`
}

// NewListTransactionsByBlockHeightRIBSB instantiates a new ListTransactionsByBlockHeightRIBSB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHeightRIBSB(locktime int64, size int32, vSize int32, version int32, vin []ListTransactionsByBlockHashRIBSBVin, vout []ListTransactionsByBlockHeightRIBSBVout) *ListTransactionsByBlockHeightRIBSB {
	this := ListTransactionsByBlockHeightRIBSB{}
	this.Locktime = locktime
	this.Size = size
	this.VSize = vSize
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHeightRIBSBWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHeightRIBSBWithDefaults() *ListTransactionsByBlockHeightRIBSB {
	this := ListTransactionsByBlockHeightRIBSB{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByBlockHeightRIBSB) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSB) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByBlockHeightRIBSB) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByBlockHeightRIBSB) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSB) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByBlockHeightRIBSB) SetSize(v int32) {
	o.Size = v
}

// GetVSize returns the VSize field value
func (o *ListTransactionsByBlockHeightRIBSB) GetVSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VSize
}

// GetVSizeOk returns a tuple with the VSize field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSB) GetVSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VSize, true
}

// SetVSize sets field value
func (o *ListTransactionsByBlockHeightRIBSB) SetVSize(v int32) {
	o.VSize = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByBlockHeightRIBSB) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSB) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByBlockHeightRIBSB) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByBlockHeightRIBSB) GetVin() []ListTransactionsByBlockHashRIBSBVin {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSBVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSB) GetVinOk() ([]ListTransactionsByBlockHashRIBSBVin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByBlockHeightRIBSB) SetVin(v []ListTransactionsByBlockHashRIBSBVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHeightRIBSB) GetVout() []ListTransactionsByBlockHeightRIBSBVout {
	if o == nil {
		var ret []ListTransactionsByBlockHeightRIBSBVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSB) GetVoutOk() ([]ListTransactionsByBlockHeightRIBSBVout, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHeightRIBSB) SetVout(v []ListTransactionsByBlockHeightRIBSBVout) {
	o.Vout = v
}

func (o ListTransactionsByBlockHeightRIBSB) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHeightRIBSB struct {
	value *ListTransactionsByBlockHeightRIBSB
	isSet bool
}

func (v NullableListTransactionsByBlockHeightRIBSB) Get() *ListTransactionsByBlockHeightRIBSB {
	return v.value
}

func (v *NullableListTransactionsByBlockHeightRIBSB) Set(val *ListTransactionsByBlockHeightRIBSB) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHeightRIBSB) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHeightRIBSB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHeightRIBSB(val *ListTransactionsByBlockHeightRIBSB) *NullableListTransactionsByBlockHeightRIBSB {
	return &NullableListTransactionsByBlockHeightRIBSB{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHeightRIBSB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHeightRIBSB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


