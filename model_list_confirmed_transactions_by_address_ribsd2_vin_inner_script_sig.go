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

// ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig Specifies the required signatures.
type ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig struct {
	// The asm strands for assembly, which is the symbolic representation of the Bitcoin's Script language op-codes.
	Asm string `json:"asm"`
	// Represents the hex of the public key of the address.
	Hex string `json:"hex"`
	// Represents the script type of the reference transaction identifier.
	Type string `json:"type"`
}

// NewListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig instantiates a new ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig(asm string, hex string, type_ string) *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig {
	this := ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig{}
	this.Asm = asm
	this.Hex = hex
	this.Type = type_
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSigWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSigWithDefaults() *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig {
	this := ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig{}
	return &this
}

// GetAsm returns the Asm field value
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) GetAsm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asm
}

// GetAsmOk returns a tuple with the Asm field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) GetAsmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asm, true
}

// SetAsm sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) SetAsm(v string) {
	o.Asm = v
}

// GetHex returns the Hex field value
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) SetHex(v string) {
	o.Hex = v
}

// GetType returns the Type field value
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) SetType(v string) {
	o.Type = v
}

func (o ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asm"] = o.Asm
	}
	if true {
		toSerialize["hex"] = o.Hex
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig struct {
	value *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) Get() *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) Set(val *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig(val *ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) *NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig {
	return &NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


