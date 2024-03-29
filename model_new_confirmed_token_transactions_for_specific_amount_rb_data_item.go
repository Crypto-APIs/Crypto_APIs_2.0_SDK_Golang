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

// NewConfirmedTokenTransactionsForSpecificAmountRBDataItem struct for NewConfirmedTokenTransactionsForSpecificAmountRBDataItem
type NewConfirmedTokenTransactionsForSpecificAmountRBDataItem struct {
	// Specifies a flag that permits or denies the creation of duplicate addresses.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// Represents a specific amount of tokens after which the system have to send a callback to customers' callbackUrl.
	AmountHigherThan int64 `json:"amountHigherThan"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs 2.0. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the contract address of the token, which controls its logic. It is not the address that holds the tokens.
	ContractAddress string `json:"contractAddress"`
}

// NewNewConfirmedTokenTransactionsForSpecificAmountRBDataItem instantiates a new NewConfirmedTokenTransactionsForSpecificAmountRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedTokenTransactionsForSpecificAmountRBDataItem(amountHigherThan int64, callbackUrl string, contractAddress string) *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem {
	this := NewConfirmedTokenTransactionsForSpecificAmountRBDataItem{}
	var allowDuplicates bool = false
	this.AllowDuplicates = &allowDuplicates
	this.AmountHigherThan = amountHigherThan
	this.CallbackUrl = callbackUrl
	this.ContractAddress = contractAddress
	return &this
}

// NewNewConfirmedTokenTransactionsForSpecificAmountRBDataItemWithDefaults instantiates a new NewConfirmedTokenTransactionsForSpecificAmountRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedTokenTransactionsForSpecificAmountRBDataItemWithDefaults() *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem {
	this := NewConfirmedTokenTransactionsForSpecificAmountRBDataItem{}
	var allowDuplicates bool = false
	this.AllowDuplicates = &allowDuplicates
	return &this
}

// GetAllowDuplicates returns the AllowDuplicates field value if set, zero value otherwise.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetAllowDuplicates() bool {
	if o == nil || o.AllowDuplicates == nil {
		var ret bool
		return ret
	}
	return *o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil || o.AllowDuplicates == nil {
		return nil, false
	}
	return o.AllowDuplicates, true
}

// HasAllowDuplicates returns a boolean if a field has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) HasAllowDuplicates() bool {
	if o != nil && o.AllowDuplicates != nil {
		return true
	}

	return false
}

// SetAllowDuplicates gets a reference to the given bool and assigns it to the AllowDuplicates field.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = &v
}

// GetAmountHigherThan returns the AmountHigherThan field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetAmountHigherThan() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AmountHigherThan
}

// GetAmountHigherThanOk returns a tuple with the AmountHigherThan field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetAmountHigherThanOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountHigherThan, true
}

// SetAmountHigherThan sets field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) SetAmountHigherThan(v int64) {
	o.AmountHigherThan = v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetContractAddress returns the ContractAddress field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowDuplicates != nil {
		toSerialize["allowDuplicates"] = o.AllowDuplicates
	}
	if true {
		toSerialize["amountHigherThan"] = o.AmountHigherThan
	}
	if o.CallbackSecretKey != nil {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem struct {
	value *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem
	isSet bool
}

func (v NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem) Get() *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem {
	return v.value
}

func (v *NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem) Set(val *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem(val *NewConfirmedTokenTransactionsForSpecificAmountRBDataItem) *NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem {
	return &NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem{value: val, isSet: true}
}

func (v NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedTokenTransactionsForSpecificAmountRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


