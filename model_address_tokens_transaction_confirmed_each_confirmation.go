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

// AddressTokensTransactionConfirmedEachConfirmation struct for AddressTokensTransactionConfirmedEachConfirmation
type AddressTokensTransactionConfirmedEachConfirmation struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc.
	ReferenceId string `json:"referenceId"`
	// Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice.
	IdempotencyKey string `json:"idempotencyKey"`
	Data AddressTokensTransactionConfirmedEachConfirmationData `json:"data"`
}

// NewAddressTokensTransactionConfirmedEachConfirmation instantiates a new AddressTokensTransactionConfirmedEachConfirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTokensTransactionConfirmedEachConfirmation(apiVersion string, referenceId string, idempotencyKey string, data AddressTokensTransactionConfirmedEachConfirmationData) *AddressTokensTransactionConfirmedEachConfirmation {
	this := AddressTokensTransactionConfirmedEachConfirmation{}
	this.ApiVersion = apiVersion
	this.ReferenceId = referenceId
	this.IdempotencyKey = idempotencyKey
	this.Data = data
	return &this
}

// NewAddressTokensTransactionConfirmedEachConfirmationWithDefaults instantiates a new AddressTokensTransactionConfirmedEachConfirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTokensTransactionConfirmedEachConfirmationWithDefaults() *AddressTokensTransactionConfirmedEachConfirmation {
	this := AddressTokensTransactionConfirmedEachConfirmation{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetReferenceId returns the ReferenceId field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetData returns the Data field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetData() AddressTokensTransactionConfirmedEachConfirmationData {
	if o == nil {
		var ret AddressTokensTransactionConfirmedEachConfirmationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmation) GetDataOk() (*AddressTokensTransactionConfirmedEachConfirmationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmation) SetData(v AddressTokensTransactionConfirmedEachConfirmationData) {
	o.Data = v
}

func (o AddressTokensTransactionConfirmedEachConfirmation) MarshalJSON() ([]byte, error) {
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

type NullableAddressTokensTransactionConfirmedEachConfirmation struct {
	value *AddressTokensTransactionConfirmedEachConfirmation
	isSet bool
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmation) Get() *AddressTokensTransactionConfirmedEachConfirmation {
	return v.value
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmation) Set(val *AddressTokensTransactionConfirmedEachConfirmation) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokensTransactionConfirmedEachConfirmation(val *AddressTokensTransactionConfirmedEachConfirmation) *NullableAddressTokensTransactionConfirmedEachConfirmation {
	return &NullableAddressTokensTransactionConfirmedEachConfirmation{value: val, isSet: true}
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


