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

// DecodeRawTransactionHexRISLVinInner struct for DecodeRawTransactionHexRISLVinInner
type DecodeRawTransactionHexRISLVinInner struct {
	// Represents the addresses which send/receive the amount.
	Address *string `json:"address,omitempty"`
	// Represents the transaction inputs' indentifier.
	InputHash *string `json:"inputHash,omitempty"`
	// Defines the output index of a transaction.
	OutputIndex *string `json:"outputIndex,omitempty"`
	ScriptSig DecodeRawTransactionHexRISLVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence *string `json:"sequence,omitempty"`
	Txinwitness []string `json:"txinwitness,omitempty"`
}

// NewDecodeRawTransactionHexRISLVinInner instantiates a new DecodeRawTransactionHexRISLVinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISLVinInner(scriptSig DecodeRawTransactionHexRISLVinInnerScriptSig) *DecodeRawTransactionHexRISLVinInner {
	this := DecodeRawTransactionHexRISLVinInner{}
	this.ScriptSig = scriptSig
	return &this
}

// NewDecodeRawTransactionHexRISLVinInnerWithDefaults instantiates a new DecodeRawTransactionHexRISLVinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISLVinInnerWithDefaults() *DecodeRawTransactionHexRISLVinInner {
	this := DecodeRawTransactionHexRISLVinInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISLVinInner) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISLVinInner) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISLVinInner) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DecodeRawTransactionHexRISLVinInner) SetAddress(v string) {
	o.Address = &v
}

// GetInputHash returns the InputHash field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISLVinInner) GetInputHash() string {
	if o == nil || o.InputHash == nil {
		var ret string
		return ret
	}
	return *o.InputHash
}

// GetInputHashOk returns a tuple with the InputHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISLVinInner) GetInputHashOk() (*string, bool) {
	if o == nil || o.InputHash == nil {
		return nil, false
	}
	return o.InputHash, true
}

// HasInputHash returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISLVinInner) HasInputHash() bool {
	if o != nil && o.InputHash != nil {
		return true
	}

	return false
}

// SetInputHash gets a reference to the given string and assigns it to the InputHash field.
func (o *DecodeRawTransactionHexRISLVinInner) SetInputHash(v string) {
	o.InputHash = &v
}

// GetOutputIndex returns the OutputIndex field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISLVinInner) GetOutputIndex() string {
	if o == nil || o.OutputIndex == nil {
		var ret string
		return ret
	}
	return *o.OutputIndex
}

// GetOutputIndexOk returns a tuple with the OutputIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISLVinInner) GetOutputIndexOk() (*string, bool) {
	if o == nil || o.OutputIndex == nil {
		return nil, false
	}
	return o.OutputIndex, true
}

// HasOutputIndex returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISLVinInner) HasOutputIndex() bool {
	if o != nil && o.OutputIndex != nil {
		return true
	}

	return false
}

// SetOutputIndex gets a reference to the given string and assigns it to the OutputIndex field.
func (o *DecodeRawTransactionHexRISLVinInner) SetOutputIndex(v string) {
	o.OutputIndex = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *DecodeRawTransactionHexRISLVinInner) GetScriptSig() DecodeRawTransactionHexRISLVinInnerScriptSig {
	if o == nil {
		var ret DecodeRawTransactionHexRISLVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISLVinInner) GetScriptSigOk() (*DecodeRawTransactionHexRISLVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *DecodeRawTransactionHexRISLVinInner) SetScriptSig(v DecodeRawTransactionHexRISLVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISLVinInner) GetSequence() string {
	if o == nil || o.Sequence == nil {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISLVinInner) GetSequenceOk() (*string, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISLVinInner) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *DecodeRawTransactionHexRISLVinInner) SetSequence(v string) {
	o.Sequence = &v
}

// GetTxinwitness returns the Txinwitness field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISLVinInner) GetTxinwitness() []string {
	if o == nil || o.Txinwitness == nil {
		var ret []string
		return ret
	}
	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISLVinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil || o.Txinwitness == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// HasTxinwitness returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISLVinInner) HasTxinwitness() bool {
	if o != nil && o.Txinwitness != nil {
		return true
	}

	return false
}

// SetTxinwitness gets a reference to the given []string and assigns it to the Txinwitness field.
func (o *DecodeRawTransactionHexRISLVinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

func (o DecodeRawTransactionHexRISLVinInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.InputHash != nil {
		toSerialize["inputHash"] = o.InputHash
	}
	if o.OutputIndex != nil {
		toSerialize["outputIndex"] = o.OutputIndex
	}
	if true {
		toSerialize["scriptSig"] = o.ScriptSig
	}
	if o.Sequence != nil {
		toSerialize["sequence"] = o.Sequence
	}
	if o.Txinwitness != nil {
		toSerialize["txinwitness"] = o.Txinwitness
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISLVinInner struct {
	value *DecodeRawTransactionHexRISLVinInner
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISLVinInner) Get() *DecodeRawTransactionHexRISLVinInner {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISLVinInner) Set(val *DecodeRawTransactionHexRISLVinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISLVinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISLVinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISLVinInner(val *DecodeRawTransactionHexRISLVinInner) *NullableDecodeRawTransactionHexRISLVinInner {
	return &NullableDecodeRawTransactionHexRISLVinInner{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISLVinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISLVinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


