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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner struct for GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner
type GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerScriptPubKey, value string) *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) GetIsSpentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) GetScriptPubKey() GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerScriptPubKey {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerScriptPubKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) SetScriptPubKey(v GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInnerScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) SetValue(v string) {
	o.Value = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSBCVoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


