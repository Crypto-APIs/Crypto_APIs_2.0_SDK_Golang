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
)

// GetZilliqaBlockDetailsByBlockHashR struct for GetZilliqaBlockDetailsByBlockHashR
type GetZilliqaBlockDetailsByBlockHashR struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data GetZilliqaBlockDetailsByBlockHashRData `json:"data"`
}

// NewGetZilliqaBlockDetailsByBlockHashR instantiates a new GetZilliqaBlockDetailsByBlockHashR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetZilliqaBlockDetailsByBlockHashR(apiVersion string, requestId string, data GetZilliqaBlockDetailsByBlockHashRData) *GetZilliqaBlockDetailsByBlockHashR {
	this := GetZilliqaBlockDetailsByBlockHashR{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewGetZilliqaBlockDetailsByBlockHashRWithDefaults instantiates a new GetZilliqaBlockDetailsByBlockHashR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetZilliqaBlockDetailsByBlockHashRWithDefaults() *GetZilliqaBlockDetailsByBlockHashR {
	this := GetZilliqaBlockDetailsByBlockHashR{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *GetZilliqaBlockDetailsByBlockHashR) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaBlockDetailsByBlockHashR) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *GetZilliqaBlockDetailsByBlockHashR) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *GetZilliqaBlockDetailsByBlockHashR) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaBlockDetailsByBlockHashR) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetZilliqaBlockDetailsByBlockHashR) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GetZilliqaBlockDetailsByBlockHashR) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetZilliqaBlockDetailsByBlockHashR) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GetZilliqaBlockDetailsByBlockHashR) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *GetZilliqaBlockDetailsByBlockHashR) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *GetZilliqaBlockDetailsByBlockHashR) GetData() GetZilliqaBlockDetailsByBlockHashRData {
	if o == nil {
		var ret GetZilliqaBlockDetailsByBlockHashRData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaBlockDetailsByBlockHashR) GetDataOk() (*GetZilliqaBlockDetailsByBlockHashRData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetZilliqaBlockDetailsByBlockHashR) SetData(v GetZilliqaBlockDetailsByBlockHashRData) {
	o.Data = v
}

func (o GetZilliqaBlockDetailsByBlockHashR) MarshalJSON() ([]byte, error) {
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

type NullableGetZilliqaBlockDetailsByBlockHashR struct {
	value *GetZilliqaBlockDetailsByBlockHashR
	isSet bool
}

func (v NullableGetZilliqaBlockDetailsByBlockHashR) Get() *GetZilliqaBlockDetailsByBlockHashR {
	return v.value
}

func (v *NullableGetZilliqaBlockDetailsByBlockHashR) Set(val *GetZilliqaBlockDetailsByBlockHashR) {
	v.value = val
	v.isSet = true
}

func (v NullableGetZilliqaBlockDetailsByBlockHashR) IsSet() bool {
	return v.isSet
}

func (v *NullableGetZilliqaBlockDetailsByBlockHashR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetZilliqaBlockDetailsByBlockHashR(val *GetZilliqaBlockDetailsByBlockHashR) *NullableGetZilliqaBlockDetailsByBlockHashR {
	return &NullableGetZilliqaBlockDetailsByBlockHashR{value: val, isSet: true}
}

func (v NullableGetZilliqaBlockDetailsByBlockHashR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetZilliqaBlockDetailsByBlockHashR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

