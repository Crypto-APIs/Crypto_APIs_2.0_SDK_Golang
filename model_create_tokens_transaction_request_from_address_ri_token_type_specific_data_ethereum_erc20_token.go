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

// CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token Ethereum Erc20 Token
type CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token struct {
	// Defines the contract address in the blockchain for an ERC20 token.
	ContractAddress string `json:"contractAddress"`
}

// NewCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token instantiates a new CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token(contractAddress string) *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token {
	this := CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token{}
	this.ContractAddress = contractAddress
	return &this
}

// NewCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20TokenWithDefaults instantiates a new CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20TokenWithDefaults() *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token {
	this := CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) GetContractAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token struct {
	value *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token
	isSet bool
}

func (v NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) Get() *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token {
	return v.value
}

func (v *NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) Set(val *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token(val *CreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) *NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token {
	return &NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token{value: val, isSet: true}
}

func (v NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokensTransactionRequestFromAddressRITokenTypeSpecificDataEthereumErc20Token) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

