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

// GetWalletTransactionDetailsByTransactionIDRIBST Tron
type GetWalletTransactionDetailsByTransactionIDRIBST struct {
	// String representation of the amount value
	Amount string `json:"amount"`
	// Numeric representation of the transaction used bandwidth
	BandwidthUsed string `json:"bandwidthUsed"`
	// Numeric representation of the transaction contract
	Contract string `json:"contract"`
	// String representation of the transaction used energy
	EnergyUsed string `json:"energyUsed"`
	HasInternalTransactions bool `json:"hasInternalTransactions"`
	HasTokenTransfers bool `json:"hasTokenTransfers"`
	// Numeric representation of the transaction input
	Input string `json:"input"`
	// String representation of the transaction status
	Status string `json:"status"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBST instantiates a new GetWalletTransactionDetailsByTransactionIDRIBST object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBST(amount string, bandwidthUsed string, contract string, energyUsed string, hasInternalTransactions bool, hasTokenTransfers bool, input string, status string) *GetWalletTransactionDetailsByTransactionIDRIBST {
	this := GetWalletTransactionDetailsByTransactionIDRIBST{}
	this.Amount = amount
	this.BandwidthUsed = bandwidthUsed
	this.Contract = contract
	this.EnergyUsed = energyUsed
	this.HasInternalTransactions = hasInternalTransactions
	this.HasTokenTransfers = hasTokenTransfers
	this.Input = input
	this.Status = status
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSTWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBST object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSTWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBST {
	this := GetWalletTransactionDetailsByTransactionIDRIBST{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetAmount(v string) {
	o.Amount = v
}

// GetBandwidthUsed returns the BandwidthUsed field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetBandwidthUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandwidthUsed
}

// GetBandwidthUsedOk returns a tuple with the BandwidthUsed field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetBandwidthUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandwidthUsed, true
}

// SetBandwidthUsed sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetBandwidthUsed(v string) {
	o.BandwidthUsed = v
}

// GetContract returns the Contract field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetContract(v string) {
	o.Contract = v
}

// GetEnergyUsed returns the EnergyUsed field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetEnergyUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnergyUsed
}

// GetEnergyUsedOk returns a tuple with the EnergyUsed field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetEnergyUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnergyUsed, true
}

// SetEnergyUsed sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetEnergyUsed(v string) {
	o.EnergyUsed = v
}

// GetHasInternalTransactions returns the HasInternalTransactions field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetHasInternalTransactions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasInternalTransactions
}

// GetHasInternalTransactionsOk returns a tuple with the HasInternalTransactions field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetHasInternalTransactionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasInternalTransactions, true
}

// SetHasInternalTransactions sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetHasInternalTransactions(v bool) {
	o.HasInternalTransactions = v
}

// GetHasTokenTransfers returns the HasTokenTransfers field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetHasTokenTransfers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasTokenTransfers
}

// GetHasTokenTransfersOk returns a tuple with the HasTokenTransfers field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetHasTokenTransfersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasTokenTransfers, true
}

// SetHasTokenTransfers sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetHasTokenTransfers(v bool) {
	o.HasTokenTransfers = v
}

// GetInput returns the Input field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetInput(v string) {
	o.Input = v
}

// GetStatus returns the Status field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBST) SetStatus(v string) {
	o.Status = v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBST) MarshalJSON() ([]byte, error) {
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
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletTransactionDetailsByTransactionIDRIBST struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBST
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBST) Get() *GetWalletTransactionDetailsByTransactionIDRIBST {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBST) Set(val *GetWalletTransactionDetailsByTransactionIDRIBST) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBST) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBST) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBST(val *GetWalletTransactionDetailsByTransactionIDRIBST) *NullableGetWalletTransactionDetailsByTransactionIDRIBST {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBST{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBST) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBST) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


