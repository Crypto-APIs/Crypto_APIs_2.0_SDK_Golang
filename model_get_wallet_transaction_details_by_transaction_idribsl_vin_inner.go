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

// GetWalletTransactionDetailsByTransactionIDRIBSLVinInner struct for GetWalletTransactionDetailsByTransactionIDRIBSLVinInner
type GetWalletTransactionDetailsByTransactionIDRIBSLVinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase string `json:"coinbase"`
	ScriptSig GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int64 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness,omitempty"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInner instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSLVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInner(addresses []string, coinbase string, scriptSig GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig, sequence int64, txid string, value string, vout int32) *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner {
	this := GetWalletTransactionDetailsByTransactionIDRIBSLVinInner{}
	this.Addresses = addresses
	this.Coinbase = coinbase
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Value = value
	this.Vout = vout
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInnerWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSLVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInnerWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner {
	this := GetWalletTransactionDetailsByTransactionIDRIBSLVinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetCoinbase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coinbase, true
}

// SetCoinbase sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetCoinbase(v string) {
	o.Coinbase = v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetScriptSig() GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig {
	if o == nil {
		var ret GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetScriptSigOk() (*GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetScriptSig(v GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetSequenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetSequence(v int64) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxinwitness() []string {
	if o == nil || o.Txinwitness == nil {
		var ret []string
		return ret
	}
	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil || o.Txinwitness == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// HasTxinwitness returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) HasTxinwitness() bool {
	if o != nil && o.Txinwitness != nil {
		return true
	}

	return false
}

// SetTxinwitness gets a reference to the given []string and assigns it to the Txinwitness field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetVout(v int32) {
	o.Vout = v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) MarshalJSON() ([]byte, error) {
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
	if o.Txinwitness != nil {
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

type NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner) Get() *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner) Set(val *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner(val *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) *NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSLVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


