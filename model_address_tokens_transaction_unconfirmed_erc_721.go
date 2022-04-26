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

// AddressTokensTransactionUnconfirmedErc721 ERC-721
type AddressTokensTransactionUnconfirmedErc721 struct {
	// Specifies the name of the token.
	Name string `json:"name"`
	// Specifies an identifier of the token, where up to five alphanumeric characters can be used for it.
	Symbol string `json:"symbol"`
	// Specifies the unique ID of the token.
	TokenId string `json:"tokenId"`
	// Specifies the address of the contract.
	ContractAddress string `json:"contractAddress"`
}

// NewAddressTokensTransactionUnconfirmedErc721 instantiates a new AddressTokensTransactionUnconfirmedErc721 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTokensTransactionUnconfirmedErc721(name string, symbol string, tokenId string, contractAddress string) *AddressTokensTransactionUnconfirmedErc721 {
	this := AddressTokensTransactionUnconfirmedErc721{}
	this.Name = name
	this.Symbol = symbol
	this.TokenId = tokenId
	this.ContractAddress = contractAddress
	return &this
}

// NewAddressTokensTransactionUnconfirmedErc721WithDefaults instantiates a new AddressTokensTransactionUnconfirmedErc721 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTokensTransactionUnconfirmedErc721WithDefaults() *AddressTokensTransactionUnconfirmedErc721 {
	this := AddressTokensTransactionUnconfirmedErc721{}
	return &this
}

// GetName returns the Name field value
func (o *AddressTokensTransactionUnconfirmedErc721) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionUnconfirmedErc721) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddressTokensTransactionUnconfirmedErc721) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *AddressTokensTransactionUnconfirmedErc721) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionUnconfirmedErc721) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *AddressTokensTransactionUnconfirmedErc721) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokenId returns the TokenId field value
func (o *AddressTokensTransactionUnconfirmedErc721) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionUnconfirmedErc721) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *AddressTokensTransactionUnconfirmedErc721) SetTokenId(v string) {
	o.TokenId = v
}

// GetContractAddress returns the ContractAddress field value
func (o *AddressTokensTransactionUnconfirmedErc721) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionUnconfirmedErc721) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *AddressTokensTransactionUnconfirmedErc721) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o AddressTokensTransactionUnconfirmedErc721) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["tokenId"] = o.TokenId
	}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableAddressTokensTransactionUnconfirmedErc721 struct {
	value *AddressTokensTransactionUnconfirmedErc721
	isSet bool
}

func (v NullableAddressTokensTransactionUnconfirmedErc721) Get() *AddressTokensTransactionUnconfirmedErc721 {
	return v.value
}

func (v *NullableAddressTokensTransactionUnconfirmedErc721) Set(val *AddressTokensTransactionUnconfirmedErc721) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokensTransactionUnconfirmedErc721) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokensTransactionUnconfirmedErc721) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokensTransactionUnconfirmedErc721(val *AddressTokensTransactionUnconfirmedErc721) *NullableAddressTokensTransactionUnconfirmedErc721 {
	return &NullableAddressTokensTransactionUnconfirmedErc721{value: val, isSet: true}
}

func (v NullableAddressTokensTransactionUnconfirmedErc721) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokensTransactionUnconfirmedErc721) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


