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

// CreateAutomaticCoinsForwardingRI struct for CreateAutomaticCoinsForwardingRI
type CreateAutomaticCoinsForwardingRI struct {
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount int32 `json:"confirmationsCount"`
	// Defines the specific time/date when the automatic forwarding was created in Unix Timestamp.
	CreatedTimestamp int32 `json:"createdTimestamp"`
	// Represents the fee priority of the automation, whether it is \"SLOW\", \"STANDARD\" OR \"FAST\".
	FeePriority string `json:"feePriority"`
	// Represents the hash of the address that forwards the currency.
	FromAddress string `json:"fromAddress"`
	// Represents the minimum transfer amount of the currency in the `fromAddress` that can be allowed for an automatic forwarding.
	MinimumTransferAmount string `json:"minimumTransferAmount"`
	// Represents a unique ID used to reference the specific callback subscription.
	ReferenceId string `json:"referenceId"`
	// Represents the hash of the address the currency is forwarded to.
	ToAddress string `json:"toAddress"`
}

// NewCreateAutomaticCoinsForwardingRI instantiates a new CreateAutomaticCoinsForwardingRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticCoinsForwardingRI(callbackUrl string, confirmationsCount int32, createdTimestamp int32, feePriority string, fromAddress string, minimumTransferAmount string, referenceId string, toAddress string) *CreateAutomaticCoinsForwardingRI {
	this := CreateAutomaticCoinsForwardingRI{}
	this.CallbackUrl = callbackUrl
	this.ConfirmationsCount = confirmationsCount
	this.CreatedTimestamp = createdTimestamp
	this.FeePriority = feePriority
	this.FromAddress = fromAddress
	this.MinimumTransferAmount = minimumTransferAmount
	this.ReferenceId = referenceId
	this.ToAddress = toAddress
	return &this
}

// NewCreateAutomaticCoinsForwardingRIWithDefaults instantiates a new CreateAutomaticCoinsForwardingRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticCoinsForwardingRIWithDefaults() *CreateAutomaticCoinsForwardingRI {
	this := CreateAutomaticCoinsForwardingRI{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *CreateAutomaticCoinsForwardingRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value
func (o *CreateAutomaticCoinsForwardingRI) GetConfirmationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationsCount, true
}

// SetConfirmationsCount sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *CreateAutomaticCoinsForwardingRI) GetCreatedTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = v
}

// GetFeePriority returns the FeePriority field value
func (o *CreateAutomaticCoinsForwardingRI) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetFeePriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetFromAddress returns the FromAddress field value
func (o *CreateAutomaticCoinsForwardingRI) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetMinimumTransferAmount returns the MinimumTransferAmount field value
func (o *CreateAutomaticCoinsForwardingRI) GetMinimumTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTransferAmount
}

// GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetMinimumTransferAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumTransferAmount, true
}

// SetMinimumTransferAmount sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetMinimumTransferAmount(v string) {
	o.MinimumTransferAmount = v
}

// GetReferenceId returns the ReferenceId field value
func (o *CreateAutomaticCoinsForwardingRI) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetToAddress returns the ToAddress field value
func (o *CreateAutomaticCoinsForwardingRI) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRI) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *CreateAutomaticCoinsForwardingRI) SetToAddress(v string) {
	o.ToAddress = v
}

func (o CreateAutomaticCoinsForwardingRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["feePriority"] = o.FeePriority
	}
	if true {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if true {
		toSerialize["minimumTransferAmount"] = o.MinimumTransferAmount
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if true {
		toSerialize["toAddress"] = o.ToAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticCoinsForwardingRI struct {
	value *CreateAutomaticCoinsForwardingRI
	isSet bool
}

func (v NullableCreateAutomaticCoinsForwardingRI) Get() *CreateAutomaticCoinsForwardingRI {
	return v.value
}

func (v *NullableCreateAutomaticCoinsForwardingRI) Set(val *CreateAutomaticCoinsForwardingRI) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticCoinsForwardingRI) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticCoinsForwardingRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticCoinsForwardingRI(val *CreateAutomaticCoinsForwardingRI) *NullableCreateAutomaticCoinsForwardingRI {
	return &NullableCreateAutomaticCoinsForwardingRI{value: val, isSet: true}
}

func (v NullableCreateAutomaticCoinsForwardingRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticCoinsForwardingRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


