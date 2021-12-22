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

// GetWalletAssetDetailsRIFungibleTokens struct for GetWalletAssetDetailsRIFungibleTokens
type GetWalletAssetDetailsRIFungibleTokens struct {
	// Defines the amount of the fungible tokens.
	ConfirmedAmount string `json:"confirmedAmount"`
	// Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract.
	Identifier string `json:"identifier"`
	// Defines the symbol of the fungible tokens.
	Symbol string `json:"symbol"`
	// Defines the specific token type.
	Type string `json:"type"`
}

// NewGetWalletAssetDetailsRIFungibleTokens instantiates a new GetWalletAssetDetailsRIFungibleTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletAssetDetailsRIFungibleTokens(confirmedAmount string, identifier string, symbol string, type_ string) *GetWalletAssetDetailsRIFungibleTokens {
	this := GetWalletAssetDetailsRIFungibleTokens{}
	this.ConfirmedAmount = confirmedAmount
	this.Identifier = identifier
	this.Symbol = symbol
	this.Type = type_
	return &this
}

// NewGetWalletAssetDetailsRIFungibleTokensWithDefaults instantiates a new GetWalletAssetDetailsRIFungibleTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletAssetDetailsRIFungibleTokensWithDefaults() *GetWalletAssetDetailsRIFungibleTokens {
	this := GetWalletAssetDetailsRIFungibleTokens{}
	return &this
}

// GetConfirmedAmount returns the ConfirmedAmount field value
func (o *GetWalletAssetDetailsRIFungibleTokens) GetConfirmedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmedAmount
}

// GetConfirmedAmountOk returns a tuple with the ConfirmedAmount field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIFungibleTokens) GetConfirmedAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmedAmount, true
}

// SetConfirmedAmount sets field value
func (o *GetWalletAssetDetailsRIFungibleTokens) SetConfirmedAmount(v string) {
	o.ConfirmedAmount = v
}

// GetIdentifier returns the Identifier field value
func (o *GetWalletAssetDetailsRIFungibleTokens) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIFungibleTokens) GetIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GetWalletAssetDetailsRIFungibleTokens) SetIdentifier(v string) {
	o.Identifier = v
}

// GetSymbol returns the Symbol field value
func (o *GetWalletAssetDetailsRIFungibleTokens) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIFungibleTokens) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetWalletAssetDetailsRIFungibleTokens) SetSymbol(v string) {
	o.Symbol = v
}

// GetType returns the Type field value
func (o *GetWalletAssetDetailsRIFungibleTokens) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIFungibleTokens) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetWalletAssetDetailsRIFungibleTokens) SetType(v string) {
	o.Type = v
}

func (o GetWalletAssetDetailsRIFungibleTokens) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["confirmedAmount"] = o.ConfirmedAmount
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletAssetDetailsRIFungibleTokens struct {
	value *GetWalletAssetDetailsRIFungibleTokens
	isSet bool
}

func (v NullableGetWalletAssetDetailsRIFungibleTokens) Get() *GetWalletAssetDetailsRIFungibleTokens {
	return v.value
}

func (v *NullableGetWalletAssetDetailsRIFungibleTokens) Set(val *GetWalletAssetDetailsRIFungibleTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletAssetDetailsRIFungibleTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletAssetDetailsRIFungibleTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletAssetDetailsRIFungibleTokens(val *GetWalletAssetDetailsRIFungibleTokens) *NullableGetWalletAssetDetailsRIFungibleTokens {
	return &NullableGetWalletAssetDetailsRIFungibleTokens{value: val, isSet: true}
}

func (v NullableGetWalletAssetDetailsRIFungibleTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletAssetDetailsRIFungibleTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

