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

// ListInternalTransactionsByAddressAndTimeRangeR struct for ListInternalTransactionsByAddressAndTimeRangeR
type ListInternalTransactionsByAddressAndTimeRangeR struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data ListInternalTransactionsByAddressAndTimeRangeRData `json:"data"`
}

// NewListInternalTransactionsByAddressAndTimeRangeR instantiates a new ListInternalTransactionsByAddressAndTimeRangeR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInternalTransactionsByAddressAndTimeRangeR(apiVersion string, requestId string, data ListInternalTransactionsByAddressAndTimeRangeRData) *ListInternalTransactionsByAddressAndTimeRangeR {
	this := ListInternalTransactionsByAddressAndTimeRangeR{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewListInternalTransactionsByAddressAndTimeRangeRWithDefaults instantiates a new ListInternalTransactionsByAddressAndTimeRangeR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInternalTransactionsByAddressAndTimeRangeRWithDefaults() *ListInternalTransactionsByAddressAndTimeRangeR {
	this := ListInternalTransactionsByAddressAndTimeRangeR{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ListInternalTransactionsByAddressAndTimeRangeR) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ListInternalTransactionsByAddressAndTimeRangeR) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListInternalTransactionsByAddressAndTimeRangeR) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *ListInternalTransactionsByAddressAndTimeRangeR) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetData() ListInternalTransactionsByAddressAndTimeRangeRData {
	if o == nil {
		var ret ListInternalTransactionsByAddressAndTimeRangeRData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListInternalTransactionsByAddressAndTimeRangeR) GetDataOk() (*ListInternalTransactionsByAddressAndTimeRangeRData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListInternalTransactionsByAddressAndTimeRangeR) SetData(v ListInternalTransactionsByAddressAndTimeRangeRData) {
	o.Data = v
}

func (o ListInternalTransactionsByAddressAndTimeRangeR) MarshalJSON() ([]byte, error) {
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

type NullableListInternalTransactionsByAddressAndTimeRangeR struct {
	value *ListInternalTransactionsByAddressAndTimeRangeR
	isSet bool
}

func (v NullableListInternalTransactionsByAddressAndTimeRangeR) Get() *ListInternalTransactionsByAddressAndTimeRangeR {
	return v.value
}

func (v *NullableListInternalTransactionsByAddressAndTimeRangeR) Set(val *ListInternalTransactionsByAddressAndTimeRangeR) {
	v.value = val
	v.isSet = true
}

func (v NullableListInternalTransactionsByAddressAndTimeRangeR) IsSet() bool {
	return v.isSet
}

func (v *NullableListInternalTransactionsByAddressAndTimeRangeR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInternalTransactionsByAddressAndTimeRangeR(val *ListInternalTransactionsByAddressAndTimeRangeR) *NullableListInternalTransactionsByAddressAndTimeRangeR {
	return &NullableListInternalTransactionsByAddressAndTimeRangeR{value: val, isSet: true}
}

func (v NullableListInternalTransactionsByAddressAndTimeRangeR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInternalTransactionsByAddressAndTimeRangeR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


