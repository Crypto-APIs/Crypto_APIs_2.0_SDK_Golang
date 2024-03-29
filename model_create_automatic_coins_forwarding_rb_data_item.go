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

// CreateAutomaticCoinsForwardingRBDataItem struct for CreateAutomaticCoinsForwardingRBDataItem
type CreateAutomaticCoinsForwardingRBDataItem struct {
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey string `json:"callbackSecretKey"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount int32 `json:"confirmationsCount"`
	// Represents the fee priority of the automation, whether it is \"slow\", \"standard\" or \"fast\".
	FeePriority string `json:"feePriority"`
	// Represents the minimum transfer amount of the currency in the `fromAddress` that can be allowed for an automatic forwarding.
	MinimumTransferAmount string `json:"minimumTransferAmount"`
	// Represents the hash of the address the currency is forwarded to.
	ToAddress string `json:"toAddress"`
}

// NewCreateAutomaticCoinsForwardingRBDataItem instantiates a new CreateAutomaticCoinsForwardingRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticCoinsForwardingRBDataItem(callbackSecretKey string, callbackUrl string, confirmationsCount int32, feePriority string, minimumTransferAmount string, toAddress string) *CreateAutomaticCoinsForwardingRBDataItem {
	this := CreateAutomaticCoinsForwardingRBDataItem{}
	this.CallbackSecretKey = callbackSecretKey
	this.CallbackUrl = callbackUrl
	this.ConfirmationsCount = confirmationsCount
	this.FeePriority = feePriority
	this.MinimumTransferAmount = minimumTransferAmount
	this.ToAddress = toAddress
	return &this
}

// NewCreateAutomaticCoinsForwardingRBDataItemWithDefaults instantiates a new CreateAutomaticCoinsForwardingRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticCoinsForwardingRBDataItemWithDefaults() *CreateAutomaticCoinsForwardingRBDataItem {
	this := CreateAutomaticCoinsForwardingRBDataItem{}
	return &this
}

// GetCallbackSecretKey returns the CallbackSecretKey field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetCallbackSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackSecretKey, true
}

// SetCallbackSecretKey sets field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetConfirmationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationsCount, true
}

// SetConfirmationsCount sets field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = v
}

// GetFeePriority returns the FeePriority field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetFeePriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetMinimumTransferAmount returns the MinimumTransferAmount field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetMinimumTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTransferAmount
}

// GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetMinimumTransferAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumTransferAmount, true
}

// SetMinimumTransferAmount sets field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) SetMinimumTransferAmount(v string) {
	o.MinimumTransferAmount = v
}

// GetToAddress returns the ToAddress field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRBDataItem) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *CreateAutomaticCoinsForwardingRBDataItem) SetToAddress(v string) {
	o.ToAddress = v
}

func (o CreateAutomaticCoinsForwardingRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["feePriority"] = o.FeePriority
	}
	if true {
		toSerialize["minimumTransferAmount"] = o.MinimumTransferAmount
	}
	if true {
		toSerialize["toAddress"] = o.ToAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticCoinsForwardingRBDataItem struct {
	value *CreateAutomaticCoinsForwardingRBDataItem
	isSet bool
}

func (v NullableCreateAutomaticCoinsForwardingRBDataItem) Get() *CreateAutomaticCoinsForwardingRBDataItem {
	return v.value
}

func (v *NullableCreateAutomaticCoinsForwardingRBDataItem) Set(val *CreateAutomaticCoinsForwardingRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticCoinsForwardingRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticCoinsForwardingRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticCoinsForwardingRBDataItem(val *CreateAutomaticCoinsForwardingRBDataItem) *NullableCreateAutomaticCoinsForwardingRBDataItem {
	return &NullableCreateAutomaticCoinsForwardingRBDataItem{value: val, isSet: true}
}

func (v NullableCreateAutomaticCoinsForwardingRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticCoinsForwardingRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


