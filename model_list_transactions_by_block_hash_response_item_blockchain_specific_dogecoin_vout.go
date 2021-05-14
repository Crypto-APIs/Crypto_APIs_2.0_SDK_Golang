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

// ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout struct for ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout
type ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout(isSpent bool, scriptPubKey ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinScriptPubKey, value string) *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout {
	this := ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVoutWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVoutWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout {
	this := ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) GetScriptPubKey() ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinScriptPubKey {
	if o == nil {
		var ret ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) GetScriptPubKeyOk() (*ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) SetScriptPubKey(v ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) SetValue(v string) {
	o.Value = v
}

func (o ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout struct {
	value *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout
	isSet bool
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) Get() *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout {
	return v.value
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) Set(val *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout(val *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout {
	return &NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

