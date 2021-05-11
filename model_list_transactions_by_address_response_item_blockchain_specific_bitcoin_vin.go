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

// ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin struct for ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin
type ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// Defines the vout of the transaction output, i.e. which output to spend.
	Vout *int32 `json:"vout,omitempty"`
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig, sequence string, txinwitness []string) *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txinwitness = txinwitness
	return &this
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetScriptSig() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetScriptSig(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetSequenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetVout(v int32) {
	o.Vout = &v
}

func (o ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Vout != nil {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin struct {
	value *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin
	isSet bool
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) Get() *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin {
	return v.value
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) Set(val *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin(val *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin {
	return &NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin{value: val, isSet: true}
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


