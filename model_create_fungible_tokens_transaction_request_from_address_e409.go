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

// CreateFungibleTokensTransactionRequestFromAddressE409 - struct for CreateFungibleTokensTransactionRequestFromAddressE409
type CreateFungibleTokensTransactionRequestFromAddressE409 struct {
	InvalidData *InvalidData
	WalletAsAServiceAddressBalanceNotEnough *WalletAsAServiceAddressBalanceNotEnough
	WalletAsAServiceTokenNotSupported *WalletAsAServiceTokenNotSupported
	WalletAsAServiceWalletBalanceNotEnough *WalletAsAServiceWalletBalanceNotEnough
}

// InvalidDataAsCreateFungibleTokensTransactionRequestFromAddressE409 is a convenience function that returns InvalidData wrapped in CreateFungibleTokensTransactionRequestFromAddressE409
func InvalidDataAsCreateFungibleTokensTransactionRequestFromAddressE409(v *InvalidData) CreateFungibleTokensTransactionRequestFromAddressE409 {
	return CreateFungibleTokensTransactionRequestFromAddressE409{
		InvalidData: v,
	}
}

// WalletAsAServiceAddressBalanceNotEnoughAsCreateFungibleTokensTransactionRequestFromAddressE409 is a convenience function that returns WalletAsAServiceAddressBalanceNotEnough wrapped in CreateFungibleTokensTransactionRequestFromAddressE409
func WalletAsAServiceAddressBalanceNotEnoughAsCreateFungibleTokensTransactionRequestFromAddressE409(v *WalletAsAServiceAddressBalanceNotEnough) CreateFungibleTokensTransactionRequestFromAddressE409 {
	return CreateFungibleTokensTransactionRequestFromAddressE409{
		WalletAsAServiceAddressBalanceNotEnough: v,
	}
}

// WalletAsAServiceTokenNotSupportedAsCreateFungibleTokensTransactionRequestFromAddressE409 is a convenience function that returns WalletAsAServiceTokenNotSupported wrapped in CreateFungibleTokensTransactionRequestFromAddressE409
func WalletAsAServiceTokenNotSupportedAsCreateFungibleTokensTransactionRequestFromAddressE409(v *WalletAsAServiceTokenNotSupported) CreateFungibleTokensTransactionRequestFromAddressE409 {
	return CreateFungibleTokensTransactionRequestFromAddressE409{
		WalletAsAServiceTokenNotSupported: v,
	}
}

// WalletAsAServiceWalletBalanceNotEnoughAsCreateFungibleTokensTransactionRequestFromAddressE409 is a convenience function that returns WalletAsAServiceWalletBalanceNotEnough wrapped in CreateFungibleTokensTransactionRequestFromAddressE409
func WalletAsAServiceWalletBalanceNotEnoughAsCreateFungibleTokensTransactionRequestFromAddressE409(v *WalletAsAServiceWalletBalanceNotEnough) CreateFungibleTokensTransactionRequestFromAddressE409 {
	return CreateFungibleTokensTransactionRequestFromAddressE409{
		WalletAsAServiceWalletBalanceNotEnough: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateFungibleTokensTransactionRequestFromAddressE409) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into WalletAsAServiceTokenNotSupported
	err = newStrictDecoder(data).Decode(&dst.WalletAsAServiceTokenNotSupported)
	if err == nil {
		jsonWalletAsAServiceTokenNotSupported, _ := json.Marshal(dst.WalletAsAServiceTokenNotSupported)
		if string(jsonWalletAsAServiceTokenNotSupported) == "{}" { // empty struct
			dst.WalletAsAServiceTokenNotSupported = nil
		} else {
			match++
		}
	} else {
		dst.WalletAsAServiceTokenNotSupported = nil
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
		dst.WalletAsAServiceTokenNotSupported = nil
		dst.WalletAsAServiceWalletBalanceNotEnough = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateFungibleTokensTransactionRequestFromAddressE409)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateFungibleTokensTransactionRequestFromAddressE409)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateFungibleTokensTransactionRequestFromAddressE409) MarshalJSON() ([]byte, error) {
	if src.InvalidData != nil {
		return json.Marshal(&src.InvalidData)
	}

	if src.WalletAsAServiceAddressBalanceNotEnough != nil {
		return json.Marshal(&src.WalletAsAServiceAddressBalanceNotEnough)
	}

	if src.WalletAsAServiceTokenNotSupported != nil {
		return json.Marshal(&src.WalletAsAServiceTokenNotSupported)
	}

	if src.WalletAsAServiceWalletBalanceNotEnough != nil {
		return json.Marshal(&src.WalletAsAServiceWalletBalanceNotEnough)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateFungibleTokensTransactionRequestFromAddressE409) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvalidData != nil {
		return obj.InvalidData
	}

	if obj.WalletAsAServiceAddressBalanceNotEnough != nil {
		return obj.WalletAsAServiceAddressBalanceNotEnough
	}

	if obj.WalletAsAServiceTokenNotSupported != nil {
		return obj.WalletAsAServiceTokenNotSupported
	}

	if obj.WalletAsAServiceWalletBalanceNotEnough != nil {
		return obj.WalletAsAServiceWalletBalanceNotEnough
	}

	// all schemas are nil
	return nil
}

type NullableCreateFungibleTokensTransactionRequestFromAddressE409 struct {
	value *CreateFungibleTokensTransactionRequestFromAddressE409
	isSet bool
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressE409) Get() *CreateFungibleTokensTransactionRequestFromAddressE409 {
	return v.value
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressE409) Set(val *CreateFungibleTokensTransactionRequestFromAddressE409) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressE409) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressE409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokensTransactionRequestFromAddressE409(val *CreateFungibleTokensTransactionRequestFromAddressE409) *NullableCreateFungibleTokensTransactionRequestFromAddressE409 {
	return &NullableCreateFungibleTokensTransactionRequestFromAddressE409{value: val, isSet: true}
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressE409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressE409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


