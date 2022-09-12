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

// CreateFungibleTokensTransactionRequestFromAddressRI struct for CreateFungibleTokensTransactionRequestFromAddressRI
type CreateFungibleTokensTransactionRequestFromAddressRI struct {
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey string `json:"callbackSecretKey"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the fee priority of the automation, whether it is \"slow\", \"standard\" or \"fast\".
	FeePriority string `json:"feePriority"`
	// Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request.
	Note *string `json:"note,omitempty"`
	// Defines the destination for the transaction, i.e. the recipient(s).
	Recipients []CreateFungibleTokensTransactionRequestFromAddressRIRecipientsInner `json:"recipients"`
	Senders CreateFungibleTokensTransactionRequestFromAddressRISenders `json:"senders"`
	TokenTypeSpecificData CreateFungibleTokensTransactionRequestFromAddressRIS `json:"tokenTypeSpecificData"`
	// Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which `referenceId` concern that specific transaction request.
	TransactionRequestId string `json:"transactionRequestId"`
}

// NewCreateFungibleTokensTransactionRequestFromAddressRI instantiates a new CreateFungibleTokensTransactionRequestFromAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFungibleTokensTransactionRequestFromAddressRI(callbackSecretKey string, callbackUrl string, feePriority string, recipients []CreateFungibleTokensTransactionRequestFromAddressRIRecipientsInner, senders CreateFungibleTokensTransactionRequestFromAddressRISenders, tokenTypeSpecificData CreateFungibleTokensTransactionRequestFromAddressRIS, transactionRequestId string) *CreateFungibleTokensTransactionRequestFromAddressRI {
	this := CreateFungibleTokensTransactionRequestFromAddressRI{}
	this.CallbackSecretKey = callbackSecretKey
	this.CallbackUrl = callbackUrl
	this.FeePriority = feePriority
	this.Recipients = recipients
	this.Senders = senders
	this.TokenTypeSpecificData = tokenTypeSpecificData
	this.TransactionRequestId = transactionRequestId
	return &this
}

// NewCreateFungibleTokensTransactionRequestFromAddressRIWithDefaults instantiates a new CreateFungibleTokensTransactionRequestFromAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFungibleTokensTransactionRequestFromAddressRIWithDefaults() *CreateFungibleTokensTransactionRequestFromAddressRI {
	this := CreateFungibleTokensTransactionRequestFromAddressRI{}
	return &this
}

// GetCallbackSecretKey returns the CallbackSecretKey field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackSecretKey, true
}

// SetCallbackSecretKey sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetFeePriority returns the FeePriority field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetFeePriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetNote(v string) {
	o.Note = &v
}

// GetRecipients returns the Recipients field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetRecipients() []CreateFungibleTokensTransactionRequestFromAddressRIRecipientsInner {
	if o == nil {
		var ret []CreateFungibleTokensTransactionRequestFromAddressRIRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetRecipientsOk() ([]CreateFungibleTokensTransactionRequestFromAddressRIRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetRecipients(v []CreateFungibleTokensTransactionRequestFromAddressRIRecipientsInner) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetSenders() CreateFungibleTokensTransactionRequestFromAddressRISenders {
	if o == nil {
		var ret CreateFungibleTokensTransactionRequestFromAddressRISenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetSendersOk() (*CreateFungibleTokensTransactionRequestFromAddressRISenders, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetSenders(v CreateFungibleTokensTransactionRequestFromAddressRISenders) {
	o.Senders = v
}

// GetTokenTypeSpecificData returns the TokenTypeSpecificData field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificData() CreateFungibleTokensTransactionRequestFromAddressRIS {
	if o == nil {
		var ret CreateFungibleTokensTransactionRequestFromAddressRIS
		return ret
	}

	return o.TokenTypeSpecificData
}

// GetTokenTypeSpecificDataOk returns a tuple with the TokenTypeSpecificData field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificDataOk() (*CreateFungibleTokensTransactionRequestFromAddressRIS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTypeSpecificData, true
}

// SetTokenTypeSpecificData sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetTokenTypeSpecificData(v CreateFungibleTokensTransactionRequestFromAddressRIS) {
	o.TokenTypeSpecificData = v
}

// GetTransactionRequestId returns the TransactionRequestId field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTransactionRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionRequestId
}

// GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTransactionRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionRequestId, true
}

// SetTransactionRequestId sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetTransactionRequestId(v string) {
	o.TransactionRequestId = v
}

func (o CreateFungibleTokensTransactionRequestFromAddressRI) MarshalJSON() ([]byte, error) {
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
	if o.Note != nil {
		toSerialize["note"] = o.Note
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
	if true {
		toSerialize["transactionRequestId"] = o.TransactionRequestId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFungibleTokensTransactionRequestFromAddressRI struct {
	value *CreateFungibleTokensTransactionRequestFromAddressRI
	isSet bool
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRI) Get() *CreateFungibleTokensTransactionRequestFromAddressRI {
	return v.value
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRI) Set(val *CreateFungibleTokensTransactionRequestFromAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokensTransactionRequestFromAddressRI(val *CreateFungibleTokensTransactionRequestFromAddressRI) *NullableCreateFungibleTokensTransactionRequestFromAddressRI {
	return &NullableCreateFungibleTokensTransactionRequestFromAddressRI{value: val, isSet: true}
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


