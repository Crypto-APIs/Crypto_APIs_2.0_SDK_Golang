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

// EstimateGasLimitE400 - struct for EstimateGasLimitE400
type EstimateGasLimitE400 struct {
	InvalidPagination *InvalidPagination
	LimitGreaterThanAllowed *LimitGreaterThanAllowed
	UriNotFound *UriNotFound
}

// InvalidPaginationAsEstimateGasLimitE400 is a convenience function that returns InvalidPagination wrapped in EstimateGasLimitE400
func InvalidPaginationAsEstimateGasLimitE400(v *InvalidPagination) EstimateGasLimitE400 {
	return EstimateGasLimitE400{
		InvalidPagination: v,
	}
}

// LimitGreaterThanAllowedAsEstimateGasLimitE400 is a convenience function that returns LimitGreaterThanAllowed wrapped in EstimateGasLimitE400
func LimitGreaterThanAllowedAsEstimateGasLimitE400(v *LimitGreaterThanAllowed) EstimateGasLimitE400 {
	return EstimateGasLimitE400{
		LimitGreaterThanAllowed: v,
	}
}

// UriNotFoundAsEstimateGasLimitE400 is a convenience function that returns UriNotFound wrapped in EstimateGasLimitE400
func UriNotFoundAsEstimateGasLimitE400(v *UriNotFound) EstimateGasLimitE400 {
	return EstimateGasLimitE400{
		UriNotFound: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EstimateGasLimitE400) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InvalidPagination
	err = newStrictDecoder(data).Decode(&dst.InvalidPagination)
	if err == nil {
		jsonInvalidPagination, _ := json.Marshal(dst.InvalidPagination)
		if string(jsonInvalidPagination) == "{}" { // empty struct
			dst.InvalidPagination = nil
		} else {
			match++
		}
	} else {
		dst.InvalidPagination = nil
	}

	// try to unmarshal data into LimitGreaterThanAllowed
	err = newStrictDecoder(data).Decode(&dst.LimitGreaterThanAllowed)
	if err == nil {
		jsonLimitGreaterThanAllowed, _ := json.Marshal(dst.LimitGreaterThanAllowed)
		if string(jsonLimitGreaterThanAllowed) == "{}" { // empty struct
			dst.LimitGreaterThanAllowed = nil
		} else {
			match++
		}
	} else {
		dst.LimitGreaterThanAllowed = nil
	}

	// try to unmarshal data into UriNotFound
	err = newStrictDecoder(data).Decode(&dst.UriNotFound)
	if err == nil {
		jsonUriNotFound, _ := json.Marshal(dst.UriNotFound)
		if string(jsonUriNotFound) == "{}" { // empty struct
			dst.UriNotFound = nil
		} else {
			match++
		}
	} else {
		dst.UriNotFound = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidPagination = nil
		dst.LimitGreaterThanAllowed = nil
		dst.UriNotFound = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(EstimateGasLimitE400)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(EstimateGasLimitE400)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EstimateGasLimitE400) MarshalJSON() ([]byte, error) {
	if src.InvalidPagination != nil {
		return json.Marshal(&src.InvalidPagination)
	}

	if src.LimitGreaterThanAllowed != nil {
		return json.Marshal(&src.LimitGreaterThanAllowed)
	}

	if src.UriNotFound != nil {
		return json.Marshal(&src.UriNotFound)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EstimateGasLimitE400) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvalidPagination != nil {
		return obj.InvalidPagination
	}

	if obj.LimitGreaterThanAllowed != nil {
		return obj.LimitGreaterThanAllowed
	}

	if obj.UriNotFound != nil {
		return obj.UriNotFound
	}

	// all schemas are nil
	return nil
}

type NullableEstimateGasLimitE400 struct {
	value *EstimateGasLimitE400
	isSet bool
}

func (v NullableEstimateGasLimitE400) Get() *EstimateGasLimitE400 {
	return v.value
}

func (v *NullableEstimateGasLimitE400) Set(val *EstimateGasLimitE400) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateGasLimitE400) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateGasLimitE400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateGasLimitE400(val *EstimateGasLimitE400) *NullableEstimateGasLimitE400 {
	return &NullableEstimateGasLimitE400{value: val, isSet: true}
}

func (v NullableEstimateGasLimitE400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateGasLimitE400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


