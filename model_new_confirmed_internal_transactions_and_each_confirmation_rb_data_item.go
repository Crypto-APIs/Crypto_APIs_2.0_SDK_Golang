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

// NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem struct for NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem
type NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem struct {
	// Defines the specific address of the internal transaction.
	Address string `json:"address"`
	// Specifies a flag that permits or denies the creation of duplicate addresses.
	AllowDuplicates bool `json:"allowDuplicates"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey string `json:"callbackSecretKey"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount int32 `json:"confirmationsCount"`
}

// NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem instantiates a new NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem(address string, allowDuplicates bool, callbackSecretKey string, callbackUrl string, confirmationsCount int32) *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem {
	this := NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem{}
	this.Address = address
	this.AllowDuplicates = allowDuplicates
	this.CallbackSecretKey = callbackSecretKey
	this.CallbackUrl = callbackUrl
	this.ConfirmationsCount = confirmationsCount
	return &this
}

// NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItemWithDefaults instantiates a new NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItemWithDefaults() *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem {
	this := NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem{}
	var allowDuplicates bool = false
	this.AllowDuplicates = allowDuplicates
	return &this
}

// GetAddress returns the Address field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetAddress(v string) {
	o.Address = v
}

// GetAllowDuplicates returns the AllowDuplicates field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAllowDuplicates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowDuplicates, true
}

// SetAllowDuplicates sets field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackSecretKey, true
}

// SetCallbackSecretKey sets field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetConfirmationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationsCount, true
}

// SetConfirmationsCount sets field value
func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = v
}

func (o NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["allowDuplicates"] = o.AllowDuplicates
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
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem struct {
	value *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem
	isSet bool
}

func (v NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) Get() *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem {
	return v.value
}

func (v *NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) Set(val *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem(val *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) *NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem {
	return &NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem{value: val, isSet: true}
}

func (v NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


