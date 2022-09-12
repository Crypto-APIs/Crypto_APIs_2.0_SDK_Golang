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

// GetFeeRecommendationsE403 - struct for GetFeeRecommendationsE403
type GetFeeRecommendationsE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsGetFeeRecommendationsE403 is a convenience function that returns BannedIpAddress wrapped in GetFeeRecommendationsE403
func BannedIpAddressAsGetFeeRecommendationsE403(v *BannedIpAddress) GetFeeRecommendationsE403 {
	return GetFeeRecommendationsE403{
		BannedIpAddress: v,
	}
}

// EndpointNotAllowedForApiKeyAsGetFeeRecommendationsE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in GetFeeRecommendationsE403
func EndpointNotAllowedForApiKeyAsGetFeeRecommendationsE403(v *EndpointNotAllowedForApiKey) GetFeeRecommendationsE403 {
	return GetFeeRecommendationsE403{
		EndpointNotAllowedForApiKey: v,
	}
}

// EndpointNotAllowedForPlanAsGetFeeRecommendationsE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in GetFeeRecommendationsE403
func EndpointNotAllowedForPlanAsGetFeeRecommendationsE403(v *EndpointNotAllowedForPlan) GetFeeRecommendationsE403 {
	return GetFeeRecommendationsE403{
		EndpointNotAllowedForPlan: v,
	}
}

// FeatureMainnetsNotAllowedForPlanAsGetFeeRecommendationsE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in GetFeeRecommendationsE403
func FeatureMainnetsNotAllowedForPlanAsGetFeeRecommendationsE403(v *FeatureMainnetsNotAllowedForPlan) GetFeeRecommendationsE403 {
	return GetFeeRecommendationsE403{
		FeatureMainnetsNotAllowedForPlan: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetFeeRecommendationsE403) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(GetFeeRecommendationsE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetFeeRecommendationsE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetFeeRecommendationsE403) MarshalJSON() ([]byte, error) {
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
func (obj *GetFeeRecommendationsE403) GetActualInstance() (interface{}) {
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

type NullableGetFeeRecommendationsE403 struct {
	value *GetFeeRecommendationsE403
	isSet bool
}

func (v NullableGetFeeRecommendationsE403) Get() *GetFeeRecommendationsE403 {
	return v.value
}

func (v *NullableGetFeeRecommendationsE403) Set(val *GetFeeRecommendationsE403) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeeRecommendationsE403) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeeRecommendationsE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeeRecommendationsE403(val *GetFeeRecommendationsE403) *NullableGetFeeRecommendationsE403 {
	return &NullableGetFeeRecommendationsE403{value: val, isSet: true}
}

func (v NullableGetFeeRecommendationsE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFeeRecommendationsE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


