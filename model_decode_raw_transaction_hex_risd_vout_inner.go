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

// DecodeRawTransactionHexRISDVoutInner struct for DecodeRawTransactionHexRISDVoutInner
type DecodeRawTransactionHexRISDVoutInner struct {
	ScriptPubKey DecodeRawTransactionHexRISDVoutInnerScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
}

// NewDecodeRawTransactionHexRISDVoutInner instantiates a new DecodeRawTransactionHexRISDVoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISDVoutInner(scriptPubKey DecodeRawTransactionHexRISDVoutInnerScriptPubKey) *DecodeRawTransactionHexRISDVoutInner {
	this := DecodeRawTransactionHexRISDVoutInner{}
	this.ScriptPubKey = scriptPubKey
	return &this
}

// NewDecodeRawTransactionHexRISDVoutInnerWithDefaults instantiates a new DecodeRawTransactionHexRISDVoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISDVoutInnerWithDefaults() *DecodeRawTransactionHexRISDVoutInner {
	this := DecodeRawTransactionHexRISDVoutInner{}
	return &this
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *DecodeRawTransactionHexRISDVoutInner) GetScriptPubKey() DecodeRawTransactionHexRISDVoutInnerScriptPubKey {
	if o == nil {
		var ret DecodeRawTransactionHexRISDVoutInnerScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVoutInner) GetScriptPubKeyOk() (*DecodeRawTransactionHexRISDVoutInnerScriptPubKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *DecodeRawTransactionHexRISDVoutInner) SetScriptPubKey(v DecodeRawTransactionHexRISDVoutInnerScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISDVoutInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVoutInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISDVoutInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DecodeRawTransactionHexRISDVoutInner) SetValue(v string) {
	o.Value = &v
}

func (o DecodeRawTransactionHexRISDVoutInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scriptPubKey"] = o.ScriptPubKey
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISDVoutInner struct {
	value *DecodeRawTransactionHexRISDVoutInner
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISDVoutInner) Get() *DecodeRawTransactionHexRISDVoutInner {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISDVoutInner) Set(val *DecodeRawTransactionHexRISDVoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISDVoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISDVoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISDVoutInner(val *DecodeRawTransactionHexRISDVoutInner) *NullableDecodeRawTransactionHexRISDVoutInner {
	return &NullableDecodeRawTransactionHexRISDVoutInner{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISDVoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISDVoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


