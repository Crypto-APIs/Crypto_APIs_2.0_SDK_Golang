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

// TransactionRequestApproval struct for TransactionRequestApproval
type TransactionRequestApproval struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc.
	ReferenceId string `json:"referenceId"`
	// Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice.
	IdempotencyKey string `json:"idempotencyKey"`
	Data TransactionRequestApprovalData `json:"data"`
}

// NewTransactionRequestApproval instantiates a new TransactionRequestApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequestApproval(apiVersion string, referenceId string, idempotencyKey string, data TransactionRequestApprovalData) *TransactionRequestApproval {
	this := TransactionRequestApproval{}
	this.ApiVersion = apiVersion
	this.ReferenceId = referenceId
	this.IdempotencyKey = idempotencyKey
	this.Data = data
	return &this
}

// NewTransactionRequestApprovalWithDefaults instantiates a new TransactionRequestApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestApprovalWithDefaults() *TransactionRequestApproval {
	this := TransactionRequestApproval{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *TransactionRequestApproval) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestApproval) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *TransactionRequestApproval) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetReferenceId returns the ReferenceId field value
func (o *TransactionRequestApproval) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestApproval) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *TransactionRequestApproval) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *TransactionRequestApproval) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestApproval) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *TransactionRequestApproval) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetData returns the Data field value
func (o *TransactionRequestApproval) GetData() TransactionRequestApprovalData {
	if o == nil {
		var ret TransactionRequestApprovalData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestApproval) GetDataOk() (*TransactionRequestApprovalData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionRequestApproval) SetData(v TransactionRequestApprovalData) {
	o.Data = v
}

func (o TransactionRequestApproval) MarshalJSON() ([]byte, error) {
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

type NullableTransactionRequestApproval struct {
	value *TransactionRequestApproval
	isSet bool
}

func (v NullableTransactionRequestApproval) Get() *TransactionRequestApproval {
	return v.value
}

func (v *NullableTransactionRequestApproval) Set(val *TransactionRequestApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestApproval(val *TransactionRequestApproval) *NullableTransactionRequestApproval {
	return &NullableTransactionRequestApproval{value: val, isSet: true}
}

func (v NullableTransactionRequestApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


