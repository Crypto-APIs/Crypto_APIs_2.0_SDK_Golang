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

// ListAllUnconfirmedTransactionsRIBSBScriptSig Specifies the required signatures.
type ListAllUnconfirmedTransactionsRIBSBScriptSig struct {
	// The asm strands for assembly, which is the symbolic representation of the Bitcoin's Script language op-codes.
	Asm string `json:"asm"`
	// Represents the hex of the public key of the address.
	Hex string `json:"hex"`
	// Represents the script type of the reference transaction identifier.
	Type string `json:"type"`
}

// NewListAllUnconfirmedTransactionsRIBSBScriptSig instantiates a new ListAllUnconfirmedTransactionsRIBSBScriptSig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllUnconfirmedTransactionsRIBSBScriptSig(asm string, hex string, type_ string) *ListAllUnconfirmedTransactionsRIBSBScriptSig {
	this := ListAllUnconfirmedTransactionsRIBSBScriptSig{}
	this.Asm = asm
	this.Hex = hex
	this.Type = type_
	return &this
}

// NewListAllUnconfirmedTransactionsRIBSBScriptSigWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSBScriptSig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllUnconfirmedTransactionsRIBSBScriptSigWithDefaults() *ListAllUnconfirmedTransactionsRIBSBScriptSig {
	this := ListAllUnconfirmedTransactionsRIBSBScriptSig{}
	return &this
}

// GetAsm returns the Asm field value
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) GetAsm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asm
}

// GetAsmOk returns a tuple with the Asm field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) GetAsmOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asm, true
}

// SetAsm sets field value
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) SetAsm(v string) {
	o.Asm = v
}

// GetHex returns the Hex field value
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) GetHexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) SetHex(v string) {
	o.Hex = v
}

// GetType returns the Type field value
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListAllUnconfirmedTransactionsRIBSBScriptSig) SetType(v string) {
	o.Type = v
}

func (o ListAllUnconfirmedTransactionsRIBSBScriptSig) MarshalJSON() ([]byte, error) {
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

type NullableListAllUnconfirmedTransactionsRIBSBScriptSig struct {
	value *ListAllUnconfirmedTransactionsRIBSBScriptSig
	isSet bool
}

func (v NullableListAllUnconfirmedTransactionsRIBSBScriptSig) Get() *ListAllUnconfirmedTransactionsRIBSBScriptSig {
	return v.value
}

func (v *NullableListAllUnconfirmedTransactionsRIBSBScriptSig) Set(val *ListAllUnconfirmedTransactionsRIBSBScriptSig) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllUnconfirmedTransactionsRIBSBScriptSig) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllUnconfirmedTransactionsRIBSBScriptSig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllUnconfirmedTransactionsRIBSBScriptSig(val *ListAllUnconfirmedTransactionsRIBSBScriptSig) *NullableListAllUnconfirmedTransactionsRIBSBScriptSig {
	return &NullableListAllUnconfirmedTransactionsRIBSBScriptSig{value: val, isSet: true}
}

func (v NullableListAllUnconfirmedTransactionsRIBSBScriptSig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllUnconfirmedTransactionsRIBSBScriptSig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

