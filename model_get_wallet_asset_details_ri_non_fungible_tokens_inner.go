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

// GetWalletAssetDetailsRINonFungibleTokensInner struct for GetWalletAssetDetailsRINonFungibleTokensInner
type GetWalletAssetDetailsRINonFungibleTokensInner struct {
	// Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract.
	Identifier string `json:"identifier"`
	// Defines the symbol of the non-fungible tokens.
	Symbol string `json:"symbol"`
	// Represents tokens' unique identifier.
	TokenId string `json:"tokenId"`
	// Defines the specific token type.
	Type string `json:"type"`
}

// NewGetWalletAssetDetailsRINonFungibleTokensInner instantiates a new GetWalletAssetDetailsRINonFungibleTokensInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletAssetDetailsRINonFungibleTokensInner(identifier string, symbol string, tokenId string, type_ string) *GetWalletAssetDetailsRINonFungibleTokensInner {
	this := GetWalletAssetDetailsRINonFungibleTokensInner{}
	this.Identifier = identifier
	this.Symbol = symbol
	this.TokenId = tokenId
	this.Type = type_
	return &this
}

// NewGetWalletAssetDetailsRINonFungibleTokensInnerWithDefaults instantiates a new GetWalletAssetDetailsRINonFungibleTokensInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletAssetDetailsRINonFungibleTokensInnerWithDefaults() *GetWalletAssetDetailsRINonFungibleTokensInner {
	this := GetWalletAssetDetailsRINonFungibleTokensInner{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) SetIdentifier(v string) {
	o.Identifier = v
}

// GetSymbol returns the Symbol field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokenId returns the TokenId field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) SetTokenId(v string) {
	o.TokenId = v
}

// GetType returns the Type field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetWalletAssetDetailsRINonFungibleTokensInner) SetType(v string) {
	o.Type = v
}

func (o GetWalletAssetDetailsRINonFungibleTokensInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
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

type NullableGetWalletAssetDetailsRINonFungibleTokensInner struct {
	value *GetWalletAssetDetailsRINonFungibleTokensInner
	isSet bool
}

func (v NullableGetWalletAssetDetailsRINonFungibleTokensInner) Get() *GetWalletAssetDetailsRINonFungibleTokensInner {
	return v.value
}

func (v *NullableGetWalletAssetDetailsRINonFungibleTokensInner) Set(val *GetWalletAssetDetailsRINonFungibleTokensInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletAssetDetailsRINonFungibleTokensInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletAssetDetailsRINonFungibleTokensInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletAssetDetailsRINonFungibleTokensInner(val *GetWalletAssetDetailsRINonFungibleTokensInner) *NullableGetWalletAssetDetailsRINonFungibleTokensInner {
	return &NullableGetWalletAssetDetailsRINonFungibleTokensInner{value: val, isSet: true}
}

func (v NullableGetWalletAssetDetailsRINonFungibleTokensInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletAssetDetailsRINonFungibleTokensInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


