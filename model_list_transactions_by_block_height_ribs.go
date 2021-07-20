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

// ListTransactionsByBlockHeightRIBS - struct for ListTransactionsByBlockHeightRIBS
type ListTransactionsByBlockHeightRIBS struct {
	ListTransactionsByBlockHeightRIBSB *ListTransactionsByBlockHeightRIBSB
	ListTransactionsByBlockHeightRIBSBC *ListTransactionsByBlockHeightRIBSBC
	ListTransactionsByBlockHeightRIBSD *ListTransactionsByBlockHeightRIBSD
	ListTransactionsByBlockHeightRIBSD2 *ListTransactionsByBlockHeightRIBSD2
	ListTransactionsByBlockHeightRIBSE *ListTransactionsByBlockHeightRIBSE
	ListTransactionsByBlockHeightRIBSEC *ListTransactionsByBlockHeightRIBSEC
	ListTransactionsByBlockHeightRIBSL *ListTransactionsByBlockHeightRIBSL
}

// ListTransactionsByBlockHeightRIBSBAsListTransactionsByBlockHeightRIBS is a convenience function that returns ListTransactionsByBlockHeightRIBSB wrapped in ListTransactionsByBlockHeightRIBS
func ListTransactionsByBlockHeightRIBSBAsListTransactionsByBlockHeightRIBS(v *ListTransactionsByBlockHeightRIBSB) ListTransactionsByBlockHeightRIBS {
	return ListTransactionsByBlockHeightRIBS{ ListTransactionsByBlockHeightRIBSB: v}
}

// ListTransactionsByBlockHeightRIBSBCAsListTransactionsByBlockHeightRIBS is a convenience function that returns ListTransactionsByBlockHeightRIBSBC wrapped in ListTransactionsByBlockHeightRIBS
func ListTransactionsByBlockHeightRIBSBCAsListTransactionsByBlockHeightRIBS(v *ListTransactionsByBlockHeightRIBSBC) ListTransactionsByBlockHeightRIBS {
	return ListTransactionsByBlockHeightRIBS{ ListTransactionsByBlockHeightRIBSBC: v}
}

// ListTransactionsByBlockHeightRIBSDAsListTransactionsByBlockHeightRIBS is a convenience function that returns ListTransactionsByBlockHeightRIBSD wrapped in ListTransactionsByBlockHeightRIBS
func ListTransactionsByBlockHeightRIBSDAsListTransactionsByBlockHeightRIBS(v *ListTransactionsByBlockHeightRIBSD) ListTransactionsByBlockHeightRIBS {
	return ListTransactionsByBlockHeightRIBS{ ListTransactionsByBlockHeightRIBSD: v}
}

// ListTransactionsByBlockHeightRIBSD2AsListTransactionsByBlockHeightRIBS is a convenience function that returns ListTransactionsByBlockHeightRIBSD2 wrapped in ListTransactionsByBlockHeightRIBS
func ListTransactionsByBlockHeightRIBSD2AsListTransactionsByBlockHeightRIBS(v *ListTransactionsByBlockHeightRIBSD2) ListTransactionsByBlockHeightRIBS {
	return ListTransactionsByBlockHeightRIBS{ ListTransactionsByBlockHeightRIBSD2: v}
}

// ListTransactionsByBlockHeightRIBSEAsListTransactionsByBlockHeightRIBS is a convenience function that returns ListTransactionsByBlockHeightRIBSE wrapped in ListTransactionsByBlockHeightRIBS
func ListTransactionsByBlockHeightRIBSEAsListTransactionsByBlockHeightRIBS(v *ListTransactionsByBlockHeightRIBSE) ListTransactionsByBlockHeightRIBS {
	return ListTransactionsByBlockHeightRIBS{ ListTransactionsByBlockHeightRIBSE: v}
}

// ListTransactionsByBlockHeightRIBSECAsListTransactionsByBlockHeightRIBS is a convenience function that returns ListTransactionsByBlockHeightRIBSEC wrapped in ListTransactionsByBlockHeightRIBS
func ListTransactionsByBlockHeightRIBSECAsListTransactionsByBlockHeightRIBS(v *ListTransactionsByBlockHeightRIBSEC) ListTransactionsByBlockHeightRIBS {
	return ListTransactionsByBlockHeightRIBS{ ListTransactionsByBlockHeightRIBSEC: v}
}

