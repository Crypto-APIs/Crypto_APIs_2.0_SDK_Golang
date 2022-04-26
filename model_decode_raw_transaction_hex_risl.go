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

// DecodeRawTransactionHexRISL Litecoin
type DecodeRawTransactionHexRISL struct {
	// Represents the time at which a particular transaction can be added to the blockchain
	Locktime int32 `json:"locktime"`
	// Represents the same as transactionId for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols hash is different from transactionId for SegWit transactions.
	TransactionHash string `json:"transactionHash"`
	// Represents the virtual size of this transaction.
	VSize int32 `json:"vSize"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []DecodeRawTransactionHexRISLVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []DecodeRawTransactionHexRISLVout `json:"vout"`
	// Represents the size of a block, measured in weight units and including the segwit discount.
	Weight *int32 `json:"weight,omitempty"`
}

// NewDecodeRawTransactionHexRISL instantiates a new DecodeRawTransactionHexRISL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISL(locktime int32, transactionHash string, vSize int32, version int32, vin []DecodeRawTransactionHexRISLVin, vout []DecodeRawTransactionHexRISLVout) *DecodeRawTransactionHexRISL {
	this := DecodeRawTransactionHexRISL{}
	this.Locktime = locktime
	this.TransactionHash = transactionHash
	this.VSize = vSize
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewDecodeRawTransactionHexRISLWithDefaults instantiates a new DecodeRawTransactionHexRISL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISLWithDefaults() *DecodeRawTransactionHexRISL {
	this := DecodeRawTransactionHexRISL{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *DecodeRawTransactionHexRISL) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISL) GetLocktimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *DecodeRawTransactionHexRISL) SetLocktime(v int32) {
	o.Locktime = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *DecodeRawTransactionHexRISL) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISL) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *DecodeRawTransactionHexRISL) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetVSize returns the VSize field value
func (o *DecodeRawTransactionHexRISL) GetVSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VSize
}

// GetVSizeOk returns a tuple with the VSize field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISL) GetVSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VSize, true
}

// SetVSize sets field value
func (o *DecodeRawTransactionHexRISL) SetVSize(v int32) {
	o.VSize = v
}

// GetVersion returns the Version field value
func (o *DecodeRawTransactionHexRISL) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISL) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DecodeRawTransactionHexRISL) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *DecodeRawTransactionHexRISL) GetVin() []DecodeRawTransactionHexRISLVin {
	if o == nil {
		var ret []DecodeRawTransactionHexRISLVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISL) GetVinOk() ([]DecodeRawTransactionHexRISLVin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *DecodeRawTransactionHexRISL) SetVin(v []DecodeRawTransactionHexRISLVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *DecodeRawTransactionHexRISL) GetVout() []DecodeRawTransactionHexRISLVout {
	if o == nil {
		var ret []DecodeRawTransactionHexRISLVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISL) GetVoutOk() ([]DecodeRawTransactionHexRISLVout, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *DecodeRawTransactionHexRISL) SetVout(v []DecodeRawTransactionHexRISLVout) {
	o.Vout = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISL) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISL) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISL) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *DecodeRawTransactionHexRISL) SetWeight(v int32) {
	o.Weight = &v
}

func (o DecodeRawTransactionHexRISL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locktime"] = o.Locktime
	}
	if true {
		toSerialize["transactionHash"] = o.TransactionHash
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
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISL struct {
	value *DecodeRawTransactionHexRISL
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISL) Get() *DecodeRawTransactionHexRISL {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISL) Set(val *DecodeRawTransactionHexRISL) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISL) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISL(val *DecodeRawTransactionHexRISL) *NullableDecodeRawTransactionHexRISL {
	return &NullableDecodeRawTransactionHexRISL{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


