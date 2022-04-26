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

// InlineResponse403106 struct for InlineResponse403106
type InlineResponse403106 struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Error GetExchangeRateByAssetSymbolsE403 `json:"error"`
}

// NewInlineResponse403106 instantiates a new InlineResponse403106 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse403106(apiVersion string, requestId string, error_ GetExchangeRateByAssetSymbolsE403) *InlineResponse403106 {
	this := InlineResponse403106{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Error = error_
	return &this
}

// NewInlineResponse403106WithDefaults instantiates a new InlineResponse403106 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse403106WithDefaults() *InlineResponse403106 {
	this := InlineResponse403106{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *InlineResponse403106) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *InlineResponse403106) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *InlineResponse403106) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *InlineResponse403106) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse403106) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InlineResponse403106) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *InlineResponse403106) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse403106) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *InlineResponse403106) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *InlineResponse403106) SetContext(v string) {
	o.Context = &v
}

// GetError returns the Error field value
func (o *InlineResponse403106) GetError() GetExchangeRateByAssetSymbolsE403 {
	if o == nil {
		var ret GetExchangeRateByAssetSymbolsE403
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *InlineResponse403106) GetErrorOk() (*GetExchangeRateByAssetSymbolsE403, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *InlineResponse403106) SetError(v GetExchangeRateByAssetSymbolsE403) {
	o.Error = v
}

func (o InlineResponse403106) MarshalJSON() ([]byte, error) {
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
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse403106 struct {
	value *InlineResponse403106
	isSet bool
}

func (v NullableInlineResponse403106) Get() *InlineResponse403106 {
	return v.value
}

func (v *NullableInlineResponse403106) Set(val *InlineResponse403106) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse403106) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse403106) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse403106(val *InlineResponse403106) *NullableInlineResponse403106 {
	return &NullableInlineResponse403106{value: val, isSet: true}
}

func (v NullableInlineResponse403106) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse403106) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


