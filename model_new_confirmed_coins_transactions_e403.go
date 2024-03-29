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

// NewConfirmedCoinsTransactionsE403 - struct for NewConfirmedCoinsTransactionsE403
type NewConfirmedCoinsTransactionsE403 struct {
	BannedIpAddress *BannedIpAddress
	BlockchainEventsCallbacksLimitReached *BlockchainEventsCallbacksLimitReached
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsNewConfirmedCoinsTransactionsE403 is a convenience function that returns BannedIpAddress wrapped in NewConfirmedCoinsTransactionsE403
func BannedIpAddressAsNewConfirmedCoinsTransactionsE403(v *BannedIpAddress) NewConfirmedCoinsTransactionsE403 {
	return NewConfirmedCoinsTransactionsE403{
		BannedIpAddress: v,
	}
}

// BlockchainEventsCallbacksLimitReachedAsNewConfirmedCoinsTransactionsE403 is a convenience function that returns BlockchainEventsCallbacksLimitReached wrapped in NewConfirmedCoinsTransactionsE403
func BlockchainEventsCallbacksLimitReachedAsNewConfirmedCoinsTransactionsE403(v *BlockchainEventsCallbacksLimitReached) NewConfirmedCoinsTransactionsE403 {
	return NewConfirmedCoinsTransactionsE403{
		BlockchainEventsCallbacksLimitReached: v,
	}
}

// EndpointNotAllowedForApiKeyAsNewConfirmedCoinsTransactionsE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in NewConfirmedCoinsTransactionsE403
func EndpointNotAllowedForApiKeyAsNewConfirmedCoinsTransactionsE403(v *EndpointNotAllowedForApiKey) NewConfirmedCoinsTransactionsE403 {
	return NewConfirmedCoinsTransactionsE403{
		EndpointNotAllowedForApiKey: v,
	}
}

// EndpointNotAllowedForPlanAsNewConfirmedCoinsTransactionsE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in NewConfirmedCoinsTransactionsE403
func EndpointNotAllowedForPlanAsNewConfirmedCoinsTransactionsE403(v *EndpointNotAllowedForPlan) NewConfirmedCoinsTransactionsE403 {
	return NewConfirmedCoinsTransactionsE403{
		EndpointNotAllowedForPlan: v,
	}
}

// FeatureMainnetsNotAllowedForPlanAsNewConfirmedCoinsTransactionsE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in NewConfirmedCoinsTransactionsE403
func FeatureMainnetsNotAllowedForPlanAsNewConfirmedCoinsTransactionsE403(v *FeatureMainnetsNotAllowedForPlan) NewConfirmedCoinsTransactionsE403 {
	return NewConfirmedCoinsTransactionsE403{
		FeatureMainnetsNotAllowedForPlan: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewConfirmedCoinsTransactionsE403) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(NewConfirmedCoinsTransactionsE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(NewConfirmedCoinsTransactionsE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewConfirmedCoinsTransactionsE403) MarshalJSON() ([]byte, error) {
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
func (obj *NewConfirmedCoinsTransactionsE403) GetActualInstance() (interface{}) {
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

type NullableNewConfirmedCoinsTransactionsE403 struct {
	value *NewConfirmedCoinsTransactionsE403
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsE403) Get() *NewConfirmedCoinsTransactionsE403 {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsE403) Set(val *NewConfirmedCoinsTransactionsE403) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsE403) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsE403(val *NewConfirmedCoinsTransactionsE403) *NullableNewConfirmedCoinsTransactionsE403 {
	return &NullableNewConfirmedCoinsTransactionsE403{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


