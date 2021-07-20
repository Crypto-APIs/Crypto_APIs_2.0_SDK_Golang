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

// NewUnconfirmedTokensTransactionsRI struct for NewUnconfirmedTokensTransactionsRI
type NewUnconfirmedTokensTransactionsRI struct {
	// Defines the type of the specific event available for the customer to subscribe to for callback notification.
	Address string `json:"address"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs.
	CallbackSecretKey string `json:"callbackSecretKey"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl string `json:"callbackUrl"`
	// Defines the specific time/date when the subscription was created in Unix Timestamp.
	CreatedTimestamp int32 `json:"createdTimestamp"`
	// Defines the type of the specific event available for the customer to subscribe to for callback notification.
	EventType string `json:"eventType"`
	// Defines whether the subscription is active or not. Set as boolean.
	IsActive bool `json:"isActive"`
	// Represents a unique ID used to reference the specific callback subscription.
	ReferenceId string `json:"referenceId"`
	// Represents the unique identification string that defines the transaction.
	TransactionId string `json:"transactionId"`
}

// NewNewUnconfirmedTokensTransactionsRI instantiates a new NewUnconfirmedTokensTransactionsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUnconfirmedTokensTransactionsRI(address string, callbackSecretKey string, callbackUrl string, createdTimestamp int32, eventType string, isActive bool, referenceId string, transactionId string) *NewUnconfirmedTokensTransactionsRI {
	this := NewUnconfirmedTokensTransactionsRI{}
	this.Address = address
	this.CallbackSecretKey = callbackSecretKey
	this.CallbackUrl = callbackUrl
	this.CreatedTimestamp = createdTimestamp
	this.EventType = eventType
	this.IsActive = isActive
	this.ReferenceId = referenceId
	this.TransactionId = transactionId
	return &this
}

// NewNewUnconfirmedTokensTransactionsRIWithDefaults instantiates a new NewUnconfirmedTokensTransactionsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUnconfirmedTokensTransactionsRIWithDefaults() *NewUnconfirmedTokensTransactionsRI {
	this := NewUnconfirmedTokensTransactionsRI{}
	return &this
}

// GetAddress returns the Address field value
func (o *NewUnconfirmedTokensTransactionsRI) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetAddress(v string) {
	o.Address = v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value
func (o *NewUnconfirmedTokensTransactionsRI) GetCallbackSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackSecretKey, true
}

// SetCallbackSecretKey sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *NewUnconfirmedTokensTransactionsRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *NewUnconfirmedTokensTransactionsRI) GetCreatedTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = v
}

// GetEventType returns the EventType field value
func (o *NewUnconfirmedTokensTransactionsRI) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetEventType(v string) {
	o.EventType = v
}

// GetIsActive returns the IsActive field value
func (o *NewUnconfirmedTokensTransactionsRI) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetIsActive(v bool) {
	o.IsActive = v
}

// GetReferenceId returns the ReferenceId field value
func (o *NewUnconfirmedTokensTransactionsRI) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetReferenceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetTransactionId returns the TransactionId field value
func (o *NewUnconfirmedTokensTransactionsRI) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRI) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *NewUnconfirmedTokensTransactionsRI) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o NewUnconfirmedTokensTransactionsRI) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableNewUnconfirmedTokensTransactionsRI struct {
	value *NewUnconfirmedTokensTransactionsRI
	isSet bool
}

func (v NullableNewUnconfirmedTokensTransactionsRI) Get() *NewUnconfirmedTokensTransactionsRI {
	return v.value
}

func (v *NullableNewUnconfirmedTokensTransactionsRI) Set(val *NewUnconfirmedTokensTransactionsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUnconfirmedTokensTransactionsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUnconfirmedTokensTransactionsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUnconfirmedTokensTransactionsRI(val *NewUnconfirmedTokensTransactionsRI) *NullableNewUnconfirmedTokensTransactionsRI {
	return &NullableNewUnconfirmedTokensTransactionsRI{value: val, isSet: true}
}

func (v NullableNewUnconfirmedTokensTransactionsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUnconfirmedTokensTransactionsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

