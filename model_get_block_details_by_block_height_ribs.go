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

// GetBlockDetailsByBlockHeightRIBS - struct for GetBlockDetailsByBlockHeightRIBS
type GetBlockDetailsByBlockHeightRIBS struct {
	GetBlockDetailsByBlockHeightRIBSB *GetBlockDetailsByBlockHeightRIBSB
	GetBlockDetailsByBlockHeightRIBSBC *GetBlockDetailsByBlockHeightRIBSBC
	GetBlockDetailsByBlockHeightRIBSBSC *GetBlockDetailsByBlockHeightRIBSBSC
	GetBlockDetailsByBlockHeightRIBSD *GetBlockDetailsByBlockHeightRIBSD
	GetBlockDetailsByBlockHeightRIBSD2 *GetBlockDetailsByBlockHeightRIBSD2
	GetBlockDetailsByBlockHeightRIBSE *GetBlockDetailsByBlockHeightRIBSE
	GetBlockDetailsByBlockHeightRIBSEC *GetBlockDetailsByBlockHeightRIBSEC
	GetBlockDetailsByBlockHeightRIBSL *GetBlockDetailsByBlockHeightRIBSL
	GetBlockDetailsByBlockHeightRIBSZ *GetBlockDetailsByBlockHeightRIBSZ
}

// GetBlockDetailsByBlockHeightRIBSBAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSB wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSBAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSB) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSB: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSBCAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSBC wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSBCAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSBC) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSBC: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSBSCAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSBSC wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSBSCAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSBSC) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSBSC: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSDAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSD wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSDAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSD) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSD: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSD2AsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSD2 wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSD2AsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSD2) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSD2: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSEAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSE wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSEAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSE) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSE: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSECAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSEC wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSECAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSEC) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSEC: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSLAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSL wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSLAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSL) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSL: v,
	}
}

// GetBlockDetailsByBlockHeightRIBSZAsGetBlockDetailsByBlockHeightRIBS is a convenience function that returns GetBlockDetailsByBlockHeightRIBSZ wrapped in GetBlockDetailsByBlockHeightRIBS
func GetBlockDetailsByBlockHeightRIBSZAsGetBlockDetailsByBlockHeightRIBS(v *GetBlockDetailsByBlockHeightRIBSZ) GetBlockDetailsByBlockHeightRIBS {
	return GetBlockDetailsByBlockHeightRIBS{
		GetBlockDetailsByBlockHeightRIBSZ: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetBlockDetailsByBlockHeightRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSB
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSB)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSB, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSB)
		if string(jsonGetBlockDetailsByBlockHeightRIBSB) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSB = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSBC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSBC)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSBC, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSBC)
		if string(jsonGetBlockDetailsByBlockHeightRIBSBC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSBC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSBSC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSBSC)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSBSC, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSBSC)
		if string(jsonGetBlockDetailsByBlockHeightRIBSBSC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSBSC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSBSC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSD
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSD)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSD, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSD)
		if string(jsonGetBlockDetailsByBlockHeightRIBSD) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSD = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSD2
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSD2)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSD2, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSD2)
		if string(jsonGetBlockDetailsByBlockHeightRIBSD2) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSD2 = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSE
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSE)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSE, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSE)
		if string(jsonGetBlockDetailsByBlockHeightRIBSE) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSE = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSE = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSEC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSEC)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSEC, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSEC)
		if string(jsonGetBlockDetailsByBlockHeightRIBSEC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSEC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSL
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSL)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSL, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSL)
		if string(jsonGetBlockDetailsByBlockHeightRIBSL) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSL = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightRIBSZ
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightRIBSZ)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightRIBSZ, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightRIBSZ)
		if string(jsonGetBlockDetailsByBlockHeightRIBSZ) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightRIBSZ = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightRIBSZ = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetBlockDetailsByBlockHeightRIBSB = nil
		dst.GetBlockDetailsByBlockHeightRIBSBC = nil
		dst.GetBlockDetailsByBlockHeightRIBSBSC = nil
		dst.GetBlockDetailsByBlockHeightRIBSD = nil
		dst.GetBlockDetailsByBlockHeightRIBSD2 = nil
		dst.GetBlockDetailsByBlockHeightRIBSE = nil
		dst.GetBlockDetailsByBlockHeightRIBSEC = nil
		dst.GetBlockDetailsByBlockHeightRIBSL = nil
		dst.GetBlockDetailsByBlockHeightRIBSZ = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetBlockDetailsByBlockHeightRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetBlockDetailsByBlockHeightRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetBlockDetailsByBlockHeightRIBS) MarshalJSON() ([]byte, error) {
	if src.GetBlockDetailsByBlockHeightRIBSB != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSB)
	}

	if src.GetBlockDetailsByBlockHeightRIBSBC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSBC)
	}

	if src.GetBlockDetailsByBlockHeightRIBSBSC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSBSC)
	}

	if src.GetBlockDetailsByBlockHeightRIBSD != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSD)
	}

	if src.GetBlockDetailsByBlockHeightRIBSD2 != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSD2)
	}

	if src.GetBlockDetailsByBlockHeightRIBSE != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSE)
	}

	if src.GetBlockDetailsByBlockHeightRIBSEC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSEC)
	}

	if src.GetBlockDetailsByBlockHeightRIBSL != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSL)
	}

	if src.GetBlockDetailsByBlockHeightRIBSZ != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightRIBSZ)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetBlockDetailsByBlockHeightRIBS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetBlockDetailsByBlockHeightRIBSB != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSB
	}

	if obj.GetBlockDetailsByBlockHeightRIBSBC != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSBC
	}

	if obj.GetBlockDetailsByBlockHeightRIBSBSC != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSBSC
	}

	if obj.GetBlockDetailsByBlockHeightRIBSD != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSD
	}

	if obj.GetBlockDetailsByBlockHeightRIBSD2 != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSD2
	}

	if obj.GetBlockDetailsByBlockHeightRIBSE != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSE
	}

	if obj.GetBlockDetailsByBlockHeightRIBSEC != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSEC
	}

	if obj.GetBlockDetailsByBlockHeightRIBSL != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSL
	}

	if obj.GetBlockDetailsByBlockHeightRIBSZ != nil {
		return obj.GetBlockDetailsByBlockHeightRIBSZ
	}

	// all schemas are nil
	return nil
}

type NullableGetBlockDetailsByBlockHeightRIBS struct {
	value *GetBlockDetailsByBlockHeightRIBS
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHeightRIBS) Get() *GetBlockDetailsByBlockHeightRIBS {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHeightRIBS) Set(val *GetBlockDetailsByBlockHeightRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHeightRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHeightRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHeightRIBS(val *GetBlockDetailsByBlockHeightRIBS) *NullableGetBlockDetailsByBlockHeightRIBS {
	return &NullableGetBlockDetailsByBlockHeightRIBS{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHeightRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHeightRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


