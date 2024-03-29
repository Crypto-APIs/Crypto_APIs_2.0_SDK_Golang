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

// AddTokensToExistingFromAddressRITSET Ethereum Token
type AddTokensToExistingFromAddressRITSET struct {
	// Token contract address to be transferred
	ContractAddress string `json:"contractAddress"`
}

// NewAddTokensToExistingFromAddressRITSET instantiates a new AddTokensToExistingFromAddressRITSET object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTokensToExistingFromAddressRITSET(contractAddress string) *AddTokensToExistingFromAddressRITSET {
	this := AddTokensToExistingFromAddressRITSET{}
	this.ContractAddress = contractAddress
	return &this
}

// NewAddTokensToExistingFromAddressRITSETWithDefaults instantiates a new AddTokensToExistingFromAddressRITSET object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTokensToExistingFromAddressRITSETWithDefaults() *AddTokensToExistingFromAddressRITSET {
	this := AddTokensToExistingFromAddressRITSET{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *AddTokensToExistingFromAddressRITSET) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRITSET) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *AddTokensToExistingFromAddressRITSET) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o AddTokensToExistingFromAddressRITSET) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableAddTokensToExistingFromAddressRITSET struct {
	value *AddTokensToExistingFromAddressRITSET
	isSet bool
}

func (v NullableAddTokensToExistingFromAddressRITSET) Get() *AddTokensToExistingFromAddressRITSET {
	return v.value
}

func (v *NullableAddTokensToExistingFromAddressRITSET) Set(val *AddTokensToExistingFromAddressRITSET) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTokensToExistingFromAddressRITSET) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTokensToExistingFromAddressRITSET) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTokensToExistingFromAddressRITSET(val *AddTokensToExistingFromAddressRITSET) *NullableAddTokensToExistingFromAddressRITSET {
	return &NullableAddTokensToExistingFromAddressRITSET{value: val, isSet: true}
}

func (v NullableAddTokensToExistingFromAddressRITSET) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTokensToExistingFromAddressRITSET) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


