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

// BlockHeightReachedE401 - struct for BlockHeightReachedE401
type BlockHeightReachedE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsBlockHeightReachedE401 is a convenience function that returns InvalidApiKey wrapped in BlockHeightReachedE401
func InvalidApiKeyAsBlockHeightReachedE401(v *InvalidApiKey) BlockHeightReachedE401 {
	return BlockHeightReachedE401{
		InvalidApiKey: v,
	}
}

// MissingApiKeyAsBlockHeightReachedE401 is a convenience function that returns MissingApiKey wrapped in BlockHeightReachedE401
func MissingApiKeyAsBlockHeightReachedE401(v *MissingApiKey) BlockHeightReachedE401 {
	return BlockHeightReachedE401{
		MissingApiKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BlockHeightReachedE401) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(BlockHeightReachedE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(BlockHeightReachedE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BlockHeightReachedE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BlockHeightReachedE401) GetActualInstance() (interface{}) {
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

type NullableBlockHeightReachedE401 struct {
	value *BlockHeightReachedE401
	isSet bool
}

func (v NullableBlockHeightReachedE401) Get() *BlockHeightReachedE401 {
	return v.value
}

func (v *NullableBlockHeightReachedE401) Set(val *BlockHeightReachedE401) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeightReachedE401) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeightReachedE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeightReachedE401(val *BlockHeightReachedE401) *NullableBlockHeightReachedE401 {
	return &NullableBlockHeightReachedE401{value: val, isSet: true}
}

func (v NullableBlockHeightReachedE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeightReachedE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

