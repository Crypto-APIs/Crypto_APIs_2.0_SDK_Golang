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

// GetWalletTransactionDetailsByTransactionIDRIBSD2 Dash
type GetWalletTransactionDetailsByTransactionIDRIBSD2 struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the transaction version number.
	Version int32 `json:"version"`
	// Object Array representation of transaction inputs
	Vin []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin `json:"vin"`
	// Object Array representation of transaction outputs
	Vout []GetTransactionDetailsByTransactionIDRIBSD2Vout `json:"vout"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSD2 instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBSD2(locktime int32, size int32, version int32, vin []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin, vout []GetTransactionDetailsByTransactionIDRIBSD2Vout) *GetWalletTransactionDetailsByTransactionIDRIBSD2 {
	this := GetWalletTransactionDetailsByTransactionIDRIBSD2{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSD2WithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSD2WithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSD2 {
	this := GetWalletTransactionDetailsByTransactionIDRIBSD2{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin {
	if o == nil {
		var ret []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVinOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSD2Vin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVout() []GetTransactionDetailsByTransactionIDRIBSD2Vout {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSD2Vout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSD2Vout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetVout(v []GetTransactionDetailsByTransactionIDRIBSD2Vout) {
	o.Vout = v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBSD2) MarshalJSON() ([]byte, error) {
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

type NullableGetWalletTransactionDetailsByTransactionIDRIBSD2 struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBSD2
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSD2) Get() *GetWalletTransactionDetailsByTransactionIDRIBSD2 {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2) Set(val *GetWalletTransactionDetailsByTransactionIDRIBSD2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSD2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBSD2(val *GetWalletTransactionDetailsByTransactionIDRIBSD2) *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2 {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBSD2{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSD2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

