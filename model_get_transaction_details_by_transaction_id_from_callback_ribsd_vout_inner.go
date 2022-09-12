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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner struct for GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner
type GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerScriptPubKey, value string) *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) GetIsSpentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) GetScriptPubKey() GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerScriptPubKey {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerScriptPubKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) SetScriptPubKey(v GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInnerScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) SetValue(v string) {
	o.Value = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

