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

// ListTokensForwardingAutomationsRI struct for ListTokensForwardingAutomationsRI
type ListTokensForwardingAutomationsRI struct {
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount *int32 `json:"confirmationsCount,omitempty"`
	// Defines the specific time/date when the automatic forwarding was created in Unix Timestamp.
	CreatedTimestamp int32 `json:"createdTimestamp"`
	// Represents the specific fee address, which is always automatically generated. Users must fund it.
	FeeAddress string `json:"feeAddress"`
	// Represents the fee priority of the automation, whether it is \"SLOW\", \"STANDARD\" or \"FAST\".
	FeePriority string `json:"feePriority"`
	// Represents the hash of the address that forwards the tokens.
	FromAddress string `json:"fromAddress"`
	// Represents the minimum transfer amount of the tokens in the `fromAddress` that can be allowed for an automatic forwarding.
	MinimumTransferAmount string `json:"minimumTransferAmount"`
	// Represents a unique ID used to reference the specific callback subscription.
	ReferenceId string `json:"referenceId"`
	// Represents the hash of the address the tokens are forwarded to.
	ToAddress string `json:"toAddress"`
	TokenData ListTokensForwardingAutomationsRITS `json:"tokenData"`
}

// NewListTokensForwardingAutomationsRI instantiates a new ListTokensForwardingAutomationsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTokensForwardingAutomationsRI(callbackUrl string, createdTimestamp int32, feeAddress string, feePriority string, fromAddress string, minimumTransferAmount string, referenceId string, toAddress string, tokenData ListTokensForwardingAutomationsRITS) *ListTokensForwardingAutomationsRI {
	this := ListTokensForwardingAutomationsRI{}
	this.CallbackUrl = callbackUrl
	this.CreatedTimestamp = createdTimestamp
	this.FeeAddress = feeAddress
	this.FeePriority = feePriority
	this.FromAddress = fromAddress
	this.MinimumTransferAmount = minimumTransferAmount
	this.ReferenceId = referenceId
	this.ToAddress = toAddress
	this.TokenData = tokenData
	return &this
}

// NewListTokensForwardingAutomationsRIWithDefaults instantiates a new ListTokensForwardingAutomationsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTokensForwardingAutomationsRIWithDefaults() *ListTokensForwardingAutomationsRI {
	this := ListTokensForwardingAutomationsRI{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *ListTokensForwardingAutomationsRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *ListTokensForwardingAutomationsRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value if set, zero value otherwise.
func (o *ListTokensForwardingAutomationsRI) GetConfirmationsCount() int32 {
	if o == nil || o.ConfirmationsCount == nil {
		var ret int32
		return ret
	}
	return *o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil || o.ConfirmationsCount == nil {
		return nil, false
	}
	return o.ConfirmationsCount, true
}

// HasConfirmationsCount returns a boolean if a field has been set.
func (o *ListTokensForwardingAutomationsRI) HasConfirmationsCount() bool {
	if o != nil && o.ConfirmationsCount != nil {
		return true
	}

	return false
}

// SetConfirmationsCount gets a reference to the given int32 and assigns it to the ConfirmationsCount field.
func (o *ListTokensForwardingAutomationsRI) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *ListTokensForwardingAutomationsRI) GetCreatedTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *ListTokensForwardingAutomationsRI) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = v
}

// GetFeeAddress returns the FeeAddress field value
func (o *ListTokensForwardingAutomationsRI) GetFeeAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeAddress
}

// GetFeeAddressOk returns a tuple with the FeeAddress field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetFeeAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAddress, true
}

// SetFeeAddress sets field value
func (o *ListTokensForwardingAutomationsRI) SetFeeAddress(v string) {
	o.FeeAddress = v
}

// GetFeePriority returns the FeePriority field value
func (o *ListTokensForwardingAutomationsRI) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetFeePriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *ListTokensForwardingAutomationsRI) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetFromAddress returns the FromAddress field value
func (o *ListTokensForwardingAutomationsRI) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *ListTokensForwardingAutomationsRI) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetMinimumTransferAmount returns the MinimumTransferAmount field value
func (o *ListTokensForwardingAutomationsRI) GetMinimumTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTransferAmount
}

// GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetMinimumTransferAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumTransferAmount, true
}

// SetMinimumTransferAmount sets field value
func (o *ListTokensForwardingAutomationsRI) SetMinimumTransferAmount(v string) {
	o.MinimumTransferAmount = v
}

// GetReferenceId returns the ReferenceId field value
func (o *ListTokensForwardingAutomationsRI) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *ListTokensForwardingAutomationsRI) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetToAddress returns the ToAddress field value
func (o *ListTokensForwardingAutomationsRI) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *ListTokensForwardingAutomationsRI) SetToAddress(v string) {
	o.ToAddress = v
}

// GetTokenData returns the TokenData field value
func (o *ListTokensForwardingAutomationsRI) GetTokenData() ListTokensForwardingAutomationsRITS {
	if o == nil {
		var ret ListTokensForwardingAutomationsRITS
		return ret
	}

	return o.TokenData
}

// GetTokenDataOk returns a tuple with the TokenData field value
// and a boolean to check if the value has been set.
func (o *ListTokensForwardingAutomationsRI) GetTokenDataOk() (*ListTokensForwardingAutomationsRITS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenData, true
}

// SetTokenData sets field value
func (o *ListTokensForwardingAutomationsRI) SetTokenData(v ListTokensForwardingAutomationsRITS) {
	o.TokenData = v
}

func (o ListTokensForwardingAutomationsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if o.ConfirmationsCount != nil {
		toSerialize["confirmationsCount"] = o.ConfirmationsCount
	}
	if true {
		toSerialize["createdTimestamp"] = o.CreatedTimestamp
	}
	if true {
		toSerialize["feeAddress"] = o.FeeAddress
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
	if true {
		toSerialize["tokenData"] = o.TokenData
	}
	return json.Marshal(toSerialize)
}

type NullableListTokensForwardingAutomationsRI struct {
	value *ListTokensForwardingAutomationsRI
	isSet bool
}

func (v NullableListTokensForwardingAutomationsRI) Get() *ListTokensForwardingAutomationsRI {
	return v.value
}

func (v *NullableListTokensForwardingAutomationsRI) Set(val *ListTokensForwardingAutomationsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokensForwardingAutomationsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokensForwardingAutomationsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokensForwardingAutomationsRI(val *ListTokensForwardingAutomationsRI) *NullableListTokensForwardingAutomationsRI {
	return &NullableListTokensForwardingAutomationsRI{value: val, isSet: true}
}

func (v NullableListTokensForwardingAutomationsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokensForwardingAutomationsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


