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

// ListWalletTransactionsRI struct for ListWalletTransactionsRI
type ListWalletTransactionsRI struct {
	// Defines the direction of the transaction, e.g. incoming.
	Direction string `json:"direction"`
	Fee ListWalletTransactionsRIFee `json:"fee"`
	// Represents fungible tokens'es detailed information
	FungibleTokens []ListWalletTransactionsRIFungibleTokensInner `json:"fungibleTokens,omitempty"`
	InternalTransactions []ListWalletTransactionsRIInternalTransactionsInner `json:"internalTransactions,omitempty"`
	// Represents non-fungible tokens'es detailed information.
	NonFungibleTokens []ListWalletTransactionsRINonFungibleTokensInner `json:"nonFungibleTokens,omitempty"`
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients []ListWalletTransactionsRIRecipientsInner `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders []ListWalletTransactionsRISendersInner `json:"senders"`
	// Defines the status of the transaction, if it is confirmed or unconfirmed.
	Status string `json:"status"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the unique TD of the transaction.
	TransactionId string `json:"transactionId"`
	Value ListWalletTransactionsRIValue `json:"value"`
}

// NewListWalletTransactionsRI instantiates a new ListWalletTransactionsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWalletTransactionsRI(direction string, fee ListWalletTransactionsRIFee, recipients []ListWalletTransactionsRIRecipientsInner, senders []ListWalletTransactionsRISendersInner, status string, timestamp int32, transactionId string, value ListWalletTransactionsRIValue) *ListWalletTransactionsRI {
	this := ListWalletTransactionsRI{}
	this.Direction = direction
	this.Fee = fee
	this.Recipients = recipients
	this.Senders = senders
	this.Status = status
	this.Timestamp = timestamp
	this.TransactionId = transactionId
	this.Value = value
	return &this
}

// NewListWalletTransactionsRIWithDefaults instantiates a new ListWalletTransactionsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWalletTransactionsRIWithDefaults() *ListWalletTransactionsRI {
	this := ListWalletTransactionsRI{}
	return &this
}

// GetDirection returns the Direction field value
func (o *ListWalletTransactionsRI) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *ListWalletTransactionsRI) SetDirection(v string) {
	o.Direction = v
}

// GetFee returns the Fee field value
func (o *ListWalletTransactionsRI) GetFee() ListWalletTransactionsRIFee {
	if o == nil {
		var ret ListWalletTransactionsRIFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetFeeOk() (*ListWalletTransactionsRIFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListWalletTransactionsRI) SetFee(v ListWalletTransactionsRIFee) {
	o.Fee = v
}

// GetFungibleTokens returns the FungibleTokens field value if set, zero value otherwise.
func (o *ListWalletTransactionsRI) GetFungibleTokens() []ListWalletTransactionsRIFungibleTokensInner {
	if o == nil || o.FungibleTokens == nil {
		var ret []ListWalletTransactionsRIFungibleTokensInner
		return ret
	}
	return o.FungibleTokens
}

// GetFungibleTokensOk returns a tuple with the FungibleTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetFungibleTokensOk() ([]ListWalletTransactionsRIFungibleTokensInner, bool) {
	if o == nil || o.FungibleTokens == nil {
		return nil, false
	}
	return o.FungibleTokens, true
}

// HasFungibleTokens returns a boolean if a field has been set.
func (o *ListWalletTransactionsRI) HasFungibleTokens() bool {
	if o != nil && o.FungibleTokens != nil {
		return true
	}

	return false
}

// SetFungibleTokens gets a reference to the given []ListWalletTransactionsRIFungibleTokensInner and assigns it to the FungibleTokens field.
func (o *ListWalletTransactionsRI) SetFungibleTokens(v []ListWalletTransactionsRIFungibleTokensInner) {
	o.FungibleTokens = v
}

// GetInternalTransactions returns the InternalTransactions field value if set, zero value otherwise.
func (o *ListWalletTransactionsRI) GetInternalTransactions() []ListWalletTransactionsRIInternalTransactionsInner {
	if o == nil || o.InternalTransactions == nil {
		var ret []ListWalletTransactionsRIInternalTransactionsInner
		return ret
	}
	return o.InternalTransactions
}

// GetInternalTransactionsOk returns a tuple with the InternalTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetInternalTransactionsOk() ([]ListWalletTransactionsRIInternalTransactionsInner, bool) {
	if o == nil || o.InternalTransactions == nil {
		return nil, false
	}
	return o.InternalTransactions, true
}

