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

// DeleteBlockchainEventSubscription400Response struct for DeleteBlockchainEventSubscription400Response
type DeleteBlockchainEventSubscription400Response struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Error DeleteBlockchainEventSubscriptionE400 `json:"error"`
}

// NewDeleteBlockchainEventSubscription400Response instantiates a new DeleteBlockchainEventSubscription400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBlockchainEventSubscription400Response(apiVersion string, requestId string, error_ DeleteBlockchainEventSubscriptionE400) *DeleteBlockchainEventSubscription400Response {
	this := DeleteBlockchainEventSubscription400Response{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Error = error_
	return &this
}

// NewDeleteBlockchainEventSubscription400ResponseWithDefaults instantiates a new DeleteBlockchainEventSubscription400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBlockchainEventSubscription400ResponseWithDefaults() *DeleteBlockchainEventSubscription400Response {
	this := DeleteBlockchainEventSubscription400Response{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *DeleteBlockchainEventSubscription400Response) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscription400Response) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *DeleteBlockchainEventSubscription400Response) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *DeleteBlockchainEventSubscription400Response) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscription400Response) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DeleteBlockchainEventSubscription400Response) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DeleteBlockchainEventSubscription400Response) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscription400Response) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DeleteBlockchainEventSubscription400Response) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *DeleteBlockchainEventSubscription400Response) SetContext(v string) {
	o.Context = &v
}

// GetError returns the Error field value
func (o *DeleteBlockchainEventSubscription400Response) GetError() DeleteBlockchainEventSubscriptionE400 {
	if o == nil {
		var ret DeleteBlockchainEventSubscriptionE400
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscription400Response) GetErrorOk() (*DeleteBlockchainEventSubscriptionE400, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *DeleteBlockchainEventSubscription400Response) SetError(v DeleteBlockchainEventSubscriptionE400) {
	o.Error = v
}

func (o DeleteBlockchainEventSubscription400Response) MarshalJSON() ([]byte, error) {
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

type NullableDeleteBlockchainEventSubscription400Response struct {
	value *DeleteBlockchainEventSubscription400Response
	isSet bool
}

func (v NullableDeleteBlockchainEventSubscription400Response) Get() *DeleteBlockchainEventSubscription400Response {
	return v.value
}

func (v *NullableDeleteBlockchainEventSubscription400Response) Set(val *DeleteBlockchainEventSubscription400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBlockchainEventSubscription400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBlockchainEventSubscription400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBlockchainEventSubscription400Response(val *DeleteBlockchainEventSubscription400Response) *NullableDeleteBlockchainEventSubscription400Response {
	return &NullableDeleteBlockchainEventSubscription400Response{value: val, isSet: true}
}

func (v NullableDeleteBlockchainEventSubscription400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBlockchainEventSubscription400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


