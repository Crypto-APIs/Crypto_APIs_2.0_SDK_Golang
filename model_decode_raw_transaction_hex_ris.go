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

// DecodeRawTransactionHexRIS - Represents the specific transaction data according to the blockchain
type DecodeRawTransactionHexRIS struct {
	DecodeRawTransactionHexRISB *DecodeRawTransactionHexRISB
	DecodeRawTransactionHexRISB2 *DecodeRawTransactionHexRISB2
	DecodeRawTransactionHexRISB22 *DecodeRawTransactionHexRISB22
	DecodeRawTransactionHexRISD *DecodeRawTransactionHexRISD
	DecodeRawTransactionHexRISD2 *DecodeRawTransactionHexRISD2
	DecodeRawTransactionHexRISE *DecodeRawTransactionHexRISE
	DecodeRawTransactionHexRISE2 *DecodeRawTransactionHexRISE2
	DecodeRawTransactionHexRISL *DecodeRawTransactionHexRISL
	DecodeRawTransactionHexRISZ *DecodeRawTransactionHexRISZ
}

// DecodeRawTransactionHexRISBAsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISB wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISBAsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISB) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISB: v,
	}
}

// DecodeRawTransactionHexRISB2AsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISB2 wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISB2AsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISB2) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISB2: v,
	}
}

// DecodeRawTransactionHexRISB22AsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISB22 wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISB22AsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISB22) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISB22: v,
	}
}

// DecodeRawTransactionHexRISDAsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISD wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISDAsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISD) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISD: v,
	}
}

// DecodeRawTransactionHexRISD2AsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISD2 wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISD2AsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISD2) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISD2: v,
	}
}

// DecodeRawTransactionHexRISEAsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISE wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISEAsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISE) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISE: v,
	}
}

// DecodeRawTransactionHexRISE2AsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISE2 wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISE2AsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISE2) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISE2: v,
	}
}

// DecodeRawTransactionHexRISLAsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISL wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISLAsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISL) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISL: v,
	}
}

