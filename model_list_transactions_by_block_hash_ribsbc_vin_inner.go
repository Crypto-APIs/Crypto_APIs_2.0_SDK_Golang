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

// ListTransactionsByBlockHashRIBSBCVinInner struct for ListTransactionsByBlockHashRIBSBCVinInner
type ListTransactionsByBlockHashRIBSBCVinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase string `json:"coinbase"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewListTransactionsByBlockHashRIBSBCVinInner instantiates a new ListTransactionsByBlockHashRIBSBCVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSBCVinInner(addresses []string, coinbase string, scriptSig GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32) *ListTransactionsByBlockHashRIBSBCVinInner {
	this := ListTransactionsByBlockHashRIBSBCVinInner{}
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

// NewListTransactionsByBlockHashRIBSBCVinInnerWithDefaults instantiates a new ListTransactionsByBlockHashRIBSBCVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSBCVinInnerWithDefaults() *ListTransactionsByBlockHashRIBSBCVinInner {
	this := ListTransactionsByBlockHashRIBSBCVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetCoinbase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coinbase, true
}

// SetCoinbase sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetCoinbase(v string) {
	o.Coinbase = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetSequenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBCVinInner) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHashRIBSBCVinInner) SetVout(v int32) {
	o.Vout = v
}

func (o ListTransactionsByBlockHashRIBSBCVinInner) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashRIBSBCVinInner struct {
	value *ListTransactionsByBlockHashRIBSBCVinInner
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSBCVinInner) Get() *ListTransactionsByBlockHashRIBSBCVinInner {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSBCVinInner) Set(val *ListTransactionsByBlockHashRIBSBCVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSBCVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSBCVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSBCVinInner(val *ListTransactionsByBlockHashRIBSBCVinInner) *NullableListTransactionsByBlockHashRIBSBCVinInner {
	return &NullableListTransactionsByBlockHashRIBSBCVinInner{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSBCVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSBCVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


