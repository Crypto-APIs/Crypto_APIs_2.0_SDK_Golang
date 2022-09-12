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

// NewConfirmedTokensTransactionsAndEachConfirmationE400 - struct for NewConfirmedTokensTransactionsAndEachConfirmationE400
type NewConfirmedTokensTransactionsAndEachConfirmationE400 struct {
	InvalidPagination *InvalidPagination
	LimitGreaterThanAllowed *LimitGreaterThanAllowed
	UriNotFound *UriNotFound
}

// InvalidPaginationAsNewConfirmedTokensTransactionsAndEachConfirmationE400 is a convenience function that returns InvalidPagination wrapped in NewConfirmedTokensTransactionsAndEachConfirmationE400
func InvalidPaginationAsNewConfirmedTokensTransactionsAndEachConfirmationE400(v *InvalidPagination) NewConfirmedTokensTransactionsAndEachConfirmationE400 {
	return NewConfirmedTokensTransactionsAndEachConfirmationE400{
		InvalidPagination: v,
	}
}

// LimitGreaterThanAllowedAsNewConfirmedTokensTransactionsAndEachConfirmationE400 is a convenience function that returns LimitGreaterThanAllowed wrapped in NewConfirmedTokensTransactionsAndEachConfirmationE400
func LimitGreaterThanAllowedAsNewConfirmedTokensTransactionsAndEachConfirmationE400(v *LimitGreaterThanAllowed) NewConfirmedTokensTransactionsAndEachConfirmationE400 {
	return NewConfirmedTokensTransactionsAndEachConfirmationE400{
		LimitGreaterThanAllowed: v,
	}
}

// UriNotFoundAsNewConfirmedTokensTransactionsAndEachConfirmationE400 is a convenience function that returns UriNotFound wrapped in NewConfirmedTokensTransactionsAndEachConfirmationE400
func UriNotFoundAsNewConfirmedTokensTransactionsAndEachConfirmationE400(v *UriNotFound) NewConfirmedTokensTransactionsAndEachConfirmationE400 {
	return NewConfirmedTokensTransactionsAndEachConfirmationE400{
		UriNotFound: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewConfirmedTokensTransactionsAndEachConfirmationE400) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(NewConfirmedTokensTransactionsAndEachConfirmationE400)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(NewConfirmedTokensTransactionsAndEachConfirmationE400)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewConfirmedTokensTransactionsAndEachConfirmationE400) MarshalJSON() ([]byte, error) {
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
func (obj *NewConfirmedTokensTransactionsAndEachConfirmationE400) GetActualInstance() (interface{}) {
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

type NullableNewConfirmedTokensTransactionsAndEachConfirmationE400 struct {
	value *NewConfirmedTokensTransactionsAndEachConfirmationE400
	isSet bool
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationE400) Get() *NewConfirmedTokensTransactionsAndEachConfirmationE400 {
	return v.value
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationE400) Set(val *NewConfirmedTokensTransactionsAndEachConfirmationE400) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationE400) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationE400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedTokensTransactionsAndEachConfirmationE400(val *NewConfirmedTokensTransactionsAndEachConfirmationE400) *NullableNewConfirmedTokensTransactionsAndEachConfirmationE400 {
	return &NullableNewConfirmedTokensTransactionsAndEachConfirmationE400{value: val, isSet: true}
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationE400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationE400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


