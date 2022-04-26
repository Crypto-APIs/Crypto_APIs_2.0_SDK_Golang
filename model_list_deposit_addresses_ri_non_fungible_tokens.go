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

// ListDepositAddressesRINonFungibleTokens struct for ListDepositAddressesRINonFungibleTokens
type ListDepositAddressesRINonFungibleTokens struct {
	// Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract.
	Identifier string `json:"identifier"`
	// Defines the token's name as a string.
	Name string `json:"name"`
	// Defines the symbol of the non-fungible tokens.
	Symbol string `json:"symbol"`
	// Represents tokens' unique identifier.
	TokenId string `json:"tokenId"`
	// Defines the specific token type.
	Type string `json:"type"`
}

// NewListDepositAddressesRINonFungibleTokens instantiates a new ListDepositAddressesRINonFungibleTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDepositAddressesRINonFungibleTokens(identifier string, name string, symbol string, tokenId string, type_ string) *ListDepositAddressesRINonFungibleTokens {
	this := ListDepositAddressesRINonFungibleTokens{}
	this.Identifier = identifier
	this.Name = name
	this.Symbol = symbol
	this.TokenId = tokenId
	this.Type = type_
	return &this
}

// NewListDepositAddressesRINonFungibleTokensWithDefaults instantiates a new ListDepositAddressesRINonFungibleTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDepositAddressesRINonFungibleTokensWithDefaults() *ListDepositAddressesRINonFungibleTokens {
	this := ListDepositAddressesRINonFungibleTokens{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *ListDepositAddressesRINonFungibleTokens) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRINonFungibleTokens) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *ListDepositAddressesRINonFungibleTokens) SetIdentifier(v string) {
	o.Identifier = v
}

// GetName returns the Name field value
func (o *ListDepositAddressesRINonFungibleTokens) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRINonFungibleTokens) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListDepositAddressesRINonFungibleTokens) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *ListDepositAddressesRINonFungibleTokens) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRINonFungibleTokens) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ListDepositAddressesRINonFungibleTokens) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokenId returns the TokenId field value
func (o *ListDepositAddressesRINonFungibleTokens) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRINonFungibleTokens) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *ListDepositAddressesRINonFungibleTokens) SetTokenId(v string) {
	o.TokenId = v
}

// GetType returns the Type field value
func (o *ListDepositAddressesRINonFungibleTokens) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRINonFungibleTokens) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListDepositAddressesRINonFungibleTokens) SetType(v string) {
	o.Type = v
}

func (o ListDepositAddressesRINonFungibleTokens) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
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
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableListDepositAddressesRINonFungibleTokens struct {
	value *ListDepositAddressesRINonFungibleTokens
	isSet bool
}

func (v NullableListDepositAddressesRINonFungibleTokens) Get() *ListDepositAddressesRINonFungibleTokens {
	return v.value
}

func (v *NullableListDepositAddressesRINonFungibleTokens) Set(val *ListDepositAddressesRINonFungibleTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableListDepositAddressesRINonFungibleTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableListDepositAddressesRINonFungibleTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDepositAddressesRINonFungibleTokens(val *ListDepositAddressesRINonFungibleTokens) *NullableListDepositAddressesRINonFungibleTokens {
	return &NullableListDepositAddressesRINonFungibleTokens{value: val, isSet: true}
}

func (v NullableListDepositAddressesRINonFungibleTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDepositAddressesRINonFungibleTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


