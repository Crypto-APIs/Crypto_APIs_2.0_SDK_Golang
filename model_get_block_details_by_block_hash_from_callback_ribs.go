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

// GetBlockDetailsByBlockHashFromCallbackRIBS - struct for GetBlockDetailsByBlockHashFromCallbackRIBS
type GetBlockDetailsByBlockHashFromCallbackRIBS struct {
	GetBlockDetailsByBlockHashFromCallbackRIBSB *GetBlockDetailsByBlockHashFromCallbackRIBSB
	GetBlockDetailsByBlockHashFromCallbackRIBSBC *GetBlockDetailsByBlockHashFromCallbackRIBSBC
	GetBlockDetailsByBlockHashFromCallbackRIBSBSC *GetBlockDetailsByBlockHashFromCallbackRIBSBSC
	GetBlockDetailsByBlockHashFromCallbackRIBSD *GetBlockDetailsByBlockHashFromCallbackRIBSD
	GetBlockDetailsByBlockHashFromCallbackRIBSD2 *GetBlockDetailsByBlockHashFromCallbackRIBSD2
	GetBlockDetailsByBlockHashFromCallbackRIBSE *GetBlockDetailsByBlockHashFromCallbackRIBSE
	GetBlockDetailsByBlockHashFromCallbackRIBSEC *GetBlockDetailsByBlockHashFromCallbackRIBSEC
	GetBlockDetailsByBlockHashFromCallbackRIBSL *GetBlockDetailsByBlockHashFromCallbackRIBSL
	GetBlockDetailsByBlockHashFromCallbackRIBST *GetBlockDetailsByBlockHashFromCallbackRIBST
	GetBlockDetailsByBlockHashFromCallbackRIBSX *GetBlockDetailsByBlockHashFromCallbackRIBSX
	GetBlockDetailsByBlockHashFromCallbackRIBSZ *GetBlockDetailsByBlockHashFromCallbackRIBSZ
	GetBlockDetailsByBlockHashFromCallbackRIBSZ2 *GetBlockDetailsByBlockHashFromCallbackRIBSZ2
}

// GetBlockDetailsByBlockHashFromCallbackRIBSBAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSB wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSBAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSB) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSB: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSBCAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSBC wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSBCAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSBC) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSBC: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSBSCAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSBSC wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSBSCAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSBSC) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSBSC: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSDAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSD wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSDAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSD) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSD: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSD2AsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSD2 wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSD2AsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSD2) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSD2: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSEAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSE wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSEAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSE) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSE: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSECAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSEC wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSECAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSEC) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSEC: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSLAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSL wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSLAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSL: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSTAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBST wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSTAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBST) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBST: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSXAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSX wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSXAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSX) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSX: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSZAsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSZ wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSZAsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSZ: v,
	}
}

// GetBlockDetailsByBlockHashFromCallbackRIBSZ2AsGetBlockDetailsByBlockHashFromCallbackRIBS is a convenience function that returns GetBlockDetailsByBlockHashFromCallbackRIBSZ2 wrapped in GetBlockDetailsByBlockHashFromCallbackRIBS
func GetBlockDetailsByBlockHashFromCallbackRIBSZ2AsGetBlockDetailsByBlockHashFromCallbackRIBS(v *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetBlockDetailsByBlockHashFromCallbackRIBS {
	return GetBlockDetailsByBlockHashFromCallbackRIBS{
		GetBlockDetailsByBlockHashFromCallbackRIBSZ2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetBlockDetailsByBlockHashFromCallbackRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSB
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSB)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSB, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSB)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSB) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSB = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSBC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSBC)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSBC, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSBC)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSBC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSBC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSBSC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSBSC)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSBSC, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSBSC)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSBSC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSBSC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSBSC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSD
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSD)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSD, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSD)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSD) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSD = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSD2
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSD2)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSD2, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSD2)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSD2) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSD2 = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSE
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSE)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSE, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSE)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSE) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSE = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSE = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSEC
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSEC)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSEC, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSEC)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSEC) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSEC = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSL
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSL)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSL, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSL)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSL) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSL = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBST
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBST)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBST, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBST)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBST) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBST = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBST = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSX
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSX)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSX, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSX)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSX) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSX = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSX = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSZ
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSZ, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSZ) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ = nil
	}

	// try to unmarshal data into GetBlockDetailsByBlockHashFromCallbackRIBSZ2
	err = newStrictDecoder(data).Decode(&dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ2)
	if err == nil {
		jsonGetBlockDetailsByBlockHashFromCallbackRIBSZ2, _ := json.Marshal(dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ2)
		if string(jsonGetBlockDetailsByBlockHashFromCallbackRIBSZ2) == "{}" { // empty struct
			dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ2 = nil
		} else {
			match++
		}
	} else {
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSB = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSBC = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSBSC = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSD = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSD2 = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSE = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSEC = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSL = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBST = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSX = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ = nil
		dst.GetBlockDetailsByBlockHashFromCallbackRIBSZ2 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetBlockDetailsByBlockHashFromCallbackRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetBlockDetailsByBlockHashFromCallbackRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetBlockDetailsByBlockHashFromCallbackRIBS) MarshalJSON() ([]byte, error) {
	if src.GetBlockDetailsByBlockHashFromCallbackRIBSB != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSB)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSBC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSBC)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSBSC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSBSC)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSD != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSD)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSD2 != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSD2)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSE != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSE)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSEC != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSEC)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSL != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSL)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBST != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBST)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSX != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSX)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSZ != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSZ)
	}

	if src.GetBlockDetailsByBlockHashFromCallbackRIBSZ2 != nil {
		return json.Marshal(&src.GetBlockDetailsByBlockHashFromCallbackRIBSZ2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetBlockDetailsByBlockHashFromCallbackRIBS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSB != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSB
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSBC != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSBC
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSBSC != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSBSC
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSD != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSD
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSD2 != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSD2
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSE != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSE
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSEC != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSEC
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSL != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSL
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBST != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBST
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSX != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSX
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSZ != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSZ
	}

	if obj.GetBlockDetailsByBlockHashFromCallbackRIBSZ2 != nil {
		return obj.GetBlockDetailsByBlockHashFromCallbackRIBSZ2
	}

	// all schemas are nil
	return nil
}

type NullableGetBlockDetailsByBlockHashFromCallbackRIBS struct {
	value *GetBlockDetailsByBlockHashFromCallbackRIBS
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBS) Get() *GetBlockDetailsByBlockHashFromCallbackRIBS {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBS) Set(val *GetBlockDetailsByBlockHashFromCallbackRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashFromCallbackRIBS(val *GetBlockDetailsByBlockHashFromCallbackRIBS) *NullableGetBlockDetailsByBlockHashFromCallbackRIBS {
	return &NullableGetBlockDetailsByBlockHashFromCallbackRIBS{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


