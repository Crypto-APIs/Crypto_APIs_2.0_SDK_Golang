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

// ListXRPRippleTransactionsByAddressRIOffer struct for ListXRPRippleTransactionsByAddressRIOffer
type ListXRPRippleTransactionsByAddressRIOffer struct {
	// Defines the amount of the offer.
	Amount string `json:"amount"`
	// Defines the unit of the offer.
	Unit string `json:"unit"`
}

// NewListXRPRippleTransactionsByAddressRIOffer instantiates a new ListXRPRippleTransactionsByAddressRIOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByAddressRIOffer(amount string, unit string) *ListXRPRippleTransactionsByAddressRIOffer {
	this := ListXRPRippleTransactionsByAddressRIOffer{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewListXRPRippleTransactionsByAddressRIOfferWithDefaults instantiates a new ListXRPRippleTransactionsByAddressRIOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByAddressRIOfferWithDefaults() *ListXRPRippleTransactionsByAddressRIOffer {
	this := ListXRPRippleTransactionsByAddressRIOffer{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListXRPRippleTransactionsByAddressRIOffer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressRIOffer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListXRPRippleTransactionsByAddressRIOffer) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ListXRPRippleTransactionsByAddressRIOffer) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressRIOffer) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ListXRPRippleTransactionsByAddressRIOffer) SetUnit(v string) {
	o.Unit = v
}

func (o ListXRPRippleTransactionsByAddressRIOffer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListXRPRippleTransactionsByAddressRIOffer struct {
	value *ListXRPRippleTransactionsByAddressRIOffer
	isSet bool
}

func (v NullableListXRPRippleTransactionsByAddressRIOffer) Get() *ListXRPRippleTransactionsByAddressRIOffer {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByAddressRIOffer) Set(val *ListXRPRippleTransactionsByAddressRIOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByAddressRIOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByAddressRIOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByAddressRIOffer(val *ListXRPRippleTransactionsByAddressRIOffer) *NullableListXRPRippleTransactionsByAddressRIOffer {
	return &NullableListXRPRippleTransactionsByAddressRIOffer{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByAddressRIOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByAddressRIOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


