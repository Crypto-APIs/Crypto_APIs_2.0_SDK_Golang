/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
	"fmt"
)

// GetBlockDetailsByBlockHashRIBS - struct for GetBlockDetailsByBlockHashRIBS
type GetBlockDetailsByBlockHashRIBS struct {
	GetBlockDetailsByBlockHashRIBSB *GetBlockDetailsByBlockHashRIBSB
	GetBlockDetailsByBlockHashRIBSBC *GetBlockDetailsByBlockHashRIBSBC
	GetBlockDetailsByBlockHashRIBSD *GetBlockDetailsByBlockHashRIBSD
	GetBlockDetailsByBlockHashRIBSD2 *GetBlockDetailsByBlockHashRIBSD2
	GetBlockDetailsByBlockHashRIBSE *GetBlockDetailsByBlockHashRIBSE
	GetBlockDetailsByBlockHashRIBSEC *GetBlockDetailsByBlockHashRIBSEC
	GetBlockDetailsByBlockHashRIBSL *GetBlockDetailsByBlockHashRIBSL
}

// GetBlockDetailsByBlockHashRIBSBAsGetBlockDetailsByBlockHashRIBS is a convenience function that returns GetBlockDetailsByBlockHashRIBSB wrapped in GetBlockDetailsByBlockHashRIBS
func GetBlockDetailsByBlockHashRIBSBAsGetBlockDetailsByBlockHashRIBS(v *GetBlockDetailsByBlockHashRIBSB) GetBlockDetailsByBlockHashRIBS {
	return GetBlockDetailsByBlockHashRIBS{ GetBlockDetailsByBlockHashRIBSB: v}
}

// GetBlockDetailsByBlockHashRIBSBCAsGetBlockDetailsByBlockHashRIBS is a convenience function that returns GetBlockDetailsByBlockHashRIBSBC wrapped in GetBlockDetailsByBlockHashRIBS
func GetBlockDetailsByBlockHashRIBSBCAsGetBlockDetailsByBlockHashRIBS(v *GetBlockDetailsByBlockHashRIBSBC) GetBlockDetailsByBlockHashRIBS {
	return GetBlockDetailsByBlockHashRIBS{ GetBlockDetailsByBlockHashRIBSBC: v}
}

// GetBlockDetailsByBlockHashRIBSDAsGetBlockDetailsByBlockHashRIBS is a convenience function that returns GetBlockDetailsByBlockHashRIBSD wrapped in GetBlockDetailsByBlockHashRIBS
func GetBlockDetailsByBlockHashRIBSDAsGetBlockDetailsByBlockHashRIBS(v *GetBlockDetailsByBlockHashRIBSD) GetBlockDetailsByBlockHashRIBS {
	return GetBlockDetailsByBlockHashRIBS{ GetBlockDetailsByBlockHashRIBSD: v}
}

// GetBlockDetailsByBlockHashRIBSD2AsGetBlockDetailsByBlockHashRIBS is a convenience function that returns GetBlockDetailsByBlockHashRIBSD2 wrapped in GetBlockDetailsByBlockHashRIBS
func GetBlockDetailsByBlockHashRIBSD2AsGetBlockDetailsByBlockHashRIBS(v *GetBlockDetailsByBlockHashRIBSD2) GetBlockDetailsByBlockHashRIBS {
	return GetBlockDetailsByBlockHashRIBS{ GetBlockDetailsByBlockHashRIBSD2: v}
}

// GetBlockDetailsByBlockHashRIBSEAsGetBlockDetailsByBlockHashRIBS is a convenience function that returns GetBlockDetailsByBlockHashRIBSE wrapped in GetBlockDetailsByBlockHashRIBS
func GetBlockDetailsByBlockHashRIBSEAsGetBlockDetailsByBlockHashRIBS(v *GetBlockDetailsByBlockHashRIBSE) GetBlockDetailsByBlockHashRIBS {
	return GetBlockDetailsByBlockHashRIBS{ GetBlockDetailsByBlockHashRIBSE: v}
}

// GetBlockDetailsByBlockHashRIBSECAsGetBlockDetailsByBlockHashRIBS is a convenience function that returns GetBlockDetailsByBlockHashRIBSEC wrapped in GetBlockDetailsByBlockHashRIBS
func GetBlockDetailsByBlockHashRIBSECAsGetBlockDetailsByBlockHashRIBS(v *GetBlockDetailsByBlockHashRIBSEC) GetBlockDetailsByBlockHashRIBS {
	return GetBlockDetailsByBlockHashRIBS{ GetBlockDetailsByBlockHashRIBSEC: v}
}

