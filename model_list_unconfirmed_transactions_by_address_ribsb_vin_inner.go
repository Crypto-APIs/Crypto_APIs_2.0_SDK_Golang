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

// ListUnconfirmedTransactionsByAddressRIBSBVinInner struct for ListUnconfirmedTransactionsByAddressRIBSBVinInner
type ListUnconfirmedTransactionsByAddressRIBSBVinInner struct {
	Addresses []string `json:"addresses"`
	ScriptSig ListUnconfirmedTransactionsByAddressRIBSBVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness,omitempty"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// Defines the vout of the transaction output, i.e. which output to spend.
	Vout *int32 `json:"vout,omitempty"`
}

// NewListUnconfirmedTransactionsByAddressRIBSBVinInner instantiates a new ListUnconfirmedTransactionsByAddressRIBSBVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTransactionsByAddressRIBSBVinInner(addresses []string, scriptSig ListUnconfirmedTransactionsByAddressRIBSBVinInnerScriptSig, sequence string) *ListUnconfirmedTransactionsByAddressRIBSBVinInner {
	this := ListUnconfirmedTransactionsByAddressRIBSBVinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	return &this
}

// NewListUnconfirmedTransactionsByAddressRIBSBVinInnerWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSBVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTransactionsByAddressRIBSBVinInnerWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSBVinInner {
	this := ListUnconfirmedTransactionsByAddressRIBSBVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetScriptSig() ListUnconfirmedTransactionsByAddressRIBSBVinInnerScriptSig {
	if o == nil {
		var ret ListUnconfirmedTransactionsByAddressRIBSBVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetScriptSigOk() (*ListUnconfirmedTransactionsByAddressRIBSBVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) SetScriptSig(v ListUnconfirmedTransactionsByAddressRIBSBVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetSequenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value if set, zero value otherwise.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetTxinwitness() []string {
	if o == nil || o.Txinwitness == nil {
		var ret []string
		return ret
	}
	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil || o.Txinwitness == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// HasTxinwitness returns a boolean if a field has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) HasTxinwitness() bool {
	if o != nil && o.Txinwitness != nil {
		return true
	}

	return false
}

// SetTxinwitness gets a reference to the given []string and assigns it to the Txinwitness field.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *ListUnconfirmedTransactionsByAddressRIBSBVinInner) SetVout(v int32) {
	o.Vout = &v
}

func (o ListUnconfirmedTransactionsByAddressRIBSBVinInner) MarshalJSON() ([]byte, error) {
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
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.Txinwitness != nil {
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

type NullableListUnconfirmedTransactionsByAddressRIBSBVinInner struct {
	value *ListUnconfirmedTransactionsByAddressRIBSBVinInner
	isSet bool
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBVinInner) Get() *ListUnconfirmedTransactionsByAddressRIBSBVinInner {
	return v.value
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBVinInner) Set(val *ListUnconfirmedTransactionsByAddressRIBSBVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTransactionsByAddressRIBSBVinInner(val *ListUnconfirmedTransactionsByAddressRIBSBVinInner) *NullableListUnconfirmedTransactionsByAddressRIBSBVinInner {
	return &NullableListUnconfirmedTransactionsByAddressRIBSBVinInner{value: val, isSet: true}
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


