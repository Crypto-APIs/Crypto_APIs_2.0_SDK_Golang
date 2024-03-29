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

// GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey Represents the script public key.
type GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey struct {
	Addresses []string `json:"addresses"`
	// Represents the assembly of the script public key of the address.
	Asm string `json:"asm"`
	// Represents the hex of the script public key of the address.
	Hex string `json:"hex"`
	// Represents the required signatures.
	ReqSigs *int32 `json:"reqSigs,omitempty"`
	// Represents the script type.
	Type string `json:"type"`
}

// NewGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey(addresses []string, asm string, hex string, type_ string) *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey {
	this := GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey{}
	this.Addresses = addresses
	this.Asm = asm
	this.Hex = hex
	this.Type = type_
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKeyWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKeyWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey {
	this := GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) SetAddresses(v []string) {
	o.Addresses = v
}

// GetAsm returns the Asm field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetAsm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asm
}

// GetAsmOk returns a tuple with the Asm field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetAsmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asm, true
}

// SetAsm sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) SetAsm(v string) {
	o.Asm = v
}

// GetHex returns the Hex field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) SetHex(v string) {
	o.Hex = v
}

// GetReqSigs returns the ReqSigs field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetReqSigs() int32 {
	if o == nil || o.ReqSigs == nil {
		var ret int32
		return ret
	}
	return *o.ReqSigs
}

// GetReqSigsOk returns a tuple with the ReqSigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetReqSigsOk() (*int32, bool) {
	if o == nil || o.ReqSigs == nil {
		return nil, false
	}
	return o.ReqSigs, true
}

// HasReqSigs returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) HasReqSigs() bool {
	if o != nil && o.ReqSigs != nil {
		return true
	}

	return false
}

// SetReqSigs gets a reference to the given int32 and assigns it to the ReqSigs field.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) SetReqSigs(v int32) {
	o.ReqSigs = &v
}

// GetType returns the Type field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) SetType(v string) {
	o.Type = v
}

func (o GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) MarshalJSON() ([]byte, error) {
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
	if o.ReqSigs != nil {
		toSerialize["reqSigs"] = o.ReqSigs
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey struct {
	value *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) Get() *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) Set(val *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey(val *GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) *NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey {
	return &NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


