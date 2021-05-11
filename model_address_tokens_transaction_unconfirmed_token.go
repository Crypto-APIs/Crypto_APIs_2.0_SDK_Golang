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

// AddressTokensTransactionUnconfirmedToken - struct for AddressTokensTransactionUnconfirmedToken
type AddressTokensTransactionUnconfirmedToken struct {
	AddressTokensTransactionUnconfirmedEthereumerc20token *AddressTokensTransactionUnconfirmedEthereumerc20token
	AddressTokensTransactionUnconfirmedEthereumerc721token *AddressTokensTransactionUnconfirmedEthereumerc721token
	AddressTokensTransactionUnconfirmedOmnilayertoken *AddressTokensTransactionUnconfirmedOmnilayertoken
}

// AddressTokensTransactionUnconfirmedEthereumerc20tokenAsAddressTokensTransactionUnconfirmedToken is a convenience function that returns AddressTokensTransactionUnconfirmedEthereumerc20token wrapped in AddressTokensTransactionUnconfirmedToken
func AddressTokensTransactionUnconfirmedEthereumerc20tokenAsAddressTokensTransactionUnconfirmedToken(v *AddressTokensTransactionUnconfirmedEthereumerc20token) AddressTokensTransactionUnconfirmedToken {
	return AddressTokensTransactionUnconfirmedToken{ AddressTokensTransactionUnconfirmedEthereumerc20token: v}
}

// AddressTokensTransactionUnconfirmedEthereumerc721tokenAsAddressTokensTransactionUnconfirmedToken is a convenience function that returns AddressTokensTransactionUnconfirmedEthereumerc721token wrapped in AddressTokensTransactionUnconfirmedToken
func AddressTokensTransactionUnconfirmedEthereumerc721tokenAsAddressTokensTransactionUnconfirmedToken(v *AddressTokensTransactionUnconfirmedEthereumerc721token) AddressTokensTransactionUnconfirmedToken {
	return AddressTokensTransactionUnconfirmedToken{ AddressTokensTransactionUnconfirmedEthereumerc721token: v}
}

// AddressTokensTransactionUnconfirmedOmnilayertokenAsAddressTokensTransactionUnconfirmedToken is a convenience function that returns AddressTokensTransactionUnconfirmedOmnilayertoken wrapped in AddressTokensTransactionUnconfirmedToken
func AddressTokensTransactionUnconfirmedOmnilayertokenAsAddressTokensTransactionUnconfirmedToken(v *AddressTokensTransactionUnconfirmedOmnilayertoken) AddressTokensTransactionUnconfirmedToken {
	return AddressTokensTransactionUnconfirmedToken{ AddressTokensTransactionUnconfirmedOmnilayertoken: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddressTokensTransactionUnconfirmedToken) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddressTokensTransactionUnconfirmedEthereumerc20token
	err = json.Unmarshal(data, &dst.AddressTokensTransactionUnconfirmedEthereumerc20token)
	if err == nil {
		jsonAddressTokensTransactionUnconfirmedEthereumerc20token, _ := json.Marshal(dst.AddressTokensTransactionUnconfirmedEthereumerc20token)
		if string(jsonAddressTokensTransactionUnconfirmedEthereumerc20token) == "{}" { // empty struct
			dst.AddressTokensTransactionUnconfirmedEthereumerc20token = nil
		} else {
			match++
		}
	} else {
		dst.AddressTokensTransactionUnconfirmedEthereumerc20token = nil
	}

	// try to unmarshal data into AddressTokensTransactionUnconfirmedEthereumerc721token
	err = json.Unmarshal(data, &dst.AddressTokensTransactionUnconfirmedEthereumerc721token)
	if err == nil {
		jsonAddressTokensTransactionUnconfirmedEthereumerc721token, _ := json.Marshal(dst.AddressTokensTransactionUnconfirmedEthereumerc721token)
		if string(jsonAddressTokensTransactionUnconfirmedEthereumerc721token) == "{}" { // empty struct
			dst.AddressTokensTransactionUnconfirmedEthereumerc721token = nil
		} else {
			match++
		}
	} else {
		dst.AddressTokensTransactionUnconfirmedEthereumerc721token = nil
	}

	// try to unmarshal data into AddressTokensTransactionUnconfirmedOmnilayertoken
	err = json.Unmarshal(data, &dst.AddressTokensTransactionUnconfirmedOmnilayertoken)
	if err == nil {
		jsonAddressTokensTransactionUnconfirmedOmnilayertoken, _ := json.Marshal(dst.AddressTokensTransactionUnconfirmedOmnilayertoken)
		if string(jsonAddressTokensTransactionUnconfirmedOmnilayertoken) == "{}" { // empty struct
			dst.AddressTokensTransactionUnconfirmedOmnilayertoken = nil
		} else {
			match++
		}
	} else {
		dst.AddressTokensTransactionUnconfirmedOmnilayertoken = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddressTokensTransactionUnconfirmedEthereumerc20token = nil
		dst.AddressTokensTransactionUnconfirmedEthereumerc721token = nil
		dst.AddressTokensTransactionUnconfirmedOmnilayertoken = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AddressTokensTransactionUnconfirmedToken)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AddressTokensTransactionUnconfirmedToken)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddressTokensTransactionUnconfirmedToken) MarshalJSON() ([]byte, error) {
	if src.AddressTokensTransactionUnconfirmedEthereumerc20token != nil {
		return json.Marshal(&src.AddressTokensTransactionUnconfirmedEthereumerc20token)
	}

	if src.AddressTokensTransactionUnconfirmedEthereumerc721token != nil {
		return json.Marshal(&src.AddressTokensTransactionUnconfirmedEthereumerc721token)
	}

	if src.AddressTokensTransactionUnconfirmedOmnilayertoken != nil {
		return json.Marshal(&src.AddressTokensTransactionUnconfirmedOmnilayertoken)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddressTokensTransactionUnconfirmedToken) GetActualInstance() (interface{}) {
	if obj.AddressTokensTransactionUnconfirmedEthereumerc20token != nil {
		return obj.AddressTokensTransactionUnconfirmedEthereumerc20token
	}

	if obj.AddressTokensTransactionUnconfirmedEthereumerc721token != nil {
		return obj.AddressTokensTransactionUnconfirmedEthereumerc721token
	}

	if obj.AddressTokensTransactionUnconfirmedOmnilayertoken != nil {
		return obj.AddressTokensTransactionUnconfirmedOmnilayertoken
	}

	// all schemas are nil
	return nil
}

type NullableAddressTokensTransactionUnconfirmedToken struct {
	value *AddressTokensTransactionUnconfirmedToken
	isSet bool
}

func (v NullableAddressTokensTransactionUnconfirmedToken) Get() *AddressTokensTransactionUnconfirmedToken {
	return v.value
}

func (v *NullableAddressTokensTransactionUnconfirmedToken) Set(val *AddressTokensTransactionUnconfirmedToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokensTransactionUnconfirmedToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokensTransactionUnconfirmedToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokensTransactionUnconfirmedToken(val *AddressTokensTransactionUnconfirmedToken) *NullableAddressTokensTransactionUnconfirmedToken {
	return &NullableAddressTokensTransactionUnconfirmedToken{value: val, isSet: true}
}

func (v NullableAddressTokensTransactionUnconfirmedToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokensTransactionUnconfirmedToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


