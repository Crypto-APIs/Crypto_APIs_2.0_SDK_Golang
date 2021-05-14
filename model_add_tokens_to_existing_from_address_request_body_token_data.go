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

// AddTokensToExistingFromAddressRequestBodyTokenData - struct for AddTokensToExistingFromAddressRequestBodyTokenData
type AddTokensToExistingFromAddressRequestBodyTokenData struct {
	AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken *AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken
	AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token *AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token
}

// AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniTokenAsAddTokensToExistingFromAddressRequestBodyTokenData is a convenience function that returns AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken wrapped in AddTokensToExistingFromAddressRequestBodyTokenData
func AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniTokenAsAddTokensToExistingFromAddressRequestBodyTokenData(v *AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken) AddTokensToExistingFromAddressRequestBodyTokenData {
	return AddTokensToExistingFromAddressRequestBodyTokenData{ AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken: v}
}

// AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20TokenAsAddTokensToExistingFromAddressRequestBodyTokenData is a convenience function that returns AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token wrapped in AddTokensToExistingFromAddressRequestBodyTokenData
func AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20TokenAsAddTokensToExistingFromAddressRequestBodyTokenData(v *AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token) AddTokensToExistingFromAddressRequestBodyTokenData {
	return AddTokensToExistingFromAddressRequestBodyTokenData{ AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddTokensToExistingFromAddressRequestBodyTokenData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken
	err = json.Unmarshal(data, &dst.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken)
	if err == nil {
		jsonAddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken, _ := json.Marshal(dst.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken)
		if string(jsonAddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken) == "{}" { // empty struct
			dst.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken = nil
		} else {
			match++
		}
	} else {
		dst.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken = nil
	}

	// try to unmarshal data into AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token
	err = json.Unmarshal(data, &dst.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token)
	if err == nil {
		jsonAddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token, _ := json.Marshal(dst.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token)
		if string(jsonAddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token) == "{}" { // empty struct
			dst.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token = nil
		} else {
			match++
		}
	} else {
		dst.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken = nil
		dst.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AddTokensToExistingFromAddressRequestBodyTokenData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AddTokensToExistingFromAddressRequestBodyTokenData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddTokensToExistingFromAddressRequestBodyTokenData) MarshalJSON() ([]byte, error) {
	if src.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken != nil {
		return json.Marshal(&src.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken)
	}

	if src.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token != nil {
		return json.Marshal(&src.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddTokensToExistingFromAddressRequestBodyTokenData) GetActualInstance() (interface{}) {
	if obj.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken != nil {
		return obj.AddTokensToExistingFromAddressRequestBodyTokenDataBitcoinOmniToken
	}

	if obj.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token != nil {
		return obj.AddTokensToExistingFromAddressRequestBodyTokenDataEthereumErc20Token
	}

	// all schemas are nil
	return nil
}

type NullableAddTokensToExistingFromAddressRequestBodyTokenData struct {
	value *AddTokensToExistingFromAddressRequestBodyTokenData
	isSet bool
}

func (v NullableAddTokensToExistingFromAddressRequestBodyTokenData) Get() *AddTokensToExistingFromAddressRequestBodyTokenData {
	return v.value
}

func (v *NullableAddTokensToExistingFromAddressRequestBodyTokenData) Set(val *AddTokensToExistingFromAddressRequestBodyTokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTokensToExistingFromAddressRequestBodyTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTokensToExistingFromAddressRequestBodyTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTokensToExistingFromAddressRequestBodyTokenData(val *AddTokensToExistingFromAddressRequestBodyTokenData) *NullableAddTokensToExistingFromAddressRequestBodyTokenData {
	return &NullableAddTokensToExistingFromAddressRequestBodyTokenData{value: val, isSet: true}
}

func (v NullableAddTokensToExistingFromAddressRequestBodyTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTokensToExistingFromAddressRequestBodyTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

