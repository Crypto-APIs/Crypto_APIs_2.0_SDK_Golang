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

// ListLatestMinedBlocksRIBS - struct for ListLatestMinedBlocksRIBS
type ListLatestMinedBlocksRIBS struct {
	ListLatestMinedBlocksRIBSB *ListLatestMinedBlocksRIBSB
	ListLatestMinedBlocksRIBSBC *ListLatestMinedBlocksRIBSBC
	ListLatestMinedBlocksRIBSBSC *ListLatestMinedBlocksRIBSBSC
	ListLatestMinedBlocksRIBSD *ListLatestMinedBlocksRIBSD
	ListLatestMinedBlocksRIBSD2 *ListLatestMinedBlocksRIBSD2
	ListLatestMinedBlocksRIBSEC *ListLatestMinedBlocksRIBSEC
	ListLatestMinedBlocksRIBSL *ListLatestMinedBlocksRIBSL
	ListLatestMinedBlocksRIBSX *ListLatestMinedBlocksRIBSX
	ListLatestMinedBlocksRIBSZ *ListLatestMinedBlocksRIBSZ
	ListLatestMinedBlocksRIBSZ2 *ListLatestMinedBlocksRIBSZ2
}

// ListLatestMinedBlocksRIBSBAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSB wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSBAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSB) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSB: v,
	}
}

// ListLatestMinedBlocksRIBSBCAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSBC wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSBCAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSBC) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSBC: v,
	}
}

// ListLatestMinedBlocksRIBSBSCAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSBSC wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSBSCAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSBSC) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSBSC: v,
	}
}

// ListLatestMinedBlocksRIBSDAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSD wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSDAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSD) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSD: v,
	}
}

// ListLatestMinedBlocksRIBSD2AsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSD2 wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSD2AsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSD2) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSD2: v,
	}
}

// ListLatestMinedBlocksRIBSECAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSEC wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSECAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSEC) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSEC: v,
	}
}

// ListLatestMinedBlocksRIBSLAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSL wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSLAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSL) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSL: v,
	}
}

// ListLatestMinedBlocksRIBSXAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSX wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSXAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSX) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSX: v,
	}
}

// ListLatestMinedBlocksRIBSZAsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSZ wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSZAsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSZ) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSZ: v,
	}
}

