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

// ListConfirmedTransactionsByAddressRIBSLVinInner struct for ListConfirmedTransactionsByAddressRIBSLVinInner
type ListConfirmedTransactionsByAddressRIBSLVinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig ListConfirmedTransactionsByAddressRIBSLVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// Defines the vout of the transaction output, i.e. which output to spend.
	Vout *int32 `json:"vout,omitempty"`
}

// NewListConfirmedTransactionsByAddressRIBSLVinInner instantiates a new ListConfirmedTransactionsByAddressRIBSLVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSLVinInner(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSLVinInnerScriptSig, sequence string, txid string, txinwitness []string) *ListConfirmedTransactionsByAddressRIBSLVinInner {
	this := ListConfirmedTransactionsByAddressRIBSLVinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSLVinInnerWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSLVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSLVinInnerWithDefaults() *ListConfirmedTransactionsByAddressRIBSLVinInner {
	this := ListConfirmedTransactionsByAddressRIBSLVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetScriptSig() ListConfirmedTransactionsByAddressRIBSLVinInnerScriptSig {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIBSLVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSLVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSLVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetSequenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *ListConfirmedTransactionsByAddressRIBSLVinInner) SetVout(v int32) {
	o.Vout = &v
}

func (o ListConfirmedTransactionsByAddressRIBSLVinInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Coinbase != nil {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Vout != nil {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSLVinInner struct {
	value *ListConfirmedTransactionsByAddressRIBSLVinInner
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSLVinInner) Get() *ListConfirmedTransactionsByAddressRIBSLVinInner {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSLVinInner) Set(val *ListConfirmedTransactionsByAddressRIBSLVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSLVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSLVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSLVinInner(val *ListConfirmedTransactionsByAddressRIBSLVinInner) *NullableListConfirmedTransactionsByAddressRIBSLVinInner {
	return &NullableListConfirmedTransactionsByAddressRIBSLVinInner{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSLVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSLVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


