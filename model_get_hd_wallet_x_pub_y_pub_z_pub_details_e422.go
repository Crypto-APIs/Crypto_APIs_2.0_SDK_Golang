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

// GetHDWalletXPubYPubZPubDetailsE422 - struct for GetHDWalletXPubYPubZPubDetailsE422
type GetHDWalletXPubYPubZPubDetailsE422 struct {
	InvalidRequestBodyStructure *InvalidRequestBodyStructure
	XpubSyncInProgress *XpubSyncInProgress
}

// InvalidRequestBodyStructureAsGetHDWalletXPubYPubZPubDetailsE422 is a convenience function that returns InvalidRequestBodyStructure wrapped in GetHDWalletXPubYPubZPubDetailsE422
func InvalidRequestBodyStructureAsGetHDWalletXPubYPubZPubDetailsE422(v *InvalidRequestBodyStructure) GetHDWalletXPubYPubZPubDetailsE422 {
	return GetHDWalletXPubYPubZPubDetailsE422{
		InvalidRequestBodyStructure: v,
	}
}

// XpubSyncInProgressAsGetHDWalletXPubYPubZPubDetailsE422 is a convenience function that returns XpubSyncInProgress wrapped in GetHDWalletXPubYPubZPubDetailsE422
func XpubSyncInProgressAsGetHDWalletXPubYPubZPubDetailsE422(v *XpubSyncInProgress) GetHDWalletXPubYPubZPubDetailsE422 {
	return GetHDWalletXPubYPubZPubDetailsE422{
		XpubSyncInProgress: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetHDWalletXPubYPubZPubDetailsE422) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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

	// try to unmarshal data into XpubSyncInProgress
	err = newStrictDecoder(data).Decode(&dst.XpubSyncInProgress)
	if err == nil {
		jsonXpubSyncInProgress, _ := json.Marshal(dst.XpubSyncInProgress)
		if string(jsonXpubSyncInProgress) == "{}" { // empty struct
			dst.XpubSyncInProgress = nil
		} else {
			match++
		}
	} else {
		dst.XpubSyncInProgress = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidRequestBodyStructure = nil
		dst.XpubSyncInProgress = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetHDWalletXPubYPubZPubDetailsE422)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetHDWalletXPubYPubZPubDetailsE422)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetHDWalletXPubYPubZPubDetailsE422) MarshalJSON() ([]byte, error) {
	if src.InvalidRequestBodyStructure != nil {
		return json.Marshal(&src.InvalidRequestBodyStructure)
	}

	if src.XpubSyncInProgress != nil {
		return json.Marshal(&src.XpubSyncInProgress)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetHDWalletXPubYPubZPubDetailsE422) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvalidRequestBodyStructure != nil {
		return obj.InvalidRequestBodyStructure
	}

	if obj.XpubSyncInProgress != nil {
		return obj.XpubSyncInProgress
	}

	// all schemas are nil
	return nil
}

type NullableGetHDWalletXPubYPubZPubDetailsE422 struct {
	value *GetHDWalletXPubYPubZPubDetailsE422
	isSet bool
}

func (v NullableGetHDWalletXPubYPubZPubDetailsE422) Get() *GetHDWalletXPubYPubZPubDetailsE422 {
	return v.value
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsE422) Set(val *GetHDWalletXPubYPubZPubDetailsE422) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHDWalletXPubYPubZPubDetailsE422) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsE422) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHDWalletXPubYPubZPubDetailsE422(val *GetHDWalletXPubYPubZPubDetailsE422) *NullableGetHDWalletXPubYPubZPubDetailsE422 {
	return &NullableGetHDWalletXPubYPubZPubDetailsE422{value: val, isSet: true}
}

func (v NullableGetHDWalletXPubYPubZPubDetailsE422) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsE422) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


