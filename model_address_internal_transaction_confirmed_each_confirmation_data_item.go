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

// AddressInternalTransactionConfirmedEachConfirmationDataItem Defines an `item` as one result.
type AddressInternalTransactionConfirmedEachConfirmationDataItem struct {
	// Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
	Blockchain string `json:"blockchain"`
	// Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
	Network string `json:"network"`
	// Defines the specific address of the internal transaction.
	Address string `json:"address"`
	MinedInBlock AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock `json:"minedInBlock"`
	// Defines the Parent Transaction's unique ID.
	ParentTransactionId string `json:"parentTransactionId"`
	// Defines the specific operation's unique ID.
	OperationId string `json:"operationId"`
	// Defines the number of currently received confirmations for the transaction.
	CurrentConfirmations int32 `json:"currentConfirmations"`
	// Defines the number of confirmation transactions requested as callbacks, i.e. the system can notify till the n-th confirmation.
	TargetConfirmations int32 `json:"targetConfirmations"`
	// Defines the amount of coins sent with the confirmed transaction.
	Amount string `json:"amount"`
	// Defines the unit of the transaction, e.g. Gwei.
	Unit string `json:"unit"`
	// Defines whether the transaction is \"incoming\" or \"outgoing\".
	Direction string `json:"direction"`
}

// NewAddressInternalTransactionConfirmedEachConfirmationDataItem instantiates a new AddressInternalTransactionConfirmedEachConfirmationDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInternalTransactionConfirmedEachConfirmationDataItem(blockchain string, network string, address string, minedInBlock AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock, parentTransactionId string, operationId string, currentConfirmations int32, targetConfirmations int32, amount string, unit string, direction string) *AddressInternalTransactionConfirmedEachConfirmationDataItem {
	this := AddressInternalTransactionConfirmedEachConfirmationDataItem{}
	this.Blockchain = blockchain
	this.Network = network
	this.Address = address
	this.MinedInBlock = minedInBlock
	this.ParentTransactionId = parentTransactionId
	this.OperationId = operationId
	this.CurrentConfirmations = currentConfirmations
	this.TargetConfirmations = targetConfirmations
	this.Amount = amount
	this.Unit = unit
	this.Direction = direction
	return &this
}

// NewAddressInternalTransactionConfirmedEachConfirmationDataItemWithDefaults instantiates a new AddressInternalTransactionConfirmedEachConfirmationDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInternalTransactionConfirmedEachConfirmationDataItemWithDefaults() *AddressInternalTransactionConfirmedEachConfirmationDataItem {
	this := AddressInternalTransactionConfirmedEachConfirmationDataItem{}
	return &this
}

// GetBlockchain returns the Blockchain field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetBlockchainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetNetwork returns the Network field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetNetwork(v string) {
	o.Network = v
}

// GetAddress returns the Address field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetAddress(v string) {
	o.Address = v
}

// GetMinedInBlock returns the MinedInBlock field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetMinedInBlock() AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock {
	if o == nil {
		var ret AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock
		return ret
	}

	return o.MinedInBlock
}

// GetMinedInBlockOk returns a tuple with the MinedInBlock field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetMinedInBlockOk() (*AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlock, true
}

// SetMinedInBlock sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetMinedInBlock(v AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) {
	o.MinedInBlock = v
}

// GetParentTransactionId returns the ParentTransactionId field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetParentTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentTransactionId
}

// GetParentTransactionIdOk returns a tuple with the ParentTransactionId field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetParentTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentTransactionId, true
}

// SetParentTransactionId sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetParentTransactionId(v string) {
	o.ParentTransactionId = v
}

// GetOperationId returns the OperationId field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetOperationId(v string) {
	o.OperationId = v
}

// GetCurrentConfirmations returns the CurrentConfirmations field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentConfirmations
}

// GetCurrentConfirmationsOk returns a tuple with the CurrentConfirmations field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentConfirmations, true
}

// SetCurrentConfirmations sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetCurrentConfirmations(v int32) {
	o.CurrentConfirmations = v
}

// GetTargetConfirmations returns the TargetConfirmations field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetConfirmations
}

// GetTargetConfirmationsOk returns a tuple with the TargetConfirmations field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetConfirmations, true
}

// SetTargetConfirmations sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetTargetConfirmations(v int32) {
	o.TargetConfirmations = v
}

// GetAmount returns the Amount field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetUnit(v string) {
	o.Unit = v
}

// GetDirection returns the Direction field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetDirection(v string) {
	o.Direction = v
}

func (o AddressInternalTransactionConfirmedEachConfirmationDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockchain"] = o.Blockchain
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["minedInBlock"] = o.MinedInBlock
	}
	if true {
		toSerialize["parentTransactionId"] = o.ParentTransactionId
	}
	if true {
		toSerialize["operationId"] = o.OperationId
	}
	if true {
		toSerialize["currentConfirmations"] = o.CurrentConfirmations
	}
	if true {
		toSerialize["targetConfirmations"] = o.TargetConfirmations
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["direction"] = o.Direction
	}
	return json.Marshal(toSerialize)
}

type NullableAddressInternalTransactionConfirmedEachConfirmationDataItem struct {
	value *AddressInternalTransactionConfirmedEachConfirmationDataItem
	isSet bool
}

func (v NullableAddressInternalTransactionConfirmedEachConfirmationDataItem) Get() *AddressInternalTransactionConfirmedEachConfirmationDataItem {
	return v.value
}

func (v *NullableAddressInternalTransactionConfirmedEachConfirmationDataItem) Set(val *AddressInternalTransactionConfirmedEachConfirmationDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInternalTransactionConfirmedEachConfirmationDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInternalTransactionConfirmedEachConfirmationDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInternalTransactionConfirmedEachConfirmationDataItem(val *AddressInternalTransactionConfirmedEachConfirmationDataItem) *NullableAddressInternalTransactionConfirmedEachConfirmationDataItem {
	return &NullableAddressInternalTransactionConfirmedEachConfirmationDataItem{value: val, isSet: true}
}

func (v NullableAddressInternalTransactionConfirmedEachConfirmationDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressInternalTransactionConfirmedEachConfirmationDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


