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

// TokensForwardingFailData Specifies all data, as attributes, included into the callback notification, which depends on the `event`.
type TokensForwardingFailData struct {
	// Represents the Crypto APIs 2.0 product which sends the callback.
	Product string `json:"product"`
	// Defines the specific event, for which a callback subscription is set.
	Event string `json:"event"`
	Item TokensForwardingFailDataItem `json:"item"`
}

// NewTokensForwardingFailData instantiates a new TokensForwardingFailData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokensForwardingFailData(product string, event string, item TokensForwardingFailDataItem) *TokensForwardingFailData {
	this := TokensForwardingFailData{}
	this.Product = product
	this.Event = event
	this.Item = item
	return &this
}

// NewTokensForwardingFailDataWithDefaults instantiates a new TokensForwardingFailData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokensForwardingFailDataWithDefaults() *TokensForwardingFailData {
	this := TokensForwardingFailData{}
	return &this
}

// GetProduct returns the Product field value
func (o *TokensForwardingFailData) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingFailData) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *TokensForwardingFailData) SetProduct(v string) {
	o.Product = v
}

// GetEvent returns the Event field value
func (o *TokensForwardingFailData) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingFailData) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *TokensForwardingFailData) SetEvent(v string) {
	o.Event = v
}

// GetItem returns the Item field value
func (o *TokensForwardingFailData) GetItem() TokensForwardingFailDataItem {
	if o == nil {
		var ret TokensForwardingFailDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingFailData) GetItemOk() (*TokensForwardingFailDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *TokensForwardingFailData) SetItem(v TokensForwardingFailDataItem) {
	o.Item = v
}

func (o TokensForwardingFailData) MarshalJSON() ([]byte, error) {
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

type NullableTokensForwardingFailData struct {
	value *TokensForwardingFailData
	isSet bool
}

func (v NullableTokensForwardingFailData) Get() *TokensForwardingFailData {
	return v.value
}

func (v *NullableTokensForwardingFailData) Set(val *TokensForwardingFailData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokensForwardingFailData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokensForwardingFailData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokensForwardingFailData(val *TokensForwardingFailData) *NullableTokensForwardingFailData {
	return &NullableTokensForwardingFailData{value: val, isSet: true}
}

func (v NullableTokensForwardingFailData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokensForwardingFailData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


