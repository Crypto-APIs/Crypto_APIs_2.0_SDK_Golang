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

// DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken Bitcoin Omni Token
type DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken struct {
	// Defines the `propertyId` of the Omni Layer token.
	PropertyId int32 `json:"propertyId"`
}

// NewDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken instantiates a new DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken(propertyId int32) *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken {
	this := DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken{}
	this.PropertyId = propertyId
	return &this
}

// NewDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniTokenWithDefaults instantiates a new DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniTokenWithDefaults() *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken {
	this := DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken{}
	return &this
}

// GetPropertyId returns the PropertyId field value
func (o *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) GetPropertyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) GetPropertyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) SetPropertyId(v int32) {
	o.PropertyId = v
}

func (o DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["propertyId"] = o.PropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken struct {
	value *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken
	isSet bool
}

func (v NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) Get() *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken {
	return v.value
}

func (v *NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) Set(val *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken(val *DeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) *NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken {
	return &NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken{value: val, isSet: true}
}

func (v NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAutomaticTokensForwardingResponseItemTokenDataBitcoinOmniToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

