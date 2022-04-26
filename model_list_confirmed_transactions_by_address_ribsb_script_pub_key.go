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

// ListConfirmedTransactionsByAddressRIBSBScriptPubKey Represents the script public key.
type ListConfirmedTransactionsByAddressRIBSBScriptPubKey struct {
	Addresses []string `json:"addresses"`
	// Represents the assembly of the script public key of the address.
	Asm string `json:"asm"`
	// Represents the hex of the script public key of the address.
	Hex string `json:"hex"`
	// Represents the required signatures.
	ReqSigs int32 `json:"reqSigs"`
	// Represents the script type.
	Type string `json:"type"`
}

// NewListConfirmedTransactionsByAddressRIBSBScriptPubKey instantiates a new ListConfirmedTransactionsByAddressRIBSBScriptPubKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSBScriptPubKey(addresses []string, asm string, hex string, reqSigs int32, type_ string) *ListConfirmedTransactionsByAddressRIBSBScriptPubKey {
	this := ListConfirmedTransactionsByAddressRIBSBScriptPubKey{}
	this.Addresses = addresses
	this.Asm = asm
	this.Hex = hex
	this.ReqSigs = reqSigs
	this.Type = type_
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSBScriptPubKeyWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSBScriptPubKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSBScriptPubKeyWithDefaults() *ListConfirmedTransactionsByAddressRIBSBScriptPubKey {
	this := ListConfirmedTransactionsByAddressRIBSBScriptPubKey{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) SetAddresses(v []string) {
	o.Addresses = v
}

// GetAsm returns the Asm field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetAsm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asm
}

// GetAsmOk returns a tuple with the Asm field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetAsmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asm, true
}

// SetAsm sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) SetAsm(v string) {
	o.Asm = v
}

// GetHex returns the Hex field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) SetHex(v string) {
	o.Hex = v
}

// GetReqSigs returns the ReqSigs field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetReqSigs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReqSigs
}

// GetReqSigsOk returns a tuple with the ReqSigs field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetReqSigsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqSigs, true
}

// SetReqSigs sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) SetReqSigs(v int32) {
	o.ReqSigs = v
}

// GetType returns the Type field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) SetType(v string) {
	o.Type = v
}

func (o ListConfirmedTransactionsByAddressRIBSBScriptPubKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if true {
		toSerialize["asm"] = o.Asm
	}
	if true {
		toSerialize["hex"] = o.Hex
	}
	if true {
		toSerialize["reqSigs"] = o.ReqSigs
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey struct {
	value *ListConfirmedTransactionsByAddressRIBSBScriptPubKey
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey) Get() *ListConfirmedTransactionsByAddressRIBSBScriptPubKey {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey) Set(val *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSBScriptPubKey(val *ListConfirmedTransactionsByAddressRIBSBScriptPubKey) *NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey {
	return &NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBScriptPubKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


