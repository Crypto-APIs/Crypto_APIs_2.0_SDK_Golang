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
	"fmt"
)

// CreateAutomaticTokensForwardingE401 - struct for CreateAutomaticTokensForwardingE401
type CreateAutomaticTokensForwardingE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsCreateAutomaticTokensForwardingE401 is a convenience function that returns InvalidApiKey wrapped in CreateAutomaticTokensForwardingE401
func InvalidApiKeyAsCreateAutomaticTokensForwardingE401(v *InvalidApiKey) CreateAutomaticTokensForwardingE401 {
	return CreateAutomaticTokensForwardingE401{
		InvalidApiKey: v,
	}
}

// MissingApiKeyAsCreateAutomaticTokensForwardingE401 is a convenience function that returns MissingApiKey wrapped in CreateAutomaticTokensForwardingE401
func MissingApiKeyAsCreateAutomaticTokensForwardingE401(v *MissingApiKey) CreateAutomaticTokensForwardingE401 {
	return CreateAutomaticTokensForwardingE401{
		MissingApiKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateAutomaticTokensForwardingE401) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InvalidApiKey
	err = newStrictDecoder(data).Decode(&dst.InvalidApiKey)
	if err == nil {
		jsonInvalidApiKey, _ := json.Marshal(dst.InvalidApiKey)
		if string(jsonInvalidApiKey) == "{}" { // empty struct
			dst.InvalidApiKey = nil
		} else {
			match++
		}
	} else {
		dst.InvalidApiKey = nil
	}

	// try to unmarshal data into MissingApiKey
	err = newStrictDecoder(data).Decode(&dst.MissingApiKey)
	if err == nil {
		jsonMissingApiKey, _ := json.Marshal(dst.MissingApiKey)
		if string(jsonMissingApiKey) == "{}" { // empty struct
			dst.MissingApiKey = nil
		} else {
			match++
		}
	} else {
		dst.MissingApiKey = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidApiKey = nil
		dst.MissingApiKey = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateAutomaticTokensForwardingE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateAutomaticTokensForwardingE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateAutomaticTokensForwardingE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateAutomaticTokensForwardingE401) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvalidApiKey != nil {
		return obj.InvalidApiKey
	}

	if obj.MissingApiKey != nil {
		return obj.MissingApiKey
	}

	// all schemas are nil
	return nil
}

type NullableCreateAutomaticTokensForwardingE401 struct {
	value *CreateAutomaticTokensForwardingE401
	isSet bool
}

func (v NullableCreateAutomaticTokensForwardingE401) Get() *CreateAutomaticTokensForwardingE401 {
	return v.value
}

func (v *NullableCreateAutomaticTokensForwardingE401) Set(val *CreateAutomaticTokensForwardingE401) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticTokensForwardingE401) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticTokensForwardingE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticTokensForwardingE401(val *CreateAutomaticTokensForwardingE401) *NullableCreateAutomaticTokensForwardingE401 {
	return &NullableCreateAutomaticTokensForwardingE401{value: val, isSet: true}
}

func (v NullableCreateAutomaticTokensForwardingE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticTokensForwardingE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


