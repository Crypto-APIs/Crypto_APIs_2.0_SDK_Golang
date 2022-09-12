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

// GetXRPRippleTransactionDetailsByTransactionIDRIOffer Defines the specific offer.
type GetXRPRippleTransactionDetailsByTransactionIDRIOffer struct {
	// Defines the amount of the offer.
	Amount string `json:"amount"`
	// Defines the unit of the offer.
	Unit string `json:"unit"`
}

// NewGetXRPRippleTransactionDetailsByTransactionIDRIOffer instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRIOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleTransactionDetailsByTransactionIDRIOffer(amount string, unit string) *GetXRPRippleTransactionDetailsByTransactionIDRIOffer {
	this := GetXRPRippleTransactionDetailsByTransactionIDRIOffer{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetXRPRippleTransactionDetailsByTransactionIDRIOfferWithDefaults instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRIOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleTransactionDetailsByTransactionIDRIOfferWithDefaults() *GetXRPRippleTransactionDetailsByTransactionIDRIOffer {
	this := GetXRPRippleTransactionDetailsByTransactionIDRIOffer{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) SetUnit(v string) {
	o.Unit = v
}

func (o GetXRPRippleTransactionDetailsByTransactionIDRIOffer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer struct {
	value *GetXRPRippleTransactionDetailsByTransactionIDRIOffer
	isSet bool
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer) Get() *GetXRPRippleTransactionDetailsByTransactionIDRIOffer {
	return v.value
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer) Set(val *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer(val *GetXRPRippleTransactionDetailsByTransactionIDRIOffer) *NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer {
	return &NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer{value: val, isSet: true}
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRIOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


