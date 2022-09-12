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

// GetTransactionDetailsByTransactionIDFromCallbackRIBST Tron
type GetTransactionDetailsByTransactionIDFromCallbackRIBST struct {
	// Defines the amount of the transaction.
	Amount string `json:"amount"`
	BandwidthUsed GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed `json:"bandwidthUsed"`
	// Represents the specific transaction contract.
	Contract string `json:"contract"`
	EnergyUsed GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed `json:"energyUsed"`
	// Defines if the transaction includes internal transactions (true) or not (false).
	HasInternalTransactions bool `json:"hasInternalTransactions"`
	// Defines if the transaction includes token transfers (true) or not (false).
	HasTokenTransfers string `json:"hasTokenTransfers"`
	// Represents the transaction's input value.
	Input string `json:"input"`
	// Represents the recipient address.
	Recipients string `json:"recipients"`
	// Represents the sender address.
	Senders string `json:"senders"`
	// Represents the status of this transaction.
	TransactionStatus string `json:"transactionStatus"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBST instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBST object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBST(amount string, bandwidthUsed GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed, contract string, energyUsed GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed, hasInternalTransactions bool, hasTokenTransfers string, input string, recipients string, senders string, transactionStatus string) *GetTransactionDetailsByTransactionIDFromCallbackRIBST {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBST{}
	this.Amount = amount
	this.BandwidthUsed = bandwidthUsed
	this.Contract = contract
	this.EnergyUsed = energyUsed
	this.HasInternalTransactions = hasInternalTransactions
	this.HasTokenTransfers = hasTokenTransfers
	this.Input = input
	this.Recipients = recipients
	this.Senders = senders
	this.TransactionStatus = transactionStatus
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSTWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBST object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSTWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBST {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBST{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetAmount(v string) {
	o.Amount = v
}

// GetBandwidthUsed returns the BandwidthUsed field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetBandwidthUsed() GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed
		return ret
	}

	return o.BandwidthUsed
}

// GetBandwidthUsedOk returns a tuple with the BandwidthUsed field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetBandwidthUsedOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandwidthUsed, true
}

// SetBandwidthUsed sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetBandwidthUsed(v GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed) {
	o.BandwidthUsed = v
}

// GetContract returns the Contract field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetContract(v string) {
	o.Contract = v
}

// GetEnergyUsed returns the EnergyUsed field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetEnergyUsed() GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed
		return ret
	}

	return o.EnergyUsed
}

// GetEnergyUsedOk returns a tuple with the EnergyUsed field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetEnergyUsedOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnergyUsed, true
}

// SetEnergyUsed sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetEnergyUsed(v GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed) {
	o.EnergyUsed = v
}

// GetHasInternalTransactions returns the HasInternalTransactions field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasInternalTransactions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasInternalTransactions
}

// GetHasInternalTransactionsOk returns a tuple with the HasInternalTransactions field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasInternalTransactionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasInternalTransactions, true
}

// SetHasInternalTransactions sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetHasInternalTransactions(v bool) {
	o.HasInternalTransactions = v
}

// GetHasTokenTransfers returns the HasTokenTransfers field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasTokenTransfers() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HasTokenTransfers
}

// GetHasTokenTransfersOk returns a tuple with the HasTokenTransfers field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasTokenTransfersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTokenTransfers, true
}

// SetHasTokenTransfers sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetHasTokenTransfers(v string) {
	o.HasTokenTransfers = v
}

// GetInput returns the Input field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetInput(v string) {
	o.Input = v
}

// GetRecipients returns the Recipients field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetRecipients() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetRecipientsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetRecipients(v string) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetSenders() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetSendersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetSenders(v string) {
	o.Senders = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBST) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["bandwidthUsed"] = o.BandwidthUsed
	}
	if true {
		toSerialize["contract"] = o.Contract
	}
	if true {
		toSerialize["energyUsed"] = o.EnergyUsed
	}
	if true {
		toSerialize["hasInternalTransactions"] = o.HasInternalTransactions
	}
	if true {
		toSerialize["hasTokenTransfers"] = o.HasTokenTransfers
	}
	if true {
		toSerialize["input"] = o.Input
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBST
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBST {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBST) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBST(val *GetTransactionDetailsByTransactionIDFromCallbackRIBST) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBST) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


