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

// ConvertBitcoinCashAddressE400 - struct for ConvertBitcoinCashAddressE400
type ConvertBitcoinCashAddressE400 struct {
	InvalidBlockchain *InvalidBlockchain
	InvalidPagination *InvalidPagination
	LimitGreaterThanAllowed *LimitGreaterThanAllowed
	UriNotFound *UriNotFound
}

// InvalidBlockchainAsConvertBitcoinCashAddressE400 is a convenience function that returns InvalidBlockchain wrapped in ConvertBitcoinCashAddressE400
func InvalidBlockchainAsConvertBitcoinCashAddressE400(v *InvalidBlockchain) ConvertBitcoinCashAddressE400 {
	return ConvertBitcoinCashAddressE400{
		InvalidBlockchain: v,
	}
}

// InvalidPaginationAsConvertBitcoinCashAddressE400 is a convenience function that returns InvalidPagination wrapped in ConvertBitcoinCashAddressE400
func InvalidPaginationAsConvertBitcoinCashAddressE400(v *InvalidPagination) ConvertBitcoinCashAddressE400 {
	return ConvertBitcoinCashAddressE400{
		InvalidPagination: v,
	}
}

// LimitGreaterThanAllowedAsConvertBitcoinCashAddressE400 is a convenience function that returns LimitGreaterThanAllowed wrapped in ConvertBitcoinCashAddressE400
func LimitGreaterThanAllowedAsConvertBitcoinCashAddressE400(v *LimitGreaterThanAllowed) ConvertBitcoinCashAddressE400 {
	return ConvertBitcoinCashAddressE400{
		LimitGreaterThanAllowed: v,
	}
}

// UriNotFoundAsConvertBitcoinCashAddressE400 is a convenience function that returns UriNotFound wrapped in ConvertBitcoinCashAddressE400
func UriNotFoundAsConvertBitcoinCashAddressE400(v *UriNotFound) ConvertBitcoinCashAddressE400 {
	return ConvertBitcoinCashAddressE400{
		UriNotFound: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConvertBitcoinCashAddressE400) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InvalidBlockchain
	err = newStrictDecoder(data).Decode(&dst.InvalidBlockchain)
	if err == nil {
		jsonInvalidBlockchain, _ := json.Marshal(dst.InvalidBlockchain)
		if string(jsonInvalidBlockchain) == "{}" { // empty struct
			dst.InvalidBlockchain = nil
		} else {
			match++
		}
	} else {
		dst.InvalidBlockchain = nil
	}

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
		dst.InvalidBlockchain = nil
		dst.InvalidPagination = nil
		dst.LimitGreaterThanAllowed = nil
		dst.UriNotFound = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ConvertBitcoinCashAddressE400)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ConvertBitcoinCashAddressE400)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConvertBitcoinCashAddressE400) MarshalJSON() ([]byte, error) {
	if src.InvalidBlockchain != nil {
		return json.Marshal(&src.InvalidBlockchain)
	}

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
func (obj *ConvertBitcoinCashAddressE400) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvalidBlockchain != nil {
		return obj.InvalidBlockchain
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

type NullableConvertBitcoinCashAddressE400 struct {
	value *ConvertBitcoinCashAddressE400
	isSet bool
}

func (v NullableConvertBitcoinCashAddressE400) Get() *ConvertBitcoinCashAddressE400 {
	return v.value
}

func (v *NullableConvertBitcoinCashAddressE400) Set(val *ConvertBitcoinCashAddressE400) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertBitcoinCashAddressE400) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertBitcoinCashAddressE400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertBitcoinCashAddressE400(val *ConvertBitcoinCashAddressE400) *NullableConvertBitcoinCashAddressE400 {
	return &NullableConvertBitcoinCashAddressE400{value: val, isSet: true}
}

func (v NullableConvertBitcoinCashAddressE400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertBitcoinCashAddressE400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


