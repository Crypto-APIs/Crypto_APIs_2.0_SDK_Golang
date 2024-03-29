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

// CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount struct for CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount
type CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount struct {
	// Defines the unit of the total amount transacted.
	Unit *string `json:"unit,omitempty"`
	// Total amount of the transaction
	Value *string `json:"value,omitempty"`
}

// NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount {
	this := CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount{}
	return &this
}

// NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmountWithDefaults instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmountWithDefaults() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount {
	this := CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) SetUnit(v string) {
	o.Unit = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) SetValue(v string) {
	o.Value = &v
}

func (o CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount struct {
	value *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount
	isSet bool
}

func (v NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) Get() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount {
	return v.value
}

func (v *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) Set(val *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount(val *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount {
	return &NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount{value: val, isSet: true}
}

func (v NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


