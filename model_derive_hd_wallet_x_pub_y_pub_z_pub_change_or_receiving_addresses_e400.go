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

// DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 - struct for DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400
type DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 struct {
	InvalidBlockchain *InvalidBlockchain
	InvalidNetwork *InvalidNetwork
	InvalidPagination *InvalidPagination
	LimitGreaterThanAllowed *LimitGreaterThanAllowed
	UriNotFound *UriNotFound
}

// InvalidBlockchainAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 is a convenience function that returns InvalidBlockchain wrapped in DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400
func InvalidBlockchainAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400(v *InvalidBlockchain) DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 {
	return DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400{
		InvalidBlockchain: v,
	}
}

// InvalidNetworkAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 is a convenience function that returns InvalidNetwork wrapped in DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400
func InvalidNetworkAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400(v *InvalidNetwork) DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 {
	return DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400{
		InvalidNetwork: v,
	}
}

// InvalidPaginationAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 is a convenience function that returns InvalidPagination wrapped in DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400
func InvalidPaginationAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400(v *InvalidPagination) DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 {
	return DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400{
		InvalidPagination: v,
	}
}

// LimitGreaterThanAllowedAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 is a convenience function that returns LimitGreaterThanAllowed wrapped in DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400
func LimitGreaterThanAllowedAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400(v *LimitGreaterThanAllowed) DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 {
	return DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400{
		LimitGreaterThanAllowed: v,
	}
}

// UriNotFoundAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 is a convenience function that returns UriNotFound wrapped in DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400
func UriNotFoundAsDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400(v *UriNotFound) DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 {
	return DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400{
		UriNotFound: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InvalidBlockchain
	err = newStrictDecoder(data).Decode(&dst.InvalidBlockchain)
	if err == nil {
		jsonInvalidBlockchain, _ := json.Marshal(dst.InvalidBlockchain)
		if string(jsonInvalidBlockchain) == "{}" { // empty struct
			dst.InvalidBlockchain = nil
		} else {
			match++
		}
	} else {
		dst.InvalidBlockchain = nil
	}

	// try to unmarshal data into InvalidNetwork
	err = newStrictDecoder(data).Decode(&dst.InvalidNetwork)
	if err == nil {
		jsonInvalidNetwork, _ := json.Marshal(dst.InvalidNetwork)
		if string(jsonInvalidNetwork) == "{}" { // empty struct
			dst.InvalidNetwork = nil
		} else {
			match++
		}
	} else {
		dst.InvalidNetwork = nil
	}

	// try to unmarshal data into InvalidPagination
	err = newStrictDecoder(data).Decode(&dst.InvalidPagination)
	if err == nil {
		jsonInvalidPagination, _ := json.Marshal(dst.InvalidPagination)
		if string(jsonInvalidPagination) == "{}" { // empty struct
			dst.InvalidPagination = nil
		} else {
			match++
		}
	} else {
		dst.InvalidPagination = nil
	}

	// try to unmarshal data into LimitGreaterThanAllowed
	err = newStrictDecoder(data).Decode(&dst.LimitGreaterThanAllowed)
	if err == nil {
		jsonLimitGreaterThanAllowed, _ := json.Marshal(dst.LimitGreaterThanAllowed)
		if string(jsonLimitGreaterThanAllowed) == "{}" { // empty struct
			dst.LimitGreaterThanAllowed = nil
		} else {
			match++
		}
	} else {
		dst.LimitGreaterThanAllowed = nil
	}

	// try to unmarshal data into UriNotFound
	err = newStrictDecoder(data).Decode(&dst.UriNotFound)
	if err == nil {
		jsonUriNotFound, _ := json.Marshal(dst.UriNotFound)
		if string(jsonUriNotFound) == "{}" { // empty struct
			dst.UriNotFound = nil
		} else {
			match++
		}
	} else {
		dst.UriNotFound = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidBlockchain = nil
		dst.InvalidNetwork = nil
		dst.InvalidPagination = nil
		dst.LimitGreaterThanAllowed = nil
		dst.UriNotFound = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) MarshalJSON() ([]byte, error) {
	if src.InvalidBlockchain != nil {
		return json.Marshal(&src.InvalidBlockchain)
	}

	if src.InvalidNetwork != nil {
		return json.Marshal(&src.InvalidNetwork)
	}

	if src.InvalidPagination != nil {
		return json.Marshal(&src.InvalidPagination)
	}

	if src.LimitGreaterThanAllowed != nil {
		return json.Marshal(&src.LimitGreaterThanAllowed)
	}

	if src.UriNotFound != nil {
		return json.Marshal(&src.UriNotFound)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvalidBlockchain != nil {
		return obj.InvalidBlockchain
	}

	if obj.InvalidNetwork != nil {
		return obj.InvalidNetwork
	}

	if obj.InvalidPagination != nil {
		return obj.InvalidPagination
	}

	if obj.LimitGreaterThanAllowed != nil {
		return obj.LimitGreaterThanAllowed
	}

	if obj.UriNotFound != nil {
		return obj.UriNotFound
	}

	// all schemas are nil
	return nil
}

type NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 struct {
	value *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400
	isSet bool
}

func (v NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) Get() *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 {
	return v.value
}

func (v *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) Set(val *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) {
	v.value = val
	v.isSet = true
}

func (v NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) IsSet() bool {
	return v.isSet
}

func (v *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400(val *DeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400 {
	return &NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400{value: val, isSet: true}
}

func (v NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeriveHDWalletXPubYPubZPubChangeOrReceivingAddressesE400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


