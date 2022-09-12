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

// ListXRPRippleTransactionsByBlockHeightRI struct for ListXRPRippleTransactionsByBlockHeightRI
type ListXRPRippleTransactionsByBlockHeightRI struct {
	AdditionalData *string `json:"additionalData,omitempty"`
	DestinationTag *int64 `json:"destinationTag,omitempty"`
	Index int32 `json:"index"`
	MinedInBlockHash string `json:"minedInBlockHash"`
	// Object Array representation of transaction receivers
	Recipients []ListXRPRippleTransactionsByBlockHeightRIRecipientsInner `json:"recipients"`
	// Object Array representation of transaction senders
	Senders []ListXRPRippleTransactionsByBlockHeightRISendersInner `json:"senders"`
	Sequence int64 `json:"sequence"`
	Status string `json:"status"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	TransactionHash string `json:"transactionHash"`
	Type string `json:"type"`
	Fee ListXRPRippleTransactionsByBlockHeightRIFee `json:"fee"`
	Offer ListXRPRippleTransactionsByBlockHeightRIOffer `json:"offer"`
	Receive ListXRPRippleTransactionsByBlockHeightRIReceive `json:"receive"`
	Value ListXRPRippleTransactionsByBlockHeightRIValue `json:"value"`
}

// NewListXRPRippleTransactionsByBlockHeightRI instantiates a new ListXRPRippleTransactionsByBlockHeightRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByBlockHeightRI(index int32, minedInBlockHash string, recipients []ListXRPRippleTransactionsByBlockHeightRIRecipientsInner, senders []ListXRPRippleTransactionsByBlockHeightRISendersInner, sequence int64, status string, timestamp int32, transactionHash string, type_ string, fee ListXRPRippleTransactionsByBlockHeightRIFee, offer ListXRPRippleTransactionsByBlockHeightRIOffer, receive ListXRPRippleTransactionsByBlockHeightRIReceive, value ListXRPRippleTransactionsByBlockHeightRIValue) *ListXRPRippleTransactionsByBlockHeightRI {
	this := ListXRPRippleTransactionsByBlockHeightRI{}
	this.Index = index
	this.MinedInBlockHash = minedInBlockHash
	this.Recipients = recipients
	this.Senders = senders
	this.Sequence = sequence
	this.Status = status
	this.Timestamp = timestamp
	this.TransactionHash = transactionHash
	this.Type = type_
	this.Fee = fee
	this.Offer = offer
	this.Receive = receive
	this.Value = value
	return &this
}

// NewListXRPRippleTransactionsByBlockHeightRIWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHeightRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByBlockHeightRIWithDefaults() *ListXRPRippleTransactionsByBlockHeightRI {
	this := ListXRPRippleTransactionsByBlockHeightRI{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetAdditionalData() string {
	if o == nil || o.AdditionalData == nil {
		var ret string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetAdditionalDataOk() (*string, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given string and assigns it to the AdditionalData field.
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetAdditionalData(v string) {
	o.AdditionalData = &v
}

// GetDestinationTag returns the DestinationTag field value if set, zero value otherwise.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetDestinationTag() int64 {
	if o == nil || o.DestinationTag == nil {
		var ret int64
		return ret
	}
	return *o.DestinationTag
}

// GetDestinationTagOk returns a tuple with the DestinationTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetDestinationTagOk() (*int64, bool) {
	if o == nil || o.DestinationTag == nil {
		return nil, false
	}
	return o.DestinationTag, true
}

// HasDestinationTag returns a boolean if a field has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) HasDestinationTag() bool {
	if o != nil && o.DestinationTag != nil {
		return true
	}

	return false
}

// SetDestinationTag gets a reference to the given int64 and assigns it to the DestinationTag field.
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetDestinationTag(v int64) {
	o.DestinationTag = &v
}

// GetIndex returns the Index field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetIndex(v int32) {
	o.Index = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetMinedInBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlockHash, true
}

// SetMinedInBlockHash sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = v
}

// GetRecipients returns the Recipients field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetRecipients() []ListXRPRippleTransactionsByBlockHeightRIRecipientsInner {
	if o == nil {
		var ret []ListXRPRippleTransactionsByBlockHeightRIRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetRecipientsOk() ([]ListXRPRippleTransactionsByBlockHeightRIRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetRecipients(v []ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSenders() []ListXRPRippleTransactionsByBlockHeightRISendersInner {
	if o == nil {
		var ret []ListXRPRippleTransactionsByBlockHeightRISendersInner
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSendersOk() ([]ListXRPRippleTransactionsByBlockHeightRISendersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Senders, true
}

// SetSenders sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetSenders(v []ListXRPRippleTransactionsByBlockHeightRISendersInner) {
	o.Senders = v
}

// GetSequence returns the Sequence field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSequence() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSequenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetSequence(v int64) {
	o.Sequence = v
}

// GetStatus returns the Status field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetType returns the Type field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetType(v string) {
	o.Type = v
}

// GetFee returns the Fee field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetFee() ListXRPRippleTransactionsByBlockHeightRIFee {
	if o == nil {
		var ret ListXRPRippleTransactionsByBlockHeightRIFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetFeeOk() (*ListXRPRippleTransactionsByBlockHeightRIFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetFee(v ListXRPRippleTransactionsByBlockHeightRIFee) {
	o.Fee = v
}

// GetOffer returns the Offer field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetOffer() ListXRPRippleTransactionsByBlockHeightRIOffer {
	if o == nil {
		var ret ListXRPRippleTransactionsByBlockHeightRIOffer
		return ret
	}

	return o.Offer
}

// GetOfferOk returns a tuple with the Offer field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetOfferOk() (*ListXRPRippleTransactionsByBlockHeightRIOffer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offer, true
}

// SetOffer sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetOffer(v ListXRPRippleTransactionsByBlockHeightRIOffer) {
	o.Offer = v
}

// GetReceive returns the Receive field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetReceive() ListXRPRippleTransactionsByBlockHeightRIReceive {
	if o == nil {
		var ret ListXRPRippleTransactionsByBlockHeightRIReceive
		return ret
	}

	return o.Receive
}

// GetReceiveOk returns a tuple with the Receive field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetReceiveOk() (*ListXRPRippleTransactionsByBlockHeightRIReceive, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Receive, true
}

// SetReceive sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetReceive(v ListXRPRippleTransactionsByBlockHeightRIReceive) {
	o.Receive = v
}

// GetValue returns the Value field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetValue() ListXRPRippleTransactionsByBlockHeightRIValue {
	if o == nil {
		var ret ListXRPRippleTransactionsByBlockHeightRIValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRI) GetValueOk() (*ListXRPRippleTransactionsByBlockHeightRIValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRI) SetValue(v ListXRPRippleTransactionsByBlockHeightRIValue) {
	o.Value = v
}

func (o ListXRPRippleTransactionsByBlockHeightRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if o.DestinationTag != nil {
		toSerialize["destinationTag"] = o.DestinationTag
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["minedInBlockHash"] = o.MinedInBlockHash
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["sequence"] = o.Sequence
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["offer"] = o.Offer
	}
	if true {
		toSerialize["receive"] = o.Receive
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableListXRPRippleTransactionsByBlockHeightRI struct {
	value *ListXRPRippleTransactionsByBlockHeightRI
	isSet bool
}

func (v NullableListXRPRippleTransactionsByBlockHeightRI) Get() *ListXRPRippleTransactionsByBlockHeightRI {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRI) Set(val *ListXRPRippleTransactionsByBlockHeightRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByBlockHeightRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByBlockHeightRI(val *ListXRPRippleTransactionsByBlockHeightRI) *NullableListXRPRippleTransactionsByBlockHeightRI {
	return &NullableListXRPRippleTransactionsByBlockHeightRI{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByBlockHeightRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