// DecodeRawTransactionHexRISZAsDecodeRawTransactionHexRIS is a convenience function that returns DecodeRawTransactionHexRISZ wrapped in DecodeRawTransactionHexRIS
func DecodeRawTransactionHexRISZAsDecodeRawTransactionHexRIS(v *DecodeRawTransactionHexRISZ) DecodeRawTransactionHexRIS {
	return DecodeRawTransactionHexRIS{
		DecodeRawTransactionHexRISZ: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DecodeRawTransactionHexRIS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DecodeRawTransactionHexRISB
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISB)
	if err == nil {
		jsonDecodeRawTransactionHexRISB, _ := json.Marshal(dst.DecodeRawTransactionHexRISB)
		if string(jsonDecodeRawTransactionHexRISB) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISB = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISB = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISB2
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISB2)
	if err == nil {
		jsonDecodeRawTransactionHexRISB2, _ := json.Marshal(dst.DecodeRawTransactionHexRISB2)
		if string(jsonDecodeRawTransactionHexRISB2) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISB2 = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISB2 = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISB22
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISB22)
	if err == nil {
		jsonDecodeRawTransactionHexRISB22, _ := json.Marshal(dst.DecodeRawTransactionHexRISB22)
		if string(jsonDecodeRawTransactionHexRISB22) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISB22 = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISB22 = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISD
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISD)
	if err == nil {
		jsonDecodeRawTransactionHexRISD, _ := json.Marshal(dst.DecodeRawTransactionHexRISD)
		if string(jsonDecodeRawTransactionHexRISD) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISD = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISD = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISD2
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISD2)
	if err == nil {
		jsonDecodeRawTransactionHexRISD2, _ := json.Marshal(dst.DecodeRawTransactionHexRISD2)
		if string(jsonDecodeRawTransactionHexRISD2) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISD2 = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISD2 = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISE
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISE)
	if err == nil {
		jsonDecodeRawTransactionHexRISE, _ := json.Marshal(dst.DecodeRawTransactionHexRISE)
		if string(jsonDecodeRawTransactionHexRISE) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISE = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISE = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISE2
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISE2)
	if err == nil {
		jsonDecodeRawTransactionHexRISE2, _ := json.Marshal(dst.DecodeRawTransactionHexRISE2)
		if string(jsonDecodeRawTransactionHexRISE2) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISE2 = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISE2 = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISL
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISL)
	if err == nil {
		jsonDecodeRawTransactionHexRISL, _ := json.Marshal(dst.DecodeRawTransactionHexRISL)
		if string(jsonDecodeRawTransactionHexRISL) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISL = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISL = nil
	}

	// try to unmarshal data into DecodeRawTransactionHexRISZ
	err = newStrictDecoder(data).Decode(&dst.DecodeRawTransactionHexRISZ)
	if err == nil {
		jsonDecodeRawTransactionHexRISZ, _ := json.Marshal(dst.DecodeRawTransactionHexRISZ)
		if string(jsonDecodeRawTransactionHexRISZ) == "{}" { // empty struct
			dst.DecodeRawTransactionHexRISZ = nil
		} else {
			match++
		}
	} else {
		dst.DecodeRawTransactionHexRISZ = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DecodeRawTransactionHexRISB = nil
		dst.DecodeRawTransactionHexRISB2 = nil
		dst.DecodeRawTransactionHexRISB22 = nil
		dst.DecodeRawTransactionHexRISD = nil
		dst.DecodeRawTransactionHexRISD2 = nil
		dst.DecodeRawTransactionHexRISE = nil
		dst.DecodeRawTransactionHexRISE2 = nil
		dst.DecodeRawTransactionHexRISL = nil
		dst.DecodeRawTransactionHexRISZ = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(DecodeRawTransactionHexRIS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(DecodeRawTransactionHexRIS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DecodeRawTransactionHexRIS) MarshalJSON() ([]byte, error) {
	if src.DecodeRawTransactionHexRISB != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISB)
	}

	if src.DecodeRawTransactionHexRISB2 != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISB2)
	}

	if src.DecodeRawTransactionHexRISB22 != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISB22)
	}

	if src.DecodeRawTransactionHexRISD != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISD)
	}

	if src.DecodeRawTransactionHexRISD2 != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISD2)
	}

	if src.DecodeRawTransactionHexRISE != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISE)
	}

	if src.DecodeRawTransactionHexRISE2 != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISE2)
	}

	if src.DecodeRawTransactionHexRISL != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISL)
	}

	if src.DecodeRawTransactionHexRISZ != nil {
		return json.Marshal(&src.DecodeRawTransactionHexRISZ)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DecodeRawTransactionHexRIS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DecodeRawTransactionHexRISB != nil {
		return obj.DecodeRawTransactionHexRISB
	}

	if obj.DecodeRawTransactionHexRISB2 != nil {
		return obj.DecodeRawTransactionHexRISB2
	}

	if obj.DecodeRawTransactionHexRISB22 != nil {
		return obj.DecodeRawTransactionHexRISB22
	}

	if obj.DecodeRawTransactionHexRISD != nil {
		return obj.DecodeRawTransactionHexRISD
	}

	if obj.DecodeRawTransactionHexRISD2 != nil {
		return obj.DecodeRawTransactionHexRISD2
	}

	if obj.DecodeRawTransactionHexRISE != nil {
		return obj.DecodeRawTransactionHexRISE
	}

	if obj.DecodeRawTransactionHexRISE2 != nil {
		return obj.DecodeRawTransactionHexRISE2
	}

	if obj.DecodeRawTransactionHexRISL != nil {
		return obj.DecodeRawTransactionHexRISL
	}

	if obj.DecodeRawTransactionHexRISZ != nil {
		return obj.DecodeRawTransactionHexRISZ
	}

	// all schemas are nil
	return nil
}

type NullableDecodeRawTransactionHexRIS struct {
	value *DecodeRawTransactionHexRIS
	isSet bool
}

func (v NullableDecodeRawTransactionHexRIS) Get() *DecodeRawTransactionHexRIS {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRIS) Set(val *DecodeRawTransactionHexRIS) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRIS) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRIS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRIS(val *DecodeRawTransactionHexRIS) *NullableDecodeRawTransactionHexRIS {
	return &NullableDecodeRawTransactionHexRIS{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRIS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRIS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


