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

// AddTokensToExistingFromAddressRequestBodyDataItem struct for AddTokensToExistingFromAddressRequestBodyDataItem
type AddTokensToExistingFromAddressRequestBodyDataItem struct {
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs.
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount int32 `json:"confirmationsCount"`
	// Represents the fee priority of the automation, whether it is \"SLOW\", \"STANDARD\" or \"FAST\".
	FeePriority string `json:"feePriority"`
	// Represents the hash of the address that forwards the tokens.
	FromAddress string `json:"fromAddress"`
	// Represents the minimum transfer amount of the currency in the `fromAddress` that can be allowed for an automatic forwarding.
	MinimumTransferAmount string `json:"minimumTransferAmount"`
	// Represents the hash of the address the currency is forwarded to.
	ToAddress string `json:"toAddress"`
	TokenData AddTokensToExistingFromAddressRequestBodyTokenData `json:"tokenData"`
}

// NewAddTokensToExistingFromAddressRequestBodyDataItem instantiates a new AddTokensToExistingFromAddressRequestBodyDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTokensToExistingFromAddressRequestBodyDataItem(callbackUrl string, confirmationsCount int32, feePriority string, fromAddress string, minimumTransferAmount string, toAddress string, tokenData AddTokensToExistingFromAddressRequestBodyTokenData) *AddTokensToExistingFromAddressRequestBodyDataItem {
	this := AddTokensToExistingFromAddressRequestBodyDataItem{}
	this.CallbackUrl = callbackUrl
	this.ConfirmationsCount = confirmationsCount
	this.FeePriority = feePriority
	this.FromAddress = fromAddress
	this.MinimumTransferAmount = minimumTransferAmount
	this.ToAddress = toAddress
	this.TokenData = tokenData
	return &this
}

// NewAddTokensToExistingFromAddressRequestBodyDataItemWithDefaults instantiates a new AddTokensToExistingFromAddressRequestBodyDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTokensToExistingFromAddressRequestBodyDataItemWithDefaults() *AddTokensToExistingFromAddressRequestBodyDataItem {
	this := AddTokensToExistingFromAddressRequestBodyDataItem{}
	return &this
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetConfirmationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmationsCount, true
}

// SetConfirmationsCount sets field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = v
}

// GetFeePriority returns the FeePriority field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetFeePriorityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetFromAddress returns the FromAddress field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetFromAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetMinimumTransferAmount returns the MinimumTransferAmount field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetMinimumTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTransferAmount
}

// GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetMinimumTransferAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinimumTransferAmount, true
}

// SetMinimumTransferAmount sets field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetMinimumTransferAmount(v string) {
	o.MinimumTransferAmount = v
}

// GetToAddress returns the ToAddress field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetToAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetToAddress(v string) {
	o.ToAddress = v
}

// GetTokenData returns the TokenData field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetTokenData() AddTokensToExistingFromAddressRequestBodyTokenData {
	if o == nil {
		var ret AddTokensToExistingFromAddressRequestBodyTokenData
		return ret
	}

	return o.TokenData
}

// GetTokenDataOk returns a tuple with the TokenData field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) GetTokenDataOk() (*AddTokensToExistingFromAddressRequestBodyTokenData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenData, true
}

// SetTokenData sets field value
func (o *AddTokensToExistingFromAddressRequestBodyDataItem) SetTokenData(v AddTokensToExistingFromAddressRequestBodyTokenData) {
	o.TokenData = v
}

func (o AddTokensToExistingFromAddressRequestBodyDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackSecretKey != nil {
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
		toSerialize["fromAddress"] = o.FromAddress
	}
	if true {
		toSerialize["minimumTransferAmount"] = o.MinimumTransferAmount
	}
	if true {
		toSerialize["toAddress"] = o.ToAddress
	}
	if true {
		toSerialize["tokenData"] = o.TokenData
	}
	return json.Marshal(toSerialize)
}

type NullableAddTokensToExistingFromAddressRequestBodyDataItem struct {
	value *AddTokensToExistingFromAddressRequestBodyDataItem
	isSet bool
}

func (v NullableAddTokensToExistingFromAddressRequestBodyDataItem) Get() *AddTokensToExistingFromAddressRequestBodyDataItem {
	return v.value
}

func (v *NullableAddTokensToExistingFromAddressRequestBodyDataItem) Set(val *AddTokensToExistingFromAddressRequestBodyDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTokensToExistingFromAddressRequestBodyDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTokensToExistingFromAddressRequestBodyDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTokensToExistingFromAddressRequestBodyDataItem(val *AddTokensToExistingFromAddressRequestBodyDataItem) *NullableAddTokensToExistingFromAddressRequestBodyDataItem {
	return &NullableAddTokensToExistingFromAddressRequestBodyDataItem{value: val, isSet: true}
}

func (v NullableAddTokensToExistingFromAddressRequestBodyDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTokensToExistingFromAddressRequestBodyDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


