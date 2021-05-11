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

// CoinsForwardingFail struct for CoinsForwardingFail
type CoinsForwardingFail struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc.
	ReferenceId string `json:"referenceId"`
	// Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice.
	IdempotencyKey string `json:"idempotencyKey"`
	Data CoinsForwardingFailData `json:"data"`
}

// NewCoinsForwardingFail instantiates a new CoinsForwardingFail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoinsForwardingFail(apiVersion string, referenceId string, idempotencyKey string, data CoinsForwardingFailData) *CoinsForwardingFail {
	this := CoinsForwardingFail{}
	this.ApiVersion = apiVersion
	this.ReferenceId = referenceId
	this.IdempotencyKey = idempotencyKey
	this.Data = data
	return &this
}

// NewCoinsForwardingFailWithDefaults instantiates a new CoinsForwardingFail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoinsForwardingFailWithDefaults() *CoinsForwardingFail {
	this := CoinsForwardingFail{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *CoinsForwardingFail) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *CoinsForwardingFail) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *CoinsForwardingFail) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetReferenceId returns the ReferenceId field value
func (o *CoinsForwardingFail) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CoinsForwardingFail) GetReferenceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CoinsForwardingFail) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *CoinsForwardingFail) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *CoinsForwardingFail) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *CoinsForwardingFail) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetData returns the Data field value
func (o *CoinsForwardingFail) GetData() CoinsForwardingFailData {
	if o == nil {
		var ret CoinsForwardingFailData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CoinsForwardingFail) GetDataOk() (*CoinsForwardingFailData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CoinsForwardingFail) SetData(v CoinsForwardingFailData) {
	o.Data = v
}

func (o CoinsForwardingFail) MarshalJSON() ([]byte, error) {
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

type NullableCoinsForwardingFail struct {
	value *CoinsForwardingFail
	isSet bool
}

func (v NullableCoinsForwardingFail) Get() *CoinsForwardingFail {
	return v.value
}

func (v *NullableCoinsForwardingFail) Set(val *CoinsForwardingFail) {
	v.value = val
	v.isSet = true
}

func (v NullableCoinsForwardingFail) IsSet() bool {
	return v.isSet
}

func (v *NullableCoinsForwardingFail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoinsForwardingFail(val *CoinsForwardingFail) *NullableCoinsForwardingFail {
	return &NullableCoinsForwardingFail{value: val, isSet: true}
}

func (v NullableCoinsForwardingFail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoinsForwardingFail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


