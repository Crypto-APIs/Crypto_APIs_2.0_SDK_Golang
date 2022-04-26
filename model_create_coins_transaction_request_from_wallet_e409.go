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

// CreateCoinsTransactionRequestFromWalletE409 - struct for CreateCoinsTransactionRequestFromWalletE409
type CreateCoinsTransactionRequestFromWalletE409 struct {
	InvalidData *InvalidData
	WalletAsAServiceAddressBalanceNotEnough *WalletAsAServiceAddressBalanceNotEnough
	WalletAsAServiceNoDepositAddressesFound *WalletAsAServiceNoDepositAddressesFound
	WalletAsAServiceWalletBalanceNotEnough *WalletAsAServiceWalletBalanceNotEnough
}

// InvalidDataAsCreateCoinsTransactionRequestFromWalletE409 is a convenience function that returns InvalidData wrapped in CreateCoinsTransactionRequestFromWalletE409
func InvalidDataAsCreateCoinsTransactionRequestFromWalletE409(v *InvalidData) CreateCoinsTransactionRequestFromWalletE409 {
	return CreateCoinsTransactionRequestFromWalletE409{
		InvalidData: v,
	}
}

// WalletAsAServiceAddressBalanceNotEnoughAsCreateCoinsTransactionRequestFromWalletE409 is a convenience function that returns WalletAsAServiceAddressBalanceNotEnough wrapped in CreateCoinsTransactionRequestFromWalletE409
func WalletAsAServiceAddressBalanceNotEnoughAsCreateCoinsTransactionRequestFromWalletE409(v *WalletAsAServiceAddressBalanceNotEnough) CreateCoinsTransactionRequestFromWalletE409 {
	return CreateCoinsTransactionRequestFromWalletE409{
		WalletAsAServiceAddressBalanceNotEnough: v,
	}
}

// WalletAsAServiceNoDepositAddressesFoundAsCreateCoinsTransactionRequestFromWalletE409 is a convenience function that returns WalletAsAServiceNoDepositAddressesFound wrapped in CreateCoinsTransactionRequestFromWalletE409
func WalletAsAServiceNoDepositAddressesFoundAsCreateCoinsTransactionRequestFromWalletE409(v *WalletAsAServiceNoDepositAddressesFound) CreateCoinsTransactionRequestFromWalletE409 {
	return CreateCoinsTransactionRequestFromWalletE409{
		WalletAsAServiceNoDepositAddressesFound: v,
	}
}

// WalletAsAServiceWalletBalanceNotEnoughAsCreateCoinsTransactionRequestFromWalletE409 is a convenience function that returns WalletAsAServiceWalletBalanceNotEnough wrapped in CreateCoinsTransactionRequestFromWalletE409
func WalletAsAServiceWalletBalanceNotEnoughAsCreateCoinsTransactionRequestFromWalletE409(v *WalletAsAServiceWalletBalanceNotEnough) CreateCoinsTransactionRequestFromWalletE409 {
	return CreateCoinsTransactionRequestFromWalletE409{
		WalletAsAServiceWalletBalanceNotEnough: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateCoinsTransactionRequestFromWalletE409) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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

	// try to unmarshal data into WalletAsAServiceAddressBalanceNotEnough
	err = newStrictDecoder(data).Decode(&dst.WalletAsAServiceAddressBalanceNotEnough)
	if err == nil {
		jsonWalletAsAServiceAddressBalanceNotEnough, _ := json.Marshal(dst.WalletAsAServiceAddressBalanceNotEnough)
		if string(jsonWalletAsAServiceAddressBalanceNotEnough) == "{}" { // empty struct
			dst.WalletAsAServiceAddressBalanceNotEnough = nil
		} else {
			match++
		}
	} else {
		dst.WalletAsAServiceAddressBalanceNotEnough = nil
	}

	// try to unmarshal data into WalletAsAServiceNoDepositAddressesFound
	err = newStrictDecoder(data).Decode(&dst.WalletAsAServiceNoDepositAddressesFound)
	if err == nil {
		jsonWalletAsAServiceNoDepositAddressesFound, _ := json.Marshal(dst.WalletAsAServiceNoDepositAddressesFound)
		if string(jsonWalletAsAServiceNoDepositAddressesFound) == "{}" { // empty struct
			dst.WalletAsAServiceNoDepositAddressesFound = nil
		} else {
			match++
		}
	} else {
		dst.WalletAsAServiceNoDepositAddressesFound = nil
	}

	// try to unmarshal data into WalletAsAServiceWalletBalanceNotEnough
	err = newStrictDecoder(data).Decode(&dst.WalletAsAServiceWalletBalanceNotEnough)
	if err == nil {
		jsonWalletAsAServiceWalletBalanceNotEnough, _ := json.Marshal(dst.WalletAsAServiceWalletBalanceNotEnough)
		if string(jsonWalletAsAServiceWalletBalanceNotEnough) == "{}" { // empty struct
			dst.WalletAsAServiceWalletBalanceNotEnough = nil
		} else {
			match++
		}
	} else {
		dst.WalletAsAServiceWalletBalanceNotEnough = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidData = nil
		dst.WalletAsAServiceAddressBalanceNotEnough = nil
		dst.WalletAsAServiceNoDepositAddressesFound = nil
		dst.WalletAsAServiceWalletBalanceNotEnough = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateCoinsTransactionRequestFromWalletE409)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateCoinsTransactionRequestFromWalletE409)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateCoinsTransactionRequestFromWalletE409) MarshalJSON() ([]byte, error) {
	if src.InvalidData != nil {
		return json.Marshal(&src.InvalidData)
	}

	if src.WalletAsAServiceAddressBalanceNotEnough != nil {
		return json.Marshal(&src.WalletAsAServiceAddressBalanceNotEnough)
	}

	if src.WalletAsAServiceNoDepositAddressesFound != nil {
		return json.Marshal(&src.WalletAsAServiceNoDepositAddressesFound)
	}

	if src.WalletAsAServiceWalletBalanceNotEnough != nil {
		return json.Marshal(&src.WalletAsAServiceWalletBalanceNotEnough)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateCoinsTransactionRequestFromWalletE409) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvalidData != nil {
		return obj.InvalidData
	}

	if obj.WalletAsAServiceAddressBalanceNotEnough != nil {
		return obj.WalletAsAServiceAddressBalanceNotEnough
	}

	if obj.WalletAsAServiceNoDepositAddressesFound != nil {
		return obj.WalletAsAServiceNoDepositAddressesFound
	}

	if obj.WalletAsAServiceWalletBalanceNotEnough != nil {
		return obj.WalletAsAServiceWalletBalanceNotEnough
	}

	// all schemas are nil
	return nil
}

type NullableCreateCoinsTransactionRequestFromWalletE409 struct {
	value *CreateCoinsTransactionRequestFromWalletE409
	isSet bool
}

func (v NullableCreateCoinsTransactionRequestFromWalletE409) Get() *CreateCoinsTransactionRequestFromWalletE409 {
	return v.value
}

func (v *NullableCreateCoinsTransactionRequestFromWalletE409) Set(val *CreateCoinsTransactionRequestFromWalletE409) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionRequestFromWalletE409) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionRequestFromWalletE409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionRequestFromWalletE409(val *CreateCoinsTransactionRequestFromWalletE409) *NullableCreateCoinsTransactionRequestFromWalletE409 {
	return &NullableCreateCoinsTransactionRequestFromWalletE409{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionRequestFromWalletE409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionRequestFromWalletE409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


