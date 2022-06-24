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

// NewConfirmedTokenTransactionsForSpecificAmountRB struct for NewConfirmedTokenTransactionsForSpecificAmountRB
type NewConfirmedTokenTransactionsForSpecificAmountRB struct {
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data NewConfirmedTokenTransactionsForSpecificAmountRBData `json:"data"`
}

// NewNewConfirmedTokenTransactionsForSpecificAmountRB instantiates a new NewConfirmedTokenTransactionsForSpecificAmountRB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedTokenTransactionsForSpecificAmountRB(data NewConfirmedTokenTransactionsForSpecificAmountRBData) *NewConfirmedTokenTransactionsForSpecificAmountRB {
	this := NewConfirmedTokenTransactionsForSpecificAmountRB{}
	this.Data = data
	return &this
}

// NewNewConfirmedTokenTransactionsForSpecificAmountRBWithDefaults instantiates a new NewConfirmedTokenTransactionsForSpecificAmountRB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedTokenTransactionsForSpecificAmountRBWithDefaults() *NewConfirmedTokenTransactionsForSpecificAmountRB {
	this := NewConfirmedTokenTransactionsForSpecificAmountRB{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRB) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRB) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRB) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRB) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRB) GetData() NewConfirmedTokenTransactionsForSpecificAmountRBData {
	if o == nil {
		var ret NewConfirmedTokenTransactionsForSpecificAmountRBData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRB) GetDataOk() (*NewConfirmedTokenTransactionsForSpecificAmountRBData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRB) SetData(v NewConfirmedTokenTransactionsForSpecificAmountRBData) {
	o.Data = v
}

func (o NewConfirmedTokenTransactionsForSpecificAmountRB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedTokenTransactionsForSpecificAmountRB struct {
	value *NewConfirmedTokenTransactionsForSpecificAmountRB
	isSet bool
}

func (v NullableNewConfirmedTokenTransactionsForSpecificAmountRB) Get() *NewConfirmedTokenTransactionsForSpecificAmountRB {
	return v.value
}

func (v *NullableNewConfirmedTokenTransactionsForSpecificAmountRB) Set(val *NewConfirmedTokenTransactionsForSpecificAmountRB) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedTokenTransactionsForSpecificAmountRB) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedTokenTransactionsForSpecificAmountRB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedTokenTransactionsForSpecificAmountRB(val *NewConfirmedTokenTransactionsForSpecificAmountRB) *NullableNewConfirmedTokenTransactionsForSpecificAmountRB {
	return &NullableNewConfirmedTokenTransactionsForSpecificAmountRB{value: val, isSet: true}
}

func (v NullableNewConfirmedTokenTransactionsForSpecificAmountRB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedTokenTransactionsForSpecificAmountRB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


