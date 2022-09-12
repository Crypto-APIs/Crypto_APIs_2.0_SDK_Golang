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

// EstimateGasLimitE403 - struct for EstimateGasLimitE403
type EstimateGasLimitE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsEstimateGasLimitE403 is a convenience function that returns BannedIpAddress wrapped in EstimateGasLimitE403
func BannedIpAddressAsEstimateGasLimitE403(v *BannedIpAddress) EstimateGasLimitE403 {
	return EstimateGasLimitE403{
		BannedIpAddress: v,
	}
}

// EndpointNotAllowedForApiKeyAsEstimateGasLimitE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in EstimateGasLimitE403
func EndpointNotAllowedForApiKeyAsEstimateGasLimitE403(v *EndpointNotAllowedForApiKey) EstimateGasLimitE403 {
	return EstimateGasLimitE403{
		EndpointNotAllowedForApiKey: v,
	}
}

// EndpointNotAllowedForPlanAsEstimateGasLimitE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in EstimateGasLimitE403
func EndpointNotAllowedForPlanAsEstimateGasLimitE403(v *EndpointNotAllowedForPlan) EstimateGasLimitE403 {
	return EstimateGasLimitE403{
		EndpointNotAllowedForPlan: v,
	}
}

// FeatureMainnetsNotAllowedForPlanAsEstimateGasLimitE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in EstimateGasLimitE403
func FeatureMainnetsNotAllowedForPlanAsEstimateGasLimitE403(v *FeatureMainnetsNotAllowedForPlan) EstimateGasLimitE403 {
	return EstimateGasLimitE403{
		FeatureMainnetsNotAllowedForPlan: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EstimateGasLimitE403) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(EstimateGasLimitE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(EstimateGasLimitE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EstimateGasLimitE403) MarshalJSON() ([]byte, error) {
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
func (obj *EstimateGasLimitE403) GetActualInstance() (interface{}) {
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

type NullableEstimateGasLimitE403 struct {
	value *EstimateGasLimitE403
	isSet bool
}

func (v NullableEstimateGasLimitE403) Get() *EstimateGasLimitE403 {
	return v.value
}

func (v *NullableEstimateGasLimitE403) Set(val *EstimateGasLimitE403) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateGasLimitE403) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateGasLimitE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateGasLimitE403(val *EstimateGasLimitE403) *NullableEstimateGasLimitE403 {
	return &NullableEstimateGasLimitE403{value: val, isSet: true}
}

func (v NullableEstimateGasLimitE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateGasLimitE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


