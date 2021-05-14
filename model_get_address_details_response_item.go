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

// GetAddressDetailsResponseItem struct for GetAddressDetailsResponseItem
type GetAddressDetailsResponseItem struct {
	// Represents the total number of confirmed coins transactions for this address, both incoming and outgoing. Applies for coins only **and not** tokens transfers e.g. for Ethereum. `transactionsCount` could result as less than incoming and outgoing transactions put together (e.g. in Bitcoin), due to the fact that one and the same address could be in senders and receivers addresses.
	TransactionsCount int32 `json:"transactionsCount"`
	ConfirmedBalance GetAddressDetailsResponseItemConfirmedBalance `json:"confirmedBalance"`
	TotalReceived GetAddressDetailsResponseItemTotalReceived `json:"totalReceived"`
	TotalSpent GetAddressDetailsResponseItemTotalSpent `json:"totalSpent"`
	// Defines the count of all confirmed incoming transactions from the address for coins. This applies to **coins** only, **not** to tokens transfers e.g. for Ethereum.
	IncomingTransactionsCount int32 `json:"incomingTransactionsCount"`
	// Defines the count of all confirmed outgoing transactions from the address for coins. This applies to **coins** only, **not** to tokens transfers e.g. for Ethereum.
	OutgoingTransactionsCount int32 `json:"outgoingTransactionsCount"`
}

// NewGetAddressDetailsResponseItem instantiates a new GetAddressDetailsResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAddressDetailsResponseItem(transactionsCount int32, confirmedBalance GetAddressDetailsResponseItemConfirmedBalance, totalReceived GetAddressDetailsResponseItemTotalReceived, totalSpent GetAddressDetailsResponseItemTotalSpent, incomingTransactionsCount int32, outgoingTransactionsCount int32) *GetAddressDetailsResponseItem {
	this := GetAddressDetailsResponseItem{}
	this.TransactionsCount = transactionsCount
	this.ConfirmedBalance = confirmedBalance
	this.TotalReceived = totalReceived
	this.TotalSpent = totalSpent
	this.IncomingTransactionsCount = incomingTransactionsCount
	this.OutgoingTransactionsCount = outgoingTransactionsCount
	return &this
}

// NewGetAddressDetailsResponseItemWithDefaults instantiates a new GetAddressDetailsResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAddressDetailsResponseItemWithDefaults() *GetAddressDetailsResponseItem {
	this := GetAddressDetailsResponseItem{}
	return &this
}

// GetTransactionsCount returns the TransactionsCount field value
func (o *GetAddressDetailsResponseItem) GetTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionsCount
}

// GetTransactionsCountOk returns a tuple with the TransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsResponseItem) GetTransactionsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionsCount, true
}

// SetTransactionsCount sets field value
func (o *GetAddressDetailsResponseItem) SetTransactionsCount(v int32) {
	o.TransactionsCount = v
}

// GetConfirmedBalance returns the ConfirmedBalance field value
func (o *GetAddressDetailsResponseItem) GetConfirmedBalance() GetAddressDetailsResponseItemConfirmedBalance {
	if o == nil {
		var ret GetAddressDetailsResponseItemConfirmedBalance
		return ret
	}

	return o.ConfirmedBalance
}

// GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsResponseItem) GetConfirmedBalanceOk() (*GetAddressDetailsResponseItemConfirmedBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmedBalance, true
}

// SetConfirmedBalance sets field value
func (o *GetAddressDetailsResponseItem) SetConfirmedBalance(v GetAddressDetailsResponseItemConfirmedBalance) {
	o.ConfirmedBalance = v
}

// GetTotalReceived returns the TotalReceived field value
func (o *GetAddressDetailsResponseItem) GetTotalReceived() GetAddressDetailsResponseItemTotalReceived {
	if o == nil {
		var ret GetAddressDetailsResponseItemTotalReceived
		return ret
	}

	return o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsResponseItem) GetTotalReceivedOk() (*GetAddressDetailsResponseItemTotalReceived, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalReceived, true
}

// SetTotalReceived sets field value
func (o *GetAddressDetailsResponseItem) SetTotalReceived(v GetAddressDetailsResponseItemTotalReceived) {
	o.TotalReceived = v
}

// GetTotalSpent returns the TotalSpent field value
func (o *GetAddressDetailsResponseItem) GetTotalSpent() GetAddressDetailsResponseItemTotalSpent {
	if o == nil {
		var ret GetAddressDetailsResponseItemTotalSpent
		return ret
	}

	return o.TotalSpent
}

// GetTotalSpentOk returns a tuple with the TotalSpent field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsResponseItem) GetTotalSpentOk() (*GetAddressDetailsResponseItemTotalSpent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalSpent, true
}

// SetTotalSpent sets field value
func (o *GetAddressDetailsResponseItem) SetTotalSpent(v GetAddressDetailsResponseItemTotalSpent) {
	o.TotalSpent = v
}

// GetIncomingTransactionsCount returns the IncomingTransactionsCount field value
func (o *GetAddressDetailsResponseItem) GetIncomingTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IncomingTransactionsCount
}

// GetIncomingTransactionsCountOk returns a tuple with the IncomingTransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsResponseItem) GetIncomingTransactionsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomingTransactionsCount, true
}

// SetIncomingTransactionsCount sets field value
func (o *GetAddressDetailsResponseItem) SetIncomingTransactionsCount(v int32) {
	o.IncomingTransactionsCount = v
}

// GetOutgoingTransactionsCount returns the OutgoingTransactionsCount field value
func (o *GetAddressDetailsResponseItem) GetOutgoingTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutgoingTransactionsCount
}

// GetOutgoingTransactionsCountOk returns a tuple with the OutgoingTransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsResponseItem) GetOutgoingTransactionsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OutgoingTransactionsCount, true
}

// SetOutgoingTransactionsCount sets field value
func (o *GetAddressDetailsResponseItem) SetOutgoingTransactionsCount(v int32) {
	o.OutgoingTransactionsCount = v
}

func (o GetAddressDetailsResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactionsCount"] = o.TransactionsCount
	}
	if true {
		toSerialize["confirmedBalance"] = o.ConfirmedBalance
	}
	if true {
		toSerialize["totalReceived"] = o.TotalReceived
	}
	if true {
		toSerialize["totalSpent"] = o.TotalSpent
	}
	if true {
		toSerialize["incomingTransactionsCount"] = o.IncomingTransactionsCount
	}
	if true {
		toSerialize["outgoingTransactionsCount"] = o.OutgoingTransactionsCount
	}
	return json.Marshal(toSerialize)
}

type NullableGetAddressDetailsResponseItem struct {
	value *GetAddressDetailsResponseItem
	isSet bool
}

func (v NullableGetAddressDetailsResponseItem) Get() *GetAddressDetailsResponseItem {
	return v.value
}

func (v *NullableGetAddressDetailsResponseItem) Set(val *GetAddressDetailsResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddressDetailsResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddressDetailsResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddressDetailsResponseItem(val *GetAddressDetailsResponseItem) *NullableGetAddressDetailsResponseItem {
	return &NullableGetAddressDetailsResponseItem{value: val, isSet: true}
}

func (v NullableGetAddressDetailsResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddressDetailsResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

