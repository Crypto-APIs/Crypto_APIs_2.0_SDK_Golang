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

// ListUnconfirmedTransactionsByAddressRIBSBScriptSig Specifies the required signatures.
type ListUnconfirmedTransactionsByAddressRIBSBScriptSig struct {
	// The asm strands for assembly, which is the symbolic representation of the Bitcoin's Script language op-codes.
	Asm string `json:"asm"`
	// Represents the hex of the public key of the address.
	Hex *string `json:"hex,omitempty"`
	// Represents the script type of the reference transaction identifier.
	Type string `json:"type"`
}

// NewListUnconfirmedTransactionsByAddressRIBSBScriptSig instantiates a new ListUnconfirmedTransactionsByAddressRIBSBScriptSig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTransactionsByAddressRIBSBScriptSig(asm string, type_ string) *ListUnconfirmedTransactionsByAddressRIBSBScriptSig {
	this := ListUnconfirmedTransactionsByAddressRIBSBScriptSig{}
	this.Asm = asm
	this.Type = type_
	return &this
}

// NewListUnconfirmedTransactionsByAddressRIBSBScriptSigWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSBScriptSig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTransactionsByAddressRIBSBScriptSigWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSBScriptSig {
	this := ListUnconfirmedTransactionsByAddressRIBSBScriptSig{}
	return &this
}

// GetAsm returns the Asm field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetAsm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asm
}

// GetAsmOk returns a tuple with the Asm field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetAsmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asm, true
}

// SetAsm sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) SetAsm(v string) {
	o.Asm = v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetHex() string {
	if o == nil || o.Hex == nil {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetHexOk() (*string, bool) {
	if o == nil || o.Hex == nil {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) HasHex() bool {
	if o != nil && o.Hex != nil {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) SetHex(v string) {
	o.Hex = &v
}

// GetType returns the Type field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) SetType(v string) {
	o.Type = v
}

func (o ListUnconfirmedTransactionsByAddressRIBSBScriptSig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asm"] = o.Asm
	}
	if o.Hex != nil {
		toSerialize["hex"] = o.Hex
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig struct {
	value *ListUnconfirmedTransactionsByAddressRIBSBScriptSig
	isSet bool
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig) Get() *ListUnconfirmedTransactionsByAddressRIBSBScriptSig {
	return v.value
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig) Set(val *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTransactionsByAddressRIBSBScriptSig(val *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) *NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig {
	return &NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig{value: val, isSet: true}
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSBScriptSig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


