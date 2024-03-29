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

// BlockHeightReachedE403 - struct for BlockHeightReachedE403
type BlockHeightReachedE403 struct {
	BannedIpAddress *BannedIpAddress
	BlockchainEventsCallbacksLimitReached *BlockchainEventsCallbacksLimitReached
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsBlockHeightReachedE403 is a convenience function that returns BannedIpAddress wrapped in BlockHeightReachedE403
func BannedIpAddressAsBlockHeightReachedE403(v *BannedIpAddress) BlockHeightReachedE403 {
	return BlockHeightReachedE403{
		BannedIpAddress: v,
	}
}

// BlockchainEventsCallbacksLimitReachedAsBlockHeightReachedE403 is a convenience function that returns BlockchainEventsCallbacksLimitReached wrapped in BlockHeightReachedE403
func BlockchainEventsCallbacksLimitReachedAsBlockHeightReachedE403(v *BlockchainEventsCallbacksLimitReached) BlockHeightReachedE403 {
	return BlockHeightReachedE403{
		BlockchainEventsCallbacksLimitReached: v,
	}
}

// EndpointNotAllowedForApiKeyAsBlockHeightReachedE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in BlockHeightReachedE403
func EndpointNotAllowedForApiKeyAsBlockHeightReachedE403(v *EndpointNotAllowedForApiKey) BlockHeightReachedE403 {
	return BlockHeightReachedE403{
		EndpointNotAllowedForApiKey: v,
	}
}

// EndpointNotAllowedForPlanAsBlockHeightReachedE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in BlockHeightReachedE403
func EndpointNotAllowedForPlanAsBlockHeightReachedE403(v *EndpointNotAllowedForPlan) BlockHeightReachedE403 {
	return BlockHeightReachedE403{
		EndpointNotAllowedForPlan: v,
	}
}

// FeatureMainnetsNotAllowedForPlanAsBlockHeightReachedE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in BlockHeightReachedE403
func FeatureMainnetsNotAllowedForPlanAsBlockHeightReachedE403(v *FeatureMainnetsNotAllowedForPlan) BlockHeightReachedE403 {
	return BlockHeightReachedE403{
		FeatureMainnetsNotAllowedForPlan: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BlockHeightReachedE403) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into BlockchainEventsCallbacksLimitReached
	err = newStrictDecoder(data).Decode(&dst.BlockchainEventsCallbacksLimitReached)
	if err == nil {
		jsonBlockchainEventsCallbacksLimitReached, _ := json.Marshal(dst.BlockchainEventsCallbacksLimitReached)
		if string(jsonBlockchainEventsCallbacksLimitReached) == "{}" { // empty struct
			dst.BlockchainEventsCallbacksLimitReached = nil
		} else {
			match++
		}
	} else {
		dst.BlockchainEventsCallbacksLimitReached = nil
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
		dst.BlockchainEventsCallbacksLimitReached = nil
		dst.EndpointNotAllowedForApiKey = nil
		dst.EndpointNotAllowedForPlan = nil
		dst.FeatureMainnetsNotAllowedForPlan = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(BlockHeightReachedE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(BlockHeightReachedE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BlockHeightReachedE403) MarshalJSON() ([]byte, error) {
	if src.BannedIpAddress != nil {
		return json.Marshal(&src.BannedIpAddress)
	}

	if src.BlockchainEventsCallbacksLimitReached != nil {
		return json.Marshal(&src.BlockchainEventsCallbacksLimitReached)
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
func (obj *BlockHeightReachedE403) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BannedIpAddress != nil {
		return obj.BannedIpAddress
	}

	if obj.BlockchainEventsCallbacksLimitReached != nil {
		return obj.BlockchainEventsCallbacksLimitReached
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

type NullableBlockHeightReachedE403 struct {
	value *BlockHeightReachedE403
	isSet bool
}

func (v NullableBlockHeightReachedE403) Get() *BlockHeightReachedE403 {
	return v.value
}

func (v *NullableBlockHeightReachedE403) Set(val *BlockHeightReachedE403) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeightReachedE403) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeightReachedE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeightReachedE403(val *BlockHeightReachedE403) *NullableBlockHeightReachedE403 {
	return &NullableBlockHeightReachedE403{value: val, isSet: true}
}

func (v NullableBlockHeightReachedE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeightReachedE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


