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

// ListSupportedTokensE403 - struct for ListSupportedTokensE403
type ListSupportedTokensE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsListSupportedTokensE403 is a convenience function that returns BannedIpAddress wrapped in ListSupportedTokensE403
func BannedIpAddressAsListSupportedTokensE403(v *BannedIpAddress) ListSupportedTokensE403 {
	return ListSupportedTokensE403{ BannedIpAddress: v}
}

// EndpointNotAllowedForApiKeyAsListSupportedTokensE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in ListSupportedTokensE403
func EndpointNotAllowedForApiKeyAsListSupportedTokensE403(v *EndpointNotAllowedForApiKey) ListSupportedTokensE403 {
	return ListSupportedTokensE403{ EndpointNotAllowedForApiKey: v}
}

// EndpointNotAllowedForPlanAsListSupportedTokensE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in ListSupportedTokensE403
func EndpointNotAllowedForPlanAsListSupportedTokensE403(v *EndpointNotAllowedForPlan) ListSupportedTokensE403 {
	return ListSupportedTokensE403{ EndpointNotAllowedForPlan: v}
}

// FeatureMainnetsNotAllowedForPlanAsListSupportedTokensE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in ListSupportedTokensE403
func FeatureMainnetsNotAllowedForPlanAsListSupportedTokensE403(v *FeatureMainnetsNotAllowedForPlan) ListSupportedTokensE403 {
	return ListSupportedTokensE403{ FeatureMainnetsNotAllowedForPlan: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListSupportedTokensE403) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BannedIpAddress
	err = json.Unmarshal(data, &dst.BannedIpAddress)
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
	err = json.Unmarshal(data, &dst.EndpointNotAllowedForApiKey)
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
	err = json.Unmarshal(data, &dst.EndpointNotAllowedForPlan)
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
	err = json.Unmarshal(data, &dst.FeatureMainnetsNotAllowedForPlan)
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

		return fmt.Errorf("Data matches more than one schema in oneOf(ListSupportedTokensE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListSupportedTokensE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListSupportedTokensE403) MarshalJSON() ([]byte, error) {
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
func (obj *ListSupportedTokensE403) GetActualInstance() (interface{}) {
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

type NullableListSupportedTokensE403 struct {
	value *ListSupportedTokensE403
	isSet bool
}

func (v NullableListSupportedTokensE403) Get() *ListSupportedTokensE403 {
	return v.value
}

func (v *NullableListSupportedTokensE403) Set(val *ListSupportedTokensE403) {
	v.value = val
	v.isSet = true
}

func (v NullableListSupportedTokensE403) IsSet() bool {
	return v.isSet
}

func (v *NullableListSupportedTokensE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSupportedTokensE403(val *ListSupportedTokensE403) *NullableListSupportedTokensE403 {
	return &NullableListSupportedTokensE403{value: val, isSet: true}
}

func (v NullableListSupportedTokensE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSupportedTokensE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

