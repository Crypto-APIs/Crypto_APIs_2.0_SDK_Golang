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

// GetXRPRippleAddressDetailsRIBalance Defines the balance of the account.
type GetXRPRippleAddressDetailsRIBalance struct {
	// Represents the total amount of the balance.
	Amount string `json:"amount"`
	// Represents the unit used for the balance.
	Unit string `json:"unit"`
}

// NewGetXRPRippleAddressDetailsRIBalance instantiates a new GetXRPRippleAddressDetailsRIBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleAddressDetailsRIBalance(amount string, unit string) *GetXRPRippleAddressDetailsRIBalance {
	this := GetXRPRippleAddressDetailsRIBalance{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetXRPRippleAddressDetailsRIBalanceWithDefaults instantiates a new GetXRPRippleAddressDetailsRIBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleAddressDetailsRIBalanceWithDefaults() *GetXRPRippleAddressDetailsRIBalance {
	this := GetXRPRippleAddressDetailsRIBalance{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetXRPRippleAddressDetailsRIBalance) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleAddressDetailsRIBalance) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetXRPRippleAddressDetailsRIBalance) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetXRPRippleAddressDetailsRIBalance) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleAddressDetailsRIBalance) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetXRPRippleAddressDetailsRIBalance) SetUnit(v string) {
	o.Unit = v
}

func (o GetXRPRippleAddressDetailsRIBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleAddressDetailsRIBalance struct {
	value *GetXRPRippleAddressDetailsRIBalance
	isSet bool
}

func (v NullableGetXRPRippleAddressDetailsRIBalance) Get() *GetXRPRippleAddressDetailsRIBalance {
	return v.value
}

func (v *NullableGetXRPRippleAddressDetailsRIBalance) Set(val *GetXRPRippleAddressDetailsRIBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleAddressDetailsRIBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleAddressDetailsRIBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleAddressDetailsRIBalance(val *GetXRPRippleAddressDetailsRIBalance) *NullableGetXRPRippleAddressDetailsRIBalance {
	return &NullableGetXRPRippleAddressDetailsRIBalance{value: val, isSet: true}
}

func (v NullableGetXRPRippleAddressDetailsRIBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleAddressDetailsRIBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


