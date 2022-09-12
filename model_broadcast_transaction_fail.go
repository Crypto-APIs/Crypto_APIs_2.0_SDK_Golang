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

// BroadcastTransactionFail struct for BroadcastTransactionFail
type BroadcastTransactionFail struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc.
	ReferenceId string `json:"referenceId"`
	// Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice.
	IdempotencyKey string `json:"idempotencyKey"`
	Data BroadcastTransactionFailData `json:"data"`
}

// NewBroadcastTransactionFail instantiates a new BroadcastTransactionFail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTransactionFail(apiVersion string, referenceId string, idempotencyKey string, data BroadcastTransactionFailData) *BroadcastTransactionFail {
	this := BroadcastTransactionFail{}
	this.ApiVersion = apiVersion
	this.ReferenceId = referenceId
	this.IdempotencyKey = idempotencyKey
	this.Data = data
	return &this
}

// NewBroadcastTransactionFailWithDefaults instantiates a new BroadcastTransactionFail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTransactionFailWithDefaults() *BroadcastTransactionFail {
	this := BroadcastTransactionFail{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *BroadcastTransactionFail) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *BroadcastTransactionFail) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *BroadcastTransactionFail) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetReferenceId returns the ReferenceId field value
func (o *BroadcastTransactionFail) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *BroadcastTransactionFail) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *BroadcastTransactionFail) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *BroadcastTransactionFail) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *BroadcastTransactionFail) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *BroadcastTransactionFail) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetData returns the Data field value
func (o *BroadcastTransactionFail) GetData() BroadcastTransactionFailData {
	if o == nil {
		var ret BroadcastTransactionFailData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BroadcastTransactionFail) GetDataOk() (*BroadcastTransactionFailData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BroadcastTransactionFail) SetData(v BroadcastTransactionFailData) {
	o.Data = v
}

func (o BroadcastTransactionFail) MarshalJSON() ([]byte, error) {
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

type NullableBroadcastTransactionFail struct {
	value *BroadcastTransactionFail
	isSet bool
}

func (v NullableBroadcastTransactionFail) Get() *BroadcastTransactionFail {
	return v.value
}

func (v *NullableBroadcastTransactionFail) Set(val *BroadcastTransactionFail) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTransactionFail) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTransactionFail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTransactionFail(val *BroadcastTransactionFail) *NullableBroadcastTransactionFail {
	return &NullableBroadcastTransactionFail{value: val, isSet: true}
}

func (v NullableBroadcastTransactionFail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTransactionFail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


