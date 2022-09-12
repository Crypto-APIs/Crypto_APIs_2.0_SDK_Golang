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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSBC Bitcoin Cash
type GetTransactionDetailsByTransactionIDFromCallbackRIBSBC struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []GetTransactionDetailsByTransactionIDRIBSBCVinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner `json:"vout"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBC instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSBC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBC(locktime int64, size int32, version int32, vin []GetTransactionDetailsByTransactionIDRIBSBCVinInner, vout []GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSBC{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBCWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSBC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBCWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSBC{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetVin() []GetTransactionDetailsByTransactionIDRIBSBCVinInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSBCVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetVinOk() ([]GetTransactionDetailsByTransactionIDRIBSBCVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) SetVin(v []GetTransactionDetailsByTransactionIDRIBSBCVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetVout() []GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) GetVoutOk() ([]GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) SetVout(v []GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) {
	o.Vout = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSBC) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


