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
)

// ListXRPRippleTransactionsByAddressR struct for ListXRPRippleTransactionsByAddressR
type ListXRPRippleTransactionsByAddressR struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data ListXRPRippleTransactionsByAddressRData `json:"data"`
}

// NewListXRPRippleTransactionsByAddressR instantiates a new ListXRPRippleTransactionsByAddressR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByAddressR(apiVersion string, requestId string, data ListXRPRippleTransactionsByAddressRData) *ListXRPRippleTransactionsByAddressR {
	this := ListXRPRippleTransactionsByAddressR{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewListXRPRippleTransactionsByAddressRWithDefaults instantiates a new ListXRPRippleTransactionsByAddressR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByAddressRWithDefaults() *ListXRPRippleTransactionsByAddressR {
	this := ListXRPRippleTransactionsByAddressR{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ListXRPRippleTransactionsByAddressR) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressR) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ListXRPRippleTransactionsByAddressR) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *ListXRPRippleTransactionsByAddressR) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressR) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ListXRPRippleTransactionsByAddressR) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListXRPRippleTransactionsByAddressR) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressR) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListXRPRippleTransactionsByAddressR) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *ListXRPRippleTransactionsByAddressR) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *ListXRPRippleTransactionsByAddressR) GetData() ListXRPRippleTransactionsByAddressRData {
	if o == nil {
		var ret ListXRPRippleTransactionsByAddressRData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressR) GetDataOk() (*ListXRPRippleTransactionsByAddressRData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListXRPRippleTransactionsByAddressR) SetData(v ListXRPRippleTransactionsByAddressRData) {
	o.Data = v
}

func (o ListXRPRippleTransactionsByAddressR) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListXRPRippleTransactionsByAddressR struct {
	value *ListXRPRippleTransactionsByAddressR
	isSet bool
}

func (v NullableListXRPRippleTransactionsByAddressR) Get() *ListXRPRippleTransactionsByAddressR {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByAddressR) Set(val *ListXRPRippleTransactionsByAddressR) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByAddressR) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByAddressR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByAddressR(val *ListXRPRippleTransactionsByAddressR) *NullableListXRPRippleTransactionsByAddressR {
	return &NullableListXRPRippleTransactionsByAddressR{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByAddressR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByAddressR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


