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

// EstimateTokenGasLimitRBDataItem struct for EstimateTokenGasLimitRBDataItem
type EstimateTokenGasLimitRBDataItem struct {
	// Represents transactions' amount.
	Amount string `json:"amount"`
	// Defines the specific token identifier.  For Ethereum-based transactions it is the contract.
	Contract string `json:"contract"`
	// Represents the ERC contract type. It can be ERC20 or ERC721
	ContractType string `json:"contractType"`
	// The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient.
	Recipient string `json:"recipient"`
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Sender string `json:"sender"`
}

// NewEstimateTokenGasLimitRBDataItem instantiates a new EstimateTokenGasLimitRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateTokenGasLimitRBDataItem(amount string, contract string, contractType string, recipient string, sender string) *EstimateTokenGasLimitRBDataItem {
	this := EstimateTokenGasLimitRBDataItem{}
	this.Amount = amount
	this.Contract = contract
	this.ContractType = contractType
	this.Recipient = recipient
	this.Sender = sender
	return &this
}

// NewEstimateTokenGasLimitRBDataItemWithDefaults instantiates a new EstimateTokenGasLimitRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateTokenGasLimitRBDataItemWithDefaults() *EstimateTokenGasLimitRBDataItem {
	this := EstimateTokenGasLimitRBDataItem{}
	return &this
}

// GetAmount returns the Amount field value
func (o *EstimateTokenGasLimitRBDataItem) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *EstimateTokenGasLimitRBDataItem) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *EstimateTokenGasLimitRBDataItem) SetAmount(v string) {
	o.Amount = v
}

// GetContract returns the Contract field value
func (o *EstimateTokenGasLimitRBDataItem) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *EstimateTokenGasLimitRBDataItem) GetContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *EstimateTokenGasLimitRBDataItem) SetContract(v string) {
	o.Contract = v
}

// GetContractType returns the ContractType field value
func (o *EstimateTokenGasLimitRBDataItem) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *EstimateTokenGasLimitRBDataItem) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *EstimateTokenGasLimitRBDataItem) SetContractType(v string) {
	o.ContractType = v
}

// GetRecipient returns the Recipient field value
func (o *EstimateTokenGasLimitRBDataItem) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *EstimateTokenGasLimitRBDataItem) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *EstimateTokenGasLimitRBDataItem) SetRecipient(v string) {
	o.Recipient = v
}

// GetSender returns the Sender field value
func (o *EstimateTokenGasLimitRBDataItem) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *EstimateTokenGasLimitRBDataItem) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *EstimateTokenGasLimitRBDataItem) SetSender(v string) {
	o.Sender = v
}

func (o EstimateTokenGasLimitRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["contract"] = o.Contract
	}
	if true {
		toSerialize["contractType"] = o.ContractType
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	return json.Marshal(toSerialize)
}

type NullableEstimateTokenGasLimitRBDataItem struct {
	value *EstimateTokenGasLimitRBDataItem
	isSet bool
}

func (v NullableEstimateTokenGasLimitRBDataItem) Get() *EstimateTokenGasLimitRBDataItem {
	return v.value
}

func (v *NullableEstimateTokenGasLimitRBDataItem) Set(val *EstimateTokenGasLimitRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateTokenGasLimitRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateTokenGasLimitRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateTokenGasLimitRBDataItem(val *EstimateTokenGasLimitRBDataItem) *NullableEstimateTokenGasLimitRBDataItem {
	return &NullableEstimateTokenGasLimitRBDataItem{value: val, isSet: true}
}

func (v NullableEstimateTokenGasLimitRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateTokenGasLimitRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


