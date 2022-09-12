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

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSB Bitcoin
type ListConfirmedTransactionsByAddressAndTimeRangeRIBSB struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Defines the transaction's virtual size.
	VSize int32 `json:"vSize"`
	// Defines the version of the transaction.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListConfirmedTransactionsByAddressRIBSBVinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListConfirmedTransactionsByAddressRIBSBVoutInner `json:"vout"`
}

// NewListConfirmedTransactionsByAddressAndTimeRangeRIBSB instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSB(locktime int64, size int32, vSize int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSBVinInner, vout []ListConfirmedTransactionsByAddressRIBSBVoutInner) *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB {
	this := ListConfirmedTransactionsByAddressAndTimeRangeRIBSB{}
	this.Locktime = locktime
	this.Size = size
	this.VSize = vSize
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBWithDefaults instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBWithDefaults() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB {
	this := ListConfirmedTransactionsByAddressAndTimeRangeRIBSB{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) SetSize(v int32) {
	o.Size = v
}

// GetVSize returns the VSize field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VSize
}

// GetVSizeOk returns a tuple with the VSize field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VSize, true
}

// SetVSize sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) SetVSize(v int32) {
	o.VSize = v
}

// GetVersion returns the Version field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVin() []ListConfirmedTransactionsByAddressRIBSBVinInner {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSBVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVinOk() ([]ListConfirmedTransactionsByAddressRIBSBVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) SetVin(v []ListConfirmedTransactionsByAddressRIBSBVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVout() []ListConfirmedTransactionsByAddressRIBSBVoutInner {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSBVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) GetVoutOk() ([]ListConfirmedTransactionsByAddressRIBSBVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) SetVout(v []ListConfirmedTransactionsByAddressRIBSBVoutInner) {
	o.Vout = v
}

func (o ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) MarshalJSON() ([]byte, error) {
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

type NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB struct {
	value *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB) Get() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB) Set(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB {
	return &NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


