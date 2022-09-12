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

// AddressTokensTransactionConfirmedEachConfirmationDataItem Defines an `item` as one result.
type AddressTokensTransactionConfirmedEachConfirmationDataItem struct {
	// Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
	Blockchain string `json:"blockchain"`
	// Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\",  are test networks.
	Network string `json:"network"`
	// Defines the specific address to which the transaction has been sent.
	Address string `json:"address"`
	MinedInBlock AddressTokensTransactionConfirmedDataItemMinedInBlock `json:"minedInBlock"`
	// Defines the unique ID of the specific transaction, i.e. its identification number.
	TransactionId string `json:"transactionId"`
	// Defines the number of currently received confirmations for the transaction.
	CurrentConfirmations int32 `json:"currentConfirmations"`
	// Defines the number of confirmation transactions requested as callbacks, i.e. the system can notify till the n-th confirmation.
	TargetConfirmations int32 `json:"targetConfirmations"`
	// Defines the type of token sent with the transaction, e.g. ERC 20.
	TokenType string `json:"tokenType"`
	Token AddressTokensTransactionConfirmedEachConfirmationToken `json:"token"`
	// Defines whether the transaction is \"incoming\" or \"outgoing\".
	Direction string `json:"direction"`
}

// NewAddressTokensTransactionConfirmedEachConfirmationDataItem instantiates a new AddressTokensTransactionConfirmedEachConfirmationDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTokensTransactionConfirmedEachConfirmationDataItem(blockchain string, network string, address string, minedInBlock AddressTokensTransactionConfirmedDataItemMinedInBlock, transactionId string, currentConfirmations int32, targetConfirmations int32, tokenType string, token AddressTokensTransactionConfirmedEachConfirmationToken, direction string) *AddressTokensTransactionConfirmedEachConfirmationDataItem {
	this := AddressTokensTransactionConfirmedEachConfirmationDataItem{}
	this.Blockchain = blockchain
	this.Network = network
	this.Address = address
	this.MinedInBlock = minedInBlock
	this.TransactionId = transactionId
	this.CurrentConfirmations = currentConfirmations
	this.TargetConfirmations = targetConfirmations
	this.TokenType = tokenType
	this.Token = token
	this.Direction = direction
	return &this
}

// NewAddressTokensTransactionConfirmedEachConfirmationDataItemWithDefaults instantiates a new AddressTokensTransactionConfirmedEachConfirmationDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTokensTransactionConfirmedEachConfirmationDataItemWithDefaults() *AddressTokensTransactionConfirmedEachConfirmationDataItem {
	this := AddressTokensTransactionConfirmedEachConfirmationDataItem{}
	return &this
}

// GetBlockchain returns the Blockchain field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetBlockchainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetNetwork returns the Network field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetNetwork(v string) {
	o.Network = v
}

// GetAddress returns the Address field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetAddress(v string) {
	o.Address = v
}

// GetMinedInBlock returns the MinedInBlock field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetMinedInBlock() AddressTokensTransactionConfirmedDataItemMinedInBlock {
	if o == nil {
		var ret AddressTokensTransactionConfirmedDataItemMinedInBlock
		return ret
	}

	return o.MinedInBlock
}

// GetMinedInBlockOk returns a tuple with the MinedInBlock field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetMinedInBlockOk() (*AddressTokensTransactionConfirmedDataItemMinedInBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlock, true
}

// SetMinedInBlock sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetMinedInBlock(v AddressTokensTransactionConfirmedDataItemMinedInBlock) {
	o.MinedInBlock = v
}

// GetTransactionId returns the TransactionId field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetCurrentConfirmations returns the CurrentConfirmations field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentConfirmations
}

// GetCurrentConfirmationsOk returns a tuple with the CurrentConfirmations field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentConfirmations, true
}

// SetCurrentConfirmations sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetCurrentConfirmations(v int32) {
	o.CurrentConfirmations = v
}

// GetTargetConfirmations returns the TargetConfirmations field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetConfirmations
}

// GetTargetConfirmationsOk returns a tuple with the TargetConfirmations field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetConfirmations, true
}

// SetTargetConfirmations sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetTargetConfirmations(v int32) {
	o.TargetConfirmations = v
}

// GetTokenType returns the TokenType field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetTokenType(v string) {
	o.TokenType = v
}

// GetToken returns the Token field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetToken() AddressTokensTransactionConfirmedEachConfirmationToken {
	if o == nil {
		var ret AddressTokensTransactionConfirmedEachConfirmationToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTokenOk() (*AddressTokensTransactionConfirmedEachConfirmationToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetToken(v AddressTokensTransactionConfirmedEachConfirmationToken) {
	o.Token = v
}

// GetDirection returns the Direction field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetDirection(v string) {
	o.Direction = v
}

func (o AddressTokensTransactionConfirmedEachConfirmationDataItem) MarshalJSON() ([]byte, error) {
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
		toSerialize["currentConfirmations"] = o.CurrentConfirmations
	}
	if true {
		toSerialize["targetConfirmations"] = o.TargetConfirmations
	}
	if true {
		toSerialize["tokenType"] = o.TokenType
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["direction"] = o.Direction
	}
	return json.Marshal(toSerialize)
}

type NullableAddressTokensTransactionConfirmedEachConfirmationDataItem struct {
	value *AddressTokensTransactionConfirmedEachConfirmationDataItem
	isSet bool
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmationDataItem) Get() *AddressTokensTransactionConfirmedEachConfirmationDataItem {
	return v.value
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmationDataItem) Set(val *AddressTokensTransactionConfirmedEachConfirmationDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmationDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmationDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTokensTransactionConfirmedEachConfirmationDataItem(val *AddressTokensTransactionConfirmedEachConfirmationDataItem) *NullableAddressTokensTransactionConfirmedEachConfirmationDataItem {
	return &NullableAddressTokensTransactionConfirmedEachConfirmationDataItem{value: val, isSet: true}
}

func (v NullableAddressTokensTransactionConfirmedEachConfirmationDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTokensTransactionConfirmedEachConfirmationDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


