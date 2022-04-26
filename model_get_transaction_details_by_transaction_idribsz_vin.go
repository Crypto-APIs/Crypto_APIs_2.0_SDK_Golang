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

// GetTransactionDetailsByTransactionIDRIBSZVin struct for GetTransactionDetailsByTransactionIDRIBSZVin
type GetTransactionDetailsByTransactionIDRIBSZVin struct {
	Addresses []string `json:"addresses"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSZScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int64 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// Defines the specific amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewGetTransactionDetailsByTransactionIDRIBSZVin instantiates a new GetTransactionDetailsByTransactionIDRIBSZVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSZVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSZScriptSig, sequence int64, txid string, txinwitness []string, value string, vout int32) *GetTransactionDetailsByTransactionIDRIBSZVin {
	this := GetTransactionDetailsByTransactionIDRIBSZVin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	this.Value = value
	this.Vout = vout
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSZVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSZVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSZVinWithDefaults() *GetTransactionDetailsByTransactionIDRIBSZVin {
	this := GetTransactionDetailsByTransactionIDRIBSZVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSZScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSZScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSZScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSZScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetSequenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetSequence(v int64) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetVout(v int32) {
	o.Vout = v
}

func (o GetTransactionDetailsByTransactionIDRIBSZVin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if true {
		toSerialize["scriptSig"] = o.ScriptSig
	}
	if true {
		toSerialize["sequence"] = o.Sequence
	}
	if true {
		toSerialize["txid"] = o.Txid
	}
	if true {
		toSerialize["txinwitness"] = o.Txinwitness
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDRIBSZVin struct {
	value *GetTransactionDetailsByTransactionIDRIBSZVin
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSZVin) Get() *GetTransactionDetailsByTransactionIDRIBSZVin {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSZVin) Set(val *GetTransactionDetailsByTransactionIDRIBSZVin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSZVin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSZVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSZVin(val *GetTransactionDetailsByTransactionIDRIBSZVin) *NullableGetTransactionDetailsByTransactionIDRIBSZVin {
	return &NullableGetTransactionDetailsByTransactionIDRIBSZVin{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSZVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSZVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


