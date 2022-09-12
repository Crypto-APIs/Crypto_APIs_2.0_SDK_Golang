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

// ListTokensForwardingAutomationsE400 - struct for ListTokensForwardingAutomationsE400
type ListTokensForwardingAutomationsE400 struct {
	InvalidPagination *InvalidPagination
	LimitGreaterThanAllowed *LimitGreaterThanAllowed
	UriNotFound *UriNotFound
}

// InvalidPaginationAsListTokensForwardingAutomationsE400 is a convenience function that returns InvalidPagination wrapped in ListTokensForwardingAutomationsE400
func InvalidPaginationAsListTokensForwardingAutomationsE400(v *InvalidPagination) ListTokensForwardingAutomationsE400 {
	return ListTokensForwardingAutomationsE400{
		InvalidPagination: v,
	}
}

// LimitGreaterThanAllowedAsListTokensForwardingAutomationsE400 is a convenience function that returns LimitGreaterThanAllowed wrapped in ListTokensForwardingAutomationsE400
func LimitGreaterThanAllowedAsListTokensForwardingAutomationsE400(v *LimitGreaterThanAllowed) ListTokensForwardingAutomationsE400 {
	return ListTokensForwardingAutomationsE400{
		LimitGreaterThanAllowed: v,
	}
}

// UriNotFoundAsListTokensForwardingAutomationsE400 is a convenience function that returns UriNotFound wrapped in ListTokensForwardingAutomationsE400
func UriNotFoundAsListTokensForwardingAutomationsE400(v *UriNotFound) ListTokensForwardingAutomationsE400 {
	return ListTokensForwardingAutomationsE400{
		UriNotFound: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListTokensForwardingAutomationsE400) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(ListTokensForwardingAutomationsE400)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListTokensForwardingAutomationsE400)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListTokensForwardingAutomationsE400) MarshalJSON() ([]byte, error) {
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
func (obj *ListTokensForwardingAutomationsE400) GetActualInstance() (interface{}) {
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

type NullableListTokensForwardingAutomationsE400 struct {
	value *ListTokensForwardingAutomationsE400
	isSet bool
}

func (v NullableListTokensForwardingAutomationsE400) Get() *ListTokensForwardingAutomationsE400 {
	return v.value
}

func (v *NullableListTokensForwardingAutomationsE400) Set(val *ListTokensForwardingAutomationsE400) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokensForwardingAutomationsE400) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokensForwardingAutomationsE400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokensForwardingAutomationsE400(val *ListTokensForwardingAutomationsE400) *NullableListTokensForwardingAutomationsE400 {
	return &NullableListTokensForwardingAutomationsE400{value: val, isSet: true}
}

func (v NullableListTokensForwardingAutomationsE400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokensForwardingAutomationsE400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


