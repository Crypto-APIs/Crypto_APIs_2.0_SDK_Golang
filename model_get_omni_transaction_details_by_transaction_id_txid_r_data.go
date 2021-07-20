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

// GetOmniTransactionDetailsByTransactionIDTxidRData struct for GetOmniTransactionDetailsByTransactionIDTxidRData
type GetOmniTransactionDetailsByTransactionIDTxidRData struct {
	Item GetOmniTransactionDetailsByTransactionIDTxidRI `json:"item"`
}

// NewGetOmniTransactionDetailsByTransactionIDTxidRData instantiates a new GetOmniTransactionDetailsByTransactionIDTxidRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOmniTransactionDetailsByTransactionIDTxidRData(item GetOmniTransactionDetailsByTransactionIDTxidRI) *GetOmniTransactionDetailsByTransactionIDTxidRData {
	this := GetOmniTransactionDetailsByTransactionIDTxidRData{}
	this.Item = item
	return &this
}

// NewGetOmniTransactionDetailsByTransactionIDTxidRDataWithDefaults instantiates a new GetOmniTransactionDetailsByTransactionIDTxidRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOmniTransactionDetailsByTransactionIDTxidRDataWithDefaults() *GetOmniTransactionDetailsByTransactionIDTxidRData {
	this := GetOmniTransactionDetailsByTransactionIDTxidRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetOmniTransactionDetailsByTransactionIDTxidRData) GetItem() GetOmniTransactionDetailsByTransactionIDTxidRI {
	if o == nil {
		var ret GetOmniTransactionDetailsByTransactionIDTxidRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetOmniTransactionDetailsByTransactionIDTxidRData) GetItemOk() (*GetOmniTransactionDetailsByTransactionIDTxidRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetOmniTransactionDetailsByTransactionIDTxidRData) SetItem(v GetOmniTransactionDetailsByTransactionIDTxidRI) {
	o.Item = v
}

func (o GetOmniTransactionDetailsByTransactionIDTxidRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetOmniTransactionDetailsByTransactionIDTxidRData struct {
	value *GetOmniTransactionDetailsByTransactionIDTxidRData
	isSet bool
}

func (v NullableGetOmniTransactionDetailsByTransactionIDTxidRData) Get() *GetOmniTransactionDetailsByTransactionIDTxidRData {
	return v.value
}

func (v *NullableGetOmniTransactionDetailsByTransactionIDTxidRData) Set(val *GetOmniTransactionDetailsByTransactionIDTxidRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOmniTransactionDetailsByTransactionIDTxidRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOmniTransactionDetailsByTransactionIDTxidRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOmniTransactionDetailsByTransactionIDTxidRData(val *GetOmniTransactionDetailsByTransactionIDTxidRData) *NullableGetOmniTransactionDetailsByTransactionIDTxidRData {
	return &NullableGetOmniTransactionDetailsByTransactionIDTxidRData{value: val, isSet: true}
}

func (v NullableGetOmniTransactionDetailsByTransactionIDTxidRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOmniTransactionDetailsByTransactionIDTxidRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

