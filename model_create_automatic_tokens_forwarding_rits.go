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

// CreateAutomaticTokensForwardingRITS - struct for CreateAutomaticTokensForwardingRITS
type CreateAutomaticTokensForwardingRITS struct {
	CreateAutomaticTokensForwardingRITSBOT *CreateAutomaticTokensForwardingRITSBOT
	CreateAutomaticTokensForwardingRITSET *CreateAutomaticTokensForwardingRITSET
}

// CreateAutomaticTokensForwardingRITSBOTAsCreateAutomaticTokensForwardingRITS is a convenience function that returns CreateAutomaticTokensForwardingRITSBOT wrapped in CreateAutomaticTokensForwardingRITS
func CreateAutomaticTokensForwardingRITSBOTAsCreateAutomaticTokensForwardingRITS(v *CreateAutomaticTokensForwardingRITSBOT) CreateAutomaticTokensForwardingRITS {
	return CreateAutomaticTokensForwardingRITS{
		CreateAutomaticTokensForwardingRITSBOT: v,
	}
}

// CreateAutomaticTokensForwardingRITSETAsCreateAutomaticTokensForwardingRITS is a convenience function that returns CreateAutomaticTokensForwardingRITSET wrapped in CreateAutomaticTokensForwardingRITS
func CreateAutomaticTokensForwardingRITSETAsCreateAutomaticTokensForwardingRITS(v *CreateAutomaticTokensForwardingRITSET) CreateAutomaticTokensForwardingRITS {
	return CreateAutomaticTokensForwardingRITS{
		CreateAutomaticTokensForwardingRITSET: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateAutomaticTokensForwardingRITS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateAutomaticTokensForwardingRITSBOT
	err = newStrictDecoder(data).Decode(&dst.CreateAutomaticTokensForwardingRITSBOT)
	if err == nil {
		jsonCreateAutomaticTokensForwardingRITSBOT, _ := json.Marshal(dst.CreateAutomaticTokensForwardingRITSBOT)
		if string(jsonCreateAutomaticTokensForwardingRITSBOT) == "{}" { // empty struct
			dst.CreateAutomaticTokensForwardingRITSBOT = nil
		} else {
			match++
		}
	} else {
		dst.CreateAutomaticTokensForwardingRITSBOT = nil
	}

	// try to unmarshal data into CreateAutomaticTokensForwardingRITSET
	err = newStrictDecoder(data).Decode(&dst.CreateAutomaticTokensForwardingRITSET)
	if err == nil {
		jsonCreateAutomaticTokensForwardingRITSET, _ := json.Marshal(dst.CreateAutomaticTokensForwardingRITSET)
		if string(jsonCreateAutomaticTokensForwardingRITSET) == "{}" { // empty struct
			dst.CreateAutomaticTokensForwardingRITSET = nil
		} else {
			match++
		}
	} else {
		dst.CreateAutomaticTokensForwardingRITSET = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateAutomaticTokensForwardingRITSBOT = nil
		dst.CreateAutomaticTokensForwardingRITSET = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateAutomaticTokensForwardingRITS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateAutomaticTokensForwardingRITS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateAutomaticTokensForwardingRITS) MarshalJSON() ([]byte, error) {
	if src.CreateAutomaticTokensForwardingRITSBOT != nil {
		return json.Marshal(&src.CreateAutomaticTokensForwardingRITSBOT)
	}

	if src.CreateAutomaticTokensForwardingRITSET != nil {
		return json.Marshal(&src.CreateAutomaticTokensForwardingRITSET)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateAutomaticTokensForwardingRITS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateAutomaticTokensForwardingRITSBOT != nil {
		return obj.CreateAutomaticTokensForwardingRITSBOT
	}

	if obj.CreateAutomaticTokensForwardingRITSET != nil {
		return obj.CreateAutomaticTokensForwardingRITSET
	}

	// all schemas are nil
	return nil
}

type NullableCreateAutomaticTokensForwardingRITS struct {
	value *CreateAutomaticTokensForwardingRITS
	isSet bool
}

func (v NullableCreateAutomaticTokensForwardingRITS) Get() *CreateAutomaticTokensForwardingRITS {
	return v.value
}

func (v *NullableCreateAutomaticTokensForwardingRITS) Set(val *CreateAutomaticTokensForwardingRITS) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticTokensForwardingRITS) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticTokensForwardingRITS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticTokensForwardingRITS(val *CreateAutomaticTokensForwardingRITS) *NullableCreateAutomaticTokensForwardingRITS {
	return &NullableCreateAutomaticTokensForwardingRITS{value: val, isSet: true}
}

func (v NullableCreateAutomaticTokensForwardingRITS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticTokensForwardingRITS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


