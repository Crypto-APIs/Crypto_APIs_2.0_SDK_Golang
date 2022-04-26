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

// ListConfirmedTokensTransfersByAddressRI struct for ListConfirmedTokensTransfersByAddressRI
type ListConfirmedTokensTransfersByAddressRI struct {
	// Represents the contract address of the token, which controls its logic. It is not the address that holds the tokens.
	ContractAddress string `json:"contractAddress"`
	// Defines the block height in which this transaction was confirmed/mined.
	MinedInBlockHeight int32 `json:"minedInBlockHeight"`
	// Defines the address to which the recipient receives the transferred tokens.
	RecipientAddress string `json:"recipientAddress"`
	// Defines the address from which the sender transfers tokens.
	SenderAddress string `json:"senderAddress"`
	// Defines the decimals of the token, i.e. the number of digits that come after the decimal coma of the token.
	TokenDecimals int32 `json:"tokenDecimals"`
	// Represents the unique token identifier.
	TokenId *string `json:"tokenId,omitempty"`
	// Defines the token's name as a string.
	TokenName string `json:"tokenName"`
	// Defines the token symbol by which the token contract is known. It is usually 3-4 characters in length.
	TokenSymbol string `json:"tokenSymbol"`
	// Defines the specific token type.
	TokenType string `json:"tokenType"`
	// Defines the token amount of the transfer.
	TokensAmount *string `json:"tokensAmount,omitempty"`
	// Represents the hash of the transaction, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	TransactionHash string `json:"transactionHash"`
	// Defines the specific time/date when the transaction was created in Unix Timestamp.
	TransactionTimestamp int32 `json:"transactionTimestamp"`
	TransactionFee ListTokensTransfersByTransactionHashRITransactionFee `json:"transactionFee"`
}

// NewListConfirmedTokensTransfersByAddressRI instantiates a new ListConfirmedTokensTransfersByAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTokensTransfersByAddressRI(contractAddress string, minedInBlockHeight int32, recipientAddress string, senderAddress string, tokenDecimals int32, tokenName string, tokenSymbol string, tokenType string, transactionHash string, transactionTimestamp int32, transactionFee ListTokensTransfersByTransactionHashRITransactionFee) *ListConfirmedTokensTransfersByAddressRI {
	this := ListConfirmedTokensTransfersByAddressRI{}
	this.ContractAddress = contractAddress
	this.MinedInBlockHeight = minedInBlockHeight
	this.RecipientAddress = recipientAddress
	this.SenderAddress = senderAddress
	this.TokenDecimals = tokenDecimals
	this.TokenName = tokenName
	this.TokenSymbol = tokenSymbol
	this.TokenType = tokenType
	this.TransactionHash = transactionHash
	this.TransactionTimestamp = transactionTimestamp
	this.TransactionFee = transactionFee
	return &this
}

// NewListConfirmedTokensTransfersByAddressRIWithDefaults instantiates a new ListConfirmedTokensTransfersByAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTokensTransfersByAddressRIWithDefaults() *ListConfirmedTokensTransfersByAddressRI {
	this := ListConfirmedTokensTransfersByAddressRI{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetMinedInBlockHeight returns the MinedInBlockHeight field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetMinedInBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinedInBlockHeight
}

// GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetMinedInBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlockHeight, true
}

// SetMinedInBlockHeight sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetMinedInBlockHeight(v int32) {
	o.MinedInBlockHeight = v
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetRecipientAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetRecipientAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetRecipientAddress(v string) {
	o.RecipientAddress = v
}

// GetSenderAddress returns the SenderAddress field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetTokenDecimals returns the TokenDecimals field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenDecimals
}

// GetTokenDecimalsOk returns a tuple with the TokenDecimals field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenDecimals, true
}

// SetTokenDecimals sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenDecimals(v int32) {
	o.TokenDecimals = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenName returns the TokenName field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenName(v string) {
	o.TokenName = v
}