// ListLatestMinedBlocksRIBSZ2AsListLatestMinedBlocksRIBS is a convenience function that returns ListLatestMinedBlocksRIBSZ2 wrapped in ListLatestMinedBlocksRIBS
func ListLatestMinedBlocksRIBSZ2AsListLatestMinedBlocksRIBS(v *ListLatestMinedBlocksRIBSZ2) ListLatestMinedBlocksRIBS {
	return ListLatestMinedBlocksRIBS{
		ListLatestMinedBlocksRIBSZ2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListLatestMinedBlocksRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListLatestMinedBlocksRIBSB
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSB)
	if err == nil {
		jsonListLatestMinedBlocksRIBSB, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSB)
		if string(jsonListLatestMinedBlocksRIBSB) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSB = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSBC
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSBC)
	if err == nil {
		jsonListLatestMinedBlocksRIBSBC, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSBC)
		if string(jsonListLatestMinedBlocksRIBSBC) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSBC = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSBSC
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSBSC)
	if err == nil {
		jsonListLatestMinedBlocksRIBSBSC, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSBSC)
		if string(jsonListLatestMinedBlocksRIBSBSC) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSBSC = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSBSC = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSD
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSD)
	if err == nil {
		jsonListLatestMinedBlocksRIBSD, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSD)
		if string(jsonListLatestMinedBlocksRIBSD) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSD = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSD2
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSD2)
	if err == nil {
		jsonListLatestMinedBlocksRIBSD2, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSD2)
		if string(jsonListLatestMinedBlocksRIBSD2) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSD2 = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSEC
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSEC)
	if err == nil {
		jsonListLatestMinedBlocksRIBSEC, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSEC)
		if string(jsonListLatestMinedBlocksRIBSEC) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSEC = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSL
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSL)
	if err == nil {
		jsonListLatestMinedBlocksRIBSL, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSL)
		if string(jsonListLatestMinedBlocksRIBSL) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSL = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSX
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSX)
	if err == nil {
		jsonListLatestMinedBlocksRIBSX, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSX)
		if string(jsonListLatestMinedBlocksRIBSX) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSX = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSX = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSZ
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSZ)
	if err == nil {
		jsonListLatestMinedBlocksRIBSZ, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSZ)
		if string(jsonListLatestMinedBlocksRIBSZ) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSZ = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSZ = nil
	}

	// try to unmarshal data into ListLatestMinedBlocksRIBSZ2
	err = newStrictDecoder(data).Decode(&dst.ListLatestMinedBlocksRIBSZ2)
	if err == nil {
		jsonListLatestMinedBlocksRIBSZ2, _ := json.Marshal(dst.ListLatestMinedBlocksRIBSZ2)
		if string(jsonListLatestMinedBlocksRIBSZ2) == "{}" { // empty struct
			dst.ListLatestMinedBlocksRIBSZ2 = nil
		} else {
			match++
		}
	} else {
		dst.ListLatestMinedBlocksRIBSZ2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListLatestMinedBlocksRIBSB = nil
		dst.ListLatestMinedBlocksRIBSBC = nil
		dst.ListLatestMinedBlocksRIBSBSC = nil
		dst.ListLatestMinedBlocksRIBSD = nil
		dst.ListLatestMinedBlocksRIBSD2 = nil
		dst.ListLatestMinedBlocksRIBSEC = nil
		dst.ListLatestMinedBlocksRIBSL = nil
		dst.ListLatestMinedBlocksRIBSX = nil
		dst.ListLatestMinedBlocksRIBSZ = nil
		dst.ListLatestMinedBlocksRIBSZ2 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListLatestMinedBlocksRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListLatestMinedBlocksRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListLatestMinedBlocksRIBS) MarshalJSON() ([]byte, error) {
	if src.ListLatestMinedBlocksRIBSB != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSB)
	}

	if src.ListLatestMinedBlocksRIBSBC != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSBC)
	}

	if src.ListLatestMinedBlocksRIBSBSC != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSBSC)
	}

	if src.ListLatestMinedBlocksRIBSD != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSD)
	}

	if src.ListLatestMinedBlocksRIBSD2 != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSD2)
	}

	if src.ListLatestMinedBlocksRIBSEC != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSEC)
	}

	if src.ListLatestMinedBlocksRIBSL != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSL)
	}

	if src.ListLatestMinedBlocksRIBSX != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSX)
	}

	if src.ListLatestMinedBlocksRIBSZ != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSZ)
	}

	if src.ListLatestMinedBlocksRIBSZ2 != nil {
		return json.Marshal(&src.ListLatestMinedBlocksRIBSZ2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListLatestMinedBlocksRIBS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ListLatestMinedBlocksRIBSB != nil {
		return obj.ListLatestMinedBlocksRIBSB
	}

	if obj.ListLatestMinedBlocksRIBSBC != nil {
		return obj.ListLatestMinedBlocksRIBSBC
	}

	if obj.ListLatestMinedBlocksRIBSBSC != nil {
		return obj.ListLatestMinedBlocksRIBSBSC
	}

	if obj.ListLatestMinedBlocksRIBSD != nil {
		return obj.ListLatestMinedBlocksRIBSD
	}

	if obj.ListLatestMinedBlocksRIBSD2 != nil {
		return obj.ListLatestMinedBlocksRIBSD2
	}

	if obj.ListLatestMinedBlocksRIBSEC != nil {
		return obj.ListLatestMinedBlocksRIBSEC
	}

	if obj.ListLatestMinedBlocksRIBSL != nil {
		return obj.ListLatestMinedBlocksRIBSL
	}

	if obj.ListLatestMinedBlocksRIBSX != nil {
		return obj.ListLatestMinedBlocksRIBSX
	}

	if obj.ListLatestMinedBlocksRIBSZ != nil {
		return obj.ListLatestMinedBlocksRIBSZ
	}

	if obj.ListLatestMinedBlocksRIBSZ2 != nil {
		return obj.ListLatestMinedBlocksRIBSZ2
	}

	// all schemas are nil
	return nil
}

type NullableListLatestMinedBlocksRIBS struct {
	value *ListLatestMinedBlocksRIBS
	isSet bool
}

func (v NullableListLatestMinedBlocksRIBS) Get() *ListLatestMinedBlocksRIBS {
	return v.value
}

func (v *NullableListLatestMinedBlocksRIBS) Set(val *ListLatestMinedBlocksRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableListLatestMinedBlocksRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableListLatestMinedBlocksRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLatestMinedBlocksRIBS(val *ListLatestMinedBlocksRIBS) *NullableListLatestMinedBlocksRIBS {
	return &NullableListLatestMinedBlocksRIBS{value: val, isSet: true}
}

func (v NullableListLatestMinedBlocksRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLatestMinedBlocksRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


