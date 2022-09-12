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

// CreateCoinsTransactionRequestFromAddressRI struct for CreateCoinsTransactionRequestFromAddressRI
type CreateCoinsTransactionRequestFromAddressRI struct {
	// Defines a specific Tag that is an additional XRP address feature. It helps identify a transaction recipient beyond a wallet address. The tag that was encoded into the x-Address along with the Source Classic Address.
	AddressTag *int32 `json:"addressTag,omitempty"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. `We support ONLY httpS type of protocol`.
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// Represents the public address, which is a compressed and shortened form of a public key. The classic address is shown when the source address is an x-Address.
	ClassicAddress *string `json:"classicAddress,omitempty"`
	// Represents the fee priority of the automation, whether it is \"slow\", \"standard\" or \"fast\".
	FeePriority string `json:"feePriority"`
	// Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request.
	Note *string `json:"note,omitempty"`
	// Defines the destination for the transaction, i.e. the recipient(s).
	Recipients []CreateCoinsTransactionRequestFromAddressRIRecipientsInner `json:"recipients"`
	Senders CreateCoinsTransactionRequestFromAddressRISenders `json:"senders"`
	// Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which `referenceId` concern that specific transaction request.
	TransactionRequestId string `json:"transactionRequestId"`
	// Defines the status of the transaction request, e.g. \"created, \"await_approval\", \"pending\", \"prepared\", \"signed\", \"broadcasted\", \"success\", \"failed\", \"rejected\", mined\".
	TransactionRequestStatus string `json:"transactionRequestStatus"`
}

// NewCreateCoinsTransactionRequestFromAddressRI instantiates a new CreateCoinsTransactionRequestFromAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCoinsTransactionRequestFromAddressRI(feePriority string, recipients []CreateCoinsTransactionRequestFromAddressRIRecipientsInner, senders CreateCoinsTransactionRequestFromAddressRISenders, transactionRequestId string, transactionRequestStatus string) *CreateCoinsTransactionRequestFromAddressRI {
	this := CreateCoinsTransactionRequestFromAddressRI{}
	this.FeePriority = feePriority
	this.Recipients = recipients
	this.Senders = senders
	this.TransactionRequestId = transactionRequestId
	this.TransactionRequestStatus = transactionRequestStatus
	return &this
}

// NewCreateCoinsTransactionRequestFromAddressRIWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCoinsTransactionRequestFromAddressRIWithDefaults() *CreateCoinsTransactionRequestFromAddressRI {
	this := CreateCoinsTransactionRequestFromAddressRI{}
	return &this
}

// GetAddressTag returns the AddressTag field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetAddressTag() int32 {
	if o == nil || o.AddressTag == nil {
		var ret int32
		return ret
	}
	return *o.AddressTag
}

// GetAddressTagOk returns a tuple with the AddressTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetAddressTagOk() (*int32, bool) {
	if o == nil || o.AddressTag == nil {
		return nil, false
	}
	return o.AddressTag, true
}

// HasAddressTag returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) HasAddressTag() bool {
	if o != nil && o.AddressTag != nil {
		return true
	}

	return false
}

// SetAddressTag gets a reference to the given int32 and assigns it to the AddressTag field.
func (o *CreateCoinsTransactionRequestFromAddressRI) SetAddressTag(v int32) {
	o.AddressTag = &v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *CreateCoinsTransactionRequestFromAddressRI) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *CreateCoinsTransactionRequestFromAddressRI) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetClassicAddress returns the ClassicAddress field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetClassicAddress() string {
	if o == nil || o.ClassicAddress == nil {
		var ret string
		return ret
	}
	return *o.ClassicAddress
}

// GetClassicAddressOk returns a tuple with the ClassicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetClassicAddressOk() (*string, bool) {
	if o == nil || o.ClassicAddress == nil {
		return nil, false
	}
	return o.ClassicAddress, true
}

// HasClassicAddress returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) HasClassicAddress() bool {
	if o != nil && o.ClassicAddress != nil {
		return true
	}

	return false
}

// SetClassicAddress gets a reference to the given string and assigns it to the ClassicAddress field.
func (o *CreateCoinsTransactionRequestFromAddressRI) SetClassicAddress(v string) {
	o.ClassicAddress = &v
}

// GetFeePriority returns the FeePriority field value
func (o *CreateCoinsTransactionRequestFromAddressRI) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetFeePriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *CreateCoinsTransactionRequestFromAddressRI) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *CreateCoinsTransactionRequestFromAddressRI) SetNote(v string) {
	o.Note = &v
}

// GetRecipients returns the Recipients field value
func (o *CreateCoinsTransactionRequestFromAddressRI) GetRecipients() []CreateCoinsTransactionRequestFromAddressRIRecipientsInner {
	if o == nil {
		var ret []CreateCoinsTransactionRequestFromAddressRIRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetRecipientsOk() ([]CreateCoinsTransactionRequestFromAddressRIRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *CreateCoinsTransactionRequestFromAddressRI) SetRecipients(v []CreateCoinsTransactionRequestFromAddressRIRecipientsInner) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *CreateCoinsTransactionRequestFromAddressRI) GetSenders() CreateCoinsTransactionRequestFromAddressRISenders {
	if o == nil {
		var ret CreateCoinsTransactionRequestFromAddressRISenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetSendersOk() (*CreateCoinsTransactionRequestFromAddressRISenders, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *CreateCoinsTransactionRequestFromAddressRI) SetSenders(v CreateCoinsTransactionRequestFromAddressRISenders) {
	o.Senders = v
}

// GetTransactionRequestId returns the TransactionRequestId field value
func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionRequestId
}

// GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionRequestId, true
}

// SetTransactionRequestId sets field value
func (o *CreateCoinsTransactionRequestFromAddressRI) SetTransactionRequestId(v string) {
	o.TransactionRequestId = v
}

// GetTransactionRequestStatus returns the TransactionRequestStatus field value
func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionRequestStatus
}

// GetTransactionRequestStatusOk returns a tuple with the TransactionRequestStatus field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionRequestStatus, true
}

// SetTransactionRequestStatus sets field value
func (o *CreateCoinsTransactionRequestFromAddressRI) SetTransactionRequestStatus(v string) {
	o.TransactionRequestStatus = v
}

func (o CreateCoinsTransactionRequestFromAddressRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressTag != nil {
		toSerialize["addressTag"] = o.AddressTag
	}
	if o.CallbackSecretKey != nil {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if o.CallbackUrl != nil {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if o.ClassicAddress != nil {
		toSerialize["classicAddress"] = o.ClassicAddress
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
		toSerialize["transactionRequestId"] = o.TransactionRequestId
	}
	if true {
		toSerialize["transactionRequestStatus"] = o.TransactionRequestStatus
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCoinsTransactionRequestFromAddressRI struct {
	value *CreateCoinsTransactionRequestFromAddressRI
	isSet bool
}

func (v NullableCreateCoinsTransactionRequestFromAddressRI) Get() *CreateCoinsTransactionRequestFromAddressRI {
	return v.value
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRI) Set(val *CreateCoinsTransactionRequestFromAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionRequestFromAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionRequestFromAddressRI(val *CreateCoinsTransactionRequestFromAddressRI) *NullableCreateCoinsTransactionRequestFromAddressRI {
	return &NullableCreateCoinsTransactionRequestFromAddressRI{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionRequestFromAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


