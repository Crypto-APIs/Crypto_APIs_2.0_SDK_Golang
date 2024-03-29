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

// NewConfirmedCoinsTransactionsAndEachConfirmationE403 - struct for NewConfirmedCoinsTransactionsAndEachConfirmationE403
type NewConfirmedCoinsTransactionsAndEachConfirmationE403 struct {
	BannedIpAddress *BannedIpAddress
	BlockchainEventsCallbacksLimitReached *BlockchainEventsCallbacksLimitReached
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsNewConfirmedCoinsTransactionsAndEachConfirmationE403 is a convenience function that returns BannedIpAddress wrapped in NewConfirmedCoinsTransactionsAndEachConfirmationE403
func BannedIpAddressAsNewConfirmedCoinsTransactionsAndEachConfirmationE403(v *BannedIpAddress) NewConfirmedCoinsTransactionsAndEachConfirmationE403 {
	return NewConfirmedCoinsTransactionsAndEachConfirmationE403{
		BannedIpAddress: v,
	}
}

// BlockchainEventsCallbacksLimitReachedAsNewConfirmedCoinsTransactionsAndEachConfirmationE403 is a convenience function that returns BlockchainEventsCallbacksLimitReached wrapped in NewConfirmedCoinsTransactionsAndEachConfirmationE403
func BlockchainEventsCallbacksLimitReachedAsNewConfirmedCoinsTransactionsAndEachConfirmationE403(v *BlockchainEventsCallbacksLimitReached) NewConfirmedCoinsTransactionsAndEachConfirmationE403 {
	return NewConfirmedCoinsTransactionsAndEachConfirmationE403{
		BlockchainEventsCallbacksLimitReached: v,
	}
}

// EndpointNotAllowedForApiKeyAsNewConfirmedCoinsTransactionsAndEachConfirmationE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in NewConfirmedCoinsTransactionsAndEachConfirmationE403
func EndpointNotAllowedForApiKeyAsNewConfirmedCoinsTransactionsAndEachConfirmationE403(v *EndpointNotAllowedForApiKey) NewConfirmedCoinsTransactionsAndEachConfirmationE403 {
	return NewConfirmedCoinsTransactionsAndEachConfirmationE403{
		EndpointNotAllowedForApiKey: v,
	}
}

// EndpointNotAllowedForPlanAsNewConfirmedCoinsTransactionsAndEachConfirmationE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in NewConfirmedCoinsTransactionsAndEachConfirmationE403
func EndpointNotAllowedForPlanAsNewConfirmedCoinsTransactionsAndEachConfirmationE403(v *EndpointNotAllowedForPlan) NewConfirmedCoinsTransactionsAndEachConfirmationE403 {
	return NewConfirmedCoinsTransactionsAndEachConfirmationE403{
		EndpointNotAllowedForPlan: v,
	}
}

// FeatureMainnetsNotAllowedForPlanAsNewConfirmedCoinsTransactionsAndEachConfirmationE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in NewConfirmedCoinsTransactionsAndEachConfirmationE403
func FeatureMainnetsNotAllowedForPlanAsNewConfirmedCoinsTransactionsAndEachConfirmationE403(v *FeatureMainnetsNotAllowedForPlan) NewConfirmedCoinsTransactionsAndEachConfirmationE403 {
	return NewConfirmedCoinsTransactionsAndEachConfirmationE403{
		FeatureMainnetsNotAllowedForPlan: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewConfirmedCoinsTransactionsAndEachConfirmationE403) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(NewConfirmedCoinsTransactionsAndEachConfirmationE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(NewConfirmedCoinsTransactionsAndEachConfirmationE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewConfirmedCoinsTransactionsAndEachConfirmationE403) MarshalJSON() ([]byte, error) {
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
func (obj *NewConfirmedCoinsTransactionsAndEachConfirmationE403) GetActualInstance() (interface{}) {
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

type NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403 struct {
	value *NewConfirmedCoinsTransactionsAndEachConfirmationE403
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403) Get() *NewConfirmedCoinsTransactionsAndEachConfirmationE403 {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403) Set(val *NewConfirmedCoinsTransactionsAndEachConfirmationE403) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsAndEachConfirmationE403(val *NewConfirmedCoinsTransactionsAndEachConfirmationE403) *NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403 {
	return &NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


