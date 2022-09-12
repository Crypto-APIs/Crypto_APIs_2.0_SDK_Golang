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

// SyncHDWalletXPubYPubZPubE409 - struct for SyncHDWalletXPubYPubZPubE409
type SyncHDWalletXPubYPubZPubE409 struct {
	AlreadyExists *AlreadyExists
	InvalidData *InvalidData
}

// AlreadyExistsAsSyncHDWalletXPubYPubZPubE409 is a convenience function that returns AlreadyExists wrapped in SyncHDWalletXPubYPubZPubE409
func AlreadyExistsAsSyncHDWalletXPubYPubZPubE409(v *AlreadyExists) SyncHDWalletXPubYPubZPubE409 {
	return SyncHDWalletXPubYPubZPubE409{
		AlreadyExists: v,
	}
}

// InvalidDataAsSyncHDWalletXPubYPubZPubE409 is a convenience function that returns InvalidData wrapped in SyncHDWalletXPubYPubZPubE409
func InvalidDataAsSyncHDWalletXPubYPubZPubE409(v *InvalidData) SyncHDWalletXPubYPubZPubE409 {
	return SyncHDWalletXPubYPubZPubE409{
		InvalidData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SyncHDWalletXPubYPubZPubE409) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(SyncHDWalletXPubYPubZPubE409)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SyncHDWalletXPubYPubZPubE409)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SyncHDWalletXPubYPubZPubE409) MarshalJSON() ([]byte, error) {
	if src.AlreadyExists != nil {
		return json.Marshal(&src.AlreadyExists)
	}

	if src.InvalidData != nil {
		return json.Marshal(&src.InvalidData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SyncHDWalletXPubYPubZPubE409) GetActualInstance() (interface{}) {
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

type NullableSyncHDWalletXPubYPubZPubE409 struct {
	value *SyncHDWalletXPubYPubZPubE409
	isSet bool
}

func (v NullableSyncHDWalletXPubYPubZPubE409) Get() *SyncHDWalletXPubYPubZPubE409 {
	return v.value
}

func (v *NullableSyncHDWalletXPubYPubZPubE409) Set(val *SyncHDWalletXPubYPubZPubE409) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncHDWalletXPubYPubZPubE409) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncHDWalletXPubYPubZPubE409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncHDWalletXPubYPubZPubE409(val *SyncHDWalletXPubYPubZPubE409) *NullableSyncHDWalletXPubYPubZPubE409 {
	return &NullableSyncHDWalletXPubYPubZPubE409{value: val, isSet: true}
}

func (v NullableSyncHDWalletXPubYPubZPubE409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncHDWalletXPubYPubZPubE409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


