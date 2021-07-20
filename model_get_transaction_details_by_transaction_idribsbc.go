/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// GetTransactionDetailsByTransactionIDRIBSBC Bitcoin Cash
type GetTransactionDetailsByTransactionIDRIBSBC struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []GetTransactionDetailsByTransactionIDRIBSBCVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []GetTransactionDetailsByTransactionIDRIBSBCVout `json:"vout"`
}

// NewGetTransactionDetailsByTransactionIDRIBSBC instantiates a new GetTransactionDetailsByTransactionIDRIBSBC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSBC(locktime int32, size int32, version int32, vin []GetTransactionDetailsByTransactionIDRIBSBCVin, vout []GetTransactionDetailsByTransactionIDRIBSBCVout) *GetTransactionDetailsByTransactionIDRIBSBC {
	this := GetTransactionDetailsByTransactionIDRIBSBC{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSBCWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSBCWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBC {
	this := GetTransactionDetailsByTransactionIDRIBSBC{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetVin() []GetTransactionDetailsByTransactionIDRIBSBCVin {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSBCVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetVinOk() (*[]GetTransactionDetailsByTransactionIDRIBSBCVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) SetVin(v []GetTransactionDetailsByTransactionIDRIBSBCVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetVout() []GetTransactionDetailsByTransactionIDRIBSBCVout {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSBCVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBC) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSBCVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBC) SetVout(v []GetTransactionDetailsByTransactionIDRIBSBCVout) {
	o.Vout = v
}

func (o GetTransactionDetailsByTransactionIDRIBSBC) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDRIBSBC struct {
	value *GetTransactionDetailsByTransactionIDRIBSBC
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBC) Get() *GetTransactionDetailsByTransactionIDRIBSBC {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBC) Set(val *GetTransactionDetailsByTransactionIDRIBSBC) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBC) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSBC(val *GetTransactionDetailsByTransactionIDRIBSBC) *NullableGetTransactionDetailsByTransactionIDRIBSBC {
	return &NullableGetTransactionDetailsByTransactionIDRIBSBC{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

