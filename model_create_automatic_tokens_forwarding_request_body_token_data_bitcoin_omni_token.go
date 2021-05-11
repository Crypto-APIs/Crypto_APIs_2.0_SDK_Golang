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

// CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken Bitcoin Omni Token
type CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken struct {
	// Represents the specific `propertyId` of the token data that will be forwarded.
	PropertyId int32 `json:"propertyId"`
}

// NewCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken instantiates a new CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken(propertyId int32) *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken {
	this := CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken{}
	this.PropertyId = propertyId
	return &this
}

// NewCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniTokenWithDefaults instantiates a new CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniTokenWithDefaults() *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken {
	this := CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken{}
	return &this
}

// GetPropertyId returns the PropertyId field value
func (o *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) GetPropertyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) GetPropertyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) SetPropertyId(v int32) {
	o.PropertyId = v
}

func (o CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["propertyId"] = o.PropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken struct {
	value *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken
	isSet bool
}

func (v NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) Get() *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken {
	return v.value
}

func (v *NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) Set(val *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken(val *CreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) *NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken {
	return &NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken{value: val, isSet: true}
}

func (v NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticTokensForwardingRequestBodyTokenDataBitcoinOmniToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


