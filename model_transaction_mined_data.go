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

// TransactionMinedData Specifies all data, as attributes, included into the callback notification, which depends on the `event`.
type TransactionMinedData struct {
	// Represents the Crypto APIs 2.0 product which sends the callback.
	Product string `json:"product"`
	// Defines the specific event, for which a callback subscription is set.
	Event string `json:"event"`
	Item TransactionMinedDataItem `json:"item"`
}

// NewTransactionMinedData instantiates a new TransactionMinedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionMinedData(product string, event string, item TransactionMinedDataItem) *TransactionMinedData {
	this := TransactionMinedData{}
	this.Product = product
	this.Event = event
	this.Item = item
	return &this
}

// NewTransactionMinedDataWithDefaults instantiates a new TransactionMinedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionMinedDataWithDefaults() *TransactionMinedData {
	this := TransactionMinedData{}
	return &this
}

// GetProduct returns the Product field value
func (o *TransactionMinedData) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *TransactionMinedData) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *TransactionMinedData) SetProduct(v string) {
	o.Product = v
}

// GetEvent returns the Event field value
func (o *TransactionMinedData) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *TransactionMinedData) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *TransactionMinedData) SetEvent(v string) {
	o.Event = v
}

// GetItem returns the Item field value
func (o *TransactionMinedData) GetItem() TransactionMinedDataItem {
	if o == nil {
		var ret TransactionMinedDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *TransactionMinedData) GetItemOk() (*TransactionMinedDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *TransactionMinedData) SetItem(v TransactionMinedDataItem) {
	o.Item = v
}

func (o TransactionMinedData) MarshalJSON() ([]byte, error) {
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

type NullableTransactionMinedData struct {
	value *TransactionMinedData
	isSet bool
}

func (v NullableTransactionMinedData) Get() *TransactionMinedData {
	return v.value
}

func (v *NullableTransactionMinedData) Set(val *TransactionMinedData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMinedData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMinedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMinedData(val *TransactionMinedData) *NullableTransactionMinedData {
	return &NullableTransactionMinedData{value: val, isSet: true}
}

func (v NullableTransactionMinedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMinedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


