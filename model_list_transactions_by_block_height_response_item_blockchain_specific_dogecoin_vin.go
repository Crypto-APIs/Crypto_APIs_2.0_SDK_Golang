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

// ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin struct for ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin
type ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig, sequence string, txinwitness []string, value string, vout int32) *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin {
	this := ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txinwitness = txinwitness
	this.Value = value
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVinWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVinWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin {
	this := ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetScriptSig() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetScriptSig(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetSequenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) GetVoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) SetVout(v int32) {
	o.Vout = v
}

func (o ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Coinbase != nil {
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
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin struct {
	value *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin
	isSet bool
}

func (v NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) Get() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin {
	return v.value
}

func (v *NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) Set(val *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin(val *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) *NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin {
	return &NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


