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

// DecodeRawTransactionHexRISDVinInnerScriptSig Specifies the required signatures.
type DecodeRawTransactionHexRISDVinInnerScriptSig struct {
	// The asm strands for assembly, which is the symbolic representation of the Bitcoin's Script language op-codes.
	Asm *string `json:"asm,omitempty"`
	// Represents the hex of the public key of the address
	Hex *string `json:"hex,omitempty"`
	// Represents the script type of the reference transaction identifier
	Type *string `json:"type,omitempty"`
}

// NewDecodeRawTransactionHexRISDVinInnerScriptSig instantiates a new DecodeRawTransactionHexRISDVinInnerScriptSig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISDVinInnerScriptSig() *DecodeRawTransactionHexRISDVinInnerScriptSig {
	this := DecodeRawTransactionHexRISDVinInnerScriptSig{}
	return &this
}

// NewDecodeRawTransactionHexRISDVinInnerScriptSigWithDefaults instantiates a new DecodeRawTransactionHexRISDVinInnerScriptSig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISDVinInnerScriptSigWithDefaults() *DecodeRawTransactionHexRISDVinInnerScriptSig {
	this := DecodeRawTransactionHexRISDVinInnerScriptSig{}
	return &this
}

// GetAsm returns the Asm field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) GetAsm() string {
	if o == nil || o.Asm == nil {
		var ret string
		return ret
	}
	return *o.Asm
}

// GetAsmOk returns a tuple with the Asm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) GetAsmOk() (*string, bool) {
	if o == nil || o.Asm == nil {
		return nil, false
	}
	return o.Asm, true
}

// HasAsm returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) HasAsm() bool {
	if o != nil && o.Asm != nil {
		return true
	}

	return false
}

// SetAsm gets a reference to the given string and assigns it to the Asm field.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) SetAsm(v string) {
	o.Asm = &v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) GetHex() string {
	if o == nil || o.Hex == nil {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) GetHexOk() (*string, bool) {
	if o == nil || o.Hex == nil {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) HasHex() bool {
	if o != nil && o.Hex != nil {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) SetHex(v string) {
	o.Hex = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DecodeRawTransactionHexRISDVinInnerScriptSig) SetType(v string) {
	o.Type = &v
}

func (o DecodeRawTransactionHexRISDVinInnerScriptSig) MarshalJSON() ([]byte, error) {
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

type NullableDecodeRawTransactionHexRISDVinInnerScriptSig struct {
	value *DecodeRawTransactionHexRISDVinInnerScriptSig
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISDVinInnerScriptSig) Get() *DecodeRawTransactionHexRISDVinInnerScriptSig {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISDVinInnerScriptSig) Set(val *DecodeRawTransactionHexRISDVinInnerScriptSig) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISDVinInnerScriptSig) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISDVinInnerScriptSig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISDVinInnerScriptSig(val *DecodeRawTransactionHexRISDVinInnerScriptSig) *NullableDecodeRawTransactionHexRISDVinInnerScriptSig {
	return &NullableDecodeRawTransactionHexRISDVinInnerScriptSig{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISDVinInnerScriptSig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISDVinInnerScriptSig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


