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

// ListAllAssetsByWalletIDRI struct for ListAllAssetsByWalletIDRI
type ListAllAssetsByWalletIDRI struct {
	Coins []ListAllAssetsFromAllWalletsRICoinsInner `json:"coins"`
	// Represents fungible tokens'es detailed information
	FungibleTokens []ListAllAssetsFromAllWalletsRIFungibleTokensInner `json:"fungibleTokens"`
	// Represents non-fungible tokens'es detailed information.
	NonFungibleTokens []ListAllAssetsFromAllWalletsRINonFungibleTokensInner `json:"nonFungibleTokens"`
	// Defines the unique ID of the Wallet.
	WalletId string `json:"walletId"`
	// Represents the name of the wallet.
	WalletName string `json:"walletName"`
}

// NewListAllAssetsByWalletIDRI instantiates a new ListAllAssetsByWalletIDRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllAssetsByWalletIDRI(coins []ListAllAssetsFromAllWalletsRICoinsInner, fungibleTokens []ListAllAssetsFromAllWalletsRIFungibleTokensInner, nonFungibleTokens []ListAllAssetsFromAllWalletsRINonFungibleTokensInner, walletId string, walletName string) *ListAllAssetsByWalletIDRI {
	this := ListAllAssetsByWalletIDRI{}
	this.Coins = coins
	this.FungibleTokens = fungibleTokens
	this.NonFungibleTokens = nonFungibleTokens
	this.WalletId = walletId
	this.WalletName = walletName
	return &this
}

// NewListAllAssetsByWalletIDRIWithDefaults instantiates a new ListAllAssetsByWalletIDRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllAssetsByWalletIDRIWithDefaults() *ListAllAssetsByWalletIDRI {
	this := ListAllAssetsByWalletIDRI{}
	return &this
}

// GetCoins returns the Coins field value
func (o *ListAllAssetsByWalletIDRI) GetCoins() []ListAllAssetsFromAllWalletsRICoinsInner {
	if o == nil {
		var ret []ListAllAssetsFromAllWalletsRICoinsInner
		return ret
	}

	return o.Coins
}

// GetCoinsOk returns a tuple with the Coins field value
// and a boolean to check if the value has been set.
func (o *ListAllAssetsByWalletIDRI) GetCoinsOk() ([]ListAllAssetsFromAllWalletsRICoinsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coins, true
}

// SetCoins sets field value
func (o *ListAllAssetsByWalletIDRI) SetCoins(v []ListAllAssetsFromAllWalletsRICoinsInner) {
	o.Coins = v
}

// GetFungibleTokens returns the FungibleTokens field value
func (o *ListAllAssetsByWalletIDRI) GetFungibleTokens() []ListAllAssetsFromAllWalletsRIFungibleTokensInner {
	if o == nil {
		var ret []ListAllAssetsFromAllWalletsRIFungibleTokensInner
		return ret
	}

	return o.FungibleTokens
}

// GetFungibleTokensOk returns a tuple with the FungibleTokens field value
// and a boolean to check if the value has been set.
func (o *ListAllAssetsByWalletIDRI) GetFungibleTokensOk() ([]ListAllAssetsFromAllWalletsRIFungibleTokensInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.FungibleTokens, true
}

// SetFungibleTokens sets field value
func (o *ListAllAssetsByWalletIDRI) SetFungibleTokens(v []ListAllAssetsFromAllWalletsRIFungibleTokensInner) {
	o.FungibleTokens = v
}

// GetNonFungibleTokens returns the NonFungibleTokens field value
func (o *ListAllAssetsByWalletIDRI) GetNonFungibleTokens() []ListAllAssetsFromAllWalletsRINonFungibleTokensInner {
	if o == nil {
		var ret []ListAllAssetsFromAllWalletsRINonFungibleTokensInner
		return ret
	}

	return o.NonFungibleTokens
}

// GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field value
// and a boolean to check if the value has been set.
func (o *ListAllAssetsByWalletIDRI) GetNonFungibleTokensOk() ([]ListAllAssetsFromAllWalletsRINonFungibleTokensInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonFungibleTokens, true
}

// SetNonFungibleTokens sets field value
func (o *ListAllAssetsByWalletIDRI) SetNonFungibleTokens(v []ListAllAssetsFromAllWalletsRINonFungibleTokensInner) {
	o.NonFungibleTokens = v
}

// GetWalletId returns the WalletId field value
func (o *ListAllAssetsByWalletIDRI) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *ListAllAssetsByWalletIDRI) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *ListAllAssetsByWalletIDRI) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletName returns the WalletName field value
func (o *ListAllAssetsByWalletIDRI) GetWalletName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletName
}

// GetWalletNameOk returns a tuple with the WalletName field value
// and a boolean to check if the value has been set.
func (o *ListAllAssetsByWalletIDRI) GetWalletNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletName, true
}

// SetWalletName sets field value
func (o *ListAllAssetsByWalletIDRI) SetWalletName(v string) {
	o.WalletName = v
}

func (o ListAllAssetsByWalletIDRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["coins"] = o.Coins
	}
	if true {
		toSerialize["fungibleTokens"] = o.FungibleTokens
	}
	if true {
		toSerialize["nonFungibleTokens"] = o.NonFungibleTokens
	}
	if true {
		toSerialize["walletId"] = o.WalletId
	}
	if true {
		toSerialize["walletName"] = o.WalletName
	}
	return json.Marshal(toSerialize)
}

type NullableListAllAssetsByWalletIDRI struct {
	value *ListAllAssetsByWalletIDRI
	isSet bool
}

func (v NullableListAllAssetsByWalletIDRI) Get() *ListAllAssetsByWalletIDRI {
	return v.value
}

func (v *NullableListAllAssetsByWalletIDRI) Set(val *ListAllAssetsByWalletIDRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllAssetsByWalletIDRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllAssetsByWalletIDRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllAssetsByWalletIDRI(val *ListAllAssetsByWalletIDRI) *NullableListAllAssetsByWalletIDRI {
	return &NullableListAllAssetsByWalletIDRI{value: val, isSet: true}
}

func (v NullableListAllAssetsByWalletIDRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllAssetsByWalletIDRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


