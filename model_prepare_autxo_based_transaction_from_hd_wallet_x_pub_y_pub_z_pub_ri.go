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

// PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI struct for PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI
type PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI struct {
	// Representation of the additional data
	AdditionalData *string `json:"additionalData,omitempty"`
	// When isConfirmed is True - Defines the amount of the transaction fee When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value.
	Fee string `json:"fee"`
	// Defines the fee per byte value
	FeePerByte *string `json:"feePerByte,omitempty"`
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int64 `json:"locktime"`
	// Representation of whether the transaction is replaceable
	Replaceable bool `json:"replaceable"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the transaction inputs.
	Vin []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner `json:"vin"`
	// Represents the transaction outputs.
	Vout []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner `json:"vout"`
}

// NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI(fee string, locktime int64, replaceable bool, size int32, vin []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner, vout []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI {
	this := PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI{}
	this.Fee = fee
	this.Locktime = locktime
	this.Replaceable = replaceable
	this.Size = size
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI {
	this := PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetAdditionalData() string {
	if o == nil || o.AdditionalData == nil {
		var ret string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetAdditionalDataOk() (*string, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given string and assigns it to the AdditionalData field.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetAdditionalData(v string) {
	o.AdditionalData = &v
}

// GetFee returns the Fee field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetFee(v string) {
	o.Fee = v
}

// GetFeePerByte returns the FeePerByte field value if set, zero value otherwise.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFeePerByte() string {
	if o == nil || o.FeePerByte == nil {
		var ret string
		return ret
	}
	return *o.FeePerByte
}

// GetFeePerByteOk returns a tuple with the FeePerByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFeePerByteOk() (*string, bool) {
	if o == nil || o.FeePerByte == nil {
		return nil, false
	}
	return o.FeePerByte, true
}

// HasFeePerByte returns a boolean if a field has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) HasFeePerByte() bool {
	if o != nil && o.FeePerByte != nil {
		return true
	}

	return false
}

// SetFeePerByte gets a reference to the given string and assigns it to the FeePerByte field.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetFeePerByte(v string) {
	o.FeePerByte = &v
}

// GetLocktime returns the Locktime field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetLocktime(v int64) {
	o.Locktime = v
}

// GetReplaceable returns the Replaceable field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetReplaceable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Replaceable
}

// GetReplaceableOk returns a tuple with the Replaceable field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetReplaceableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replaceable, true
}

// SetReplaceable sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetReplaceable(v bool) {
	o.Replaceable = v
}

// GetSize returns the Size field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetSize(v int32) {
	o.Size = v
}

// GetVin returns the Vin field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVin() []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner {
	if o == nil {
		var ret []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVinOk() ([]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetVin(v []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVout() []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner {
	if o == nil {
		var ret []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVoutOk() ([]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetVout(v []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner) {
	o.Vout = v
}

func (o PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if o.FeePerByte != nil {
		toSerialize["feePerByte"] = o.FeePerByte
	}
	if true {
		toSerialize["locktime"] = o.Locktime
	}
	if true {
		toSerialize["replaceable"] = o.Replaceable
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["vin"] = o.Vin
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI struct {
	value *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI
	isSet bool
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) Get() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI {
	return v.value
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) Set(val *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI(val *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI {
	return &NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI{value: val, isSet: true}
}

func (v NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


