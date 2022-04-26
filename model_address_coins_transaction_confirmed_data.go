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

// AddressCoinsTransactionConfirmedData Specifies all data, as attributes, included into the callback notification, which depends on the `event`.
type AddressCoinsTransactionConfirmedData struct {
	// Represents the Crypto APIs 2.0 product which sends the callback.
	Product string `json:"product"`
	// Defines the specific event, for which a callback subscription is set.
	Event string `json:"event"`
	Item AddressCoinsTransactionConfirmedDataItem `json:"item"`
}

// NewAddressCoinsTransactionConfirmedData instantiates a new AddressCoinsTransactionConfirmedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCoinsTransactionConfirmedData(product string, event string, item AddressCoinsTransactionConfirmedDataItem) *AddressCoinsTransactionConfirmedData {
	this := AddressCoinsTransactionConfirmedData{}
	this.Product = product
	this.Event = event
	this.Item = item
	return &this
}

// NewAddressCoinsTransactionConfirmedDataWithDefaults instantiates a new AddressCoinsTransactionConfirmedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCoinsTransactionConfirmedDataWithDefaults() *AddressCoinsTransactionConfirmedData {
	this := AddressCoinsTransactionConfirmedData{}
	return &this
}

// GetProduct returns the Product field value
func (o *AddressCoinsTransactionConfirmedData) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedData) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *AddressCoinsTransactionConfirmedData) SetProduct(v string) {
	o.Product = v
}

// GetEvent returns the Event field value
func (o *AddressCoinsTransactionConfirmedData) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedData) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AddressCoinsTransactionConfirmedData) SetEvent(v string) {
	o.Event = v
}

// GetItem returns the Item field value
func (o *AddressCoinsTransactionConfirmedData) GetItem() AddressCoinsTransactionConfirmedDataItem {
	if o == nil {
		var ret AddressCoinsTransactionConfirmedDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedData) GetItemOk() (*AddressCoinsTransactionConfirmedDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *AddressCoinsTransactionConfirmedData) SetItem(v AddressCoinsTransactionConfirmedDataItem) {
	o.Item = v
}

func (o AddressCoinsTransactionConfirmedData) MarshalJSON() ([]byte, error) {
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

type NullableAddressCoinsTransactionConfirmedData struct {
	value *AddressCoinsTransactionConfirmedData
	isSet bool
}

func (v NullableAddressCoinsTransactionConfirmedData) Get() *AddressCoinsTransactionConfirmedData {
	return v.value
}

func (v *NullableAddressCoinsTransactionConfirmedData) Set(val *AddressCoinsTransactionConfirmedData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCoinsTransactionConfirmedData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCoinsTransactionConfirmedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCoinsTransactionConfirmedData(val *AddressCoinsTransactionConfirmedData) *NullableAddressCoinsTransactionConfirmedData {
	return &NullableAddressCoinsTransactionConfirmedData{value: val, isSet: true}
}

func (v NullableAddressCoinsTransactionConfirmedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCoinsTransactionConfirmedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


