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

// ListOmniTransactionsByBlockHeightRI struct for ListOmniTransactionsByBlockHeightRI
type ListOmniTransactionsByBlockHeightRI struct {
	// Defines the amount of the sent tokens.
	Amount string `json:"amount"`
	// Defines whether the attribute can be divisible or not, as boolean. E.g., if it is \"true\", the attribute is divisible.
	Divisible bool `json:"divisible"`
	// Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	MinedInBlockHash string `json:"minedInBlockHash"`
	// Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block.
	MinedInBlockHeight int32 `json:"minedInBlockHeight"`
	// Represents the index position of the transaction in the specific block.
	PositionInBlock int32 `json:"positionInBlock"`
	// Represents the identifier of the tokens to send.
	PropertyId int32 `json:"propertyId"`
	// Represents an object of addresses that receive the transactions.
	Recipients []ListOmniTransactionsByBlockHeightRIRecipients `json:"recipients"`
	// Represents an object of addresses that provide the funds.
	Senders []ListOmniTransactionsByBlockHeightRISenders `json:"senders"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
	TransactionId string `json:"transactionId"`
	// Defines the type of the transaction as a string.
	Type string `json:"type"`
	// Defines the type of the transaction as a number.
	TypeInt int32 `json:"typeInt"`
	// Defines whether the transaction is valid or not, as boolean. E.g., if it is \"true\", the transaction is valid.
	Valid bool `json:"valid"`
	// Defines the specific version.
	Version int32 `json:"version"`
	Fee ListOmniTransactionsByBlockHeightRIFee `json:"fee"`
}

// NewListOmniTransactionsByBlockHeightRI instantiates a new ListOmniTransactionsByBlockHeightRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOmniTransactionsByBlockHeightRI(amount string, divisible bool, minedInBlockHash string, minedInBlockHeight int32, positionInBlock int32, propertyId int32, recipients []ListOmniTransactionsByBlockHeightRIRecipients, senders []ListOmniTransactionsByBlockHeightRISenders, timestamp int32, transactionId string, type_ string, typeInt int32, valid bool, version int32, fee ListOmniTransactionsByBlockHeightRIFee) *ListOmniTransactionsByBlockHeightRI {
	this := ListOmniTransactionsByBlockHeightRI{}
	this.Amount = amount
	this.Divisible = divisible
	this.MinedInBlockHash = minedInBlockHash
	this.MinedInBlockHeight = minedInBlockHeight
	this.PositionInBlock = positionInBlock
	this.PropertyId = propertyId
	this.Recipients = recipients
	this.Senders = senders
	this.Timestamp = timestamp
	this.TransactionId = transactionId
	this.Type = type_
	this.TypeInt = typeInt
	this.Valid = valid
	this.Version = version
	this.Fee = fee
	return &this
}

// NewListOmniTransactionsByBlockHeightRIWithDefaults instantiates a new ListOmniTransactionsByBlockHeightRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOmniTransactionsByBlockHeightRIWithDefaults() *ListOmniTransactionsByBlockHeightRI {
	this := ListOmniTransactionsByBlockHeightRI{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListOmniTransactionsByBlockHeightRI) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetAmount(v string) {
	o.Amount = v
}

// GetDivisible returns the Divisible field value
func (o *ListOmniTransactionsByBlockHeightRI) GetDivisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Divisible
}

// GetDivisibleOk returns a tuple with the Divisible field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetDivisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Divisible, true
}

// SetDivisible sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetDivisible(v bool) {
	o.Divisible = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value
func (o *ListOmniTransactionsByBlockHeightRI) GetMinedInBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlockHash, true
}

// SetMinedInBlockHash sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = v
}

// GetMinedInBlockHeight returns the MinedInBlockHeight field value
func (o *ListOmniTransactionsByBlockHeightRI) GetMinedInBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinedInBlockHeight
}

// GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetMinedInBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlockHeight, true
}

// SetMinedInBlockHeight sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetMinedInBlockHeight(v int32) {
	o.MinedInBlockHeight = v
}

// GetPositionInBlock returns the PositionInBlock field value
func (o *ListOmniTransactionsByBlockHeightRI) GetPositionInBlock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PositionInBlock
}

// GetPositionInBlockOk returns a tuple with the PositionInBlock field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetPositionInBlockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PositionInBlock, true
}

// SetPositionInBlock sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetPositionInBlock(v int32) {
	o.PositionInBlock = v
}

// GetPropertyId returns the PropertyId field value
func (o *ListOmniTransactionsByBlockHeightRI) GetPropertyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetPropertyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetPropertyId(v int32) {
	o.PropertyId = v
}

// GetRecipients returns the Recipients field value
func (o *ListOmniTransactionsByBlockHeightRI) GetRecipients() []ListOmniTransactionsByBlockHeightRIRecipients {
	if o == nil {
		var ret []ListOmniTransactionsByBlockHeightRIRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetRecipientsOk() ([]ListOmniTransactionsByBlockHeightRIRecipients, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetRecipients(v []ListOmniTransactionsByBlockHeightRIRecipients) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListOmniTransactionsByBlockHeightRI) GetSenders() []ListOmniTransactionsByBlockHeightRISenders {
	if o == nil {
		var ret []ListOmniTransactionsByBlockHeightRISenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetSendersOk() ([]ListOmniTransactionsByBlockHeightRISenders, bool) {
	if o == nil {
		return nil, false
	}
	return o.Senders, true
}

// SetSenders sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetSenders(v []ListOmniTransactionsByBlockHeightRISenders) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListOmniTransactionsByBlockHeightRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListOmniTransactionsByBlockHeightRI) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetType returns the Type field value
func (o *ListOmniTransactionsByBlockHeightRI) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetType(v string) {
	o.Type = v
}

// GetTypeInt returns the TypeInt field value
func (o *ListOmniTransactionsByBlockHeightRI) GetTypeInt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TypeInt
}

// GetTypeIntOk returns a tuple with the TypeInt field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetTypeIntOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeInt, true
}

// SetTypeInt sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetTypeInt(v int32) {
	o.TypeInt = v
}

// GetValid returns the Valid field value
func (o *ListOmniTransactionsByBlockHeightRI) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetValid(v bool) {
	o.Valid = v
}

// GetVersion returns the Version field value
func (o *ListOmniTransactionsByBlockHeightRI) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetVersion(v int32) {
	o.Version = v
}

// GetFee returns the Fee field value
func (o *ListOmniTransactionsByBlockHeightRI) GetFee() ListOmniTransactionsByBlockHeightRIFee {
	if o == nil {
		var ret ListOmniTransactionsByBlockHeightRIFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHeightRI) GetFeeOk() (*ListOmniTransactionsByBlockHeightRIFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListOmniTransactionsByBlockHeightRI) SetFee(v ListOmniTransactionsByBlockHeightRIFee) {
	o.Fee = v
}

func (o ListOmniTransactionsByBlockHeightRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["divisible"] = o.Divisible
	}
	if true {
		toSerialize["minedInBlockHash"] = o.MinedInBlockHash
	}
	if true {
		toSerialize["minedInBlockHeight"] = o.MinedInBlockHeight
	}
	if true {
		toSerialize["positionInBlock"] = o.PositionInBlock
	}
	if true {
		toSerialize["propertyId"] = o.PropertyId
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["typeInt"] = o.TypeInt
	}
	if true {
		toSerialize["valid"] = o.Valid
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	return json.Marshal(toSerialize)
}

type NullableListOmniTransactionsByBlockHeightRI struct {
	value *ListOmniTransactionsByBlockHeightRI
	isSet bool
}

func (v NullableListOmniTransactionsByBlockHeightRI) Get() *ListOmniTransactionsByBlockHeightRI {
	return v.value
}

func (v *NullableListOmniTransactionsByBlockHeightRI) Set(val *ListOmniTransactionsByBlockHeightRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByBlockHeightRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByBlockHeightRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByBlockHeightRI(val *ListOmniTransactionsByBlockHeightRI) *NullableListOmniTransactionsByBlockHeightRI {
	return &NullableListOmniTransactionsByBlockHeightRI{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByBlockHeightRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByBlockHeightRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


