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

// GetFeeAddressDetailsResponseItemBalance Specifies the balance of the fee address.
type GetFeeAddressDetailsResponseItemBalance struct {
	// Represents the amount of the units in fee address.
	Amount string `json:"amount"`
	// Represents the unit of the fee spent for the forwarded tokens, e.g. BTC.
	Unit string `json:"unit"`
}

// NewGetFeeAddressDetailsResponseItemBalance instantiates a new GetFeeAddressDetailsResponseItemBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeeAddressDetailsResponseItemBalance(amount string, unit string) *GetFeeAddressDetailsResponseItemBalance {
	this := GetFeeAddressDetailsResponseItemBalance{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetFeeAddressDetailsResponseItemBalanceWithDefaults instantiates a new GetFeeAddressDetailsResponseItemBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeeAddressDetailsResponseItemBalanceWithDefaults() *GetFeeAddressDetailsResponseItemBalance {
	this := GetFeeAddressDetailsResponseItemBalance{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetFeeAddressDetailsResponseItemBalance) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetFeeAddressDetailsResponseItemBalance) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetFeeAddressDetailsResponseItemBalance) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetFeeAddressDetailsResponseItemBalance) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetFeeAddressDetailsResponseItemBalance) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetFeeAddressDetailsResponseItemBalance) SetUnit(v string) {
	o.Unit = v
}

func (o GetFeeAddressDetailsResponseItemBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetFeeAddressDetailsResponseItemBalance struct {
	value *GetFeeAddressDetailsResponseItemBalance
	isSet bool
}

func (v NullableGetFeeAddressDetailsResponseItemBalance) Get() *GetFeeAddressDetailsResponseItemBalance {
	return v.value
}

func (v *NullableGetFeeAddressDetailsResponseItemBalance) Set(val *GetFeeAddressDetailsResponseItemBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeeAddressDetailsResponseItemBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeeAddressDetailsResponseItemBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeeAddressDetailsResponseItemBalance(val *GetFeeAddressDetailsResponseItemBalance) *NullableGetFeeAddressDetailsResponseItemBalance {
	return &NullableGetFeeAddressDetailsResponseItemBalance{value: val, isSet: true}
}

func (v NullableGetFeeAddressDetailsResponseItemBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFeeAddressDetailsResponseItemBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


