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
	"fmt"
)

// ListXRPRippleTransactionsByAddressAndTimeRangeE401 - struct for ListXRPRippleTransactionsByAddressAndTimeRangeE401
type ListXRPRippleTransactionsByAddressAndTimeRangeE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsListXRPRippleTransactionsByAddressAndTimeRangeE401 is a convenience function that returns InvalidApiKey wrapped in ListXRPRippleTransactionsByAddressAndTimeRangeE401
func InvalidApiKeyAsListXRPRippleTransactionsByAddressAndTimeRangeE401(v *InvalidApiKey) ListXRPRippleTransactionsByAddressAndTimeRangeE401 {
	return ListXRPRippleTransactionsByAddressAndTimeRangeE401{
		InvalidApiKey: v,
	}
}

// MissingApiKeyAsListXRPRippleTransactionsByAddressAndTimeRangeE401 is a convenience function that returns MissingApiKey wrapped in ListXRPRippleTransactionsByAddressAndTimeRangeE401
func MissingApiKeyAsListXRPRippleTransactionsByAddressAndTimeRangeE401(v *MissingApiKey) ListXRPRippleTransactionsByAddressAndTimeRangeE401 {
	return ListXRPRippleTransactionsByAddressAndTimeRangeE401{
		MissingApiKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListXRPRippleTransactionsByAddressAndTimeRangeE401) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(ListXRPRippleTransactionsByAddressAndTimeRangeE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListXRPRippleTransactionsByAddressAndTimeRangeE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListXRPRippleTransactionsByAddressAndTimeRangeE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListXRPRippleTransactionsByAddressAndTimeRangeE401) GetActualInstance() (interface{}) {
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

type NullableListXRPRippleTransactionsByAddressAndTimeRangeE401 struct {
	value *ListXRPRippleTransactionsByAddressAndTimeRangeE401
	isSet bool
}

func (v NullableListXRPRippleTransactionsByAddressAndTimeRangeE401) Get() *ListXRPRippleTransactionsByAddressAndTimeRangeE401 {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByAddressAndTimeRangeE401) Set(val *ListXRPRippleTransactionsByAddressAndTimeRangeE401) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByAddressAndTimeRangeE401) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByAddressAndTimeRangeE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByAddressAndTimeRangeE401(val *ListXRPRippleTransactionsByAddressAndTimeRangeE401) *NullableListXRPRippleTransactionsByAddressAndTimeRangeE401 {
	return &NullableListXRPRippleTransactionsByAddressAndTimeRangeE401{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByAddressAndTimeRangeE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByAddressAndTimeRangeE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


