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

// ListUnconfirmedTransactionsByAddressRIBSBCVoutInner struct for ListUnconfirmedTransactionsByAddressRIBSBCVoutInner
type ListUnconfirmedTransactionsByAddressRIBSBCVoutInner struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListUnconfirmedTransactionsByAddressRIBSBCVoutInnerScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewListUnconfirmedTransactionsByAddressRIBSBCVoutInner instantiates a new ListUnconfirmedTransactionsByAddressRIBSBCVoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTransactionsByAddressRIBSBCVoutInner(isSpent bool, scriptPubKey ListUnconfirmedTransactionsByAddressRIBSBCVoutInnerScriptPubKey, value string) *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner {
	this := ListUnconfirmedTransactionsByAddressRIBSBCVoutInner{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListUnconfirmedTransactionsByAddressRIBSBCVoutInnerWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSBCVoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTransactionsByAddressRIBSBCVoutInnerWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner {
	this := ListUnconfirmedTransactionsByAddressRIBSBCVoutInner{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) GetIsSpentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) GetScriptPubKey() ListUnconfirmedTransactionsByAddressRIBSBCVoutInnerScriptPubKey {
	if o == nil {
		var ret ListUnconfirmedTransactionsByAddressRIBSBCVoutInnerScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) GetScriptPubKeyOk() (*ListUnconfirmedTransactionsByAddressRIBSBCVoutInnerScriptPubKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) SetScriptPubKey(v ListUnconfirmedTransactionsByAddressRIBSBCVoutInnerScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) SetValue(v string) {
	o.Value = v
}

func (o ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isSpent"] = o.IsSpent
	}
	if true {
		toSerialize["scriptPubKey"] = o.ScriptPubKey
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner struct {
	value *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner
	isSet bool
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner) Get() *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner {
	return v.value
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner) Set(val *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner(val *ListUnconfirmedTransactionsByAddressRIBSBCVoutInner) *NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner {
	return &NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner{value: val, isSet: true}
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBCVoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


