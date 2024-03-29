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

// GetWalletAssetDetailsRIConfirmedBalance Specifies the confirmed balance.
type GetWalletAssetDetailsRIConfirmedBalance struct {
	// Specifies the amount of the confirmed balance.
	Amount string `json:"amount"`
	// Specifies the unit of the amount of the confirmed balance.
	Unit string `json:"unit"`
}

// NewGetWalletAssetDetailsRIConfirmedBalance instantiates a new GetWalletAssetDetailsRIConfirmedBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletAssetDetailsRIConfirmedBalance(amount string, unit string) *GetWalletAssetDetailsRIConfirmedBalance {
	this := GetWalletAssetDetailsRIConfirmedBalance{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetWalletAssetDetailsRIConfirmedBalanceWithDefaults instantiates a new GetWalletAssetDetailsRIConfirmedBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletAssetDetailsRIConfirmedBalanceWithDefaults() *GetWalletAssetDetailsRIConfirmedBalance {
	this := GetWalletAssetDetailsRIConfirmedBalance{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetWalletAssetDetailsRIConfirmedBalance) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIConfirmedBalance) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetWalletAssetDetailsRIConfirmedBalance) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetWalletAssetDetailsRIConfirmedBalance) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIConfirmedBalance) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetWalletAssetDetailsRIConfirmedBalance) SetUnit(v string) {
	o.Unit = v
}

func (o GetWalletAssetDetailsRIConfirmedBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletAssetDetailsRIConfirmedBalance struct {
	value *GetWalletAssetDetailsRIConfirmedBalance
	isSet bool
}

func (v NullableGetWalletAssetDetailsRIConfirmedBalance) Get() *GetWalletAssetDetailsRIConfirmedBalance {
	return v.value
}

func (v *NullableGetWalletAssetDetailsRIConfirmedBalance) Set(val *GetWalletAssetDetailsRIConfirmedBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletAssetDetailsRIConfirmedBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletAssetDetailsRIConfirmedBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletAssetDetailsRIConfirmedBalance(val *GetWalletAssetDetailsRIConfirmedBalance) *NullableGetWalletAssetDetailsRIConfirmedBalance {
	return &NullableGetWalletAssetDetailsRIConfirmedBalance{value: val, isSet: true}
}

func (v NullableGetWalletAssetDetailsRIConfirmedBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletAssetDetailsRIConfirmedBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


