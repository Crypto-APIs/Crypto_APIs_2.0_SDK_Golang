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

// EstimateGasLimitRBDataItem struct for EstimateGasLimitRBDataItem
type EstimateGasLimitRBDataItem struct {
	// Represents an optional field to add additonal data.
	AdditionalData *string `json:"additionalData,omitempty"`
	// Represents transactions' amount.
	Amount string `json:"amount"`
	// The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient.
	Recipient string `json:"recipient"`
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Sender string `json:"sender"`
}

// NewEstimateGasLimitRBDataItem instantiates a new EstimateGasLimitRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateGasLimitRBDataItem(amount string, recipient string, sender string) *EstimateGasLimitRBDataItem {
	this := EstimateGasLimitRBDataItem{}
	this.Amount = amount
	this.Recipient = recipient
	this.Sender = sender
	return &this
}

// NewEstimateGasLimitRBDataItemWithDefaults instantiates a new EstimateGasLimitRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateGasLimitRBDataItemWithDefaults() *EstimateGasLimitRBDataItem {
	this := EstimateGasLimitRBDataItem{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *EstimateGasLimitRBDataItem) GetAdditionalData() string {
	if o == nil || o.AdditionalData == nil {
		var ret string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateGasLimitRBDataItem) GetAdditionalDataOk() (*string, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *EstimateGasLimitRBDataItem) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given string and assigns it to the AdditionalData field.
func (o *EstimateGasLimitRBDataItem) SetAdditionalData(v string) {
	o.AdditionalData = &v
}

// GetAmount returns the Amount field value
func (o *EstimateGasLimitRBDataItem) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *EstimateGasLimitRBDataItem) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *EstimateGasLimitRBDataItem) SetAmount(v string) {
	o.Amount = v
}

// GetRecipient returns the Recipient field value
func (o *EstimateGasLimitRBDataItem) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *EstimateGasLimitRBDataItem) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *EstimateGasLimitRBDataItem) SetRecipient(v string) {
	o.Recipient = v
}

// GetSender returns the Sender field value
func (o *EstimateGasLimitRBDataItem) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *EstimateGasLimitRBDataItem) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *EstimateGasLimitRBDataItem) SetSender(v string) {
	o.Sender = v
}

func (o EstimateGasLimitRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	return json.Marshal(toSerialize)
}

type NullableEstimateGasLimitRBDataItem struct {
	value *EstimateGasLimitRBDataItem
	isSet bool
}

func (v NullableEstimateGasLimitRBDataItem) Get() *EstimateGasLimitRBDataItem {
	return v.value
}

func (v *NullableEstimateGasLimitRBDataItem) Set(val *EstimateGasLimitRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateGasLimitRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateGasLimitRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateGasLimitRBDataItem(val *EstimateGasLimitRBDataItem) *NullableEstimateGasLimitRBDataItem {
	return &NullableEstimateGasLimitRBDataItem{value: val, isSet: true}
}

func (v NullableEstimateGasLimitRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateGasLimitRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


