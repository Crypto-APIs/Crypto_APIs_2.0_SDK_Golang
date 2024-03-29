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

// GetEIP1559FeeRecommendationsR struct for GetEIP1559FeeRecommendationsR
type GetEIP1559FeeRecommendationsR struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data GetEIP1559FeeRecommendationsRData `json:"data"`
}

// NewGetEIP1559FeeRecommendationsR instantiates a new GetEIP1559FeeRecommendationsR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEIP1559FeeRecommendationsR(apiVersion string, requestId string, data GetEIP1559FeeRecommendationsRData) *GetEIP1559FeeRecommendationsR {
	this := GetEIP1559FeeRecommendationsR{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewGetEIP1559FeeRecommendationsRWithDefaults instantiates a new GetEIP1559FeeRecommendationsR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEIP1559FeeRecommendationsRWithDefaults() *GetEIP1559FeeRecommendationsR {
	this := GetEIP1559FeeRecommendationsR{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *GetEIP1559FeeRecommendationsR) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsR) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *GetEIP1559FeeRecommendationsR) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *GetEIP1559FeeRecommendationsR) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsR) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetEIP1559FeeRecommendationsR) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GetEIP1559FeeRecommendationsR) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsR) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GetEIP1559FeeRecommendationsR) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *GetEIP1559FeeRecommendationsR) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *GetEIP1559FeeRecommendationsR) GetData() GetEIP1559FeeRecommendationsRData {
	if o == nil {
		var ret GetEIP1559FeeRecommendationsRData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEIP1559FeeRecommendationsR) GetDataOk() (*GetEIP1559FeeRecommendationsRData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEIP1559FeeRecommendationsR) SetData(v GetEIP1559FeeRecommendationsRData) {
	o.Data = v
}

func (o GetEIP1559FeeRecommendationsR) MarshalJSON() ([]byte, error) {
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

type NullableGetEIP1559FeeRecommendationsR struct {
	value *GetEIP1559FeeRecommendationsR
	isSet bool
}

func (v NullableGetEIP1559FeeRecommendationsR) Get() *GetEIP1559FeeRecommendationsR {
	return v.value
}

func (v *NullableGetEIP1559FeeRecommendationsR) Set(val *GetEIP1559FeeRecommendationsR) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEIP1559FeeRecommendationsR) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEIP1559FeeRecommendationsR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEIP1559FeeRecommendationsR(val *GetEIP1559FeeRecommendationsR) *NullableGetEIP1559FeeRecommendationsR {
	return &NullableGetEIP1559FeeRecommendationsR{value: val, isSet: true}
}

func (v NullableGetEIP1559FeeRecommendationsR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEIP1559FeeRecommendationsR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


