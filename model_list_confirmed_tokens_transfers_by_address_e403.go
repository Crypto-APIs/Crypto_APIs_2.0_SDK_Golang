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

// ListConfirmedTokensTransfersByAddressE403 - struct for ListConfirmedTokensTransfersByAddressE403
type ListConfirmedTokensTransfersByAddressE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsListConfirmedTokensTransfersByAddressE403 is a convenience function that returns BannedIpAddress wrapped in ListConfirmedTokensTransfersByAddressE403
func BannedIpAddressAsListConfirmedTokensTransfersByAddressE403(v *BannedIpAddress) ListConfirmedTokensTransfersByAddressE403 {
	return ListConfirmedTokensTransfersByAddressE403{
		BannedIpAddress: v,
	}
}

// EndpointNotAllowedForApiKeyAsListConfirmedTokensTransfersByAddressE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in ListConfirmedTokensTransfersByAddressE403
func EndpointNotAllowedForApiKeyAsListConfirmedTokensTransfersByAddressE403(v *EndpointNotAllowedForApiKey) ListConfirmedTokensTransfersByAddressE403 {
	return ListConfirmedTokensTransfersByAddressE403{
		EndpointNotAllowedForApiKey: v,
	}
}

// EndpointNotAllowedForPlanAsListConfirmedTokensTransfersByAddressE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in ListConfirmedTokensTransfersByAddressE403
func EndpointNotAllowedForPlanAsListConfirmedTokensTransfersByAddressE403(v *EndpointNotAllowedForPlan) ListConfirmedTokensTransfersByAddressE403 {
	return ListConfirmedTokensTransfersByAddressE403{
		EndpointNotAllowedForPlan: v,
	}
}

// FeatureMainnetsNotAllowedForPlanAsListConfirmedTokensTransfersByAddressE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in ListConfirmedTokensTransfersByAddressE403
func FeatureMainnetsNotAllowedForPlanAsListConfirmedTokensTransfersByAddressE403(v *FeatureMainnetsNotAllowedForPlan) ListConfirmedTokensTransfersByAddressE403 {
	return ListConfirmedTokensTransfersByAddressE403{
		FeatureMainnetsNotAllowedForPlan: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListConfirmedTokensTransfersByAddressE403) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(ListConfirmedTokensTransfersByAddressE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListConfirmedTokensTransfersByAddressE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListConfirmedTokensTransfersByAddressE403) MarshalJSON() ([]byte, error) {
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
func (obj *ListConfirmedTokensTransfersByAddressE403) GetActualInstance() (interface{}) {
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

type NullableListConfirmedTokensTransfersByAddressE403 struct {
	value *ListConfirmedTokensTransfersByAddressE403
	isSet bool
}

func (v NullableListConfirmedTokensTransfersByAddressE403) Get() *ListConfirmedTokensTransfersByAddressE403 {
	return v.value
}

func (v *NullableListConfirmedTokensTransfersByAddressE403) Set(val *ListConfirmedTokensTransfersByAddressE403) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTokensTransfersByAddressE403) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTokensTransfersByAddressE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTokensTransfersByAddressE403(val *ListConfirmedTokensTransfersByAddressE403) *NullableListConfirmedTokensTransfersByAddressE403 {
	return &NullableListConfirmedTokensTransfersByAddressE403{value: val, isSet: true}
}

func (v NullableListConfirmedTokensTransfersByAddressE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTokensTransfersByAddressE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


