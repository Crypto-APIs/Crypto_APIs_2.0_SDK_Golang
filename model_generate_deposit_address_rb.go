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

// GenerateDepositAddressRB struct for GenerateDepositAddressRB
type GenerateDepositAddressRB struct {
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data GenerateDepositAddressRBData `json:"data"`
}

// NewGenerateDepositAddressRB instantiates a new GenerateDepositAddressRB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDepositAddressRB(data GenerateDepositAddressRBData) *GenerateDepositAddressRB {
	this := GenerateDepositAddressRB{}
	this.Data = data
	return &this
}

// NewGenerateDepositAddressRBWithDefaults instantiates a new GenerateDepositAddressRB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDepositAddressRBWithDefaults() *GenerateDepositAddressRB {
	this := GenerateDepositAddressRB{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GenerateDepositAddressRB) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateDepositAddressRB) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GenerateDepositAddressRB) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *GenerateDepositAddressRB) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *GenerateDepositAddressRB) GetData() GenerateDepositAddressRBData {
	if o == nil {
		var ret GenerateDepositAddressRBData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GenerateDepositAddressRB) GetDataOk() (*GenerateDepositAddressRBData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GenerateDepositAddressRB) SetData(v GenerateDepositAddressRBData) {
	o.Data = v
}

func (o GenerateDepositAddressRB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateDepositAddressRB struct {
	value *GenerateDepositAddressRB
	isSet bool
}

func (v NullableGenerateDepositAddressRB) Get() *GenerateDepositAddressRB {
	return v.value
}

func (v *NullableGenerateDepositAddressRB) Set(val *GenerateDepositAddressRB) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDepositAddressRB) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDepositAddressRB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDepositAddressRB(val *GenerateDepositAddressRB) *NullableGenerateDepositAddressRB {
	return &NullableGenerateDepositAddressRB{value: val, isSet: true}
}

func (v NullableGenerateDepositAddressRB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateDepositAddressRB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


