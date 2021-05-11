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

// ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin struct for ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin
type ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase string `json:"coinbase"`
	ScriptSig ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout *int32 `json:"vout,omitempty"`
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin(addresses []string, coinbase string, scriptSig ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinScriptSig, sequence string, txinwitness []string, value string) *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin{}
	this.Addresses = addresses
	this.Coinbase = coinbase
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txinwitness = txinwitness
	this.Value = value
	return &this
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetCoinbase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetCoinbaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Coinbase, true
}

// SetCoinbase sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetCoinbase(v string) {
	o.Coinbase = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetScriptSig() ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinScriptSig {
	if o == nil {
		var ret ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetScriptSigOk() (*ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetScriptSig(v ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetSequenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) SetVout(v int32) {
	o.Vout = &v
}

func (o ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if true {
		toSerialize["coinbase"] = o.Coinbase
	}
	if true {
		toSerialize["scriptSig"] = o.ScriptSig
	}
	if true {
		toSerialize["sequence"] = o.Sequence
	}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if true {
		toSerialize["txinwitness"] = o.Txinwitness
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.Vout != nil {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin struct {
	value *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin
	isSet bool
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) Get() *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin {
	return v.value
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) Set(val *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin(val *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) *NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin {
	return &NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin{value: val, isSet: true}
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


