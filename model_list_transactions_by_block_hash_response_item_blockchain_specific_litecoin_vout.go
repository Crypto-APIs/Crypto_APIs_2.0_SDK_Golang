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

// ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout struct for ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout
type ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout(isSpent bool, scriptPubKey ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey, value string) *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout {
	this := ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVoutWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVoutWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout {
	this := ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetScriptPubKey() ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey {
	if o == nil {
		var ret ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetScriptPubKeyOk() (*ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) SetScriptPubKey(v ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) SetValue(v string) {
	o.Value = v
}

func (o ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout struct {
	value *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout
	isSet bool
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) Get() *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout {
	return v.value
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) Set(val *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout(val *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout {
	return &NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


