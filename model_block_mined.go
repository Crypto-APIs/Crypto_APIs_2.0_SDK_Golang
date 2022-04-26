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

// BlockMined struct for BlockMined
type BlockMined struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc.
	ReferenceId string `json:"referenceId"`
	// Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice.
	IdempotencyKey string `json:"idempotencyKey"`
	Data BlockMinedData `json:"data"`
}

// NewBlockMined instantiates a new BlockMined object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockMined(apiVersion string, referenceId string, idempotencyKey string, data BlockMinedData) *BlockMined {
	this := BlockMined{}
	this.ApiVersion = apiVersion
	this.ReferenceId = referenceId
	this.IdempotencyKey = idempotencyKey
	this.Data = data
	return &this
}

// NewBlockMinedWithDefaults instantiates a new BlockMined object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockMinedWithDefaults() *BlockMined {
	this := BlockMined{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *BlockMined) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *BlockMined) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *BlockMined) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetReferenceId returns the ReferenceId field value
func (o *BlockMined) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *BlockMined) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *BlockMined) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *BlockMined) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *BlockMined) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *BlockMined) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetData returns the Data field value
func (o *BlockMined) GetData() BlockMinedData {
	if o == nil {
		var ret BlockMinedData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BlockMined) GetDataOk() (*BlockMinedData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BlockMined) SetData(v BlockMinedData) {
	o.Data = v
}

func (o BlockMined) MarshalJSON() ([]byte, error) {
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

type NullableBlockMined struct {
	value *BlockMined
	isSet bool
}

func (v NullableBlockMined) Get() *BlockMined {
	return v.value
}

func (v *NullableBlockMined) Set(val *BlockMined) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockMined) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockMined) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockMined(val *BlockMined) *NullableBlockMined {
	return &NullableBlockMined{value: val, isSet: true}
}

func (v NullableBlockMined) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockMined) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


