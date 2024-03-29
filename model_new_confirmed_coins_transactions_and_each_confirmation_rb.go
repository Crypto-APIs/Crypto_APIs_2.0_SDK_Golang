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

// NewConfirmedCoinsTransactionsAndEachConfirmationRB struct for NewConfirmedCoinsTransactionsAndEachConfirmationRB
type NewConfirmedCoinsTransactionsAndEachConfirmationRB struct {
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data NewConfirmedCoinsTransactionsAndEachConfirmationRBData `json:"data"`
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRB instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRB(data NewConfirmedCoinsTransactionsAndEachConfirmationRBData) *NewConfirmedCoinsTransactionsAndEachConfirmationRB {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRB{}
	this.Data = data
	return &this
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRBWithDefaults instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRBWithDefaults() *NewConfirmedCoinsTransactionsAndEachConfirmationRB {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRB{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRB) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRB) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRB) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRB) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRB) GetData() NewConfirmedCoinsTransactionsAndEachConfirmationRBData {
	if o == nil {
		var ret NewConfirmedCoinsTransactionsAndEachConfirmationRBData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRB) GetDataOk() (*NewConfirmedCoinsTransactionsAndEachConfirmationRBData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRB) SetData(v NewConfirmedCoinsTransactionsAndEachConfirmationRBData) {
	o.Data = v
}

func (o NewConfirmedCoinsTransactionsAndEachConfirmationRB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB struct {
	value *NewConfirmedCoinsTransactionsAndEachConfirmationRB
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB) Get() *NewConfirmedCoinsTransactionsAndEachConfirmationRB {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB) Set(val *NewConfirmedCoinsTransactionsAndEachConfirmationRB) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsAndEachConfirmationRB(val *NewConfirmedCoinsTransactionsAndEachConfirmationRB) *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB {
	return &NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


