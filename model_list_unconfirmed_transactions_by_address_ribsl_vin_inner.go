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

// ListUnconfirmedTransactionsByAddressRIBSLVinInner struct for ListUnconfirmedTransactionsByAddressRIBSLVinInner
type ListUnconfirmedTransactionsByAddressRIBSLVinInner struct {
	Addresses []string `json:"addresses"`
	ScriptSig ListUnconfirmedTransactionsByAddressRIBSLVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// Defines the vout of the transaction output, i.e. which output to spend.
	Vout *int32 `json:"vout,omitempty"`
}

// NewListUnconfirmedTransactionsByAddressRIBSLVinInner instantiates a new ListUnconfirmedTransactionsByAddressRIBSLVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTransactionsByAddressRIBSLVinInner(addresses []string, scriptSig ListUnconfirmedTransactionsByAddressRIBSLVinInnerScriptSig, sequence string, txid string, txinwitness []string, value string) *ListUnconfirmedTransactionsByAddressRIBSLVinInner {
	this := ListUnconfirmedTransactionsByAddressRIBSLVinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	this.Value = value
	return &this
}

// NewListUnconfirmedTransactionsByAddressRIBSLVinInnerWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSLVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTransactionsByAddressRIBSLVinInnerWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSLVinInner {
	this := ListUnconfirmedTransactionsByAddressRIBSLVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetScriptSig() ListUnconfirmedTransactionsByAddressRIBSLVinInnerScriptSig {
	if o == nil {
		var ret ListUnconfirmedTransactionsByAddressRIBSLVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetScriptSigOk() (*ListUnconfirmedTransactionsByAddressRIBSLVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) SetScriptSig(v ListUnconfirmedTransactionsByAddressRIBSLVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetSequenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *ListUnconfirmedTransactionsByAddressRIBSLVinInner) SetVout(v int32) {
	o.Vout = &v
}

func (o ListUnconfirmedTransactionsByAddressRIBSLVinInner) MarshalJSON() ([]byte, error) {
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
	if o.Vout != nil {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListUnconfirmedTransactionsByAddressRIBSLVinInner struct {
	value *ListUnconfirmedTransactionsByAddressRIBSLVinInner
	isSet bool
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSLVinInner) Get() *ListUnconfirmedTransactionsByAddressRIBSLVinInner {
	return v.value
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSLVinInner) Set(val *ListUnconfirmedTransactionsByAddressRIBSLVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSLVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSLVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTransactionsByAddressRIBSLVinInner(val *ListUnconfirmedTransactionsByAddressRIBSLVinInner) *NullableListUnconfirmedTransactionsByAddressRIBSLVinInner {
	return &NullableListUnconfirmedTransactionsByAddressRIBSLVinInner{value: val, isSet: true}
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSLVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSLVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


