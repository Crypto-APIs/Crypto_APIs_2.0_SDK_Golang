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

// ListWalletTransactionsRIFungibleTokensInner struct for ListWalletTransactionsRIFungibleTokensInner
type ListWalletTransactionsRIFungibleTokensInner struct {
	// Defines the amount of the fungible tokens.
	Amount string `json:"amount"`
	// Defines the tokens' converted amount value.
	ConvertedAmount string `json:"convertedAmount"`
	// Represents token's exchange rate unit.
	ExchangeRateUnit string `json:"exchangeRateUnit"`
	// Defines the token's name as a string.
	Name string `json:"name"`
	// The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one
	Recipient string `json:"recipient"`
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Sender string `json:"sender"`
	// Defines the symbol of the fungible tokens.
	Symbol string `json:"symbol"`
	// Defines the decimals of the token, i.e. the number of digits that come after the decimal coma of the token.
	TokenDecimals int32 `json:"tokenDecimals"`
	// Defines the specific token type.
	Type string `json:"type"`
}

// NewListWalletTransactionsRIFungibleTokensInner instantiates a new ListWalletTransactionsRIFungibleTokensInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWalletTransactionsRIFungibleTokensInner(amount string, convertedAmount string, exchangeRateUnit string, name string, recipient string, sender string, symbol string, tokenDecimals int32, type_ string) *ListWalletTransactionsRIFungibleTokensInner {
	this := ListWalletTransactionsRIFungibleTokensInner{}
	this.Amount = amount
	this.ConvertedAmount = convertedAmount
	this.ExchangeRateUnit = exchangeRateUnit
	this.Name = name
	this.Recipient = recipient
	this.Sender = sender
	this.Symbol = symbol
	this.TokenDecimals = tokenDecimals
	this.Type = type_
	return &this
}

// NewListWalletTransactionsRIFungibleTokensInnerWithDefaults instantiates a new ListWalletTransactionsRIFungibleTokensInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWalletTransactionsRIFungibleTokensInnerWithDefaults() *ListWalletTransactionsRIFungibleTokensInner {
	this := ListWalletTransactionsRIFungibleTokensInner{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetAmount(v string) {
	o.Amount = v
}

// GetConvertedAmount returns the ConvertedAmount field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetConvertedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConvertedAmount
}

// GetConvertedAmountOk returns a tuple with the ConvertedAmount field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetConvertedAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConvertedAmount, true
}

// SetConvertedAmount sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetConvertedAmount(v string) {
	o.ConvertedAmount = v
}

// GetExchangeRateUnit returns the ExchangeRateUnit field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetExchangeRateUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeRateUnit
}

// GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetExchangeRateUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeRateUnit, true
}

// SetExchangeRateUnit sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetExchangeRateUnit(v string) {
	o.ExchangeRateUnit = v
}

// GetName returns the Name field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetName(v string) {
	o.Name = v
}

// GetRecipient returns the Recipient field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetRecipient(v string) {
	o.Recipient = v
}

// GetSender returns the Sender field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetSender(v string) {
	o.Sender = v
}

// GetSymbol returns the Symbol field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokenDecimals returns the TokenDecimals field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetTokenDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenDecimals
}

// GetTokenDecimalsOk returns a tuple with the TokenDecimals field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetTokenDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenDecimals, true
}

// SetTokenDecimals sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetTokenDecimals(v int32) {
	o.TokenDecimals = v
}

// GetType returns the Type field value
func (o *ListWalletTransactionsRIFungibleTokensInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRIFungibleTokensInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListWalletTransactionsRIFungibleTokensInner) SetType(v string) {
	o.Type = v
}

func (o ListWalletTransactionsRIFungibleTokensInner) MarshalJSON() ([]byte, error) {
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
		toSerialize["name"] = o.Name
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
	if true {
		toSerialize["tokenDecimals"] = o.TokenDecimals
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableListWalletTransactionsRIFungibleTokensInner struct {
	value *ListWalletTransactionsRIFungibleTokensInner
	isSet bool
}

func (v NullableListWalletTransactionsRIFungibleTokensInner) Get() *ListWalletTransactionsRIFungibleTokensInner {
	return v.value
}

func (v *NullableListWalletTransactionsRIFungibleTokensInner) Set(val *ListWalletTransactionsRIFungibleTokensInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWalletTransactionsRIFungibleTokensInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWalletTransactionsRIFungibleTokensInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWalletTransactionsRIFungibleTokensInner(val *ListWalletTransactionsRIFungibleTokensInner) *NullableListWalletTransactionsRIFungibleTokensInner {
	return &NullableListWalletTransactionsRIFungibleTokensInner{value: val, isSet: true}
}

func (v NullableListWalletTransactionsRIFungibleTokensInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWalletTransactionsRIFungibleTokensInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

