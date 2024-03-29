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

// ListConfirmedTransactionsByAddressRIBS - struct for ListConfirmedTransactionsByAddressRIBS
type ListConfirmedTransactionsByAddressRIBS struct {
	ListConfirmedTransactionsByAddressRIBSB *ListConfirmedTransactionsByAddressRIBSB
	ListConfirmedTransactionsByAddressRIBSBC *ListConfirmedTransactionsByAddressRIBSBC
	ListConfirmedTransactionsByAddressRIBSBSC *ListConfirmedTransactionsByAddressRIBSBSC
	ListConfirmedTransactionsByAddressRIBSD *ListConfirmedTransactionsByAddressRIBSD
	ListConfirmedTransactionsByAddressRIBSD2 *ListConfirmedTransactionsByAddressRIBSD2
	ListConfirmedTransactionsByAddressRIBSE *ListConfirmedTransactionsByAddressRIBSE
	ListConfirmedTransactionsByAddressRIBSEC *ListConfirmedTransactionsByAddressRIBSEC
	ListConfirmedTransactionsByAddressRIBSL *ListConfirmedTransactionsByAddressRIBSL
	ListConfirmedTransactionsByAddressRIBSZ *ListConfirmedTransactionsByAddressRIBSZ
}

// ListConfirmedTransactionsByAddressRIBSBAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSB wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSBAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSB) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSB: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSBCAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSBC wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSBCAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSBC) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSBC: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSBSCAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSBSC wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSBSCAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSBSC) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSBSC: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSDAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSD wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSDAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSD) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSD: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSD2AsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSD2 wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSD2AsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSD2) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSD2: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSEAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSE wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSEAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSE) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSE: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSECAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSEC wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSECAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSEC) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSEC: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSLAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSL wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSLAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSL) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSL: v,
	}
}

// ListConfirmedTransactionsByAddressRIBSZAsListConfirmedTransactionsByAddressRIBS is a convenience function that returns ListConfirmedTransactionsByAddressRIBSZ wrapped in ListConfirmedTransactionsByAddressRIBS
func ListConfirmedTransactionsByAddressRIBSZAsListConfirmedTransactionsByAddressRIBS(v *ListConfirmedTransactionsByAddressRIBSZ) ListConfirmedTransactionsByAddressRIBS {
	return ListConfirmedTransactionsByAddressRIBS{
		ListConfirmedTransactionsByAddressRIBSZ: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListConfirmedTransactionsByAddressRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSB
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSB)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSB, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSB)
		if string(jsonListConfirmedTransactionsByAddressRIBSB) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSB = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSBC
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSBC)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSBC, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSBC)
		if string(jsonListConfirmedTransactionsByAddressRIBSBC) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSBC = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSBSC
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSBSC)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSBSC, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSBSC)
		if string(jsonListConfirmedTransactionsByAddressRIBSBSC) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSBSC = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSBSC = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSD
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSD)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSD, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSD)
		if string(jsonListConfirmedTransactionsByAddressRIBSD) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSD = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSD2
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSD2)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSD2, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSD2)
		if string(jsonListConfirmedTransactionsByAddressRIBSD2) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSD2 = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSE
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSE)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSE, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSE)
		if string(jsonListConfirmedTransactionsByAddressRIBSE) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSE = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSE = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSEC
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSEC)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSEC, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSEC)
		if string(jsonListConfirmedTransactionsByAddressRIBSEC) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSEC = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSL
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSL)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSL, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSL)
		if string(jsonListConfirmedTransactionsByAddressRIBSL) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSL = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressRIBSZ
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressRIBSZ)
	if err == nil {
		jsonListConfirmedTransactionsByAddressRIBSZ, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressRIBSZ)
		if string(jsonListConfirmedTransactionsByAddressRIBSZ) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressRIBSZ = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressRIBSZ = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListConfirmedTransactionsByAddressRIBSB = nil
		dst.ListConfirmedTransactionsByAddressRIBSBC = nil
		dst.ListConfirmedTransactionsByAddressRIBSBSC = nil
		dst.ListConfirmedTransactionsByAddressRIBSD = nil
		dst.ListConfirmedTransactionsByAddressRIBSD2 = nil
		dst.ListConfirmedTransactionsByAddressRIBSE = nil
		dst.ListConfirmedTransactionsByAddressRIBSEC = nil
		dst.ListConfirmedTransactionsByAddressRIBSL = nil
		dst.ListConfirmedTransactionsByAddressRIBSZ = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListConfirmedTransactionsByAddressRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListConfirmedTransactionsByAddressRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListConfirmedTransactionsByAddressRIBS) MarshalJSON() ([]byte, error) {
	if src.ListConfirmedTransactionsByAddressRIBSB != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSB)
	}

	if src.ListConfirmedTransactionsByAddressRIBSBC != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSBC)
	}

	if src.ListConfirmedTransactionsByAddressRIBSBSC != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSBSC)
	}

	if src.ListConfirmedTransactionsByAddressRIBSD != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSD)
	}

	if src.ListConfirmedTransactionsByAddressRIBSD2 != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSD2)
	}

	if src.ListConfirmedTransactionsByAddressRIBSE != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSE)
	}

	if src.ListConfirmedTransactionsByAddressRIBSEC != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSEC)
	}

	if src.ListConfirmedTransactionsByAddressRIBSL != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSL)
	}

	if src.ListConfirmedTransactionsByAddressRIBSZ != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressRIBSZ)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListConfirmedTransactionsByAddressRIBS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ListConfirmedTransactionsByAddressRIBSB != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSB
	}

	if obj.ListConfirmedTransactionsByAddressRIBSBC != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSBC
	}

	if obj.ListConfirmedTransactionsByAddressRIBSBSC != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSBSC
	}

	if obj.ListConfirmedTransactionsByAddressRIBSD != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSD
	}

	if obj.ListConfirmedTransactionsByAddressRIBSD2 != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSD2
	}

	if obj.ListConfirmedTransactionsByAddressRIBSE != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSE
	}

	if obj.ListConfirmedTransactionsByAddressRIBSEC != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSEC
	}

	if obj.ListConfirmedTransactionsByAddressRIBSL != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSL
	}

	if obj.ListConfirmedTransactionsByAddressRIBSZ != nil {
		return obj.ListConfirmedTransactionsByAddressRIBSZ
	}

	// all schemas are nil
	return nil
}

type NullableListConfirmedTransactionsByAddressRIBS struct {
	value *ListConfirmedTransactionsByAddressRIBS
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBS) Get() *ListConfirmedTransactionsByAddressRIBS {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBS) Set(val *ListConfirmedTransactionsByAddressRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBS(val *ListConfirmedTransactionsByAddressRIBS) *NullableListConfirmedTransactionsByAddressRIBS {
	return &NullableListConfirmedTransactionsByAddressRIBS{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


