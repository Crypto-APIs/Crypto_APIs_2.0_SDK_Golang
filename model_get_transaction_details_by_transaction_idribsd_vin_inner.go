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

// GetTransactionDetailsByTransactionIDRIBSDVinInner struct for GetTransactionDetailsByTransactionIDRIBSDVinInner
type GetTransactionDetailsByTransactionIDRIBSDVinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int64 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout *int32 `json:"vout,omitempty"`
}

// NewGetTransactionDetailsByTransactionIDRIBSDVinInner instantiates a new GetTransactionDetailsByTransactionIDRIBSDVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSDVinInner(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig, sequence int64, txinwitness []string, value string) *GetTransactionDetailsByTransactionIDRIBSDVinInner {
	this := GetTransactionDetailsByTransactionIDRIBSDVinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txinwitness = txinwitness
	this.Value = value
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSDVinInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSDVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSDVinInnerWithDefaults() *GetTransactionDetailsByTransactionIDRIBSDVinInner {
	this := GetTransactionDetailsByTransactionIDRIBSDVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetSequenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetSequence(v int64) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *GetTransactionDetailsByTransactionIDRIBSDVinInner) SetVout(v int32) {
	o.Vout = &v
}

func (o GetTransactionDetailsByTransactionIDRIBSDVinInner) MarshalJSON() ([]byte, error) {
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
	if o.Txid != nil {
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

type NullableGetTransactionDetailsByTransactionIDRIBSDVinInner struct {
	value *GetTransactionDetailsByTransactionIDRIBSDVinInner
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSDVinInner) Get() *GetTransactionDetailsByTransactionIDRIBSDVinInner {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSDVinInner) Set(val *GetTransactionDetailsByTransactionIDRIBSDVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSDVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSDVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSDVinInner(val *GetTransactionDetailsByTransactionIDRIBSDVinInner) *NullableGetTransactionDetailsByTransactionIDRIBSDVinInner {
	return &NullableGetTransactionDetailsByTransactionIDRIBSDVinInner{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSDVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSDVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


