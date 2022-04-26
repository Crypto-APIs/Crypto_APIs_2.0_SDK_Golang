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

// ListWalletTransactionsRIValue struct for ListWalletTransactionsRIValue
type ListWalletTransactionsRIValue struct {
	// Defines the transaction amount.
	Amount string `json:"amount"`
	// Defines the converted amount of the transaction as a string.
	ConvertedAmount string `json:"convertedAmount"`
	// Defines the exchange rate's unit.
	ExchangeRateUnit string `json:"exchangeRateUnit"`
	// Defines the unit of the transaction's amount.
	Symbol string `json:"symbol"`
}

// NewListWalletTransactionsRIValue instantiates a new ListWalletTransactionsRIValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWalletTransactionsRIValue(amount string, convertedAmount string, exchangeRateUnit string, symbol string) *ListWalletTransactionsRIValue {
	this := ListWalletTransactionsRIValue{}
	this.Amount = amount
	this.ConvertedAmount = convertedAmount
	this.ExchangeRateUnit = exchangeRateUnit
	this.Symbol = symbol
	return &this
}

// NewListWalletTransactionsRIValueWithDefaults instantiates a new ListWalletTransactionsRIValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWalletTransactionsRIValueWithDefaults() *ListWalletTransactionsRIValue {
	this := ListWalletTransactionsRIValue{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListWalletTransactionsRIValue) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIValue) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListWalletTransactionsRIValue) SetAmount(v string) {
	o.Amount = v
}

// GetConvertedAmount returns the ConvertedAmount field value
func (o *ListWalletTransactionsRIValue) GetConvertedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConvertedAmount
}

// GetConvertedAmountOk returns a tuple with the ConvertedAmount field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIValue) GetConvertedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConvertedAmount, true
}

// SetConvertedAmount sets field value
func (o *ListWalletTransactionsRIValue) SetConvertedAmount(v string) {
	o.ConvertedAmount = v
}

// GetExchangeRateUnit returns the ExchangeRateUnit field value
func (o *ListWalletTransactionsRIValue) GetExchangeRateUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeRateUnit
}

// GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIValue) GetExchangeRateUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeRateUnit, true
}

// SetExchangeRateUnit sets field value
func (o *ListWalletTransactionsRIValue) SetExchangeRateUnit(v string) {
	o.ExchangeRateUnit = v
}

// GetSymbol returns the Symbol field value
func (o *ListWalletTransactionsRIValue) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIValue) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ListWalletTransactionsRIValue) SetSymbol(v string) {
	o.Symbol = v
}

func (o ListWalletTransactionsRIValue) MarshalJSON() ([]byte, error) {
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
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableListWalletTransactionsRIValue struct {
	value *ListWalletTransactionsRIValue
	isSet bool
}

func (v NullableListWalletTransactionsRIValue) Get() *ListWalletTransactionsRIValue {
	return v.value
}

func (v *NullableListWalletTransactionsRIValue) Set(val *ListWalletTransactionsRIValue) {
	v.value = val
	v.isSet = true
}

func (v NullableListWalletTransactionsRIValue) IsSet() bool {
	return v.isSet
}

func (v *NullableListWalletTransactionsRIValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWalletTransactionsRIValue(val *ListWalletTransactionsRIValue) *NullableListWalletTransactionsRIValue {
	return &NullableListWalletTransactionsRIValue{value: val, isSet: true}
}

func (v NullableListWalletTransactionsRIValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWalletTransactionsRIValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


