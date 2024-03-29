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

// GetExchangeRateByAssetSymbolsE422 - struct for GetExchangeRateByAssetSymbolsE422
type GetExchangeRateByAssetSymbolsE422 struct {
	CouldNotCalculateRateForPair *CouldNotCalculateRateForPair
	InvalidRequestBodyStructure *InvalidRequestBodyStructure
}

// CouldNotCalculateRateForPairAsGetExchangeRateByAssetSymbolsE422 is a convenience function that returns CouldNotCalculateRateForPair wrapped in GetExchangeRateByAssetSymbolsE422
func CouldNotCalculateRateForPairAsGetExchangeRateByAssetSymbolsE422(v *CouldNotCalculateRateForPair) GetExchangeRateByAssetSymbolsE422 {
	return GetExchangeRateByAssetSymbolsE422{
		CouldNotCalculateRateForPair: v,
	}
}

// InvalidRequestBodyStructureAsGetExchangeRateByAssetSymbolsE422 is a convenience function that returns InvalidRequestBodyStructure wrapped in GetExchangeRateByAssetSymbolsE422
func InvalidRequestBodyStructureAsGetExchangeRateByAssetSymbolsE422(v *InvalidRequestBodyStructure) GetExchangeRateByAssetSymbolsE422 {
	return GetExchangeRateByAssetSymbolsE422{
		InvalidRequestBodyStructure: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetExchangeRateByAssetSymbolsE422) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CouldNotCalculateRateForPair
	err = newStrictDecoder(data).Decode(&dst.CouldNotCalculateRateForPair)
	if err == nil {
		jsonCouldNotCalculateRateForPair, _ := json.Marshal(dst.CouldNotCalculateRateForPair)
		if string(jsonCouldNotCalculateRateForPair) == "{}" { // empty struct
			dst.CouldNotCalculateRateForPair = nil
		} else {
			match++
		}
	} else {
		dst.CouldNotCalculateRateForPair = nil
	}

	// try to unmarshal data into InvalidRequestBodyStructure
	err = newStrictDecoder(data).Decode(&dst.InvalidRequestBodyStructure)
	if err == nil {
		jsonInvalidRequestBodyStructure, _ := json.Marshal(dst.InvalidRequestBodyStructure)
		if string(jsonInvalidRequestBodyStructure) == "{}" { // empty struct
			dst.InvalidRequestBodyStructure = nil
		} else {
			match++
		}
	} else {
		dst.InvalidRequestBodyStructure = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CouldNotCalculateRateForPair = nil
		dst.InvalidRequestBodyStructure = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetExchangeRateByAssetSymbolsE422)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetExchangeRateByAssetSymbolsE422)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetExchangeRateByAssetSymbolsE422) MarshalJSON() ([]byte, error) {
	if src.CouldNotCalculateRateForPair != nil {
		return json.Marshal(&src.CouldNotCalculateRateForPair)
	}

	if src.InvalidRequestBodyStructure != nil {
		return json.Marshal(&src.InvalidRequestBodyStructure)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetExchangeRateByAssetSymbolsE422) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CouldNotCalculateRateForPair != nil {
		return obj.CouldNotCalculateRateForPair
	}

	if obj.InvalidRequestBodyStructure != nil {
		return obj.InvalidRequestBodyStructure
	}

	// all schemas are nil
	return nil
}

type NullableGetExchangeRateByAssetSymbolsE422 struct {
	value *GetExchangeRateByAssetSymbolsE422
	isSet bool
}

func (v NullableGetExchangeRateByAssetSymbolsE422) Get() *GetExchangeRateByAssetSymbolsE422 {
	return v.value
}

func (v *NullableGetExchangeRateByAssetSymbolsE422) Set(val *GetExchangeRateByAssetSymbolsE422) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeRateByAssetSymbolsE422) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeRateByAssetSymbolsE422) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeRateByAssetSymbolsE422(val *GetExchangeRateByAssetSymbolsE422) *NullableGetExchangeRateByAssetSymbolsE422 {
	return &NullableGetExchangeRateByAssetSymbolsE422{value: val, isSet: true}
}

func (v NullableGetExchangeRateByAssetSymbolsE422) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeRateByAssetSymbolsE422) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


