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

// SyncNewHDWalletXPubYPubZPubE400 - struct for SyncNewHDWalletXPubYPubZPubE400
type SyncNewHDWalletXPubYPubZPubE400 struct {
	InvalidBlockchain *InvalidBlockchain
	InvalidNetwork *InvalidNetwork
	InvalidPagination *InvalidPagination
	InvalidXpub *InvalidXpub
	LimitGreaterThanAllowed *LimitGreaterThanAllowed
	UriNotFound *UriNotFound
	XpubNotSynced *XpubNotSynced
}

// InvalidBlockchainAsSyncNewHDWalletXPubYPubZPubE400 is a convenience function that returns InvalidBlockchain wrapped in SyncNewHDWalletXPubYPubZPubE400
func InvalidBlockchainAsSyncNewHDWalletXPubYPubZPubE400(v *InvalidBlockchain) SyncNewHDWalletXPubYPubZPubE400 {
	return SyncNewHDWalletXPubYPubZPubE400{
		InvalidBlockchain: v,
	}
}

// InvalidNetworkAsSyncNewHDWalletXPubYPubZPubE400 is a convenience function that returns InvalidNetwork wrapped in SyncNewHDWalletXPubYPubZPubE400
func InvalidNetworkAsSyncNewHDWalletXPubYPubZPubE400(v *InvalidNetwork) SyncNewHDWalletXPubYPubZPubE400 {
	return SyncNewHDWalletXPubYPubZPubE400{
		InvalidNetwork: v,
	}
}

// InvalidPaginationAsSyncNewHDWalletXPubYPubZPubE400 is a convenience function that returns InvalidPagination wrapped in SyncNewHDWalletXPubYPubZPubE400
func InvalidPaginationAsSyncNewHDWalletXPubYPubZPubE400(v *InvalidPagination) SyncNewHDWalletXPubYPubZPubE400 {
	return SyncNewHDWalletXPubYPubZPubE400{
		InvalidPagination: v,
	}
}

// InvalidXpubAsSyncNewHDWalletXPubYPubZPubE400 is a convenience function that returns InvalidXpub wrapped in SyncNewHDWalletXPubYPubZPubE400
func InvalidXpubAsSyncNewHDWalletXPubYPubZPubE400(v *InvalidXpub) SyncNewHDWalletXPubYPubZPubE400 {
	return SyncNewHDWalletXPubYPubZPubE400{
		InvalidXpub: v,
	}
}

// LimitGreaterThanAllowedAsSyncNewHDWalletXPubYPubZPubE400 is a convenience function that returns LimitGreaterThanAllowed wrapped in SyncNewHDWalletXPubYPubZPubE400
func LimitGreaterThanAllowedAsSyncNewHDWalletXPubYPubZPubE400(v *LimitGreaterThanAllowed) SyncNewHDWalletXPubYPubZPubE400 {
	return SyncNewHDWalletXPubYPubZPubE400{
		LimitGreaterThanAllowed: v,
	}
}

// UriNotFoundAsSyncNewHDWalletXPubYPubZPubE400 is a convenience function that returns UriNotFound wrapped in SyncNewHDWalletXPubYPubZPubE400
func UriNotFoundAsSyncNewHDWalletXPubYPubZPubE400(v *UriNotFound) SyncNewHDWalletXPubYPubZPubE400 {
	return SyncNewHDWalletXPubYPubZPubE400{
		UriNotFound: v,
	}
}

// XpubNotSyncedAsSyncNewHDWalletXPubYPubZPubE400 is a convenience function that returns XpubNotSynced wrapped in SyncNewHDWalletXPubYPubZPubE400
func XpubNotSyncedAsSyncNewHDWalletXPubYPubZPubE400(v *XpubNotSynced) SyncNewHDWalletXPubYPubZPubE400 {
	return SyncNewHDWalletXPubYPubZPubE400{
		XpubNotSynced: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SyncNewHDWalletXPubYPubZPubE400) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into InvalidXpub
	err = newStrictDecoder(data).Decode(&dst.InvalidXpub)
	if err == nil {
		jsonInvalidXpub, _ := json.Marshal(dst.InvalidXpub)
		if string(jsonInvalidXpub) == "{}" { // empty struct
			dst.InvalidXpub = nil
		} else {
			match++
		}
	} else {
		dst.InvalidXpub = nil
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

	// try to unmarshal data into XpubNotSynced
	err = newStrictDecoder(data).Decode(&dst.XpubNotSynced)
	if err == nil {
		jsonXpubNotSynced, _ := json.Marshal(dst.XpubNotSynced)
		if string(jsonXpubNotSynced) == "{}" { // empty struct
			dst.XpubNotSynced = nil
		} else {
			match++
		}
	} else {
		dst.XpubNotSynced = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidBlockchain = nil
		dst.InvalidNetwork = nil
		dst.InvalidPagination = nil
		dst.InvalidXpub = nil
		dst.LimitGreaterThanAllowed = nil
		dst.UriNotFound = nil
		dst.XpubNotSynced = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SyncNewHDWalletXPubYPubZPubE400)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SyncNewHDWalletXPubYPubZPubE400)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SyncNewHDWalletXPubYPubZPubE400) MarshalJSON() ([]byte, error) {
	if src.InvalidBlockchain != nil {
		return json.Marshal(&src.InvalidBlockchain)
	}

	if src.InvalidNetwork != nil {
		return json.Marshal(&src.InvalidNetwork)
	}

	if src.InvalidPagination != nil {
		return json.Marshal(&src.InvalidPagination)
	}

	if src.InvalidXpub != nil {
		return json.Marshal(&src.InvalidXpub)
	}

	if src.LimitGreaterThanAllowed != nil {
		return json.Marshal(&src.LimitGreaterThanAllowed)
	}

	if src.UriNotFound != nil {
		return json.Marshal(&src.UriNotFound)
	}

	if src.XpubNotSynced != nil {
		return json.Marshal(&src.XpubNotSynced)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SyncNewHDWalletXPubYPubZPubE400) GetActualInstance() (interface{}) {
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

	if obj.InvalidXpub != nil {
		return obj.InvalidXpub
	}

	if obj.LimitGreaterThanAllowed != nil {
		return obj.LimitGreaterThanAllowed
	}

	if obj.UriNotFound != nil {
		return obj.UriNotFound
	}

	if obj.XpubNotSynced != nil {
		return obj.XpubNotSynced
	}

	// all schemas are nil
	return nil
}

type NullableSyncNewHDWalletXPubYPubZPubE400 struct {
	value *SyncNewHDWalletXPubYPubZPubE400
	isSet bool
}

func (v NullableSyncNewHDWalletXPubYPubZPubE400) Get() *SyncNewHDWalletXPubYPubZPubE400 {
	return v.value
}

func (v *NullableSyncNewHDWalletXPubYPubZPubE400) Set(val *SyncNewHDWalletXPubYPubZPubE400) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncNewHDWalletXPubYPubZPubE400) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncNewHDWalletXPubYPubZPubE400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncNewHDWalletXPubYPubZPubE400(val *SyncNewHDWalletXPubYPubZPubE400) *NullableSyncNewHDWalletXPubYPubZPubE400 {
	return &NullableSyncNewHDWalletXPubYPubZPubE400{value: val, isSet: true}
}

func (v NullableSyncNewHDWalletXPubYPubZPubE400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncNewHDWalletXPubYPubZPubE400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

