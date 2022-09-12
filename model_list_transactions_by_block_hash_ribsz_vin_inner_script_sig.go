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

// ListTransactionsByBlockHashRIBSZVinInnerScriptSig Specifies the required signatures.
type ListTransactionsByBlockHashRIBSZVinInnerScriptSig struct {
	// The asm strands for assembly, which is the symbolic representation of the Bitcoin's Script language op-codes.
	Asm string `json:"asm"`
	// Represents the hex of the public key of the address.
	Hex string `json:"hex"`
	// Represents the script type of the reference transaction identifier.
	Type string `json:"type"`
}

// NewListTransactionsByBlockHashRIBSZVinInnerScriptSig instantiates a new ListTransactionsByBlockHashRIBSZVinInnerScriptSig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSZVinInnerScriptSig(asm string, hex string, type_ string) *ListTransactionsByBlockHashRIBSZVinInnerScriptSig {
	this := ListTransactionsByBlockHashRIBSZVinInnerScriptSig{}
	this.Asm = asm
	this.Hex = hex
	this.Type = type_
	return &this
}

// NewListTransactionsByBlockHashRIBSZVinInnerScriptSigWithDefaults instantiates a new ListTransactionsByBlockHashRIBSZVinInnerScriptSig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSZVinInnerScriptSigWithDefaults() *ListTransactionsByBlockHashRIBSZVinInnerScriptSig {
	this := ListTransactionsByBlockHashRIBSZVinInnerScriptSig{}
	return &this
}

// GetAsm returns the Asm field value
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) GetAsm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asm
}

// GetAsmOk returns a tuple with the Asm field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) GetAsmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asm, true
}

// SetAsm sets field value
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) SetAsm(v string) {
	o.Asm = v
}

// GetHex returns the Hex field value
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) SetHex(v string) {
	o.Hex = v
}

// GetType returns the Type field value
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) SetType(v string) {
	o.Type = v
}

func (o ListTransactionsByBlockHashRIBSZVinInnerScriptSig) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig struct {
	value *ListTransactionsByBlockHashRIBSZVinInnerScriptSig
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig) Get() *ListTransactionsByBlockHashRIBSZVinInnerScriptSig {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig) Set(val *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSZVinInnerScriptSig(val *ListTransactionsByBlockHashRIBSZVinInnerScriptSig) *NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig {
	return &NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSZVinInnerScriptSig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


