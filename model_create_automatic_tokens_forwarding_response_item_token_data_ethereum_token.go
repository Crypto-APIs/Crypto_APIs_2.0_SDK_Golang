/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken Ethereum Token
type CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken struct {
	// Represents the specific `contractAddress` of the Token that will be forwarded.
	ContractAddress string `json:"contractAddress"`
}

// NewCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken instantiates a new CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken(contractAddress string) *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken {
	this := CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken{}
	this.ContractAddress = contractAddress
	return &this
}

// NewCreateAutomaticTokensForwardingResponseItemTokenDataEthereumTokenWithDefaults instantiates a new CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticTokensForwardingResponseItemTokenDataEthereumTokenWithDefaults() *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken {
	this := CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) GetContractAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken struct {
	value *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken
	isSet bool
}

func (v NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) Get() *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken {
	return v.value
}

func (v *NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) Set(val *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken(val *CreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) *NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken {
	return &NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken{value: val, isSet: true}
}

func (v NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticTokensForwardingResponseItemTokenDataEthereumToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


