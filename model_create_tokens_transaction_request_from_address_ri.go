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

// CreateTokensTransactionRequestFromAddressRI struct for CreateTokensTransactionRequestFromAddressRI
type CreateTokensTransactionRequestFromAddressRI struct {
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs.
	CallbackSecretKey string `json:"callbackSecretKey"`
	// Verified URL for sending callbacks
	CallbackUrl string `json:"callbackUrl"`
	// Represents the fee priority of the automation, whether it is \"slow\", \"standard\" or \"fast\".
	FeePriority string `json:"feePriority"`
	// Defines the destination for the transaction, i.e. the recipient(s).
	Recipients []CreateTokensTransactionRequestFromAddressRIRecipients `json:"recipients"`
	Senders CreateTokensTransactionRequestFromAddressRISenders `json:"senders"`
	TokenTypeSpecificData CreateTokensTransactionRequestFromAddressRIS `json:"tokenTypeSpecificData"`
}

// NewCreateTokensTransactionRequestFromAddressRI instantiates a new CreateTokensTransactionRequestFromAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokensTransactionRequestFromAddressRI(callbackSecretKey string, callbackUrl string, feePriority string, recipients []CreateTokensTransactionRequestFromAddressRIRecipients, senders CreateTokensTransactionRequestFromAddressRISenders, tokenTypeSpecificData CreateTokensTransactionRequestFromAddressRIS) *CreateTokensTransactionRequestFromAddressRI {
	this := CreateTokensTransactionRequestFromAddressRI{}
	this.CallbackSecretKey = callbackSecretKey
	this.CallbackUrl = callbackUrl
	this.FeePriority = feePriority
	this.Recipients = recipients
	this.Senders = senders
	this.TokenTypeSpecificData = tokenTypeSpecificData
	return &this
}

// NewCreateTokensTransactionRequestFromAddressRIWithDefaults instantiates a new CreateTokensTransactionRequestFromAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokensTransactionRequestFromAddressRIWithDefaults() *CreateTokensTransactionRequestFromAddressRI {
	this := CreateTokensTransactionRequestFromAddressRI{}
	return &this
}

// GetCallbackSecretKey returns the CallbackSecretKey field value
func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackSecretKey, true
}

// SetCallbackSecretKey sets field value
func (o *CreateTokensTransactionRequestFromAddressRI) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *CreateTokensTransactionRequestFromAddressRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetFeePriority returns the FeePriority field value
func (o *CreateTokensTransactionRequestFromAddressRI) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRI) GetFeePriorityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *CreateTokensTransactionRequestFromAddressRI) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetRecipients returns the Recipients field value
func (o *CreateTokensTransactionRequestFromAddressRI) GetRecipients() []CreateTokensTransactionRequestFromAddressRIRecipients {
	if o == nil {
		var ret []CreateTokensTransactionRequestFromAddressRIRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRI) GetRecipientsOk() (*[]CreateTokensTransactionRequestFromAddressRIRecipients, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *CreateTokensTransactionRequestFromAddressRI) SetRecipients(v []CreateTokensTransactionRequestFromAddressRIRecipients) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *CreateTokensTransactionRequestFromAddressRI) GetSenders() CreateTokensTransactionRequestFromAddressRISenders {
	if o == nil {
		var ret CreateTokensTransactionRequestFromAddressRISenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRI) GetSendersOk() (*CreateTokensTransactionRequestFromAddressRISenders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *CreateTokensTransactionRequestFromAddressRI) SetSenders(v CreateTokensTransactionRequestFromAddressRISenders) {
	o.Senders = v
}

// GetTokenTypeSpecificData returns the TokenTypeSpecificData field value
func (o *CreateTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificData() CreateTokensTransactionRequestFromAddressRIS {
	if o == nil {
		var ret CreateTokensTransactionRequestFromAddressRIS
		return ret
	}

	return o.TokenTypeSpecificData
}

// GetTokenTypeSpecificDataOk returns a tuple with the TokenTypeSpecificData field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificDataOk() (*CreateTokensTransactionRequestFromAddressRIS, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenTypeSpecificData, true
}

// SetTokenTypeSpecificData sets field value
func (o *CreateTokensTransactionRequestFromAddressRI) SetTokenTypeSpecificData(v CreateTokensTransactionRequestFromAddressRIS) {
	o.TokenTypeSpecificData = v
}

func (o CreateTokensTransactionRequestFromAddressRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if true {
		toSerialize["feePriority"] = o.FeePriority
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["tokenTypeSpecificData"] = o.TokenTypeSpecificData
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTokensTransactionRequestFromAddressRI struct {
	value *CreateTokensTransactionRequestFromAddressRI
	isSet bool
}

func (v NullableCreateTokensTransactionRequestFromAddressRI) Get() *CreateTokensTransactionRequestFromAddressRI {
	return v.value
}

func (v *NullableCreateTokensTransactionRequestFromAddressRI) Set(val *CreateTokensTransactionRequestFromAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokensTransactionRequestFromAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokensTransactionRequestFromAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokensTransactionRequestFromAddressRI(val *CreateTokensTransactionRequestFromAddressRI) *NullableCreateTokensTransactionRequestFromAddressRI {
	return &NullableCreateTokensTransactionRequestFromAddressRI{value: val, isSet: true}
}

func (v NullableCreateTokensTransactionRequestFromAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokensTransactionRequestFromAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


