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

// AddressTokensTransactionConfirmed struct for AddressTokensTransactionConfirmed
type AddressTokensTransactionConfirmed struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc.
	ReferenceId string `json:"referenceId"`
	// Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice.
	IdempotencyKey string `json:"idempotencyKey"`
	Data AddressTokensTransactionConfirmedData `json:"data"`
}

// NewAddressTokensTransactionConfirmed instantiates a new AddressTokensTransactionConfirmed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTokensTransactionConfirmed(apiVersion string, referenceId string, idempotencyKey string, data AddressTokensTransactionConfirmedData) *AddressTokensTransactionConfirmed {
	this := AddressTokensTransactionConfirmed{}
	this.ApiVersion = apiVersion
	this.ReferenceId = referenceId
	this.IdempotencyKey = idempotencyKey
	this.Data = data
	return &this
}

// NewAddressTokensTransactionConfirmedWithDefaults instantiates a new AddressTokensTransactionConfirmed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTokensTransactionConfirmedWithDefaults() *AddressTokensTransactionConfirmed {
	this := AddressTokensTransactionConfirmed{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *AddressTokensTransactionConfirmed) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmed) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *AddressTokensTransactionConfirmed) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetReferenceId returns the ReferenceId field value
func (o *AddressTokensTransactionConfirmed) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmed) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *AddressTokensTransactionConfirmed) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *AddressTokensTransactionConfirmed) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmed) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *AddressTokensTransactionConfirmed) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetData returns the Data field value
func (o *AddressTokensTransactionConfirmed) GetData() AddressTokensTransactionConfirmedData {
	if o == nil {
		var ret AddressTokensTransactionConfirmedData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmed) GetDataOk() (*AddressTokensTransactionConfirmedData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressTokensTransactionConfirmed) SetData(v AddressTokensTransactionConfirmedData) {
	o.Data = v
}

func (o AddressTokensTransactionConfirmed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if true {
		toSerialize["idempotencyKey"] = o.IdempotencyKey
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAddressTokensTransactionConfirmed struct {
	value *AddressTokensTransactionConfirmed
	isSet bool
}

func (v NullableAddressTokensTransactionConfirmed) Get() *AddressTokensTransactionConfirmed {
	return v.value
}

func (v *NullableAddressTokensTransactionConfirmed) Set(val *AddressTokensTransactionConfirmed) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokensTransactionConfirmed) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokensTransactionConfirmed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokensTransactionConfirmed(val *AddressTokensTransactionConfirmed) *NullableAddressTokensTransactionConfirmed {
	return &NullableAddressTokensTransactionConfirmed{value: val, isSet: true}
}

func (v NullableAddressTokensTransactionConfirmed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokensTransactionConfirmed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


