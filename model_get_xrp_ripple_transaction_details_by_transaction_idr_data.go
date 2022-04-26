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

// GetXRPRippleTransactionDetailsByTransactionIDRData struct for GetXRPRippleTransactionDetailsByTransactionIDRData
type GetXRPRippleTransactionDetailsByTransactionIDRData struct {
	Item GetXRPRippleTransactionDetailsByTransactionIDRI `json:"item"`
}

// NewGetXRPRippleTransactionDetailsByTransactionIDRData instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleTransactionDetailsByTransactionIDRData(item GetXRPRippleTransactionDetailsByTransactionIDRI) *GetXRPRippleTransactionDetailsByTransactionIDRData {
	this := GetXRPRippleTransactionDetailsByTransactionIDRData{}
	this.Item = item
	return &this
}

// NewGetXRPRippleTransactionDetailsByTransactionIDRDataWithDefaults instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleTransactionDetailsByTransactionIDRDataWithDefaults() *GetXRPRippleTransactionDetailsByTransactionIDRData {
	this := GetXRPRippleTransactionDetailsByTransactionIDRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRData) GetItem() GetXRPRippleTransactionDetailsByTransactionIDRI {
	if o == nil {
		var ret GetXRPRippleTransactionDetailsByTransactionIDRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleTransactionDetailsByTransactionIDRData) GetItemOk() (*GetXRPRippleTransactionDetailsByTransactionIDRI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRData) SetItem(v GetXRPRippleTransactionDetailsByTransactionIDRI) {
	o.Item = v
}

func (o GetXRPRippleTransactionDetailsByTransactionIDRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleTransactionDetailsByTransactionIDRData struct {
	value *GetXRPRippleTransactionDetailsByTransactionIDRData
	isSet bool
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRData) Get() *GetXRPRippleTransactionDetailsByTransactionIDRData {
	return v.value
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRData) Set(val *GetXRPRippleTransactionDetailsByTransactionIDRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleTransactionDetailsByTransactionIDRData(val *GetXRPRippleTransactionDetailsByTransactionIDRData) *NullableGetXRPRippleTransactionDetailsByTransactionIDRData {
	return &NullableGetXRPRippleTransactionDetailsByTransactionIDRData{value: val, isSet: true}
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


