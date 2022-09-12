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

// GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner struct for GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner
type GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSD2VinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int64 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness,omitempty"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout *int32 `json:"vout,omitempty"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSD2VinInnerScriptSig, sequence int64) *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner {
	this := GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSD2VinInnerWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSD2VinInnerWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner {
	this := GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSD2VinInnerScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSD2VinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSD2VinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSD2VinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetSequenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetSequence(v int64) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetTxinwitness() []string {
	if o == nil || o.Txinwitness == nil {
		var ret []string
		return ret
	}
	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil || o.Txinwitness == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// HasTxinwitness returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) HasTxinwitness() bool {
	if o != nil && o.Txinwitness != nil {
		return true
	}

	return false
}

// SetTxinwitness gets a reference to the given []string and assigns it to the Txinwitness field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) SetVout(v int32) {
	o.Vout = &v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) MarshalJSON() ([]byte, error) {
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

type NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) Get() *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) Set(val *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner(val *GetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSD2VinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


