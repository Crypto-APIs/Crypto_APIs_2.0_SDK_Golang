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

// GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash Bitcoin Cash
type GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVout `json:"vout"`
}

// NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash(locktime int32, size int32, version int32, vin []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVin, vout []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVout) *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash {
	this := GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash {
	this := GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetVin() []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVin {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetVinOk() (*[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) SetVin(v []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetVout() []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVout {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) SetVout(v []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCashVout) {
	o.Vout = v
}

func (o GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash struct {
	value *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) Get() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) Set(val *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash(val *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) *NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash {
	return &NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinCash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

