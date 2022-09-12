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

// ListXRPRippleTransactionsByBlockHashRIOffer struct for ListXRPRippleTransactionsByBlockHashRIOffer
type ListXRPRippleTransactionsByBlockHashRIOffer struct {
	// Defines the amount of the offer.
	Amount string `json:"amount"`
	// Defines the unit of the offer.
	Unit string `json:"unit"`
}

// NewListXRPRippleTransactionsByBlockHashRIOffer instantiates a new ListXRPRippleTransactionsByBlockHashRIOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByBlockHashRIOffer(amount string, unit string) *ListXRPRippleTransactionsByBlockHashRIOffer {
	this := ListXRPRippleTransactionsByBlockHashRIOffer{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewListXRPRippleTransactionsByBlockHashRIOfferWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHashRIOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByBlockHashRIOfferWithDefaults() *ListXRPRippleTransactionsByBlockHashRIOffer {
	this := ListXRPRippleTransactionsByBlockHashRIOffer{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListXRPRippleTransactionsByBlockHashRIOffer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHashRIOffer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListXRPRippleTransactionsByBlockHashRIOffer) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ListXRPRippleTransactionsByBlockHashRIOffer) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHashRIOffer) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ListXRPRippleTransactionsByBlockHashRIOffer) SetUnit(v string) {
	o.Unit = v
}

func (o ListXRPRippleTransactionsByBlockHashRIOffer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListXRPRippleTransactionsByBlockHashRIOffer struct {
	value *ListXRPRippleTransactionsByBlockHashRIOffer
	isSet bool
}

func (v NullableListXRPRippleTransactionsByBlockHashRIOffer) Get() *ListXRPRippleTransactionsByBlockHashRIOffer {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByBlockHashRIOffer) Set(val *ListXRPRippleTransactionsByBlockHashRIOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByBlockHashRIOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByBlockHashRIOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByBlockHashRIOffer(val *ListXRPRippleTransactionsByBlockHashRIOffer) *NullableListXRPRippleTransactionsByBlockHashRIOffer {
	return &NullableListXRPRippleTransactionsByBlockHashRIOffer{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByBlockHashRIOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByBlockHashRIOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


