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

// ListUnconfirmedTokensTransfersByAddressRI struct for ListUnconfirmedTokensTransfersByAddressRI
type ListUnconfirmedTokensTransfersByAddressRI struct {
	// Represents the contract address of the token, which controls its logic. It is not the address that holds the tokens.
	ContractAddress string `json:"contractAddress"`
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
}

// NewListUnconfirmedTokensTransfersByAddressRI instantiates a new ListUnconfirmedTokensTransfersByAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTokensTransfersByAddressRI(contractAddress string, recipientAddress string, senderAddress string, tokenDecimals int32, tokenName string, tokenSymbol string, tokenType string, transactionHash string, transactionTimestamp int32) *ListUnconfirmedTokensTransfersByAddressRI {
	this := ListUnconfirmedTokensTransfersByAddressRI{}
	this.ContractAddress = contractAddress
	this.RecipientAddress = recipientAddress
	this.SenderAddress = senderAddress
	this.TokenDecimals = tokenDecimals
	this.TokenName = tokenName
	this.TokenSymbol = tokenSymbol
	this.TokenType = tokenType
	this.TransactionHash = transactionHash
	this.TransactionTimestamp = transactionTimestamp
	return &this
}

// NewListUnconfirmedTokensTransfersByAddressRIWithDefaults instantiates a new ListUnconfirmedTokensTransfersByAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTokensTransfersByAddressRIWithDefaults() *ListUnconfirmedTokensTransfersByAddressRI {
	this := ListUnconfirmedTokensTransfersByAddressRI{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetRecipientAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetRecipientAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetRecipientAddress(v string) {
	o.RecipientAddress = v
}

// GetSenderAddress returns the SenderAddress field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetSenderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAddress, true
}

// SetSenderAddress sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetSenderAddress(v string) {
	o.SenderAddress = v
}

// GetTokenDecimals returns the TokenDecimals field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenDecimals
}

// GetTokenDecimalsOk returns a tuple with the TokenDecimals field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenDecimals, true
}

// SetTokenDecimals sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTokenDecimals(v int32) {
	o.TokenDecimals = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenName returns the TokenName field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTokenName(v string) {
	o.TokenName = v
}

// GetTokenSymbol returns the TokenSymbol field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenSymbol
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenSymbol, true
}

// SetTokenSymbol sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTokenSymbol(v string) {
	o.TokenSymbol = v
}

// GetTokenType returns the TokenType field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTokenType(v string) {
	o.TokenType = v
}

// GetTokensAmount returns the TokensAmount field value if set, zero value otherwise.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokensAmount() string {
	if o == nil || o.TokensAmount == nil {
		var ret string
		return ret
	}
	return *o.TokensAmount
}

// GetTokensAmountOk returns a tuple with the TokensAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTokensAmountOk() (*string, bool) {
	if o == nil || o.TokensAmount == nil {
		return nil, false
	}
	return o.TokensAmount, true
}

// HasTokensAmount returns a boolean if a field has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) HasTokensAmount() bool {
	if o != nil && o.TokensAmount != nil {
		return true
	}

	return false
}

// SetTokensAmount gets a reference to the given string and assigns it to the TokensAmount field.
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTokensAmount(v string) {
	o.TokensAmount = &v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionTimestamp returns the TransactionTimestamp field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTransactionTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionTimestamp
}

// GetTransactionTimestampOk returns a tuple with the TransactionTimestamp field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTokensTransfersByAddressRI) GetTransactionTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTimestamp, true
}

// SetTransactionTimestamp sets field value
func (o *ListUnconfirmedTokensTransfersByAddressRI) SetTransactionTimestamp(v int32) {
	o.TransactionTimestamp = v
}

func (o ListUnconfirmedTokensTransfersByAddressRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
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
	return json.Marshal(toSerialize)
}

type NullableListUnconfirmedTokensTransfersByAddressRI struct {
	value *ListUnconfirmedTokensTransfersByAddressRI
	isSet bool
}

func (v NullableListUnconfirmedTokensTransfersByAddressRI) Get() *ListUnconfirmedTokensTransfersByAddressRI {
	return v.value
}

func (v *NullableListUnconfirmedTokensTransfersByAddressRI) Set(val *ListUnconfirmedTokensTransfersByAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTokensTransfersByAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTokensTransfersByAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTokensTransfersByAddressRI(val *ListUnconfirmedTokensTransfersByAddressRI) *NullableListUnconfirmedTokensTransfersByAddressRI {
	return &NullableListUnconfirmedTokensTransfersByAddressRI{value: val, isSet: true}
}

func (v NullableListUnconfirmedTokensTransfersByAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTokensTransfersByAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


