/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// GetTransactionDetailsByTransactionIDRIBSD2Vout struct for GetTransactionDetailsByTransactionIDRIBSD2Vout
type GetTransactionDetailsByTransactionIDRIBSD2Vout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewGetTransactionDetailsByTransactionIDRIBSD2Vout instantiates a new GetTransactionDetailsByTransactionIDRIBSD2Vout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSD2Vout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey, value string) *GetTransactionDetailsByTransactionIDRIBSD2Vout {
	this := GetTransactionDetailsByTransactionIDRIBSD2Vout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSD2VoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSD2Vout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSD2VoutWithDefaults() *GetTransactionDetailsByTransactionIDRIBSD2Vout {
	this := GetTransactionDetailsByTransactionIDRIBSD2Vout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) SetValue(v string) {
	o.Value = v
}

func (o GetTransactionDetailsByTransactionIDRIBSD2Vout) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDRIBSD2Vout struct {
	value *GetTransactionDetailsByTransactionIDRIBSD2Vout
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSD2Vout) Get() *GetTransactionDetailsByTransactionIDRIBSD2Vout {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSD2Vout) Set(val *GetTransactionDetailsByTransactionIDRIBSD2Vout) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSD2Vout) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSD2Vout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSD2Vout(val *GetTransactionDetailsByTransactionIDRIBSD2Vout) *NullableGetTransactionDetailsByTransactionIDRIBSD2Vout {
	return &NullableGetTransactionDetailsByTransactionIDRIBSD2Vout{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSD2Vout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSD2Vout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

