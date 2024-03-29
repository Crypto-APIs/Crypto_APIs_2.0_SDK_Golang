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

// GetBlockDetailsByBlockHeightFromCallbackRIBS - struct for GetBlockDetailsByBlockHeightFromCallbackRIBS
type GetBlockDetailsByBlockHeightFromCallbackRIBS struct {
	GetBlockDetailsByBlockHeightFromCallbackRIBSB *GetBlockDetailsByBlockHeightFromCallbackRIBSB
	GetBlockDetailsByBlockHeightFromCallbackRIBSBC *GetBlockDetailsByBlockHeightFromCallbackRIBSBC
	GetBlockDetailsByBlockHeightFromCallbackRIBSBSC *GetBlockDetailsByBlockHeightFromCallbackRIBSBSC
	GetBlockDetailsByBlockHeightFromCallbackRIBSD *GetBlockDetailsByBlockHeightFromCallbackRIBSD
	GetBlockDetailsByBlockHeightFromCallbackRIBSD2 *GetBlockDetailsByBlockHeightFromCallbackRIBSD2
	GetBlockDetailsByBlockHeightFromCallbackRIBSE *GetBlockDetailsByBlockHeightFromCallbackRIBSE
	GetBlockDetailsByBlockHeightFromCallbackRIBSEC *GetBlockDetailsByBlockHeightFromCallbackRIBSEC
	GetBlockDetailsByBlockHeightFromCallbackRIBSL *GetBlockDetailsByBlockHeightFromCallbackRIBSL
	GetBlockDetailsByBlockHeightFromCallbackRIBSP *GetBlockDetailsByBlockHeightFromCallbackRIBSP
	GetBlockDetailsByBlockHeightFromCallbackRIBST *GetBlockDetailsByBlockHeightFromCallbackRIBST
	GetBlockDetailsByBlockHeightFromCallbackRIBSX *GetBlockDetailsByBlockHeightFromCallbackRIBSX
	GetBlockDetailsByBlockHeightFromCallbackRIBSZ *GetBlockDetailsByBlockHeightFromCallbackRIBSZ
	GetBlockDetailsByBlockHeightFromCallbackRIBSZ2 *GetBlockDetailsByBlockHeightFromCallbackRIBSZ2
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSBAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSB wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSBAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSB) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSB: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSBCAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSBC wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSBCAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSBC) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSBC: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSBSCAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSBSC wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSBSCAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSBSC) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSBSC: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSDAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSD wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSDAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSD) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSD: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSD2AsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSD2 wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSD2AsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSD2: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSEAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSE wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSEAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSE) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSE: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSECAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSEC wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSECAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSEC) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSEC: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSLAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSL wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSLAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSL) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSL: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSPAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSP wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSPAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSP) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSP: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSTAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBST wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSTAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBST: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSXAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSX wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSXAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSX) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSX: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSZAsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSZ wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSZAsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSZ) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSZ: v,
	}
}

// GetBlockDetailsByBlockHeightFromCallbackRIBSZ2AsGetBlockDetailsByBlockHeightFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHeightFromCallbackRIBSZ2 wrapped in GetBlockDetailsByBlockHeightFromCallbackRIBS
func GetBlockDetailsByBlockHeightFromCallbackRIBSZ2AsGetBlockDetailsByBlockHeightFromCallbackRIBS(v *GetBlockDetailsByBlockHeightFromCallbackRIBSZ2) GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return GetBlockDetailsByBlockHeightFromCallbackRIBS{
		GetBlockDetailsByBlockHeightFromCallbackRIBSZ2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetBlockDetailsByBlockHeightFromCallbackRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSB
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSB)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSB, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSB)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSB) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSB = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSBC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBC)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSBC, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBC)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSBC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSBSC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSBSC, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSBSC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSD
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSD, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSD) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSD2
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD2)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSD2, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD2)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSD2) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD2 = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSE
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSE)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSE, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSE)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSE) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSE = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSE = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSEC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSEC)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSEC, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSEC)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSEC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSEC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSL
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSL)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSL, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSL)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSL) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSL = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSP
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSP)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSP, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSP)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSP) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSP = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSP = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBST
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBST)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBST, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBST)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBST) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBST = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBST = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSX
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSX)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSX, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSX)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSX) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSX = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSX = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSZ
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSZ, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSZ) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHeightFromCallbackRIBSZ2
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2)
	if err == nil {
		jsonGetBlockDetailsByBlockHeightFromCallbackRIBSZ2, _ := json.Marshal(dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2)
		if string(jsonGetBlockDetailsByBlockHeightFromCallbackRIBSZ2) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2 = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSB = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBC = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSD2 = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSE = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSEC = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSL = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSP = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBST = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSX = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ = nil
		dst.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetBlockDetailsByBlockHeightFromCallbackRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetBlockDetailsByBlockHeightFromCallbackRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetBlockDetailsByBlockHeightFromCallbackRIBS) MarshalJSON() ([]byte, error) {
	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSB != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSB)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSBC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSBC)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSD != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSD)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSD2 != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSD2)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSE != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSE)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSEC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSEC)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSL != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSL)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSP != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSP)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBST != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBST)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSX != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSX)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSZ != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSZ)
	}

	if src.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2 != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSB != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSB
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSBC != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSBC
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSBSC
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSD != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSD
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSD2 != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSD2
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSE != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSE
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSEC != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSEC
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSL != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSL
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSP != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSP
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBST != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBST
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSX != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSX
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSZ != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSZ
	}

	if obj.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2 != nil {
		return obj.GetBlockDetailsByBlockHeightFromCallbackRIBSZ2
	}

	// all schemas are nil
	return nil
}

type NullableGetBlockDetailsByBlockHeightFromCallbackRIBS struct {
	value *GetBlockDetailsByBlockHeightFromCallbackRIBS
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHeightFromCallbackRIBS) Get() *GetBlockDetailsByBlockHeightFromCallbackRIBS {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHeightFromCallbackRIBS) Set(val *GetBlockDetailsByBlockHeightFromCallbackRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHeightFromCallbackRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHeightFromCallbackRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHeightFromCallbackRIBS(val *GetBlockDetailsByBlockHeightFromCallbackRIBS) *NullableGetBlockDetailsByBlockHeightFromCallbackRIBS {
	return &NullableGetBlockDetailsByBlockHeightFromCallbackRIBS{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHeightFromCallbackRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHeightFromCallbackRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


