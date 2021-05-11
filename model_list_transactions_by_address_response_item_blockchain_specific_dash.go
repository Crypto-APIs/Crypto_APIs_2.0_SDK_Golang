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

// ListTransactionsByAddressResponseItemBlockchainSpecificDash Dash
type ListTransactionsByAddressResponseItemBlockchainSpecificDash struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the transaction's version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout `json:"vout"`
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificDash instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByAddressResponseItemBlockchainSpecificDash(locktime int32, size int32, version int32, vin []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin, vout []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout) *ListTransactionsByAddressResponseItemBlockchainSpecificDash {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificDash{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificDashWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByAddressResponseItemBlockchainSpecificDashWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificDash {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificDash{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVin() []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin {
	if o == nil {
		var ret []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVinOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetVin(v []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVout() []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout {
	if o == nil {
		var ret []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVoutOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetVout(v []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout) {
	o.Vout = v
}

func (o ListTransactionsByAddressResponseItemBlockchainSpecificDash) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByAddressResponseItemBlockchainSpecificDash struct {
	value *ListTransactionsByAddressResponseItemBlockchainSpecificDash
	isSet bool
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificDash) Get() *ListTransactionsByAddressResponseItemBlockchainSpecificDash {
	return v.value
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificDash) Set(val *ListTransactionsByAddressResponseItemBlockchainSpecificDash) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificDash) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificDash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByAddressResponseItemBlockchainSpecificDash(val *ListTransactionsByAddressResponseItemBlockchainSpecificDash) *NullableListTransactionsByAddressResponseItemBlockchainSpecificDash {
	return &NullableListTransactionsByAddressResponseItemBlockchainSpecificDash{value: val, isSet: true}
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificDash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificDash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


