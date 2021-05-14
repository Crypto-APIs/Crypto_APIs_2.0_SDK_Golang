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

// AddressCoinsTransactionConfirmedDataItem Defines an `item` as one result.
type AddressCoinsTransactionConfirmedDataItem struct {
	// Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
	Blockchain string `json:"blockchain"`
	// Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
	Network string `json:"network"`
	// Defines the specific address to which the coin transaction has been sent and is confirmed.
	Address string `json:"address"`
	MinedInBlock AddressCoinsTransactionConfirmedDataItemMinedInBlock `json:"minedInBlock"`
	// Defines the unique ID of the specific transaction, i.e. its identification number.
	TransactionId string `json:"transactionId"`
	// Defines the amount of coins sent with the confirmed transaction.
	Amount string `json:"amount"`
	// Defines the unit of the transaction, e.g. BTC.
	Unit string `json:"unit"`
	// Defines whether the transaction is \"incoming\" or \"outgoing\".
	Direction string `json:"direction"`
}

// NewAddressCoinsTransactionConfirmedDataItem instantiates a new AddressCoinsTransactionConfirmedDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCoinsTransactionConfirmedDataItem(blockchain string, network string, address string, minedInBlock AddressCoinsTransactionConfirmedDataItemMinedInBlock, transactionId string, amount string, unit string, direction string) *AddressCoinsTransactionConfirmedDataItem {
	this := AddressCoinsTransactionConfirmedDataItem{}
	this.Blockchain = blockchain
	this.Network = network
	this.Address = address
	this.MinedInBlock = minedInBlock
	this.TransactionId = transactionId
	this.Amount = amount
	this.Unit = unit
	this.Direction = direction
	return &this
}

// NewAddressCoinsTransactionConfirmedDataItemWithDefaults instantiates a new AddressCoinsTransactionConfirmedDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCoinsTransactionConfirmedDataItemWithDefaults() *AddressCoinsTransactionConfirmedDataItem {
	this := AddressCoinsTransactionConfirmedDataItem{}
	return &this
}

// GetBlockchain returns the Blockchain field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetBlockchainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetNetwork returns the Network field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetNetwork(v string) {
	o.Network = v
}

// GetAddress returns the Address field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetAddress(v string) {
	o.Address = v
}

// GetMinedInBlock returns the MinedInBlock field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetMinedInBlock() AddressCoinsTransactionConfirmedDataItemMinedInBlock {
	if o == nil {
		var ret AddressCoinsTransactionConfirmedDataItemMinedInBlock
		return ret
	}

	return o.MinedInBlock
}

// GetMinedInBlockOk returns a tuple with the MinedInBlock field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetMinedInBlockOk() (*AddressCoinsTransactionConfirmedDataItemMinedInBlock, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinedInBlock, true
}

// SetMinedInBlock sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetMinedInBlock(v AddressCoinsTransactionConfirmedDataItemMinedInBlock) {
	o.MinedInBlock = v
}

// GetTransactionId returns the TransactionId field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetAmount returns the Amount field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetUnit(v string) {
	o.Unit = v
}

// GetDirection returns the Direction field value
func (o *AddressCoinsTransactionConfirmedDataItem) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItem) GetDirectionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *AddressCoinsTransactionConfirmedDataItem) SetDirection(v string) {
	o.Direction = v
}

func (o AddressCoinsTransactionConfirmedDataItem) MarshalJSON() ([]byte, error) {
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
		toSerialize["transactionId"] = o.TransactionId
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

type NullableAddressCoinsTransactionConfirmedDataItem struct {
	value *AddressCoinsTransactionConfirmedDataItem
	isSet bool
}

func (v NullableAddressCoinsTransactionConfirmedDataItem) Get() *AddressCoinsTransactionConfirmedDataItem {
	return v.value
}

func (v *NullableAddressCoinsTransactionConfirmedDataItem) Set(val *AddressCoinsTransactionConfirmedDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCoinsTransactionConfirmedDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCoinsTransactionConfirmedDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCoinsTransactionConfirmedDataItem(val *AddressCoinsTransactionConfirmedDataItem) *NullableAddressCoinsTransactionConfirmedDataItem {
	return &NullableAddressCoinsTransactionConfirmedDataItem{value: val, isSet: true}
}

func (v NullableAddressCoinsTransactionConfirmedDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCoinsTransactionConfirmedDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

