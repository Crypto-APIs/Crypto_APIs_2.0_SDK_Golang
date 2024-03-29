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

// GetWalletAssetDetailsRIRecievedConfirmedAmount Specifies the confirmed amount that has been received.
type GetWalletAssetDetailsRIRecievedConfirmedAmount struct {
	// Specifies the confirmed amount that has been received.
	Amount string `json:"amount"`
	// Specifies the unit of the confirmed amount that has been received.
	Unit string `json:"unit"`
}

// NewGetWalletAssetDetailsRIRecievedConfirmedAmount instantiates a new GetWalletAssetDetailsRIRecievedConfirmedAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletAssetDetailsRIRecievedConfirmedAmount(amount string, unit string) *GetWalletAssetDetailsRIRecievedConfirmedAmount {
	this := GetWalletAssetDetailsRIRecievedConfirmedAmount{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetWalletAssetDetailsRIRecievedConfirmedAmountWithDefaults instantiates a new GetWalletAssetDetailsRIRecievedConfirmedAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletAssetDetailsRIRecievedConfirmedAmountWithDefaults() *GetWalletAssetDetailsRIRecievedConfirmedAmount {
	this := GetWalletAssetDetailsRIRecievedConfirmedAmount{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetWalletAssetDetailsRIRecievedConfirmedAmount) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIRecievedConfirmedAmount) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetWalletAssetDetailsRIRecievedConfirmedAmount) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetWalletAssetDetailsRIRecievedConfirmedAmount) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRIRecievedConfirmedAmount) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetWalletAssetDetailsRIRecievedConfirmedAmount) SetUnit(v string) {
	o.Unit = v
}

func (o GetWalletAssetDetailsRIRecievedConfirmedAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletAssetDetailsRIRecievedConfirmedAmount struct {
	value *GetWalletAssetDetailsRIRecievedConfirmedAmount
	isSet bool
}

func (v NullableGetWalletAssetDetailsRIRecievedConfirmedAmount) Get() *GetWalletAssetDetailsRIRecievedConfirmedAmount {
	return v.value
}

func (v *NullableGetWalletAssetDetailsRIRecievedConfirmedAmount) Set(val *GetWalletAssetDetailsRIRecievedConfirmedAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletAssetDetailsRIRecievedConfirmedAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletAssetDetailsRIRecievedConfirmedAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletAssetDetailsRIRecievedConfirmedAmount(val *GetWalletAssetDetailsRIRecievedConfirmedAmount) *NullableGetWalletAssetDetailsRIRecievedConfirmedAmount {
	return &NullableGetWalletAssetDetailsRIRecievedConfirmedAmount{value: val, isSet: true}
}

func (v NullableGetWalletAssetDetailsRIRecievedConfirmedAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletAssetDetailsRIRecievedConfirmedAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


