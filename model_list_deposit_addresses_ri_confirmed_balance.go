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

// ListDepositAddressesRIConfirmedBalance struct for ListDepositAddressesRIConfirmedBalance
type ListDepositAddressesRIConfirmedBalance struct {
	// Defines the total balance of the address that is confirmed. It doesn't include unconfirmed transactions.
	Amount string `json:"amount"`
	// Represents the unit of the confirmed balance.
	Unit string `json:"unit"`
}

// NewListDepositAddressesRIConfirmedBalance instantiates a new ListDepositAddressesRIConfirmedBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDepositAddressesRIConfirmedBalance(amount string, unit string) *ListDepositAddressesRIConfirmedBalance {
	this := ListDepositAddressesRIConfirmedBalance{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewListDepositAddressesRIConfirmedBalanceWithDefaults instantiates a new ListDepositAddressesRIConfirmedBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDepositAddressesRIConfirmedBalanceWithDefaults() *ListDepositAddressesRIConfirmedBalance {
	this := ListDepositAddressesRIConfirmedBalance{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListDepositAddressesRIConfirmedBalance) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRIConfirmedBalance) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListDepositAddressesRIConfirmedBalance) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ListDepositAddressesRIConfirmedBalance) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRIConfirmedBalance) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ListDepositAddressesRIConfirmedBalance) SetUnit(v string) {
	o.Unit = v
}

func (o ListDepositAddressesRIConfirmedBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListDepositAddressesRIConfirmedBalance struct {
	value *ListDepositAddressesRIConfirmedBalance
	isSet bool
}

func (v NullableListDepositAddressesRIConfirmedBalance) Get() *ListDepositAddressesRIConfirmedBalance {
	return v.value
}

func (v *NullableListDepositAddressesRIConfirmedBalance) Set(val *ListDepositAddressesRIConfirmedBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableListDepositAddressesRIConfirmedBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableListDepositAddressesRIConfirmedBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDepositAddressesRIConfirmedBalance(val *ListDepositAddressesRIConfirmedBalance) *NullableListDepositAddressesRIConfirmedBalance {
	return &NullableListDepositAddressesRIConfirmedBalance{value: val, isSet: true}
}

func (v NullableListDepositAddressesRIConfirmedBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDepositAddressesRIConfirmedBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


