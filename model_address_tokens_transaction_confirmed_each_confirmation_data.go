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

// AddressTokensTransactionConfirmedEachConfirmationData Specifies all data, as attributes, included into the callback notification, which depends on the `event`.
type AddressTokensTransactionConfirmedEachConfirmationData struct {
	// Represents the Crypto APIs 2.0 product which sends the callback.
	Product string `json:"product"`
	// Defines the specific event, for which a callback subscription is set.
	Event string `json:"event"`
	Item AddressTokensTransactionConfirmedEachConfirmationDataItem `json:"item"`
}

// NewAddressTokensTransactionConfirmedEachConfirmationData instantiates a new AddressTokensTransactionConfirmedEachConfirmationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTokensTransactionConfirmedEachConfirmationData(product string, event string, item AddressTokensTransactionConfirmedEachConfirmationDataItem) *AddressTokensTransactionConfirmedEachConfirmationData {
	this := AddressTokensTransactionConfirmedEachConfirmationData{}
	this.Product = product
	this.Event = event
	this.Item = item
	return &this
}

// NewAddressTokensTransactionConfirmedEachConfirmationDataWithDefaults instantiates a new AddressTokensTransactionConfirmedEachConfirmationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTokensTransactionConfirmedEachConfirmationDataWithDefaults() *AddressTokensTransactionConfirmedEachConfirmationData {
	this := AddressTokensTransactionConfirmedEachConfirmationData{}
	return &this
}

// GetProduct returns the Product field value
func (o *AddressTokensTransactionConfirmedEachConfirmationData) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationData) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationData) SetProduct(v string) {
	o.Product = v
}

// GetEvent returns the Event field value
func (o *AddressTokensTransactionConfirmedEachConfirmationData) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationData) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationData) SetEvent(v string) {
	o.Event = v
}

// GetItem returns the Item field value
func (o *AddressTokensTransactionConfirmedEachConfirmationData) GetItem() AddressTokensTransactionConfirmedEachConfirmationDataItem {
	if o == nil {
		var ret AddressTokensTransactionConfirmedEachConfirmationDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationData) GetItemOk() (*AddressTokensTransactionConfirmedEachConfirmationDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationData) SetItem(v AddressTokensTransactionConfirmedEachConfirmationDataItem) {
	o.Item = v
}

func (o AddressTokensTransactionConfirmedEachConfirmationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["product"] = o.Product
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableAddressTokensTransactionConfirmedEachConfirmationData struct {
	value *AddressTokensTransactionConfirmedEachConfirmationData
	isSet bool
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmationData) Get() *AddressTokensTransactionConfirmedEachConfirmationData {
	return v.value
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmationData) Set(val *AddressTokensTransactionConfirmedEachConfirmationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokensTransactionConfirmedEachConfirmationData(val *AddressTokensTransactionConfirmedEachConfirmationData) *NullableAddressTokensTransactionConfirmedEachConfirmationData {
	return &NullableAddressTokensTransactionConfirmedEachConfirmationData{value: val, isSet: true}
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


