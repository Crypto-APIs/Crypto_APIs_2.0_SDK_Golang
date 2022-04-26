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

// DecodeRawTransactionHexRISBScriptSig Specifies the required signatures.
type DecodeRawTransactionHexRISBScriptSig struct {
	// The asm strands for assembly, which is the symbolic representation of the Bitcoin's Script language op-codes.
	Asm *string `json:"asm,omitempty"`
	// Represents the hex of the public key of the address.
	Hex *string `json:"hex,omitempty"`
	// Represents the script type of the reference transaction identifier.
	Type *string `json:"type,omitempty"`
}

// NewDecodeRawTransactionHexRISBScriptSig instantiates a new DecodeRawTransactionHexRISBScriptSig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISBScriptSig() *DecodeRawTransactionHexRISBScriptSig {
	this := DecodeRawTransactionHexRISBScriptSig{}
	return &this
}

// NewDecodeRawTransactionHexRISBScriptSigWithDefaults instantiates a new DecodeRawTransactionHexRISBScriptSig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISBScriptSigWithDefaults() *DecodeRawTransactionHexRISBScriptSig {
	this := DecodeRawTransactionHexRISBScriptSig{}
	return &this
}

// GetAsm returns the Asm field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISBScriptSig) GetAsm() string {
	if o == nil || o.Asm == nil {
		var ret string
		return ret
	}
	return *o.Asm
}

// GetAsmOk returns a tuple with the Asm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISBScriptSig) GetAsmOk() (*string, bool) {
	if o == nil || o.Asm == nil {
		return nil, false
	}
	return o.Asm, true
}

// HasAsm returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISBScriptSig) HasAsm() bool {
	if o != nil && o.Asm != nil {
		return true
	}

	return false
}

// SetAsm gets a reference to the given string and assigns it to the Asm field.
func (o *DecodeRawTransactionHexRISBScriptSig) SetAsm(v string) {
	o.Asm = &v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISBScriptSig) GetHex() string {
	if o == nil || o.Hex == nil {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISBScriptSig) GetHexOk() (*string, bool) {
	if o == nil || o.Hex == nil {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISBScriptSig) HasHex() bool {
	if o != nil && o.Hex != nil {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *DecodeRawTransactionHexRISBScriptSig) SetHex(v string) {
	o.Hex = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISBScriptSig) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISBScriptSig) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISBScriptSig) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DecodeRawTransactionHexRISBScriptSig) SetType(v string) {
	o.Type = &v
}

func (o DecodeRawTransactionHexRISBScriptSig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Asm != nil {
		toSerialize["asm"] = o.Asm
	}
	if o.Hex != nil {
		toSerialize["hex"] = o.Hex
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISBScriptSig struct {
	value *DecodeRawTransactionHexRISBScriptSig
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISBScriptSig) Get() *DecodeRawTransactionHexRISBScriptSig {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISBScriptSig) Set(val *DecodeRawTransactionHexRISBScriptSig) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISBScriptSig) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISBScriptSig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISBScriptSig(val *DecodeRawTransactionHexRISBScriptSig) *NullableDecodeRawTransactionHexRISBScriptSig {
	return &NullableDecodeRawTransactionHexRISBScriptSig{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISBScriptSig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISBScriptSig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

