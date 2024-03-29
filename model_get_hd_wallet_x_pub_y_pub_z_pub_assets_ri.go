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

// GetHDWalletXPubYPubZPubAssetsRI struct for GetHDWalletXPubYPubZPubAssetsRI
type GetHDWalletXPubYPubZPubAssetsRI struct {
	// Represents fungible tokens'es detailed information
	FungibleTokens []GetHDWalletXPubYPubZPubAssetsRIFungibleTokensInner `json:"fungibleTokens,omitempty"`
	// Represents non-fungible tokens'es detailed information.
	NonFungibleTokens []GetHDWalletXPubYPubZPubAssetsRINonFungibleTokensInner `json:"nonFungibleTokens,omitempty"`
	ConfirmedBalance GetHDWalletXPubYPubZPubAssetsRIConfirmedBalance `json:"confirmedBalance"`
}

// NewGetHDWalletXPubYPubZPubAssetsRI instantiates a new GetHDWalletXPubYPubZPubAssetsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHDWalletXPubYPubZPubAssetsRI(confirmedBalance GetHDWalletXPubYPubZPubAssetsRIConfirmedBalance) *GetHDWalletXPubYPubZPubAssetsRI {
	this := GetHDWalletXPubYPubZPubAssetsRI{}
	this.ConfirmedBalance = confirmedBalance
	return &this
}

// NewGetHDWalletXPubYPubZPubAssetsRIWithDefaults instantiates a new GetHDWalletXPubYPubZPubAssetsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHDWalletXPubYPubZPubAssetsRIWithDefaults() *GetHDWalletXPubYPubZPubAssetsRI {
	this := GetHDWalletXPubYPubZPubAssetsRI{}
	return &this
}

// GetFungibleTokens returns the FungibleTokens field value if set, zero value otherwise.
func (o *GetHDWalletXPubYPubZPubAssetsRI) GetFungibleTokens() []GetHDWalletXPubYPubZPubAssetsRIFungibleTokensInner {
	if o == nil || o.FungibleTokens == nil {
		var ret []GetHDWalletXPubYPubZPubAssetsRIFungibleTokensInner
		return ret
	}
	return o.FungibleTokens
}

// GetFungibleTokensOk returns a tuple with the FungibleTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRI) GetFungibleTokensOk() ([]GetHDWalletXPubYPubZPubAssetsRIFungibleTokensInner, bool) {
	if o == nil || o.FungibleTokens == nil {
		return nil, false
	}
	return o.FungibleTokens, true
}

// HasFungibleTokens returns a boolean if a field has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRI) HasFungibleTokens() bool {
	if o != nil && o.FungibleTokens != nil {
		return true
	}

	return false
}

// SetFungibleTokens gets a reference to the given []GetHDWalletXPubYPubZPubAssetsRIFungibleTokensInner and assigns it to the FungibleTokens field.
func (o *GetHDWalletXPubYPubZPubAssetsRI) SetFungibleTokens(v []GetHDWalletXPubYPubZPubAssetsRIFungibleTokensInner) {
	o.FungibleTokens = v
}

// GetNonFungibleTokens returns the NonFungibleTokens field value if set, zero value otherwise.
func (o *GetHDWalletXPubYPubZPubAssetsRI) GetNonFungibleTokens() []GetHDWalletXPubYPubZPubAssetsRINonFungibleTokensInner {
	if o == nil || o.NonFungibleTokens == nil {
		var ret []GetHDWalletXPubYPubZPubAssetsRINonFungibleTokensInner
		return ret
	}
	return o.NonFungibleTokens
}

// GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRI) GetNonFungibleTokensOk() ([]GetHDWalletXPubYPubZPubAssetsRINonFungibleTokensInner, bool) {
	if o == nil || o.NonFungibleTokens == nil {
		return nil, false
	}
	return o.NonFungibleTokens, true
}

// HasNonFungibleTokens returns a boolean if a field has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRI) HasNonFungibleTokens() bool {
	if o != nil && o.NonFungibleTokens != nil {
		return true
	}

	return false
}

// SetNonFungibleTokens gets a reference to the given []GetHDWalletXPubYPubZPubAssetsRINonFungibleTokensInner and assigns it to the NonFungibleTokens field.
func (o *GetHDWalletXPubYPubZPubAssetsRI) SetNonFungibleTokens(v []GetHDWalletXPubYPubZPubAssetsRINonFungibleTokensInner) {
	o.NonFungibleTokens = v
}

// GetConfirmedBalance returns the ConfirmedBalance field value
func (o *GetHDWalletXPubYPubZPubAssetsRI) GetConfirmedBalance() GetHDWalletXPubYPubZPubAssetsRIConfirmedBalance {
	if o == nil {
		var ret GetHDWalletXPubYPubZPubAssetsRIConfirmedBalance
		return ret
	}

	return o.ConfirmedBalance
}

// GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field value
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRI) GetConfirmedBalanceOk() (*GetHDWalletXPubYPubZPubAssetsRIConfirmedBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmedBalance, true
}

// SetConfirmedBalance sets field value
func (o *GetHDWalletXPubYPubZPubAssetsRI) SetConfirmedBalance(v GetHDWalletXPubYPubZPubAssetsRIConfirmedBalance) {
	o.ConfirmedBalance = v
}

func (o GetHDWalletXPubYPubZPubAssetsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FungibleTokens != nil {
		toSerialize["fungibleTokens"] = o.FungibleTokens
	}
	if o.NonFungibleTokens != nil {
		toSerialize["nonFungibleTokens"] = o.NonFungibleTokens
	}
	if true {
		toSerialize["confirmedBalance"] = o.ConfirmedBalance
	}
	return json.Marshal(toSerialize)
}

type NullableGetHDWalletXPubYPubZPubAssetsRI struct {
	value *GetHDWalletXPubYPubZPubAssetsRI
	isSet bool
}

func (v NullableGetHDWalletXPubYPubZPubAssetsRI) Get() *GetHDWalletXPubYPubZPubAssetsRI {
	return v.value
}

func (v *NullableGetHDWalletXPubYPubZPubAssetsRI) Set(val *GetHDWalletXPubYPubZPubAssetsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHDWalletXPubYPubZPubAssetsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHDWalletXPubYPubZPubAssetsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHDWalletXPubYPubZPubAssetsRI(val *GetHDWalletXPubYPubZPubAssetsRI) *NullableGetHDWalletXPubYPubZPubAssetsRI {
	return &NullableGetHDWalletXPubYPubZPubAssetsRI{value: val, isSet: true}
}

func (v NullableGetHDWalletXPubYPubZPubAssetsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHDWalletXPubYPubZPubAssetsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


