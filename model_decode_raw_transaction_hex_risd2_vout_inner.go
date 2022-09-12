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

// DecodeRawTransactionHexRISD2VoutInner struct for DecodeRawTransactionHexRISD2VoutInner
type DecodeRawTransactionHexRISD2VoutInner struct {
	ScriptPubKey DecodeRawTransactionHexRISD2VoutInnerScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
}

// NewDecodeRawTransactionHexRISD2VoutInner instantiates a new DecodeRawTransactionHexRISD2VoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISD2VoutInner(scriptPubKey DecodeRawTransactionHexRISD2VoutInnerScriptPubKey) *DecodeRawTransactionHexRISD2VoutInner {
	this := DecodeRawTransactionHexRISD2VoutInner{}
	this.ScriptPubKey = scriptPubKey
	return &this
}

// NewDecodeRawTransactionHexRISD2VoutInnerWithDefaults instantiates a new DecodeRawTransactionHexRISD2VoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISD2VoutInnerWithDefaults() *DecodeRawTransactionHexRISD2VoutInner {
	this := DecodeRawTransactionHexRISD2VoutInner{}
	return &this
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *DecodeRawTransactionHexRISD2VoutInner) GetScriptPubKey() DecodeRawTransactionHexRISD2VoutInnerScriptPubKey {
	if o == nil {
		var ret DecodeRawTransactionHexRISD2VoutInnerScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISD2VoutInner) GetScriptPubKeyOk() (*DecodeRawTransactionHexRISD2VoutInnerScriptPubKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *DecodeRawTransactionHexRISD2VoutInner) SetScriptPubKey(v DecodeRawTransactionHexRISD2VoutInnerScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISD2VoutInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISD2VoutInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISD2VoutInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DecodeRawTransactionHexRISD2VoutInner) SetValue(v string) {
	o.Value = &v
}

func (o DecodeRawTransactionHexRISD2VoutInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scriptPubKey"] = o.ScriptPubKey
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISD2VoutInner struct {
	value *DecodeRawTransactionHexRISD2VoutInner
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISD2VoutInner) Get() *DecodeRawTransactionHexRISD2VoutInner {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISD2VoutInner) Set(val *DecodeRawTransactionHexRISD2VoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISD2VoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISD2VoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISD2VoutInner(val *DecodeRawTransactionHexRISD2VoutInner) *NullableDecodeRawTransactionHexRISD2VoutInner {
	return &NullableDecodeRawTransactionHexRISD2VoutInner{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISD2VoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISD2VoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


