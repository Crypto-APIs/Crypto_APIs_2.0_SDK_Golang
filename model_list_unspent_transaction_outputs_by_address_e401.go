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

// ListUnspentTransactionOutputsByAddressE401 - struct for ListUnspentTransactionOutputsByAddressE401
type ListUnspentTransactionOutputsByAddressE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsListUnspentTransactionOutputsByAddressE401 is a convenience function that returns InvalidApiKey wrapped in ListUnspentTransactionOutputsByAddressE401
func InvalidApiKeyAsListUnspentTransactionOutputsByAddressE401(v *InvalidApiKey) ListUnspentTransactionOutputsByAddressE401 {
	return ListUnspentTransactionOutputsByAddressE401{
		InvalidApiKey: v,
	}
}

// MissingApiKeyAsListUnspentTransactionOutputsByAddressE401 is a convenience function that returns MissingApiKey wrapped in ListUnspentTransactionOutputsByAddressE401
func MissingApiKeyAsListUnspentTransactionOutputsByAddressE401(v *MissingApiKey) ListUnspentTransactionOutputsByAddressE401 {
	return ListUnspentTransactionOutputsByAddressE401{
		MissingApiKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListUnspentTransactionOutputsByAddressE401) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(ListUnspentTransactionOutputsByAddressE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListUnspentTransactionOutputsByAddressE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListUnspentTransactionOutputsByAddressE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListUnspentTransactionOutputsByAddressE401) GetActualInstance() (interface{}) {
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

type NullableListUnspentTransactionOutputsByAddressE401 struct {
	value *ListUnspentTransactionOutputsByAddressE401
	isSet bool
}

func (v NullableListUnspentTransactionOutputsByAddressE401) Get() *ListUnspentTransactionOutputsByAddressE401 {
	return v.value
}

func (v *NullableListUnspentTransactionOutputsByAddressE401) Set(val *ListUnspentTransactionOutputsByAddressE401) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnspentTransactionOutputsByAddressE401) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnspentTransactionOutputsByAddressE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnspentTransactionOutputsByAddressE401(val *ListUnspentTransactionOutputsByAddressE401) *NullableListUnspentTransactionOutputsByAddressE401 {
	return &NullableListUnspentTransactionOutputsByAddressE401{value: val, isSet: true}
}

func (v NullableListUnspentTransactionOutputsByAddressE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnspentTransactionOutputsByAddressE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


