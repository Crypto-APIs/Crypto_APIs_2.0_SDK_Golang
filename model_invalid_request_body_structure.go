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

// InvalidRequestBodyStructure invalid_request_body_structure
type InvalidRequestBodyStructure struct {
	// Specifies an error code, e.g. error 404.
	Code string `json:"code"`
	// Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”.
	Message string `json:"message"`
	Details []BannedIpAddressDetailsInner `json:"details,omitempty"`
}

// NewInvalidRequestBodyStructure instantiates a new InvalidRequestBodyStructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidRequestBodyStructure(code string, message string) *InvalidRequestBodyStructure {
	this := InvalidRequestBodyStructure{}
	this.Code = code
	this.Message = message
	return &this
}

// NewInvalidRequestBodyStructureWithDefaults instantiates a new InvalidRequestBodyStructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidRequestBodyStructureWithDefaults() *InvalidRequestBodyStructure {
	this := InvalidRequestBodyStructure{}
	return &this
}

// GetCode returns the Code field value
func (o *InvalidRequestBodyStructure) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InvalidRequestBodyStructure) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InvalidRequestBodyStructure) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *InvalidRequestBodyStructure) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidRequestBodyStructure) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidRequestBodyStructure) SetMessage(v string) {
	o.Message = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InvalidRequestBodyStructure) GetDetails() []BannedIpAddressDetailsInner {
	if o == nil || o.Details == nil {
		var ret []BannedIpAddressDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidRequestBodyStructure) GetDetailsOk() ([]BannedIpAddressDetailsInner, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InvalidRequestBodyStructure) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []BannedIpAddressDetailsInner and assigns it to the Details field.
func (o *InvalidRequestBodyStructure) SetDetails(v []BannedIpAddressDetailsInner) {
	o.Details = v
}

func (o InvalidRequestBodyStructure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidRequestBodyStructure struct {
	value *InvalidRequestBodyStructure
	isSet bool
}

func (v NullableInvalidRequestBodyStructure) Get() *InvalidRequestBodyStructure {
	return v.value
}

func (v *NullableInvalidRequestBodyStructure) Set(val *InvalidRequestBodyStructure) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidRequestBodyStructure) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidRequestBodyStructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidRequestBodyStructure(val *InvalidRequestBodyStructure) *NullableInvalidRequestBodyStructure {
	return &NullableInvalidRequestBodyStructure{value: val, isSet: true}
}

func (v NullableInvalidRequestBodyStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidRequestBodyStructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


