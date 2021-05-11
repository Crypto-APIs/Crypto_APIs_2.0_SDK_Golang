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

// NewConfirmedTokensTransactionsAndEachConfirmationResponseData struct for NewConfirmedTokensTransactionsAndEachConfirmationResponseData
type NewConfirmedTokensTransactionsAndEachConfirmationResponseData struct {
	Item NewConfirmedTokensTransactionsAndEachConfirmationResponseItem `json:"item"`
}

// NewNewConfirmedTokensTransactionsAndEachConfirmationResponseData instantiates a new NewConfirmedTokensTransactionsAndEachConfirmationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedTokensTransactionsAndEachConfirmationResponseData(item NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) *NewConfirmedTokensTransactionsAndEachConfirmationResponseData {
	this := NewConfirmedTokensTransactionsAndEachConfirmationResponseData{}
	this.Item = item
	return &this
}

// NewNewConfirmedTokensTransactionsAndEachConfirmationResponseDataWithDefaults instantiates a new NewConfirmedTokensTransactionsAndEachConfirmationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedTokensTransactionsAndEachConfirmationResponseDataWithDefaults() *NewConfirmedTokensTransactionsAndEachConfirmationResponseData {
	this := NewConfirmedTokensTransactionsAndEachConfirmationResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseData) GetItem() NewConfirmedTokensTransactionsAndEachConfirmationResponseItem {
	if o == nil {
		var ret NewConfirmedTokensTransactionsAndEachConfirmationResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseData) GetItemOk() (*NewConfirmedTokensTransactionsAndEachConfirmationResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseData) SetItem(v NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) {
	o.Item = v
}

func (o NewConfirmedTokensTransactionsAndEachConfirmationResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData struct {
	value *NewConfirmedTokensTransactionsAndEachConfirmationResponseData
	isSet bool
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData) Get() *NewConfirmedTokensTransactionsAndEachConfirmationResponseData {
	return v.value
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData) Set(val *NewConfirmedTokensTransactionsAndEachConfirmationResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData(val *NewConfirmedTokensTransactionsAndEachConfirmationResponseData) *NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData {
	return &NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData{value: val, isSet: true}
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


