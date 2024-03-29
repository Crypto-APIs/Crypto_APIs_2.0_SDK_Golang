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

// NewBlockE409 - struct for NewBlockE409
type NewBlockE409 struct {
	AlreadyExists *AlreadyExists
	InvalidData *InvalidData
}

// AlreadyExistsAsNewBlockE409 is a convenience function that returns AlreadyExists wrapped in NewBlockE409
func AlreadyExistsAsNewBlockE409(v *AlreadyExists) NewBlockE409 {
	return NewBlockE409{
		AlreadyExists: v,
	}
}

// InvalidDataAsNewBlockE409 is a convenience function that returns InvalidData wrapped in NewBlockE409
func InvalidDataAsNewBlockE409(v *InvalidData) NewBlockE409 {
	return NewBlockE409{
		InvalidData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewBlockE409) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AlreadyExists
	err = newStrictDecoder(data).Decode(&dst.AlreadyExists)
	if err == nil {
		jsonAlreadyExists, _ := json.Marshal(dst.AlreadyExists)
		if string(jsonAlreadyExists) == "{}" { // empty struct
			dst.AlreadyExists = nil
		} else {
			match++
		}
	} else {
		dst.AlreadyExists = nil
	}

	// try to unmarshal data into InvalidData
	err = newStrictDecoder(data).Decode(&dst.InvalidData)
	if err == nil {
		jsonInvalidData, _ := json.Marshal(dst.InvalidData)
		if string(jsonInvalidData) == "{}" { // empty struct
			dst.InvalidData = nil
		} else {
			match++
		}
	} else {
		dst.InvalidData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AlreadyExists = nil
		dst.InvalidData = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(NewBlockE409)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(NewBlockE409)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewBlockE409) MarshalJSON() ([]byte, error) {
	if src.AlreadyExists != nil {
		return json.Marshal(&src.AlreadyExists)
	}

	if src.InvalidData != nil {
		return json.Marshal(&src.InvalidData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NewBlockE409) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AlreadyExists != nil {
		return obj.AlreadyExists
	}

	if obj.InvalidData != nil {
		return obj.InvalidData
	}

	// all schemas are nil
	return nil
}

type NullableNewBlockE409 struct {
	value *NewBlockE409
	isSet bool
}

func (v NullableNewBlockE409) Get() *NewBlockE409 {
	return v.value
}

func (v *NullableNewBlockE409) Set(val *NewBlockE409) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockE409) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockE409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockE409(val *NewBlockE409) *NullableNewBlockE409 {
	return &NullableNewBlockE409{value: val, isSet: true}
}

func (v NullableNewBlockE409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockE409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


