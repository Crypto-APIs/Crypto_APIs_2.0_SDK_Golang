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

// AddTokensToExistingFromAddressResponseItemTokenData - struct for AddTokensToExistingFromAddressResponseItemTokenData
type AddTokensToExistingFromAddressResponseItemTokenData struct {
	AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken *AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken
	AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token *AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token
}

// AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniTokenAsAddTokensToExistingFromAddressResponseItemTokenData is a convenience function that returns AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken wrapped in AddTokensToExistingFromAddressResponseItemTokenData
func AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniTokenAsAddTokensToExistingFromAddressResponseItemTokenData(v *AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken) AddTokensToExistingFromAddressResponseItemTokenData {
	return AddTokensToExistingFromAddressResponseItemTokenData{ AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken: v}
}

// AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20TokenAsAddTokensToExistingFromAddressResponseItemTokenData is a convenience function that returns AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token wrapped in AddTokensToExistingFromAddressResponseItemTokenData
func AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20TokenAsAddTokensToExistingFromAddressResponseItemTokenData(v *AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token) AddTokensToExistingFromAddressResponseItemTokenData {
	return AddTokensToExistingFromAddressResponseItemTokenData{ AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddTokensToExistingFromAddressResponseItemTokenData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken
	err = json.Unmarshal(data, &dst.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken)
	if err == nil {
		jsonAddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken, _ := json.Marshal(dst.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken)
		if string(jsonAddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken) == "{}" { // empty struct
			dst.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken = nil
		} else {
			match++
		}
	} else {
		dst.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken = nil
	}

	// try to unmarshal data into AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token
	err = json.Unmarshal(data, &dst.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token)
	if err == nil {
		jsonAddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token, _ := json.Marshal(dst.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token)
		if string(jsonAddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token) == "{}" { // empty struct
			dst.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token = nil
		} else {
			match++
		}
	} else {
		dst.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken = nil
		dst.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AddTokensToExistingFromAddressResponseItemTokenData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AddTokensToExistingFromAddressResponseItemTokenData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddTokensToExistingFromAddressResponseItemTokenData) MarshalJSON() ([]byte, error) {
	if src.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken != nil {
		return json.Marshal(&src.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken)
	}

	if src.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token != nil {
		return json.Marshal(&src.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddTokensToExistingFromAddressResponseItemTokenData) GetActualInstance() (interface{}) {
	if obj.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken != nil {
		return obj.AddTokensToExistingFromAddressResponseItemTokenDataBitcoinOmniToken
	}

	if obj.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token != nil {
		return obj.AddTokensToExistingFromAddressResponseItemTokenDataEthereumErc20Token
	}

	// all schemas are nil
	return nil
}

type NullableAddTokensToExistingFromAddressResponseItemTokenData struct {
	value *AddTokensToExistingFromAddressResponseItemTokenData
	isSet bool
}

func (v NullableAddTokensToExistingFromAddressResponseItemTokenData) Get() *AddTokensToExistingFromAddressResponseItemTokenData {
	return v.value
}

func (v *NullableAddTokensToExistingFromAddressResponseItemTokenData) Set(val *AddTokensToExistingFromAddressResponseItemTokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTokensToExistingFromAddressResponseItemTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTokensToExistingFromAddressResponseItemTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTokensToExistingFromAddressResponseItemTokenData(val *AddTokensToExistingFromAddressResponseItemTokenData) *NullableAddTokensToExistingFromAddressResponseItemTokenData {
	return &NullableAddTokensToExistingFromAddressResponseItemTokenData{value: val, isSet: true}
}

func (v NullableAddTokensToExistingFromAddressResponseItemTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTokensToExistingFromAddressResponseItemTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

