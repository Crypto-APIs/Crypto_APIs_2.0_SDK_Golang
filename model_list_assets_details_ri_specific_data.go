/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
	"fmt"
)

// ListAssetsDetailsRISpecificData - Represents a specific asset's data depending on its type (whether it is \"crypto\" or \"fiat\").
type ListAssetsDetailsRISpecificData struct {
	ListAssetsDetailsRISpecificDataCryptoTypeData *ListAssetsDetailsRISpecificDataCryptoTypeData
}

// ListAssetsDetailsRISpecificDataCryptoTypeDataAsListAssetsDetailsRISpecificData is a convenience function that returns ListAssetsDetailsRISpecificDataCryptoTypeData wrapped in ListAssetsDetailsRISpecificData
func ListAssetsDetailsRISpecificDataCryptoTypeDataAsListAssetsDetailsRISpecificData(v *ListAssetsDetailsRISpecificDataCryptoTypeData) ListAssetsDetailsRISpecificData {
	return ListAssetsDetailsRISpecificData{ ListAssetsDetailsRISpecificDataCryptoTypeData: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListAssetsDetailsRISpecificData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListAssetsDetailsRISpecificDataCryptoTypeData
	err = json.Unmarshal(data, &dst.ListAssetsDetailsRISpecificDataCryptoTypeData)
	if err == nil {
		jsonListAssetsDetailsRISpecificDataCryptoTypeData, _ := json.Marshal(dst.ListAssetsDetailsRISpecificDataCryptoTypeData)
		if string(jsonListAssetsDetailsRISpecificDataCryptoTypeData) == "{}" { // empty struct
			dst.ListAssetsDetailsRISpecificDataCryptoTypeData = nil
		} else {
			match++
		}
	} else {
		dst.ListAssetsDetailsRISpecificDataCryptoTypeData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListAssetsDetailsRISpecificDataCryptoTypeData = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListAssetsDetailsRISpecificData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListAssetsDetailsRISpecificData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAssetsDetailsRISpecificData) MarshalJSON() ([]byte, error) {
	if src.ListAssetsDetailsRISpecificDataCryptoTypeData != nil {
		return json.Marshal(&src.ListAssetsDetailsRISpecificDataCryptoTypeData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAssetsDetailsRISpecificData) GetActualInstance() (interface{}) {
	if obj.ListAssetsDetailsRISpecificDataCryptoTypeData != nil {
		return obj.ListAssetsDetailsRISpecificDataCryptoTypeData
	}

	// all schemas are nil
	return nil
}

type NullableListAssetsDetailsRISpecificData struct {
	value *ListAssetsDetailsRISpecificData
	isSet bool
}

func (v NullableListAssetsDetailsRISpecificData) Get() *ListAssetsDetailsRISpecificData {
	return v.value
}

func (v *NullableListAssetsDetailsRISpecificData) Set(val *ListAssetsDetailsRISpecificData) {
	v.value = val
	v.isSet = true
}

func (v NullableListAssetsDetailsRISpecificData) IsSet() bool {
	return v.isSet
}

func (v *NullableListAssetsDetailsRISpecificData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAssetsDetailsRISpecificData(val *ListAssetsDetailsRISpecificData) *NullableListAssetsDetailsRISpecificData {
	return &NullableListAssetsDetailsRISpecificData{value: val, isSet: true}
}

func (v NullableListAssetsDetailsRISpecificData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAssetsDetailsRISpecificData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

