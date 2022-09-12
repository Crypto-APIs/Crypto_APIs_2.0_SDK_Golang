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

// ListTokensTransfersByTransactionHashE403 - struct for ListTokensTransfersByTransactionHashE403
type ListTokensTransfersByTransactionHashE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsListTokensTransfersByTransactionHashE403 is a convenience function that returns BannedIpAddress wrapped in ListTokensTransfersByTransactionHashE403
func BannedIpAddressAsListTokensTransfersByTransactionHashE403(v *BannedIpAddress) ListTokensTransfersByTransactionHashE403 {
	return ListTokensTransfersByTransactionHashE403{
		BannedIpAddress: v,
	}
}

// EndpointNotAllowedForApiKeyAsListTokensTransfersByTransactionHashE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in ListTokensTransfersByTransactionHashE403
func EndpointNotAllowedForApiKeyAsListTokensTransfersByTransactionHashE403(v *EndpointNotAllowedForApiKey) ListTokensTransfersByTransactionHashE403 {
	return ListTokensTransfersByTransactionHashE403{
		EndpointNotAllowedForApiKey: v,
	}
}

// EndpointNotAllowedForPlanAsListTokensTransfersByTransactionHashE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in ListTokensTransfersByTransactionHashE403
func EndpointNotAllowedForPlanAsListTokensTransfersByTransactionHashE403(v *EndpointNotAllowedForPlan) ListTokensTransfersByTransactionHashE403 {
	return ListTokensTransfersByTransactionHashE403{
		EndpointNotAllowedForPlan: v,
	}
}

// FeatureMainnetsNotAllowedForPlanAsListTokensTransfersByTransactionHashE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in ListTokensTransfersByTransactionHashE403
func FeatureMainnetsNotAllowedForPlanAsListTokensTransfersByTransactionHashE403(v *FeatureMainnetsNotAllowedForPlan) ListTokensTransfersByTransactionHashE403 {
	return ListTokensTransfersByTransactionHashE403{
		FeatureMainnetsNotAllowedForPlan: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListTokensTransfersByTransactionHashE403) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BannedIpAddress
	err = newStrictDecoder(data).Decode(&dst.BannedIpAddress)
	if err == nil {
		jsonBannedIpAddress, _ := json.Marshal(dst.BannedIpAddress)
		if string(jsonBannedIpAddress) == "{}" { // empty struct
			dst.BannedIpAddress = nil
		} else {
			match++
		}
	} else {
		dst.BannedIpAddress = nil
	}

	// try to unmarshal data into EndpointNotAllowedForApiKey
	err = newStrictDecoder(data).Decode(&dst.EndpointNotAllowedForApiKey)
	if err == nil {
		jsonEndpointNotAllowedForApiKey, _ := json.Marshal(dst.EndpointNotAllowedForApiKey)
		if string(jsonEndpointNotAllowedForApiKey) == "{}" { // empty struct
			dst.EndpointNotAllowedForApiKey = nil
		} else {
			match++
		}
	} else {
		dst.EndpointNotAllowedForApiKey = nil
	}

	// try to unmarshal data into EndpointNotAllowedForPlan
	err = newStrictDecoder(data).Decode(&dst.EndpointNotAllowedForPlan)
	if err == nil {
		jsonEndpointNotAllowedForPlan, _ := json.Marshal(dst.EndpointNotAllowedForPlan)
		if string(jsonEndpointNotAllowedForPlan) == "{}" { // empty struct
			dst.EndpointNotAllowedForPlan = nil
		} else {
			match++
		}
	} else {
		dst.EndpointNotAllowedForPlan = nil
	}

	// try to unmarshal data into FeatureMainnetsNotAllowedForPlan
	err = newStrictDecoder(data).Decode(&dst.FeatureMainnetsNotAllowedForPlan)
	if err == nil {
		jsonFeatureMainnetsNotAllowedForPlan, _ := json.Marshal(dst.FeatureMainnetsNotAllowedForPlan)
		if string(jsonFeatureMainnetsNotAllowedForPlan) == "{}" { // empty struct
			dst.FeatureMainnetsNotAllowedForPlan = nil
		} else {
			match++
		}
	} else {
		dst.FeatureMainnetsNotAllowedForPlan = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BannedIpAddress = nil
		dst.EndpointNotAllowedForApiKey = nil
		dst.EndpointNotAllowedForPlan = nil
		dst.FeatureMainnetsNotAllowedForPlan = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListTokensTransfersByTransactionHashE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListTokensTransfersByTransactionHashE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListTokensTransfersByTransactionHashE403) MarshalJSON() ([]byte, error) {
	if src.BannedIpAddress != nil {
		return json.Marshal(&src.BannedIpAddress)
	}

	if src.EndpointNotAllowedForApiKey != nil {
		return json.Marshal(&src.EndpointNotAllowedForApiKey)
	}

	if src.EndpointNotAllowedForPlan != nil {
		return json.Marshal(&src.EndpointNotAllowedForPlan)
	}

	if src.FeatureMainnetsNotAllowedForPlan != nil {
		return json.Marshal(&src.FeatureMainnetsNotAllowedForPlan)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListTokensTransfersByTransactionHashE403) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BannedIpAddress != nil {
		return obj.BannedIpAddress
	}

	if obj.EndpointNotAllowedForApiKey != nil {
		return obj.EndpointNotAllowedForApiKey
	}

	if obj.EndpointNotAllowedForPlan != nil {
		return obj.EndpointNotAllowedForPlan
	}

	if obj.FeatureMainnetsNotAllowedForPlan != nil {
		return obj.FeatureMainnetsNotAllowedForPlan
	}

	// all schemas are nil
	return nil
}

type NullableListTokensTransfersByTransactionHashE403 struct {
	value *ListTokensTransfersByTransactionHashE403
	isSet bool
}

func (v NullableListTokensTransfersByTransactionHashE403) Get() *ListTokensTransfersByTransactionHashE403 {
	return v.value
}

func (v *NullableListTokensTransfersByTransactionHashE403) Set(val *ListTokensTransfersByTransactionHashE403) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokensTransfersByTransactionHashE403) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokensTransfersByTransactionHashE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokensTransfersByTransactionHashE403(val *ListTokensTransfersByTransactionHashE403) *NullableListTokensTransfersByTransactionHashE403 {
	return &NullableListTokensTransfersByTransactionHashE403{value: val, isSet: true}
}

func (v NullableListTokensTransfersByTransactionHashE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokensTransfersByTransactionHashE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


