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

// ListWalletTransactionsRIInternalTransactions struct for ListWalletTransactionsRIInternalTransactions
type ListWalletTransactionsRIInternalTransactions struct {
	// Defines the specific amount of the transaction.
	Amount string `json:"amount"`
	// Represents the converted amount.
	ConvertedAmount string `json:"convertedAmount"`
	// Defines the base asset symbol to get a rate for.
	ExchangeRateUnit string `json:"exchangeRateUnit"`
	// Represents the unique internal transaction ID in regards to the parent transaction (type trace address).
	OperationId string `json:"operationId"`
	// Represents the recipient address with the respective amount.
	Recipient string `json:"recipient"`
	// Represents the sender address with the respective amount.
	Sender string `json:"sender"`
	// Represents the unique unit symbol.
	Symbol string `json:"symbol"`
}

// NewListWalletTransactionsRIInternalTransactions instantiates a new ListWalletTransactionsRIInternalTransactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWalletTransactionsRIInternalTransactions(amount string, convertedAmount string, exchangeRateUnit string, operationId string, recipient string, sender string, symbol string) *ListWalletTransactionsRIInternalTransactions {
	this := ListWalletTransactionsRIInternalTransactions{}
	this.Amount = amount
	this.ConvertedAmount = convertedAmount
	this.ExchangeRateUnit = exchangeRateUnit
	this.OperationId = operationId
	this.Recipient = recipient
	this.Sender = sender
	this.Symbol = symbol
	return &this
}

// NewListWalletTransactionsRIInternalTransactionsWithDefaults instantiates a new ListWalletTransactionsRIInternalTransactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWalletTransactionsRIInternalTransactionsWithDefaults() *ListWalletTransactionsRIInternalTransactions {
	this := ListWalletTransactionsRIInternalTransactions{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListWalletTransactionsRIInternalTransactions) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIInternalTransactions) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListWalletTransactionsRIInternalTransactions) SetAmount(v string) {
	o.Amount = v
}

// GetConvertedAmount returns the ConvertedAmount field value
func (o *ListWalletTransactionsRIInternalTransactions) GetConvertedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConvertedAmount
}

// GetConvertedAmountOk returns a tuple with the ConvertedAmount field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIInternalTransactions) GetConvertedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConvertedAmount, true
}

// SetConvertedAmount sets field value
func (o *ListWalletTransactionsRIInternalTransactions) SetConvertedAmount(v string) {
	o.ConvertedAmount = v
}

// GetExchangeRateUnit returns the ExchangeRateUnit field value
func (o *ListWalletTransactionsRIInternalTransactions) GetExchangeRateUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeRateUnit
}

// GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIInternalTransactions) GetExchangeRateUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeRateUnit, true
}

// SetExchangeRateUnit sets field value
func (o *ListWalletTransactionsRIInternalTransactions) SetExchangeRateUnit(v string) {
	o.ExchangeRateUnit = v
}

// GetOperationId returns the OperationId field value
func (o *ListWalletTransactionsRIInternalTransactions) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIInternalTransactions) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ListWalletTransactionsRIInternalTransactions) SetOperationId(v string) {
	o.OperationId = v
}

// GetRecipient returns the Recipient field value
func (o *ListWalletTransactionsRIInternalTransactions) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIInternalTransactions) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *ListWalletTransactionsRIInternalTransactions) SetRecipient(v string) {
	o.Recipient = v
}

// GetSender returns the Sender field value
func (o *ListWalletTransactionsRIInternalTransactions) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIInternalTransactions) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *ListWalletTransactionsRIInternalTransactions) SetSender(v string) {
	o.Sender = v
}

// GetSymbol returns the Symbol field value
func (o *ListWalletTransactionsRIInternalTransactions) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIInternalTransactions) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ListWalletTransactionsRIInternalTransactions) SetSymbol(v string) {
	o.Symbol = v
}

func (o ListWalletTransactionsRIInternalTransactions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["convertedAmount"] = o.ConvertedAmount
	}
	if true {
		toSerialize["exchangeRateUnit"] = o.ExchangeRateUnit
	}
	if true {
		toSerialize["operationId"] = o.OperationId
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableListWalletTransactionsRIInternalTransactions struct {
	value *ListWalletTransactionsRIInternalTransactions
	isSet bool
}

func (v NullableListWalletTransactionsRIInternalTransactions) Get() *ListWalletTransactionsRIInternalTransactions {
	return v.value
}

func (v *NullableListWalletTransactionsRIInternalTransactions) Set(val *ListWalletTransactionsRIInternalTransactions) {
	v.value = val
	v.isSet = true
}

func (v NullableListWalletTransactionsRIInternalTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullableListWalletTransactionsRIInternalTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWalletTransactionsRIInternalTransactions(val *ListWalletTransactionsRIInternalTransactions) *NullableListWalletTransactionsRIInternalTransactions {
	return &NullableListWalletTransactionsRIInternalTransactions{value: val, isSet: true}
}

func (v NullableListWalletTransactionsRIInternalTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWalletTransactionsRIInternalTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


