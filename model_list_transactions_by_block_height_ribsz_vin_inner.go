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

// ListTransactionsByBlockHeightRIBSZVinInner struct for ListTransactionsByBlockHeightRIBSZVinInner
type ListTransactionsByBlockHeightRIBSZVinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase string `json:"coinbase"`
	ScriptSig ListTransactionsByBlockHeightRIBSZVinInnerScriptSig `json:"scriptSig"`
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

// NewListTransactionsByBlockHeightRIBSZVinInner instantiates a new ListTransactionsByBlockHeightRIBSZVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHeightRIBSZVinInner(addresses []string, coinbase string, scriptSig ListTransactionsByBlockHeightRIBSZVinInnerScriptSig, sequence int64, txid string, txinwitness []string, value string, vout int32) *ListTransactionsByBlockHeightRIBSZVinInner {
	this := ListTransactionsByBlockHeightRIBSZVinInner{}
	this.Addresses = addresses
	this.Coinbase = coinbase
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	this.Value = value
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHeightRIBSZVinInnerWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSZVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHeightRIBSZVinInnerWithDefaults() *ListTransactionsByBlockHeightRIBSZVinInner {
	this := ListTransactionsByBlockHeightRIBSZVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetCoinbase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coinbase, true
}

// SetCoinbase sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetCoinbase(v string) {
	o.Coinbase = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetScriptSig() ListTransactionsByBlockHeightRIBSZVinInnerScriptSig {
	if o == nil {
		var ret ListTransactionsByBlockHeightRIBSZVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetScriptSigOk() (*ListTransactionsByBlockHeightRIBSZVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetScriptSig(v ListTransactionsByBlockHeightRIBSZVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetSequenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetSequence(v int64) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSZVinInner) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHeightRIBSZVinInner) SetVout(v int32) {
	o.Vout = v
}

func (o ListTransactionsByBlockHeightRIBSZVinInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if true {
		toSerialize["coinbase"] = o.Coinbase
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

type NullableListTransactionsByBlockHeightRIBSZVinInner struct {
	value *ListTransactionsByBlockHeightRIBSZVinInner
	isSet bool
}

func (v NullableListTransactionsByBlockHeightRIBSZVinInner) Get() *ListTransactionsByBlockHeightRIBSZVinInner {
	return v.value
}

func (v *NullableListTransactionsByBlockHeightRIBSZVinInner) Set(val *ListTransactionsByBlockHeightRIBSZVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHeightRIBSZVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHeightRIBSZVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHeightRIBSZVinInner(val *ListTransactionsByBlockHeightRIBSZVinInner) *NullableListTransactionsByBlockHeightRIBSZVinInner {
	return &NullableListTransactionsByBlockHeightRIBSZVinInner{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHeightRIBSZVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHeightRIBSZVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


