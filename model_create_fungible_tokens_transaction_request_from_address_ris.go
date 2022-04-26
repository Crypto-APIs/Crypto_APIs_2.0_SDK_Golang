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

// CreateFungibleTokensTransactionRequestFromAddressRIS - Represents the specific token data which depends on its type - if it is a Coin or Token.
type CreateFungibleTokensTransactionRequestFromAddressRIS struct {
	CreateFungibleTokensTransactionRequestFromAddressRISE *CreateFungibleTokensTransactionRequestFromAddressRISE
}

// CreateFungibleTokensTransactionRequestFromAddressRISEAsCreateFungibleTokensTransactionRequestFromAddressRIS is a convenience function that returns CreateFungibleTokensTransactionRequestFromAddressRISE wrapped in CreateFungibleTokensTransactionRequestFromAddressRIS
func CreateFungibleTokensTransactionRequestFromAddressRISEAsCreateFungibleTokensTransactionRequestFromAddressRIS(v *CreateFungibleTokensTransactionRequestFromAddressRISE) CreateFungibleTokensTransactionRequestFromAddressRIS {
	return CreateFungibleTokensTransactionRequestFromAddressRIS{
		CreateFungibleTokensTransactionRequestFromAddressRISE: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateFungibleTokensTransactionRequestFromAddressRIS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateFungibleTokensTransactionRequestFromAddressRISE
	err = newStrictDecoder(data).Decode(&dst.CreateFungibleTokensTransactionRequestFromAddressRISE)
	if err == nil {
		jsonCreateFungibleTokensTransactionRequestFromAddressRISE, _ := json.Marshal(dst.CreateFungibleTokensTransactionRequestFromAddressRISE)
		if string(jsonCreateFungibleTokensTransactionRequestFromAddressRISE) == "{}" { // empty struct
			dst.CreateFungibleTokensTransactionRequestFromAddressRISE = nil
		} else {
			match++
		}
	} else {
		dst.CreateFungibleTokensTransactionRequestFromAddressRISE = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateFungibleTokensTransactionRequestFromAddressRISE = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateFungibleTokensTransactionRequestFromAddressRIS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateFungibleTokensTransactionRequestFromAddressRIS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateFungibleTokensTransactionRequestFromAddressRIS) MarshalJSON() ([]byte, error) {
	if src.CreateFungibleTokensTransactionRequestFromAddressRISE != nil {
		return json.Marshal(&src.CreateFungibleTokensTransactionRequestFromAddressRISE)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateFungibleTokensTransactionRequestFromAddressRIS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateFungibleTokensTransactionRequestFromAddressRISE != nil {
		return obj.CreateFungibleTokensTransactionRequestFromAddressRISE
	}

	// all schemas are nil
	return nil
}

type NullableCreateFungibleTokensTransactionRequestFromAddressRIS struct {
	value *CreateFungibleTokensTransactionRequestFromAddressRIS
	isSet bool
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRIS) Get() *CreateFungibleTokensTransactionRequestFromAddressRIS {
	return v.value
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRIS) Set(val *CreateFungibleTokensTransactionRequestFromAddressRIS) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRIS) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRIS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokensTransactionRequestFromAddressRIS(val *CreateFungibleTokensTransactionRequestFromAddressRIS) *NullableCreateFungibleTokensTransactionRequestFromAddressRIS {
	return &NullableCreateFungibleTokensTransactionRequestFromAddressRIS{value: val, isSet: true}
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRIS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRIS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


