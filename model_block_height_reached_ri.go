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

// BlockHeightReachedRI struct for BlockHeightReachedRI
type BlockHeightReachedRI struct {
	// Represents the specified value of block height for which the callback will be received.
	BlockHeightReached int64 `json:"blockHeightReached"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey string `json:"callbackSecretKey"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. We support ONLY httpS type of protocol.
	CallbackUrl string `json:"callbackUrl"`
	// Defines the specific time/date when the subscription was created in Unix Timestamp.
	CreatedTimestamp int32 `json:"createdTimestamp"`
	// Defines whether the subscription is active or not. Set as boolean.
	IsActive bool `json:"isActive"`
	// Represents a unique ID used to reference the specific callback subscription.
	ReferenceId string `json:"referenceId"`
}

// NewBlockHeightReachedRI instantiates a new BlockHeightReachedRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockHeightReachedRI(blockHeightReached int64, callbackSecretKey string, callbackUrl string, createdTimestamp int32, isActive bool, referenceId string) *BlockHeightReachedRI {
	this := BlockHeightReachedRI{}
	this.BlockHeightReached = blockHeightReached
	this.CallbackSecretKey = callbackSecretKey
	this.CallbackUrl = callbackUrl
	this.CreatedTimestamp = createdTimestamp
	this.IsActive = isActive
	this.ReferenceId = referenceId
	return &this
}

// NewBlockHeightReachedRIWithDefaults instantiates a new BlockHeightReachedRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockHeightReachedRIWithDefaults() *BlockHeightReachedRI {
	this := BlockHeightReachedRI{}
	return &this
}

// GetBlockHeightReached returns the BlockHeightReached field value
func (o *BlockHeightReachedRI) GetBlockHeightReached() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BlockHeightReached
}

// GetBlockHeightReachedOk returns a tuple with the BlockHeightReached field value
// and a boolean to check if the value has been set.
func (o *BlockHeightReachedRI) GetBlockHeightReachedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHeightReached, true
}

// SetBlockHeightReached sets field value
func (o *BlockHeightReachedRI) SetBlockHeightReached(v int64) {
	o.BlockHeightReached = v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value
func (o *BlockHeightReachedRI) GetCallbackSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value
// and a boolean to check if the value has been set.
func (o *BlockHeightReachedRI) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackSecretKey, true
}

// SetCallbackSecretKey sets field value
func (o *BlockHeightReachedRI) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *BlockHeightReachedRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *BlockHeightReachedRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *BlockHeightReachedRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *BlockHeightReachedRI) GetCreatedTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *BlockHeightReachedRI) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *BlockHeightReachedRI) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = v
}

// GetIsActive returns the IsActive field value
func (o *BlockHeightReachedRI) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *BlockHeightReachedRI) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *BlockHeightReachedRI) SetIsActive(v bool) {
	o.IsActive = v
}

// GetReferenceId returns the ReferenceId field value
func (o *BlockHeightReachedRI) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *BlockHeightReachedRI) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *BlockHeightReachedRI) SetReferenceId(v string) {
	o.ReferenceId = v
}

func (o BlockHeightReachedRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockHeightReached"] = o.BlockHeightReached
	}
	if true {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if true {
		toSerialize["createdTimestamp"] = o.CreatedTimestamp
	}
	if true {
		toSerialize["isActive"] = o.IsActive
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	return json.Marshal(toSerialize)
}

type NullableBlockHeightReachedRI struct {
	value *BlockHeightReachedRI
	isSet bool
}

func (v NullableBlockHeightReachedRI) Get() *BlockHeightReachedRI {
	return v.value
}

func (v *NullableBlockHeightReachedRI) Set(val *BlockHeightReachedRI) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeightReachedRI) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeightReachedRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeightReachedRI(val *BlockHeightReachedRI) *NullableBlockHeightReachedRI {
	return &NullableBlockHeightReachedRI{value: val, isSet: true}
}

func (v NullableBlockHeightReachedRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeightReachedRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


