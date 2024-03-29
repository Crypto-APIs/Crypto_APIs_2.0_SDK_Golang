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

// ListTokensForwardingAutomationsRITSET Ethereum Token
type ListTokensForwardingAutomationsRITSET struct {
	// Defines the token contract address.
	ContractAddress string `json:"contractAddress"`
}

// NewListTokensForwardingAutomationsRITSET instantiates a new ListTokensForwardingAutomationsRITSET object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTokensForwardingAutomationsRITSET(contractAddress string) *ListTokensForwardingAutomationsRITSET {
	this := ListTokensForwardingAutomationsRITSET{}
	this.ContractAddress = contractAddress
	return &this
}

// NewListTokensForwardingAutomationsRITSETWithDefaults instantiates a new ListTokensForwardingAutomationsRITSET object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTokensForwardingAutomationsRITSETWithDefaults() *ListTokensForwardingAutomationsRITSET {
	this := ListTokensForwardingAutomationsRITSET{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *ListTokensForwardingAutomationsRITSET) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRITSET) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ListTokensForwardingAutomationsRITSET) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o ListTokensForwardingAutomationsRITSET) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableListTokensForwardingAutomationsRITSET struct {
	value *ListTokensForwardingAutomationsRITSET
	isSet bool
}

func (v NullableListTokensForwardingAutomationsRITSET) Get() *ListTokensForwardingAutomationsRITSET {
	return v.value
}

func (v *NullableListTokensForwardingAutomationsRITSET) Set(val *ListTokensForwardingAutomationsRITSET) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokensForwardingAutomationsRITSET) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokensForwardingAutomationsRITSET) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokensForwardingAutomationsRITSET(val *ListTokensForwardingAutomationsRITSET) *NullableListTokensForwardingAutomationsRITSET {
	return &NullableListTokensForwardingAutomationsRITSET{value: val, isSet: true}
}

func (v NullableListTokensForwardingAutomationsRITSET) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokensForwardingAutomationsRITSET) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