// HasInternalTransactions returns a boolean if a field has been set.
func (o *ListWalletTransactionsRI) HasInternalTransactions() bool {
	if o != nil && o.InternalTransactions != nil {
		return true
	}

	return false
}

// SetInternalTransactions gets a reference to the given []ListWalletTransactionsRIInternalTransactionsInner and assigns it to the InternalTransactions field.
func (o *ListWalletTransactionsRI) SetInternalTransactions(v []ListWalletTransactionsRIInternalTransactionsInner) {
	o.InternalTransactions = v
}

// GetNonFungibleTokens returns the NonFungibleTokens field value if set, zero value otherwise.
func (o *ListWalletTransactionsRI) GetNonFungibleTokens() []ListWalletTransactionsRINonFungibleTokensInner {
	if o == nil || o.NonFungibleTokens == nil {
		var ret []ListWalletTransactionsRINonFungibleTokensInner
		return ret
	}
	return o.NonFungibleTokens
}

// GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetNonFungibleTokensOk() ([]ListWalletTransactionsRINonFungibleTokensInner, bool) {
	if o == nil || o.NonFungibleTokens == nil {
		return nil, false
	}
	return o.NonFungibleTokens, true
}

// HasNonFungibleTokens returns a boolean if a field has been set.
func (o *ListWalletTransactionsRI) HasNonFungibleTokens() bool {
	if o != nil && o.NonFungibleTokens != nil {
		return true
	}

	return false
}

// SetNonFungibleTokens gets a reference to the given []ListWalletTransactionsRINonFungibleTokensInner and assigns it to the NonFungibleTokens field.
func (o *ListWalletTransactionsRI) SetNonFungibleTokens(v []ListWalletTransactionsRINonFungibleTokensInner) {
	o.NonFungibleTokens = v
}

// GetRecipients returns the Recipients field value
func (o *ListWalletTransactionsRI) GetRecipients() []ListWalletTransactionsRIRecipientsInner {
	if o == nil {
		var ret []ListWalletTransactionsRIRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetRecipientsOk() ([]ListWalletTransactionsRIRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ListWalletTransactionsRI) SetRecipients(v []ListWalletTransactionsRIRecipientsInner) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListWalletTransactionsRI) GetSenders() []ListWalletTransactionsRISendersInner {
	if o == nil {
		var ret []ListWalletTransactionsRISendersInner
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetSendersOk() ([]ListWalletTransactionsRISendersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Senders, true
}

// SetSenders sets field value
func (o *ListWalletTransactionsRI) SetSenders(v []ListWalletTransactionsRISendersInner) {
	o.Senders = v
}

// GetStatus returns the Status field value
func (o *ListWalletTransactionsRI) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListWalletTransactionsRI) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListWalletTransactionsRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListWalletTransactionsRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListWalletTransactionsRI) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListWalletTransactionsRI) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetValue returns the Value field value
func (o *ListWalletTransactionsRI) GetValue() ListWalletTransactionsRIValue {
	if o == nil {
		var ret ListWalletTransactionsRIValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRI) GetValueOk() (*ListWalletTransactionsRIValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListWalletTransactionsRI) SetValue(v ListWalletTransactionsRIValue) {
	o.Value = v
}

func (o ListWalletTransactionsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["direction"] = o.Direction
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if o.FungibleTokens != nil {
		toSerialize["fungibleTokens"] = o.FungibleTokens
	}
	if o.InternalTransactions != nil {
		toSerialize["internalTransactions"] = o.InternalTransactions
	}
	if o.NonFungibleTokens != nil {
		toSerialize["nonFungibleTokens"] = o.NonFungibleTokens
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableListWalletTransactionsRI struct {
	value *ListWalletTransactionsRI
	isSet bool
}

func (v NullableListWalletTransactionsRI) Get() *ListWalletTransactionsRI {
	return v.value
}

func (v *NullableListWalletTransactionsRI) Set(val *ListWalletTransactionsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListWalletTransactionsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListWalletTransactionsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWalletTransactionsRI(val *ListWalletTransactionsRI) *NullableListWalletTransactionsRI {
	return &NullableListWalletTransactionsRI{value: val, isSet: true}
}

func (v NullableListWalletTransactionsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWalletTransactionsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


