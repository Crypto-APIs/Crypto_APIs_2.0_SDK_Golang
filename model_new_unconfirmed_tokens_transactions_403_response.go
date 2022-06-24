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

// NewUnconfirmedTokensTransactions403Response struct for NewUnconfirmedTokensTransactions403Response
type NewUnconfirmedTokensTransactions403Response struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Error NewUnconfirmedTokensTransactionsE403 `json:"error"`
}

// NewNewUnconfirmedTokensTransactions403Response instantiates a new NewUnconfirmedTokensTransactions403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUnconfirmedTokensTransactions403Response(apiVersion string, requestId string, error_ NewUnconfirmedTokensTransactionsE403) *NewUnconfirmedTokensTransactions403Response {
	this := NewUnconfirmedTokensTransactions403Response{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Error = error_
	return &this
}

// NewNewUnconfirmedTokensTransactions403ResponseWithDefaults instantiates a new NewUnconfirmedTokensTransactions403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUnconfirmedTokensTransactions403ResponseWithDefaults() *NewUnconfirmedTokensTransactions403Response {
	this := NewUnconfirmedTokensTransactions403Response{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *NewUnconfirmedTokensTransactions403Response) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactions403Response) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *NewUnconfirmedTokensTransactions403Response) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *NewUnconfirmedTokensTransactions403Response) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactions403Response) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *NewUnconfirmedTokensTransactions403Response) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *NewUnconfirmedTokensTransactions403Response) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactions403Response) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *NewUnconfirmedTokensTransactions403Response) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *NewUnconfirmedTokensTransactions403Response) SetContext(v string) {
	o.Context = &v
}

// GetError returns the Error field value
func (o *NewUnconfirmedTokensTransactions403Response) GetError() NewUnconfirmedTokensTransactionsE403 {
	if o == nil {
		var ret NewUnconfirmedTokensTransactionsE403
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactions403Response) GetErrorOk() (*NewUnconfirmedTokensTransactionsE403, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *NewUnconfirmedTokensTransactions403Response) SetError(v NewUnconfirmedTokensTransactionsE403) {
	o.Error = v
}

func (o NewUnconfirmedTokensTransactions403Response) MarshalJSON() ([]byte, error) {
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

type NullableNewUnconfirmedTokensTransactions403Response struct {
	value *NewUnconfirmedTokensTransactions403Response
	isSet bool
}

func (v NullableNewUnconfirmedTokensTransactions403Response) Get() *NewUnconfirmedTokensTransactions403Response {
	return v.value
}

func (v *NullableNewUnconfirmedTokensTransactions403Response) Set(val *NewUnconfirmedTokensTransactions403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUnconfirmedTokensTransactions403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUnconfirmedTokensTransactions403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUnconfirmedTokensTransactions403Response(val *NewUnconfirmedTokensTransactions403Response) *NullableNewUnconfirmedTokensTransactions403Response {
	return &NullableNewUnconfirmedTokensTransactions403Response{value: val, isSet: true}
}

func (v NullableNewUnconfirmedTokensTransactions403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUnconfirmedTokensTransactions403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


