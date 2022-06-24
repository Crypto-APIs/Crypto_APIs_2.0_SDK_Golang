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

// ListWalletTransactionsRISendersInner struct for ListWalletTransactionsRISendersInner
type ListWalletTransactionsRISendersInner struct {
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Address string `json:"address"`
	// Represents the total amount sent by this address including the fee.
	Amount string `json:"amount"`
	// Defines a plain text string as a label for the sender.
	Label *string `json:"label,omitempty"`
}

// NewListWalletTransactionsRISendersInner instantiates a new ListWalletTransactionsRISendersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWalletTransactionsRISendersInner(address string, amount string) *ListWalletTransactionsRISendersInner {
	this := ListWalletTransactionsRISendersInner{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListWalletTransactionsRISendersInnerWithDefaults instantiates a new ListWalletTransactionsRISendersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWalletTransactionsRISendersInnerWithDefaults() *ListWalletTransactionsRISendersInner {
	this := ListWalletTransactionsRISendersInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListWalletTransactionsRISendersInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRISendersInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListWalletTransactionsRISendersInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListWalletTransactionsRISendersInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRISendersInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListWalletTransactionsRISendersInner) SetAmount(v string) {
	o.Amount = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ListWalletTransactionsRISendersInner) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWalletTransactionsRISendersInner) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ListWalletTransactionsRISendersInner) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ListWalletTransactionsRISendersInner) SetLabel(v string) {
	o.Label = &v
}

func (o ListWalletTransactionsRISendersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableListWalletTransactionsRISendersInner struct {
	value *ListWalletTransactionsRISendersInner
	isSet bool
}

func (v NullableListWalletTransactionsRISendersInner) Get() *ListWalletTransactionsRISendersInner {
	return v.value
}

func (v *NullableListWalletTransactionsRISendersInner) Set(val *ListWalletTransactionsRISendersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWalletTransactionsRISendersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWalletTransactionsRISendersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWalletTransactionsRISendersInner(val *ListWalletTransactionsRISendersInner) *NullableListWalletTransactionsRISendersInner {
	return &NullableListWalletTransactionsRISendersInner{value: val, isSet: true}
}

func (v NullableListWalletTransactionsRISendersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWalletTransactionsRISendersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