// ListTransactionsByBlockHeightRIBSLAsListTransactionsByBlockHeightRIBS is a convenience function that returns ListTransactionsByBlockHeightRIBSL wrapped in ListTransactionsByBlockHeightRIBS
func ListTransactionsByBlockHeightRIBSLAsListTransactionsByBlockHeightRIBS(v *ListTransactionsByBlockHeightRIBSL) ListTransactionsByBlockHeightRIBS {
	return ListTransactionsByBlockHeightRIBS{ ListTransactionsByBlockHeightRIBSL: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListTransactionsByBlockHeightRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListTransactionsByBlockHeightRIBSB
	err = json.Unmarshal(data, &dst.ListTransactionsByBlockHeightRIBSB)
	if err == nil {
		jsonListTransactionsByBlockHeightRIBSB, _ := json.Marshal(dst.ListTransactionsByBlockHeightRIBSB)
		if string(jsonListTransactionsByBlockHeightRIBSB) == "{}" { // empty struct
			dst.ListTransactionsByBlockHeightRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.ListTransactionsByBlockHeightRIBSB = nil
	}

	// try to unmarshal data into ListTransactionsByBlockHeightRIBSBC
	err = json.Unmarshal(data, &dst.ListTransactionsByBlockHeightRIBSBC)
	if err == nil {
		jsonListTransactionsByBlockHeightRIBSBC, _ := json.Marshal(dst.ListTransactionsByBlockHeightRIBSBC)
		if string(jsonListTransactionsByBlockHeightRIBSBC) == "{}" { // empty struct
			dst.ListTransactionsByBlockHeightRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.ListTransactionsByBlockHeightRIBSBC = nil
	}

	// try to unmarshal data into ListTransactionsByBlockHeightRIBSD
	err = json.Unmarshal(data, &dst.ListTransactionsByBlockHeightRIBSD)
	if err == nil {
		jsonListTransactionsByBlockHeightRIBSD, _ := json.Marshal(dst.ListTransactionsByBlockHeightRIBSD)
		if string(jsonListTransactionsByBlockHeightRIBSD) == "{}" { // empty struct
			dst.ListTransactionsByBlockHeightRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.ListTransactionsByBlockHeightRIBSD = nil
	}

	// try to unmarshal data into ListTransactionsByBlockHeightRIBSD2
	err = json.Unmarshal(data, &dst.ListTransactionsByBlockHeightRIBSD2)
	if err == nil {
		jsonListTransactionsByBlockHeightRIBSD2, _ := json.Marshal(dst.ListTransactionsByBlockHeightRIBSD2)
		if string(jsonListTransactionsByBlockHeightRIBSD2) == "{}" { // empty struct
			dst.ListTransactionsByBlockHeightRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.ListTransactionsByBlockHeightRIBSD2 = nil
	}

	// try to unmarshal data into ListTransactionsByBlockHeightRIBSE
	err = json.Unmarshal(data, &dst.ListTransactionsByBlockHeightRIBSE)
	if err == nil {
		jsonListTransactionsByBlockHeightRIBSE, _ := json.Marshal(dst.ListTransactionsByBlockHeightRIBSE)
		if string(jsonListTransactionsByBlockHeightRIBSE) == "{}" { // empty struct
			dst.ListTransactionsByBlockHeightRIBSE = nil
		} else {
			match++
		}
	} else {
		dst.ListTransactionsByBlockHeightRIBSE = nil
	}

	// try to unmarshal data into ListTransactionsByBlockHeightRIBSEC
	err = json.Unmarshal(data, &dst.ListTransactionsByBlockHeightRIBSEC)
	if err == nil {
		jsonListTransactionsByBlockHeightRIBSEC, _ := json.Marshal(dst.ListTransactionsByBlockHeightRIBSEC)
		if string(jsonListTransactionsByBlockHeightRIBSEC) == "{}" { // empty struct
			dst.ListTransactionsByBlockHeightRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.ListTransactionsByBlockHeightRIBSEC = nil
	}

	// try to unmarshal data into ListTransactionsByBlockHeightRIBSL
	err = json.Unmarshal(data, &dst.ListTransactionsByBlockHeightRIBSL)
	if err == nil {
		jsonListTransactionsByBlockHeightRIBSL, _ := json.Marshal(dst.ListTransactionsByBlockHeightRIBSL)
		if string(jsonListTransactionsByBlockHeightRIBSL) == "{}" { // empty struct
			dst.ListTransactionsByBlockHeightRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.ListTransactionsByBlockHeightRIBSL = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListTransactionsByBlockHeightRIBSB = nil
		dst.ListTransactionsByBlockHeightRIBSBC = nil
		dst.ListTransactionsByBlockHeightRIBSD = nil
		dst.ListTransactionsByBlockHeightRIBSD2 = nil
		dst.ListTransactionsByBlockHeightRIBSE = nil
		dst.ListTransactionsByBlockHeightRIBSEC = nil
		dst.ListTransactionsByBlockHeightRIBSL = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListTransactionsByBlockHeightRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListTransactionsByBlockHeightRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListTransactionsByBlockHeightRIBS) MarshalJSON() ([]byte, error) {
	if src.ListTransactionsByBlockHeightRIBSB != nil {
		return json.Marshal(&src.ListTransactionsByBlockHeightRIBSB)
	}

	if src.ListTransactionsByBlockHeightRIBSBC != nil {
		return json.Marshal(&src.ListTransactionsByBlockHeightRIBSBC)
	}

	if src.ListTransactionsByBlockHeightRIBSD != nil {
		return json.Marshal(&src.ListTransactionsByBlockHeightRIBSD)
	}

	if src.ListTransactionsByBlockHeightRIBSD2 != nil {
		return json.Marshal(&src.ListTransactionsByBlockHeightRIBSD2)
	}

	if src.ListTransactionsByBlockHeightRIBSE != nil {
		return json.Marshal(&src.ListTransactionsByBlockHeightRIBSE)
	}

	if src.ListTransactionsByBlockHeightRIBSEC != nil {
		return json.Marshal(&src.ListTransactionsByBlockHeightRIBSEC)
	}

	if src.ListTransactionsByBlockHeightRIBSL != nil {
		return json.Marshal(&src.ListTransactionsByBlockHeightRIBSL)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListTransactionsByBlockHeightRIBS) GetActualInstance() (interface{}) {
	if obj.ListTransactionsByBlockHeightRIBSB != nil {
		return obj.ListTransactionsByBlockHeightRIBSB
	}

	if obj.ListTransactionsByBlockHeightRIBSBC != nil {
		return obj.ListTransactionsByBlockHeightRIBSBC
	}

	if obj.ListTransactionsByBlockHeightRIBSD != nil {
		return obj.ListTransactionsByBlockHeightRIBSD
	}

	if obj.ListTransactionsByBlockHeightRIBSD2 != nil {
		return obj.ListTransactionsByBlockHeightRIBSD2
	}

	if obj.ListTransactionsByBlockHeightRIBSE != nil {
		return obj.ListTransactionsByBlockHeightRIBSE
	}

	if obj.ListTransactionsByBlockHeightRIBSEC != nil {
		return obj.ListTransactionsByBlockHeightRIBSEC
	}

	if obj.ListTransactionsByBlockHeightRIBSL != nil {
		return obj.ListTransactionsByBlockHeightRIBSL
	}

	// all schemas are nil
	return nil
}

type NullableListTransactionsByBlockHeightRIBS struct {
	value *ListTransactionsByBlockHeightRIBS
	isSet bool
}

func (v NullableListTransactionsByBlockHeightRIBS) Get() *ListTransactionsByBlockHeightRIBS {
	return v.value
}

func (v *NullableListTransactionsByBlockHeightRIBS) Set(val *ListTransactionsByBlockHeightRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHeightRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHeightRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHeightRIBS(val *ListTransactionsByBlockHeightRIBS) *NullableListTransactionsByBlockHeightRIBS {
	return &NullableListTransactionsByBlockHeightRIBS{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHeightRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHeightRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

