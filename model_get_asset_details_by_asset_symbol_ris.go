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

// GetAssetDetailsByAssetSymbolRIS - Represents a specific asset's data depending on its type (whether it is \"crypto\" or \"fiat\").
type GetAssetDetailsByAssetSymbolRIS struct {
	GetAssetDetailsByAssetSymbolRISC *GetAssetDetailsByAssetSymbolRISC
}

// GetAssetDetailsByAssetSymbolRISCAsGetAssetDetailsByAssetSymbolRIS is a convenience function that returns GetAssetDetailsByAssetSymbolRISC wrapped in GetAssetDetailsByAssetSymbolRIS
func GetAssetDetailsByAssetSymbolRISCAsGetAssetDetailsByAssetSymbolRIS(v *GetAssetDetailsByAssetSymbolRISC) GetAssetDetailsByAssetSymbolRIS {
	return GetAssetDetailsByAssetSymbolRIS{
		GetAssetDetailsByAssetSymbolRISC: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetAssetDetailsByAssetSymbolRIS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetAssetDetailsByAssetSymbolRISC
	err = newStrictDecoder(data).Decode(&dst.GetAssetDetailsByAssetSymbolRISC)
	if err == nil {
		jsonGetAssetDetailsByAssetSymbolRISC, _ := json.Marshal(dst.GetAssetDetailsByAssetSymbolRISC)
		if string(jsonGetAssetDetailsByAssetSymbolRISC) == "{}" { // empty struct
			dst.GetAssetDetailsByAssetSymbolRISC = nil
		} else {
			match++
		}
	} else {
		dst.GetAssetDetailsByAssetSymbolRISC = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetAssetDetailsByAssetSymbolRISC = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetAssetDetailsByAssetSymbolRIS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetAssetDetailsByAssetSymbolRIS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetAssetDetailsByAssetSymbolRIS) MarshalJSON() ([]byte, error) {
	if src.GetAssetDetailsByAssetSymbolRISC != nil {
		return json.Marshal(&src.GetAssetDetailsByAssetSymbolRISC)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetAssetDetailsByAssetSymbolRIS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetAssetDetailsByAssetSymbolRISC != nil {
		return obj.GetAssetDetailsByAssetSymbolRISC
	}

	// all schemas are nil
	return nil
}

type NullableGetAssetDetailsByAssetSymbolRIS struct {
	value *GetAssetDetailsByAssetSymbolRIS
	isSet bool
}

func (v NullableGetAssetDetailsByAssetSymbolRIS) Get() *GetAssetDetailsByAssetSymbolRIS {
	return v.value
}

func (v *NullableGetAssetDetailsByAssetSymbolRIS) Set(val *GetAssetDetailsByAssetSymbolRIS) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetDetailsByAssetSymbolRIS) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetDetailsByAssetSymbolRIS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetDetailsByAssetSymbolRIS(val *GetAssetDetailsByAssetSymbolRIS) *NullableGetAssetDetailsByAssetSymbolRIS {
	return &NullableGetAssetDetailsByAssetSymbolRIS{value: val, isSet: true}
}

func (v NullableGetAssetDetailsByAssetSymbolRIS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetDetailsByAssetSymbolRIS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