// GetTokenSymbol returns the TokenSymbol field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenSymbol
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenSymbol, true
}

// SetTokenSymbol sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenSymbol(v string) {
	o.TokenSymbol = v
}

// GetTokenType returns the TokenType field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenType(v string) {
	o.TokenType = v
}

// GetTokensAmount returns the TokensAmount field value if set, zero value otherwise.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokensAmount() string {
	if o == nil || o.TokensAmount == nil {
		var ret string
		return ret
	}
	return *o.TokensAmount
}

// GetTokensAmountOk returns a tuple with the TokensAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTokensAmountOk() (*string, bool) {
	if o == nil || o.TokensAmount == nil {
		return nil, false
	}
	return o.TokensAmount, true
}

// HasTokensAmount returns a boolean if a field has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) HasTokensAmount() bool {
	if o != nil && o.TokensAmount != nil {
		return true
	}

	return false
}

// SetTokensAmount gets a reference to the given string and assigns it to the TokensAmount field.
func (o *ListConfirmedTokensTransfersByAddressRI) SetTokensAmount(v string) {
	o.TokensAmount = &v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionTimestamp returns the TransactionTimestamp field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionTimestamp
}

// GetTransactionTimestampOk returns a tuple with the TransactionTimestamp field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTimestamp, true
}

// SetTransactionTimestamp sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetTransactionTimestamp(v int32) {
	o.TransactionTimestamp = v
}

// GetTransactionFee returns the TransactionFee field value
func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionFee() ListTokensTransfersByTransactionHashRITransactionFee {
	if o == nil {
		var ret ListTokensTransfersByTransactionHashRITransactionFee
		return ret
	}

	return o.TransactionFee
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionFeeOk() (*ListTokensTransfersByTransactionHashRITransactionFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionFee, true
}

// SetTransactionFee sets field value
func (o *ListConfirmedTokensTransfersByAddressRI) SetTransactionFee(v ListTokensTransfersByTransactionHashRITransactionFee) {
	o.TransactionFee = v
}

func (o ListConfirmedTokensTransfersByAddressRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if true {
		toSerialize["minedInBlockHeight"] = o.MinedInBlockHeight
	}
	if true {
		toSerialize["recipientAddress"] = o.RecipientAddress
	}
	if true {
		toSerialize["senderAddress"] = o.SenderAddress
	}
	if true {
		toSerialize["tokenDecimals"] = o.TokenDecimals
	}
	if o.TokenId != nil {
		toSerialize["tokenId"] = o.TokenId
	}
	if true {
		toSerialize["tokenName"] = o.TokenName
	}
	if true {
		toSerialize["tokenSymbol"] = o.TokenSymbol
	}
	if true {
		toSerialize["tokenType"] = o.TokenType
	}
	if o.TokensAmount != nil {
		toSerialize["tokensAmount"] = o.TokensAmount
	}
	if true {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	if true {
		toSerialize["transactionTimestamp"] = o.TransactionTimestamp
	}
	if true {
		toSerialize["transactionFee"] = o.TransactionFee
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTokensTransfersByAddressRI struct {
	value *ListConfirmedTokensTransfersByAddressRI
	isSet bool
}

func (v NullableListConfirmedTokensTransfersByAddressRI) Get() *ListConfirmedTokensTransfersByAddressRI {
	return v.value
}

func (v *NullableListConfirmedTokensTransfersByAddressRI) Set(val *ListConfirmedTokensTransfersByAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTokensTransfersByAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTokensTransfersByAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTokensTransfersByAddressRI(val *ListConfirmedTokensTransfersByAddressRI) *NullableListConfirmedTokensTransfersByAddressRI {
	return &NullableListConfirmedTokensTransfersByAddressRI{value: val, isSet: true}
}

func (v NullableListConfirmedTokensTransfersByAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTokensTransfersByAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


