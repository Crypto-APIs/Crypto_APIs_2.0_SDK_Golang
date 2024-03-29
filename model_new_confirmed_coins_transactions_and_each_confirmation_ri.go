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

// NewConfirmedCoinsTransactionsAndEachConfirmationRI struct for NewConfirmedCoinsTransactionsAndEachConfirmationRI
type NewConfirmedCoinsTransactionsAndEachConfirmationRI struct {
	// Represents the address of the transaction, per which the result is returned.
	Address string `json:"address"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey string `json:"callbackSecretKey"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount int32 `json:"confirmationsCount"`
	// Defines the specific time/date when the subscription was created in Unix Timestamp.
	CreatedTimestamp int32 `json:"createdTimestamp"`
	// Defines the type of the specific event available for the customer to subscribe to for callback notification.
	EventType string `json:"eventType"`
	// Defines whether the subscription is active or not. Set as boolean.
	IsActive bool `json:"isActive"`
	// Represents a unique ID used to reference the specific callback subscription.
	ReferenceId string `json:"referenceId"`
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRI instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRI(address string, callbackSecretKey string, callbackUrl string, confirmationsCount int32, createdTimestamp int32, eventType string, isActive bool, referenceId string) *NewConfirmedCoinsTransactionsAndEachConfirmationRI {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRI{}
	this.Address = address
	this.CallbackSecretKey = callbackSecretKey
	this.CallbackUrl = callbackUrl
	this.ConfirmationsCount = confirmationsCount
	this.CreatedTimestamp = createdTimestamp
	this.EventType = eventType
	this.IsActive = isActive
	this.ReferenceId = referenceId
	return &this
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRIWithDefaults instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRIWithDefaults() *NewConfirmedCoinsTransactionsAndEachConfirmationRI {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRI{}
	return &this
}

// GetAddress returns the Address field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetAddress(v string) {
	o.Address = v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetCallbackSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackSecretKey, true
}

// SetCallbackSecretKey sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetConfirmationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationsCount, true
}

// SetConfirmationsCount sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetCreatedTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = v
}

// GetEventType returns the EventType field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetEventType(v string) {
	o.EventType = v
}

// GetIsActive returns the IsActive field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetIsActive(v bool) {
	o.IsActive = v
}

// GetReferenceId returns the ReferenceId field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRI) SetReferenceId(v string) {
	o.ReferenceId = v
}

func (o NewConfirmedCoinsTransactionsAndEachConfirmationRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if true {
		toSerialize["confirmationsCount"] = o.ConfirmationsCount
	}
	if true {
		toSerialize["createdTimestamp"] = o.CreatedTimestamp
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["isActive"] = o.IsActive
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI struct {
	value *NewConfirmedCoinsTransactionsAndEachConfirmationRI
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI) Get() *NewConfirmedCoinsTransactionsAndEachConfirmationRI {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI) Set(val *NewConfirmedCoinsTransactionsAndEachConfirmationRI) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsAndEachConfirmationRI(val *NewConfirmedCoinsTransactionsAndEachConfirmationRI) *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI {
	return &NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


