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

// NewUnconfirmedCoinsTransactionsRequestBodyDataItem struct for NewUnconfirmedCoinsTransactionsRequestBodyDataItem
type NewUnconfirmedCoinsTransactionsRequestBodyDataItem struct {
	// Represents the address of the transaction, per which the result is returned.
	Address string `json:"address"`
	// Specifies a flag that permits or denies the creation of duplicate addresses.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs 2.0.
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl string `json:"callbackUrl"`
}

// NewNewUnconfirmedCoinsTransactionsRequestBodyDataItem instantiates a new NewUnconfirmedCoinsTransactionsRequestBodyDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUnconfirmedCoinsTransactionsRequestBodyDataItem(address string, callbackUrl string) *NewUnconfirmedCoinsTransactionsRequestBodyDataItem {
	this := NewUnconfirmedCoinsTransactionsRequestBodyDataItem{}
	this.Address = address
	var allowDuplicates bool = false
	this.AllowDuplicates = &allowDuplicates
	this.CallbackUrl = callbackUrl
	return &this
}

// NewNewUnconfirmedCoinsTransactionsRequestBodyDataItemWithDefaults instantiates a new NewUnconfirmedCoinsTransactionsRequestBodyDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUnconfirmedCoinsTransactionsRequestBodyDataItemWithDefaults() *NewUnconfirmedCoinsTransactionsRequestBodyDataItem {
	this := NewUnconfirmedCoinsTransactionsRequestBodyDataItem{}
	var allowDuplicates bool = false
	this.AllowDuplicates = &allowDuplicates
	return &this
}

// GetAddress returns the Address field value
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetAddress(v string) {
	o.Address = v
}

// GetAllowDuplicates returns the AllowDuplicates field value if set, zero value otherwise.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAllowDuplicates() bool {
	if o == nil || o.AllowDuplicates == nil {
		var ret bool
		return ret
	}
	return *o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil || o.AllowDuplicates == nil {
		return nil, false
	}
	return o.AllowDuplicates, true
}

// HasAllowDuplicates returns a boolean if a field has been set.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) HasAllowDuplicates() bool {
	if o != nil && o.AllowDuplicates != nil {
		return true
	}

	return false
}

// SetAllowDuplicates gets a reference to the given bool and assigns it to the AllowDuplicates field.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = &v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

func (o NewUnconfirmedCoinsTransactionsRequestBodyDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.AllowDuplicates != nil {
		toSerialize["allowDuplicates"] = o.AllowDuplicates
	}
	if o.CallbackSecretKey != nil {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	return json.Marshal(toSerialize)
}

type NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem struct {
	value *NewUnconfirmedCoinsTransactionsRequestBodyDataItem
	isSet bool
}

func (v NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem) Get() *NewUnconfirmedCoinsTransactionsRequestBodyDataItem {
	return v.value
}

func (v *NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem) Set(val *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem(val *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) *NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem {
	return &NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem{value: val, isSet: true}
}

func (v NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUnconfirmedCoinsTransactionsRequestBodyDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


