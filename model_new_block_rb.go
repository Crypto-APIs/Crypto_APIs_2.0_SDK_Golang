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

// NewBlockRB struct for NewBlockRB
type NewBlockRB struct {
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data NewBlockRBData `json:"data"`
}

// NewNewBlockRB instantiates a new NewBlockRB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewBlockRB(data NewBlockRBData) *NewBlockRB {
	this := NewBlockRB{}
	this.Data = data
	return &this
}

// NewNewBlockRBWithDefaults instantiates a new NewBlockRB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewBlockRBWithDefaults() *NewBlockRB {
	this := NewBlockRB{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *NewBlockRB) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewBlockRB) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *NewBlockRB) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *NewBlockRB) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *NewBlockRB) GetData() NewBlockRBData {
	if o == nil {
		var ret NewBlockRBData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NewBlockRB) GetDataOk() (*NewBlockRBData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NewBlockRB) SetData(v NewBlockRBData) {
	o.Data = v
}

func (o NewBlockRB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableNewBlockRB struct {
	value *NewBlockRB
	isSet bool
}

func (v NullableNewBlockRB) Get() *NewBlockRB {
	return v.value
}

func (v *NullableNewBlockRB) Set(val *NewBlockRB) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockRB) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockRB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockRB(val *NewBlockRB) *NullableNewBlockRB {
	return &NullableNewBlockRB{value: val, isSet: true}
}

func (v NullableNewBlockRB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockRB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


