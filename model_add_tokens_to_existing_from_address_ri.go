/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// AddTokensToExistingFromAddressRI struct for AddTokensToExistingFromAddressRI
type AddTokensToExistingFromAddressRI struct {
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount int32 `json:"confirmationsCount"`
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
	TokenData AddTokensToExistingFromAddressRITS `json:"tokenData"`
}

// NewAddTokensToExistingFromAddressRI instantiates a new AddTokensToExistingFromAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTokensToExistingFromAddressRI(callbackUrl string, confirmationsCount int32, createdTimestamp int32, feeAddress string, feePriority string, fromAddress string, minimumTransferAmount string, referenceId string, toAddress string, tokenData AddTokensToExistingFromAddressRITS) *AddTokensToExistingFromAddressRI {
	this := AddTokensToExistingFromAddressRI{}
	this.CallbackUrl = callbackUrl
	this.ConfirmationsCount = confirmationsCount
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

// NewAddTokensToExistingFromAddressRIWithDefaults instantiates a new AddTokensToExistingFromAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTokensToExistingFromAddressRIWithDefaults() *AddTokensToExistingFromAddressRI {
	this := AddTokensToExistingFromAddressRI{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *AddTokensToExistingFromAddressRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *AddTokensToExistingFromAddressRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value
func (o *AddTokensToExistingFromAddressRI) GetConfirmationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationsCount, true
}

// SetConfirmationsCount sets field value
func (o *AddTokensToExistingFromAddressRI) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *AddTokensToExistingFromAddressRI) GetCreatedTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *AddTokensToExistingFromAddressRI) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = v
}

// GetFeeAddress returns the FeeAddress field value
func (o *AddTokensToExistingFromAddressRI) GetFeeAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeAddress
}

// GetFeeAddressOk returns a tuple with the FeeAddress field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetFeeAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAddress, true
}

// SetFeeAddress sets field value
func (o *AddTokensToExistingFromAddressRI) SetFeeAddress(v string) {
	o.FeeAddress = v
}

// GetFeePriority returns the FeePriority field value
func (o *AddTokensToExistingFromAddressRI) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetFeePriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *AddTokensToExistingFromAddressRI) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetFromAddress returns the FromAddress field value
func (o *AddTokensToExistingFromAddressRI) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *AddTokensToExistingFromAddressRI) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetMinimumTransferAmount returns the MinimumTransferAmount field value
func (o *AddTokensToExistingFromAddressRI) GetMinimumTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTransferAmount
}

// GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetMinimumTransferAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumTransferAmount, true
}

// SetMinimumTransferAmount sets field value
func (o *AddTokensToExistingFromAddressRI) SetMinimumTransferAmount(v string) {
	o.MinimumTransferAmount = v
}

// GetReferenceId returns the ReferenceId field value
func (o *AddTokensToExistingFromAddressRI) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *AddTokensToExistingFromAddressRI) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetToAddress returns the ToAddress field value
func (o *AddTokensToExistingFromAddressRI) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *AddTokensToExistingFromAddressRI) SetToAddress(v string) {
	o.ToAddress = v
}

// GetTokenData returns the TokenData field value
func (o *AddTokensToExistingFromAddressRI) GetTokenData() AddTokensToExistingFromAddressRITS {
	if o == nil {
		var ret AddTokensToExistingFromAddressRITS
		return ret
	}

	return o.TokenData
}

// GetTokenDataOk returns a tuple with the TokenData field value
// and a boolean to check if the value has been set.
func (o *AddTokensToExistingFromAddressRI) GetTokenDataOk() (*AddTokensToExistingFromAddressRITS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenData, true
}

// SetTokenData sets field value
func (o *AddTokensToExistingFromAddressRI) SetTokenData(v AddTokensToExistingFromAddressRITS) {
	o.TokenData = v
}

func (o AddTokensToExistingFromAddressRI) MarshalJSON() ([]byte, error) {
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

type NullableAddTokensToExistingFromAddressRI struct {
	value *AddTokensToExistingFromAddressRI
	isSet bool
}

func (v NullableAddTokensToExistingFromAddressRI) Get() *AddTokensToExistingFromAddressRI {
	return v.value
}

func (v *NullableAddTokensToExistingFromAddressRI) Set(val *AddTokensToExistingFromAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTokensToExistingFromAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTokensToExistingFromAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTokensToExistingFromAddressRI(val *AddTokensToExistingFromAddressRI) *NullableAddTokensToExistingFromAddressRI {
	return &NullableAddTokensToExistingFromAddressRI{value: val, isSet: true}
}

func (v NullableAddTokensToExistingFromAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTokensToExistingFromAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


