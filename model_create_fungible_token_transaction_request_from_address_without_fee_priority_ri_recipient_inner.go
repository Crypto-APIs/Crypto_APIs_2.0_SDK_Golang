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

// CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner struct for CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner
type CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner struct {
	// Defines the destination address.
	Address string `json:"address"`
	// Defines the amount sent to the destination address.
	Amount string `json:"amount"`
	// Represents the unit of the tokens amount recieved.
	Unit string `json:"unit"`
}

// NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner(address string, amount string, unit string) *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner {
	this := CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner{}
	this.Address = address
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInnerWithDefaults instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInnerWithDefaults() *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner {
	this := CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) SetUnit(v string) {
	o.Unit = v
}

func (o CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner struct {
	value *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner
	isSet bool
}

func (v NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) Get() *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner {
	return v.value
}

func (v *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) Set(val *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner(val *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner {
	return &NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner{value: val, isSet: true}
}

func (v NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


