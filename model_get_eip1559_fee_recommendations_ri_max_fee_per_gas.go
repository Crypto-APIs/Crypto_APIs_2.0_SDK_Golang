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
)

// GetEIP1559FeeRecommendationsRIMaxFeePerGas struct for GetEIP1559FeeRecommendationsRIMaxFeePerGas
type GetEIP1559FeeRecommendationsRIMaxFeePerGas struct {
	// Represents the fast maximum fee per gas, calculated from unconfirmed transactions.
	Fast string `json:"fast"`
	// Represents the slow maximum fee per gas, calculated from unconfirmed transactions.
	Slow string `json:"slow"`
	// Represents the standard maximum fee per gas, calculated from unconfirmed transactions.
	Standard string `json:"standard"`
	// Represents the unit of the maximum fee per gas.
	Unit string `json:"unit"`
}

// NewGetEIP1559FeeRecommendationsRIMaxFeePerGas instantiates a new GetEIP1559FeeRecommendationsRIMaxFeePerGas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEIP1559FeeRecommendationsRIMaxFeePerGas(fast string, slow string, standard string, unit string) *GetEIP1559FeeRecommendationsRIMaxFeePerGas {
	this := GetEIP1559FeeRecommendationsRIMaxFeePerGas{}
	this.Fast = fast
	this.Slow = slow
	this.Standard = standard
	this.Unit = unit
	return &this
}

// NewGetEIP1559FeeRecommendationsRIMaxFeePerGasWithDefaults instantiates a new GetEIP1559FeeRecommendationsRIMaxFeePerGas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEIP1559FeeRecommendationsRIMaxFeePerGasWithDefaults() *GetEIP1559FeeRecommendationsRIMaxFeePerGas {
	this := GetEIP1559FeeRecommendationsRIMaxFeePerGas{}
	return &this
}

// GetFast returns the Fast field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetFast() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fast
}

// GetFastOk returns a tuple with the Fast field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetFastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fast, true
}

// SetFast sets field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) SetFast(v string) {
	o.Fast = v
}

// GetSlow returns the Slow field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetSlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slow
}

// GetSlowOk returns a tuple with the Slow field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetSlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slow, true
}

// SetSlow sets field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) SetSlow(v string) {
	o.Slow = v
}

// GetStandard returns the Standard field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetStandard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Standard
}

// GetStandardOk returns a tuple with the Standard field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetStandardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Standard, true
}

// SetStandard sets field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) SetStandard(v string) {
	o.Standard = v
}

// GetUnit returns the Unit field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetEIP1559FeeRecommendationsRIMaxFeePerGas) SetUnit(v string) {
	o.Unit = v
}

func (o GetEIP1559FeeRecommendationsRIMaxFeePerGas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fast"] = o.Fast
	}
	if true {
		toSerialize["slow"] = o.Slow
	}
	if true {
		toSerialize["standard"] = o.Standard
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas struct {
	value *GetEIP1559FeeRecommendationsRIMaxFeePerGas
	isSet bool
}

func (v NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas) Get() *GetEIP1559FeeRecommendationsRIMaxFeePerGas {
	return v.value
}

func (v *NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas) Set(val *GetEIP1559FeeRecommendationsRIMaxFeePerGas) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEIP1559FeeRecommendationsRIMaxFeePerGas(val *GetEIP1559FeeRecommendationsRIMaxFeePerGas) *NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas {
	return &NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas{value: val, isSet: true}
}

func (v NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEIP1559FeeRecommendationsRIMaxFeePerGas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


