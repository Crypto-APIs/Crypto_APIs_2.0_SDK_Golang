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

// PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee struct for PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee
type PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee struct {
	// Representation of the exact amount value
	ExactAmount *string `json:"exactAmount,omitempty"`
	// Represents the fee priority
	Priority *string `json:"priority,omitempty"`
}

// NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee() *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee {
	this := PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee{}
	return &this
}

// NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFeeWithDefaults instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFeeWithDefaults() *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee {
	this := PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee{}
	return &this
}

// GetExactAmount returns the ExactAmount field value if set, zero value otherwise.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) GetExactAmount() string {
	if o == nil || o.ExactAmount == nil {
		var ret string
		return ret
	}
	return *o.ExactAmount
}

// GetExactAmountOk returns a tuple with the ExactAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) GetExactAmountOk() (*string, bool) {
	if o == nil || o.ExactAmount == nil {
		return nil, false
	}
	return o.ExactAmount, true
}

// HasExactAmount returns a boolean if a field has been set.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) HasExactAmount() bool {
	if o != nil && o.ExactAmount != nil {
		return true
	}

	return false
}

// SetExactAmount gets a reference to the given string and assigns it to the ExactAmount field.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) SetExactAmount(v string) {
	o.ExactAmount = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) SetPriority(v string) {
	o.Priority = &v
}

func (o PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExactAmount != nil {
		toSerialize["exactAmount"] = o.ExactAmount
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee struct {
	value *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee
	isSet bool
}

func (v NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) Get() *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee {
	return v.value
}

func (v *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) Set(val *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee(val *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee {
	return &NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee{value: val, isSet: true}
}

func (v NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


