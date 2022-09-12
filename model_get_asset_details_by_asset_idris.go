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

// GetAssetDetailsByAssetIDRIS - Represents a specific asset's data depending on its type (whether it is \"crypto\" or \"fiat\").
type GetAssetDetailsByAssetIDRIS struct {
	GetAssetDetailsByAssetIDRISC *GetAssetDetailsByAssetIDRISC
}

// GetAssetDetailsByAssetIDRISCAsGetAssetDetailsByAssetIDRIS is a convenience function that returns GetAssetDetailsByAssetIDRISC wrapped in GetAssetDetailsByAssetIDRIS
func GetAssetDetailsByAssetIDRISCAsGetAssetDetailsByAssetIDRIS(v *GetAssetDetailsByAssetIDRISC) GetAssetDetailsByAssetIDRIS {
	return GetAssetDetailsByAssetIDRIS{
		GetAssetDetailsByAssetIDRISC: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetAssetDetailsByAssetIDRIS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetAssetDetailsByAssetIDRISC
	err = newStrictDecoder(data).Decode(&dst.GetAssetDetailsByAssetIDRISC)
	if err == nil {
		jsonGetAssetDetailsByAssetIDRISC, _ := json.Marshal(dst.GetAssetDetailsByAssetIDRISC)
		if string(jsonGetAssetDetailsByAssetIDRISC) == "{}" { // empty struct
			dst.GetAssetDetailsByAssetIDRISC = nil
		} else {
			match++
		}
	} else {
		dst.GetAssetDetailsByAssetIDRISC = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetAssetDetailsByAssetIDRISC = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetAssetDetailsByAssetIDRIS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetAssetDetailsByAssetIDRIS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetAssetDetailsByAssetIDRIS) MarshalJSON() ([]byte, error) {
	if src.GetAssetDetailsByAssetIDRISC != nil {
		return json.Marshal(&src.GetAssetDetailsByAssetIDRISC)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetAssetDetailsByAssetIDRIS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetAssetDetailsByAssetIDRISC != nil {
		return obj.GetAssetDetailsByAssetIDRISC
	}

	// all schemas are nil
	return nil
}

type NullableGetAssetDetailsByAssetIDRIS struct {
	value *GetAssetDetailsByAssetIDRIS
	isSet bool
}

func (v NullableGetAssetDetailsByAssetIDRIS) Get() *GetAssetDetailsByAssetIDRIS {
	return v.value
}

func (v *NullableGetAssetDetailsByAssetIDRIS) Set(val *GetAssetDetailsByAssetIDRIS) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetDetailsByAssetIDRIS) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetDetailsByAssetIDRIS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetDetailsByAssetIDRIS(val *GetAssetDetailsByAssetIDRIS) *NullableGetAssetDetailsByAssetIDRIS {
	return &NullableGetAssetDetailsByAssetIDRIS{value: val, isSet: true}
}

func (v NullableGetAssetDetailsByAssetIDRIS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetDetailsByAssetIDRIS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


