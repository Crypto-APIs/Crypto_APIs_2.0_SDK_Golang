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

// ListAllUnconfirmedTransactionsRIBSEFee Object representation of the transaction fee
type ListAllUnconfirmedTransactionsRIBSEFee struct {
	// String representation of the fee value
	Amount string `json:"amount"`
}

// NewListAllUnconfirmedTransactionsRIBSEFee instantiates a new ListAllUnconfirmedTransactionsRIBSEFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllUnconfirmedTransactionsRIBSEFee(amount string) *ListAllUnconfirmedTransactionsRIBSEFee {
	this := ListAllUnconfirmedTransactionsRIBSEFee{}
	this.Amount = amount
	return &this
}

// NewListAllUnconfirmedTransactionsRIBSEFeeWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSEFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllUnconfirmedTransactionsRIBSEFeeWithDefaults() *ListAllUnconfirmedTransactionsRIBSEFee {
	this := ListAllUnconfirmedTransactionsRIBSEFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListAllUnconfirmedTransactionsRIBSEFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSEFee) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListAllUnconfirmedTransactionsRIBSEFee) SetAmount(v string) {
	o.Amount = v
}

func (o ListAllUnconfirmedTransactionsRIBSEFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListAllUnconfirmedTransactionsRIBSEFee struct {
	value *ListAllUnconfirmedTransactionsRIBSEFee
	isSet bool
}

func (v NullableListAllUnconfirmedTransactionsRIBSEFee) Get() *ListAllUnconfirmedTransactionsRIBSEFee {
	return v.value
}

func (v *NullableListAllUnconfirmedTransactionsRIBSEFee) Set(val *ListAllUnconfirmedTransactionsRIBSEFee) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllUnconfirmedTransactionsRIBSEFee) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllUnconfirmedTransactionsRIBSEFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllUnconfirmedTransactionsRIBSEFee(val *ListAllUnconfirmedTransactionsRIBSEFee) *NullableListAllUnconfirmedTransactionsRIBSEFee {
	return &NullableListAllUnconfirmedTransactionsRIBSEFee{value: val, isSet: true}
}

func (v NullableListAllUnconfirmedTransactionsRIBSEFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllUnconfirmedTransactionsRIBSEFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


