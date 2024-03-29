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

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSL Litecoin
type ListConfirmedTransactionsByAddressAndTimeRangeRIBSL struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the virtual size of this transaction.
	VSize int32 `json:"vSize"`
	// Represents the transaction's version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListConfirmedTransactionsByAddressRIBSLVinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []GetTransactionDetailsByTransactionIDRIBSLVoutInner `json:"vout"`
}

// NewListConfirmedTransactionsByAddressAndTimeRangeRIBSL instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSL(locktime int64, size int32, vSize int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSLVinInner, vout []GetTransactionDetailsByTransactionIDRIBSLVoutInner) *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL {
	this := ListConfirmedTransactionsByAddressAndTimeRangeRIBSL{}
	this.Locktime = locktime
	this.Size = size
	this.VSize = vSize
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListConfirmedTransactionsByAddressAndTimeRangeRIBSLWithDefaults instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSLWithDefaults() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL {
	this := ListConfirmedTransactionsByAddressAndTimeRangeRIBSL{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) SetSize(v int32) {
	o.Size = v
}

// GetVSize returns the VSize field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VSize
}

// GetVSizeOk returns a tuple with the VSize field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VSize, true
}

// SetVSize sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) SetVSize(v int32) {
	o.VSize = v
}

// GetVersion returns the Version field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVin() []ListConfirmedTransactionsByAddressRIBSLVinInner {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSLVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVinOk() ([]ListConfirmedTransactionsByAddressRIBSLVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) SetVin(v []ListConfirmedTransactionsByAddressRIBSLVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVout() []GetTransactionDetailsByTransactionIDRIBSLVoutInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSLVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) GetVoutOk() ([]GetTransactionDetailsByTransactionIDRIBSLVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) SetVout(v []GetTransactionDetailsByTransactionIDRIBSLVoutInner) {
	o.Vout = v
}

func (o ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) MarshalJSON() ([]byte, error) {
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

type NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL struct {
	value *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL) Get() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL) Set(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL {
	return &NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


