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

// DecodeRawTransactionHexRISZVoutInner struct for DecodeRawTransactionHexRISZVoutInner
type DecodeRawTransactionHexRISZVoutInner struct {
	ScriptPubKey DecodeRawTransactionHexRISZVoutInnerScriptPubKey `json:"scriptPubKey"`
	// Defines the specific amount.
	Value *string `json:"value,omitempty"`
}

// NewDecodeRawTransactionHexRISZVoutInner instantiates a new DecodeRawTransactionHexRISZVoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISZVoutInner(scriptPubKey DecodeRawTransactionHexRISZVoutInnerScriptPubKey) *DecodeRawTransactionHexRISZVoutInner {
	this := DecodeRawTransactionHexRISZVoutInner{}
	this.ScriptPubKey = scriptPubKey
	return &this
}

// NewDecodeRawTransactionHexRISZVoutInnerWithDefaults instantiates a new DecodeRawTransactionHexRISZVoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISZVoutInnerWithDefaults() *DecodeRawTransactionHexRISZVoutInner {
	this := DecodeRawTransactionHexRISZVoutInner{}
	return &this
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *DecodeRawTransactionHexRISZVoutInner) GetScriptPubKey() DecodeRawTransactionHexRISZVoutInnerScriptPubKey {
	if o == nil {
		var ret DecodeRawTransactionHexRISZVoutInnerScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISZVoutInner) GetScriptPubKeyOk() (*DecodeRawTransactionHexRISZVoutInnerScriptPubKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *DecodeRawTransactionHexRISZVoutInner) SetScriptPubKey(v DecodeRawTransactionHexRISZVoutInnerScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISZVoutInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISZVoutInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISZVoutInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DecodeRawTransactionHexRISZVoutInner) SetValue(v string) {
	o.Value = &v
}

func (o DecodeRawTransactionHexRISZVoutInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scriptPubKey"] = o.ScriptPubKey
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISZVoutInner struct {
	value *DecodeRawTransactionHexRISZVoutInner
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISZVoutInner) Get() *DecodeRawTransactionHexRISZVoutInner {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISZVoutInner) Set(val *DecodeRawTransactionHexRISZVoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISZVoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISZVoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISZVoutInner(val *DecodeRawTransactionHexRISZVoutInner) *NullableDecodeRawTransactionHexRISZVoutInner {
	return &NullableDecodeRawTransactionHexRISZVoutInner{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISZVoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISZVoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


