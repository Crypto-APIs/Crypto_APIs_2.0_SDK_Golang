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

// AddTokensToExistingFromAddressRBTokenData - struct for AddTokensToExistingFromAddressRBTokenData
type AddTokensToExistingFromAddressRBTokenData struct {
	AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken *AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken
	AddTokensToExistingFromAddressRBTokenDataEthereumToken *AddTokensToExistingFromAddressRBTokenDataEthereumToken
}

// AddTokensToExistingFromAddressRBTokenDataBitcoinOmniTokenAsAddTokensToExistingFromAddressRBTokenData is a convenience function that returns AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken wrapped in AddTokensToExistingFromAddressRBTokenData
func AddTokensToExistingFromAddressRBTokenDataBitcoinOmniTokenAsAddTokensToExistingFromAddressRBTokenData(v *AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken) AddTokensToExistingFromAddressRBTokenData {
	return AddTokensToExistingFromAddressRBTokenData{
		AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken: v,
	}
}

// AddTokensToExistingFromAddressRBTokenDataEthereumTokenAsAddTokensToExistingFromAddressRBTokenData is a convenience function that returns AddTokensToExistingFromAddressRBTokenDataEthereumToken wrapped in AddTokensToExistingFromAddressRBTokenData
func AddTokensToExistingFromAddressRBTokenDataEthereumTokenAsAddTokensToExistingFromAddressRBTokenData(v *AddTokensToExistingFromAddressRBTokenDataEthereumToken) AddTokensToExistingFromAddressRBTokenData {
	return AddTokensToExistingFromAddressRBTokenData{
		AddTokensToExistingFromAddressRBTokenDataEthereumToken: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddTokensToExistingFromAddressRBTokenData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken
	err = newStrictDecoder(data).Decode(&dst.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken)
	if err == nil {
		jsonAddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken, _ := json.Marshal(dst.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken)
		if string(jsonAddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken) == "{}" { // empty struct
			dst.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken = nil
		} else {
			match++
		}
	} else {
		dst.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken = nil
	}

	// try to unmarshal data into AddTokensToExistingFromAddressRBTokenDataEthereumToken
	err = newStrictDecoder(data).Decode(&dst.AddTokensToExistingFromAddressRBTokenDataEthereumToken)
	if err == nil {
		jsonAddTokensToExistingFromAddressRBTokenDataEthereumToken, _ := json.Marshal(dst.AddTokensToExistingFromAddressRBTokenDataEthereumToken)
		if string(jsonAddTokensToExistingFromAddressRBTokenDataEthereumToken) == "{}" { // empty struct
			dst.AddTokensToExistingFromAddressRBTokenDataEthereumToken = nil
		} else {
			match++
		}
	} else {
		dst.AddTokensToExistingFromAddressRBTokenDataEthereumToken = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken = nil
		dst.AddTokensToExistingFromAddressRBTokenDataEthereumToken = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AddTokensToExistingFromAddressRBTokenData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AddTokensToExistingFromAddressRBTokenData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddTokensToExistingFromAddressRBTokenData) MarshalJSON() ([]byte, error) {
	if src.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken != nil {
		return json.Marshal(&src.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken)
	}

	if src.AddTokensToExistingFromAddressRBTokenDataEthereumToken != nil {
		return json.Marshal(&src.AddTokensToExistingFromAddressRBTokenDataEthereumToken)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddTokensToExistingFromAddressRBTokenData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken != nil {
		return obj.AddTokensToExistingFromAddressRBTokenDataBitcoinOmniToken
	}

	if obj.AddTokensToExistingFromAddressRBTokenDataEthereumToken != nil {
		return obj.AddTokensToExistingFromAddressRBTokenDataEthereumToken
	}

	// all schemas are nil
	return nil
}

type NullableAddTokensToExistingFromAddressRBTokenData struct {
	value *AddTokensToExistingFromAddressRBTokenData
	isSet bool
}

func (v NullableAddTokensToExistingFromAddressRBTokenData) Get() *AddTokensToExistingFromAddressRBTokenData {
	return v.value
}

func (v *NullableAddTokensToExistingFromAddressRBTokenData) Set(val *AddTokensToExistingFromAddressRBTokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTokensToExistingFromAddressRBTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTokensToExistingFromAddressRBTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTokensToExistingFromAddressRBTokenData(val *AddTokensToExistingFromAddressRBTokenData) *NullableAddTokensToExistingFromAddressRBTokenData {
	return &NullableAddTokensToExistingFromAddressRBTokenData{value: val, isSet: true}
}

func (v NullableAddTokensToExistingFromAddressRBTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTokensToExistingFromAddressRBTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


