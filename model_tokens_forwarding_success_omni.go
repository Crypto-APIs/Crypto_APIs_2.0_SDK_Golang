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

// TokensForwardingSuccessOmni OMNI
type TokensForwardingSuccessOmni struct {
	// Specifies the name of the token.
	Name string `json:"name"`
	// Defines the ID of the property for Omni Layer.
	PropertyId string `json:"propertyId"`
	// Defines the type of the transaction.
	TransactionType string `json:"transactionType"`
	// The transaction ID used to create the token.
	CreatedByTransactionId string `json:"createdByTransactionId"`
	// Defines the amount of tokens sent with the confirmed transaction.
	Amount string `json:"amount"`
}

// NewTokensForwardingSuccessOmni instantiates a new TokensForwardingSuccessOmni object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokensForwardingSuccessOmni(name string, propertyId string, transactionType string, createdByTransactionId string, amount string) *TokensForwardingSuccessOmni {
	this := TokensForwardingSuccessOmni{}
	this.Name = name
	this.PropertyId = propertyId
	this.TransactionType = transactionType
	this.CreatedByTransactionId = createdByTransactionId
	this.Amount = amount
	return &this
}

// NewTokensForwardingSuccessOmniWithDefaults instantiates a new TokensForwardingSuccessOmni object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokensForwardingSuccessOmniWithDefaults() *TokensForwardingSuccessOmni {
	this := TokensForwardingSuccessOmni{}
	return &this
}

// GetName returns the Name field value
func (o *TokensForwardingSuccessOmni) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessOmni) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokensForwardingSuccessOmni) SetName(v string) {
	o.Name = v
}

// GetPropertyId returns the PropertyId field value
func (o *TokensForwardingSuccessOmni) GetPropertyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessOmni) GetPropertyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *TokensForwardingSuccessOmni) SetPropertyId(v string) {
	o.PropertyId = v
}

// GetTransactionType returns the TransactionType field value
func (o *TokensForwardingSuccessOmni) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessOmni) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *TokensForwardingSuccessOmni) SetTransactionType(v string) {
	o.TransactionType = v
}

// GetCreatedByTransactionId returns the CreatedByTransactionId field value
func (o *TokensForwardingSuccessOmni) GetCreatedByTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedByTransactionId
}

// GetCreatedByTransactionIdOk returns a tuple with the CreatedByTransactionId field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessOmni) GetCreatedByTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByTransactionId, true
}

// SetCreatedByTransactionId sets field value
func (o *TokensForwardingSuccessOmni) SetCreatedByTransactionId(v string) {
	o.CreatedByTransactionId = v
}

// GetAmount returns the Amount field value
func (o *TokensForwardingSuccessOmni) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessOmni) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokensForwardingSuccessOmni) SetAmount(v string) {
	o.Amount = v
}

func (o TokensForwardingSuccessOmni) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["propertyId"] = o.PropertyId
	}
	if true {
		toSerialize["transactionType"] = o.TransactionType
	}
	if true {
		toSerialize["createdByTransactionId"] = o.CreatedByTransactionId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableTokensForwardingSuccessOmni struct {
	value *TokensForwardingSuccessOmni
	isSet bool
}

func (v NullableTokensForwardingSuccessOmni) Get() *TokensForwardingSuccessOmni {
	return v.value
}

func (v *NullableTokensForwardingSuccessOmni) Set(val *TokensForwardingSuccessOmni) {
	v.value = val
	v.isSet = true
}

func (v NullableTokensForwardingSuccessOmni) IsSet() bool {
	return v.isSet
}

func (v *NullableTokensForwardingSuccessOmni) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokensForwardingSuccessOmni(val *TokensForwardingSuccessOmni) *NullableTokensForwardingSuccessOmni {
	return &NullableTokensForwardingSuccessOmni{value: val, isSet: true}
}

func (v NullableTokensForwardingSuccessOmni) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokensForwardingSuccessOmni) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