// GetBlockDetailsByBlockHashRIBSLAsGetBlockDetailsByBlockHashRIBS is a convenience function that returns GetBlockDetailsByBlockHashRIBSL wrapped in GetBlockDetailsByBlockHashRIBS
func GetBlockDetailsByBlockHashRIBSLAsGetBlockDetailsByBlockHashRIBS(v *GetBlockDetailsByBlockHashRIBSL) GetBlockDetailsByBlockHashRIBS {
	return GetBlockDetailsByBlockHashRIBS{ GetBlockDetailsByBlockHashRIBSL: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetBlockDetailsByBlockHashRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetBlockDetailsByBlockHashRIBSB
	err = json.Unmarshal(data, &dst.GetBlockDetailsByBlockHashRIBSB)
	if err == nil {
		jsonGetBlockDetailsByBlockHashRIBSB, _ := json.Marshal(dst.GetBlockDetailsByBlockHashRIBSB)
		if string(jsonGetBlockDetailsByBlockHashRIBSB) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashRIBSB = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashRIBSBC
	err = json.Unmarshal(data, &dst.GetBlockDetailsByBlockHashRIBSBC)
	if err == nil {
		jsonGetBlockDetailsByBlockHashRIBSBC, _ := json.Marshal(dst.GetBlockDetailsByBlockHashRIBSBC)
		if string(jsonGetBlockDetailsByBlockHashRIBSBC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashRIBSBC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashRIBSD
	err = json.Unmarshal(data, &dst.GetBlockDetailsByBlockHashRIBSD)
	if err == nil {
		jsonGetBlockDetailsByBlockHashRIBSD, _ := json.Marshal(dst.GetBlockDetailsByBlockHashRIBSD)
		if string(jsonGetBlockDetailsByBlockHashRIBSD) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashRIBSD = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashRIBSD2
	err = json.Unmarshal(data, &dst.GetBlockDetailsByBlockHashRIBSD2)
	if err == nil {
		jsonGetBlockDetailsByBlockHashRIBSD2, _ := json.Marshal(dst.GetBlockDetailsByBlockHashRIBSD2)
		if string(jsonGetBlockDetailsByBlockHashRIBSD2) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashRIBSD2 = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashRIBSE
	err = json.Unmarshal(data, &dst.GetBlockDetailsByBlockHashRIBSE)
	if err == nil {
		jsonGetBlockDetailsByBlockHashRIBSE, _ := json.Marshal(dst.GetBlockDetailsByBlockHashRIBSE)
		if string(jsonGetBlockDetailsByBlockHashRIBSE) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashRIBSE = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashRIBSE = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashRIBSEC
	err = json.Unmarshal(data, &dst.GetBlockDetailsByBlockHashRIBSEC)
	if err == nil {
		jsonGetBlockDetailsByBlockHashRIBSEC, _ := json.Marshal(dst.GetBlockDetailsByBlockHashRIBSEC)
		if string(jsonGetBlockDetailsByBlockHashRIBSEC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashRIBSEC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashRIBSL
	err = json.Unmarshal(data, &dst.GetBlockDetailsByBlockHashRIBSL)
	if err == nil {
		jsonGetBlockDetailsByBlockHashRIBSL, _ := json.Marshal(dst.GetBlockDetailsByBlockHashRIBSL)
		if string(jsonGetBlockDetailsByBlockHashRIBSL) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashRIBSL = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetBlockDetailsByBlockHashRIBSB = nil
		dst.GetBlockDetailsByBlockHashRIBSBC = nil
		dst.GetBlockDetailsByBlockHashRIBSD = nil
		dst.GetBlockDetailsByBlockHashRIBSD2 = nil
		dst.GetBlockDetailsByBlockHashRIBSE = nil
		dst.GetBlockDetailsByBlockHashRIBSEC = nil
		dst.GetBlockDetailsByBlockHashRIBSL = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetBlockDetailsByBlockHashRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetBlockDetailsByBlockHashRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetBlockDetailsByBlockHashRIBS) MarshalJSON() ([]byte, error) {
	if src.GetBlockDetailsByBlockHashRIBSB != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashRIBSB)
	}

	if src.GetBlockDetailsByBlockHashRIBSBC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashRIBSBC)
	}

	if src.GetBlockDetailsByBlockHashRIBSD != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashRIBSD)
	}

	if src.GetBlockDetailsByBlockHashRIBSD2 != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashRIBSD2)
	}

	if src.GetBlockDetailsByBlockHashRIBSE != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashRIBSE)
	}

	if src.GetBlockDetailsByBlockHashRIBSEC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashRIBSEC)
	}

	if src.GetBlockDetailsByBlockHashRIBSL != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashRIBSL)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetBlockDetailsByBlockHashRIBS) GetActualInstance() (interface{}) {
	if obj.GetBlockDetailsByBlockHashRIBSB != nil {
		return obj.GetBlockDetailsByBlockHashRIBSB
	}

	if obj.GetBlockDetailsByBlockHashRIBSBC != nil {
		return obj.GetBlockDetailsByBlockHashRIBSBC
	}

	if obj.GetBlockDetailsByBlockHashRIBSD != nil {
		return obj.GetBlockDetailsByBlockHashRIBSD
	}

	if obj.GetBlockDetailsByBlockHashRIBSD2 != nil {
		return obj.GetBlockDetailsByBlockHashRIBSD2
	}

	if obj.GetBlockDetailsByBlockHashRIBSE != nil {
		return obj.GetBlockDetailsByBlockHashRIBSE
	}

	if obj.GetBlockDetailsByBlockHashRIBSEC != nil {
		return obj.GetBlockDetailsByBlockHashRIBSEC
	}

	if obj.GetBlockDetailsByBlockHashRIBSL != nil {
		return obj.GetBlockDetailsByBlockHashRIBSL
	}

	// all schemas are nil
	return nil
}

type NullableGetBlockDetailsByBlockHashRIBS struct {
	value *GetBlockDetailsByBlockHashRIBS
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashRIBS) Get() *GetBlockDetailsByBlockHashRIBS {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashRIBS) Set(val *GetBlockDetailsByBlockHashRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashRIBS(val *GetBlockDetailsByBlockHashRIBS) *NullableGetBlockDetailsByBlockHashRIBS {
	return &NullableGetBlockDetailsByBlockHashRIBS{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

