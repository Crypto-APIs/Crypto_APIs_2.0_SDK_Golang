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

// ListConfirmedTransactionsByAddressRIBSBCVinInner struct for ListConfirmedTransactionsByAddressRIBSBCVinInner
type ListConfirmedTransactionsByAddressRIBSBCVinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// Defines the vout of the transaction output, i.e. which output to spend.
	Vout int32 `json:"vout"`
}

// NewListConfirmedTransactionsByAddressRIBSBCVinInner instantiates a new ListConfirmedTransactionsByAddressRIBSBCVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSBCVinInner(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32) *ListConfirmedTransactionsByAddressRIBSBCVinInner {
	this := ListConfirmedTransactionsByAddressRIBSBCVinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	this.Value = value
	this.Vout = vout
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSBCVinInnerWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSBCVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSBCVinInnerWithDefaults() *ListConfirmedTransactionsByAddressRIBSBCVinInner {
	this := ListConfirmedTransactionsByAddressRIBSBCVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetSequenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBCVinInner) SetVout(v int32) {
	o.Vout = v
}

func (o ListConfirmedTransactionsByAddressRIBSBCVinInner) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSBCVinInner struct {
	value *ListConfirmedTransactionsByAddressRIBSBCVinInner
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSBCVinInner) Get() *ListConfirmedTransactionsByAddressRIBSBCVinInner {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBCVinInner) Set(val *ListConfirmedTransactionsByAddressRIBSBCVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSBCVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBCVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSBCVinInner(val *ListConfirmedTransactionsByAddressRIBSBCVinInner) *NullableListConfirmedTransactionsByAddressRIBSBCVinInner {
	return &NullableListConfirmedTransactionsByAddressRIBSBCVinInner{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSBCVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBCVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

