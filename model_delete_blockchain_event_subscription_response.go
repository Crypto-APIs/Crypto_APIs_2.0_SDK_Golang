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

// DeleteBlockchainEventSubscriptionResponse struct for DeleteBlockchainEventSubscriptionResponse
type DeleteBlockchainEventSubscriptionResponse struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data DeleteBlockchainEventSubscriptionResponseData `json:"data"`
}

// NewDeleteBlockchainEventSubscriptionResponse instantiates a new DeleteBlockchainEventSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBlockchainEventSubscriptionResponse(apiVersion string, requestId string, data DeleteBlockchainEventSubscriptionResponseData) *DeleteBlockchainEventSubscriptionResponse {
	this := DeleteBlockchainEventSubscriptionResponse{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewDeleteBlockchainEventSubscriptionResponseWithDefaults instantiates a new DeleteBlockchainEventSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBlockchainEventSubscriptionResponseWithDefaults() *DeleteBlockchainEventSubscriptionResponse {
	this := DeleteBlockchainEventSubscriptionResponse{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *DeleteBlockchainEventSubscriptionResponse) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscriptionResponse) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *DeleteBlockchainEventSubscriptionResponse) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *DeleteBlockchainEventSubscriptionResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscriptionResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DeleteBlockchainEventSubscriptionResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DeleteBlockchainEventSubscriptionResponse) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscriptionResponse) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DeleteBlockchainEventSubscriptionResponse) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *DeleteBlockchainEventSubscriptionResponse) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *DeleteBlockchainEventSubscriptionResponse) GetData() DeleteBlockchainEventSubscriptionResponseData {
	if o == nil {
		var ret DeleteBlockchainEventSubscriptionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscriptionResponse) GetDataOk() (*DeleteBlockchainEventSubscriptionResponseData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteBlockchainEventSubscriptionResponse) SetData(v DeleteBlockchainEventSubscriptionResponseData) {
	o.Data = v
}

func (o DeleteBlockchainEventSubscriptionResponse) MarshalJSON() ([]byte, error) {
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

type NullableDeleteBlockchainEventSubscriptionResponse struct {
	value *DeleteBlockchainEventSubscriptionResponse
	isSet bool
}

func (v NullableDeleteBlockchainEventSubscriptionResponse) Get() *DeleteBlockchainEventSubscriptionResponse {
	return v.value
}

func (v *NullableDeleteBlockchainEventSubscriptionResponse) Set(val *DeleteBlockchainEventSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBlockchainEventSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBlockchainEventSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBlockchainEventSubscriptionResponse(val *DeleteBlockchainEventSubscriptionResponse) *NullableDeleteBlockchainEventSubscriptionResponse {
	return &NullableDeleteBlockchainEventSubscriptionResponse{value: val, isSet: true}
}

func (v NullableDeleteBlockchainEventSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBlockchainEventSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

