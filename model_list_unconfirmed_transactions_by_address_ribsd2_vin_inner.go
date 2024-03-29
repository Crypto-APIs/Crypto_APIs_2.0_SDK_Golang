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

// ListUnconfirmedTransactionsByAddressRIBSD2VinInner struct for ListUnconfirmedTransactionsByAddressRIBSD2VinInner
type ListUnconfirmedTransactionsByAddressRIBSD2VinInner struct {
	Addresses []string `json:"addresses"`
	ScriptSig ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// String representation of the amount
	Value *string `json:"value,omitempty"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewListUnconfirmedTransactionsByAddressRIBSD2VinInner instantiates a new ListUnconfirmedTransactionsByAddressRIBSD2VinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTransactionsByAddressRIBSD2VinInner(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig, sequence string, txid string, txinwitness []string, vout int32) *ListUnconfirmedTransactionsByAddressRIBSD2VinInner {
	this := ListUnconfirmedTransactionsByAddressRIBSD2VinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	this.Vout = vout
	return &this
}

// NewListUnconfirmedTransactionsByAddressRIBSD2VinInnerWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSD2VinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTransactionsByAddressRIBSD2VinInnerWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSD2VinInner {
	this := ListUnconfirmedTransactionsByAddressRIBSD2VinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetScriptSig() ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetSequenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) SetVout(v int32) {
	o.Vout = v
}

func (o ListUnconfirmedTransactionsByAddressRIBSD2VinInner) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner struct {
	value *ListUnconfirmedTransactionsByAddressRIBSD2VinInner
	isSet bool
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner) Get() *ListUnconfirmedTransactionsByAddressRIBSD2VinInner {
	return v.value
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner) Set(val *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTransactionsByAddressRIBSD2VinInner(val *ListUnconfirmedTransactionsByAddressRIBSD2VinInner) *NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner {
	return &NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner{value: val, isSet: true}
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD2VinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


