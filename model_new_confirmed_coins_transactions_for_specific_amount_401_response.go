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

// NewConfirmedCoinsTransactionsForSpecificAmount401Response struct for NewConfirmedCoinsTransactionsForSpecificAmount401Response
type NewConfirmedCoinsTransactionsForSpecificAmount401Response struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Error NewConfirmedCoinsTransactionsForSpecificAmountE401 `json:"error"`
}

// NewNewConfirmedCoinsTransactionsForSpecificAmount401Response instantiates a new NewConfirmedCoinsTransactionsForSpecificAmount401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedCoinsTransactionsForSpecificAmount401Response(apiVersion string, requestId string, error_ NewConfirmedCoinsTransactionsForSpecificAmountE401) *NewConfirmedCoinsTransactionsForSpecificAmount401Response {
	this := NewConfirmedCoinsTransactionsForSpecificAmount401Response{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Error = error_
	return &this
}

// NewNewConfirmedCoinsTransactionsForSpecificAmount401ResponseWithDefaults instantiates a new NewConfirmedCoinsTransactionsForSpecificAmount401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedCoinsTransactionsForSpecificAmount401ResponseWithDefaults() *NewConfirmedCoinsTransactionsForSpecificAmount401Response {
	this := NewConfirmedCoinsTransactionsForSpecificAmount401Response{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) SetContext(v string) {
	o.Context = &v
}

// GetError returns the Error field value
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetError() NewConfirmedCoinsTransactionsForSpecificAmountE401 {
	if o == nil {
		var ret NewConfirmedCoinsTransactionsForSpecificAmountE401
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) GetErrorOk() (*NewConfirmedCoinsTransactionsForSpecificAmountE401, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *NewConfirmedCoinsTransactionsForSpecificAmount401Response) SetError(v NewConfirmedCoinsTransactionsForSpecificAmountE401) {
	o.Error = v
}

func (o NewConfirmedCoinsTransactionsForSpecificAmount401Response) MarshalJSON() ([]byte, error) {
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

type NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response struct {
	value *NewConfirmedCoinsTransactionsForSpecificAmount401Response
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response) Get() *NewConfirmedCoinsTransactionsForSpecificAmount401Response {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response) Set(val *NewConfirmedCoinsTransactionsForSpecificAmount401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsForSpecificAmount401Response(val *NewConfirmedCoinsTransactionsForSpecificAmount401Response) *NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response {
	return &NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsForSpecificAmount401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


