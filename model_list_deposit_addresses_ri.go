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

// ListDepositAddressesRI struct for ListDepositAddressesRI
type ListDepositAddressesRI struct {
	// Specifies the specific address's unique string value.
	Address string `json:"address"`
	ConfirmedBalance ListDepositAddressesRIConfirmedBalance `json:"confirmedBalance"`
	// Defines the specific UNIX time when the deposit address was created.
	CreatedTimestamp int32 `json:"createdTimestamp"`
	// Represents fungible tokens'es detailed information
	FungibleTokens []ListDepositAddressesRIFungibleTokensInner `json:"fungibleTokens"`
	// Represents the index of the address in the wallet.
	Index string `json:"index"`
	// Represents a custom tag that customers can set up for their Wallets and addresses. E.g. custom label named \"Special addresses\".
	Label string `json:"label"`
	// Represents non-fungible tokens'es detailed information.
	NonFungibleTokens []ListDepositAddressesRINonFungibleTokensInner `json:"nonFungibleTokens"`
}

// NewListDepositAddressesRI instantiates a new ListDepositAddressesRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDepositAddressesRI(address string, confirmedBalance ListDepositAddressesRIConfirmedBalance, createdTimestamp int32, fungibleTokens []ListDepositAddressesRIFungibleTokensInner, index string, label string, nonFungibleTokens []ListDepositAddressesRINonFungibleTokensInner) *ListDepositAddressesRI {
	this := ListDepositAddressesRI{}
	this.Address = address
	this.ConfirmedBalance = confirmedBalance
	this.CreatedTimestamp = createdTimestamp
	this.FungibleTokens = fungibleTokens
	this.Index = index
	this.Label = label
	this.NonFungibleTokens = nonFungibleTokens
	return &this
}

// NewListDepositAddressesRIWithDefaults instantiates a new ListDepositAddressesRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDepositAddressesRIWithDefaults() *ListDepositAddressesRI {
	this := ListDepositAddressesRI{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListDepositAddressesRI) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRI) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListDepositAddressesRI) SetAddress(v string) {
	o.Address = v
}

// GetConfirmedBalance returns the ConfirmedBalance field value
func (o *ListDepositAddressesRI) GetConfirmedBalance() ListDepositAddressesRIConfirmedBalance {
	if o == nil {
		var ret ListDepositAddressesRIConfirmedBalance
		return ret
	}

	return o.ConfirmedBalance
}

// GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRI) GetConfirmedBalanceOk() (*ListDepositAddressesRIConfirmedBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmedBalance, true
}

// SetConfirmedBalance sets field value
func (o *ListDepositAddressesRI) SetConfirmedBalance(v ListDepositAddressesRIConfirmedBalance) {
	o.ConfirmedBalance = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *ListDepositAddressesRI) GetCreatedTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRI) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *ListDepositAddressesRI) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = v
}

// GetFungibleTokens returns the FungibleTokens field value
func (o *ListDepositAddressesRI) GetFungibleTokens() []ListDepositAddressesRIFungibleTokensInner {
	if o == nil {
		var ret []ListDepositAddressesRIFungibleTokensInner
		return ret
	}

	return o.FungibleTokens
}

// GetFungibleTokensOk returns a tuple with the FungibleTokens field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRI) GetFungibleTokensOk() ([]ListDepositAddressesRIFungibleTokensInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.FungibleTokens, true
}

// SetFungibleTokens sets field value
func (o *ListDepositAddressesRI) SetFungibleTokens(v []ListDepositAddressesRIFungibleTokensInner) {
	o.FungibleTokens = v
}

// GetIndex returns the Index field value
func (o *ListDepositAddressesRI) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRI) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ListDepositAddressesRI) SetIndex(v string) {
	o.Index = v
}

// GetLabel returns the Label field value
func (o *ListDepositAddressesRI) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRI) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ListDepositAddressesRI) SetLabel(v string) {
	o.Label = v
}

// GetNonFungibleTokens returns the NonFungibleTokens field value
func (o *ListDepositAddressesRI) GetNonFungibleTokens() []ListDepositAddressesRINonFungibleTokensInner {
	if o == nil {
		var ret []ListDepositAddressesRINonFungibleTokensInner
		return ret
	}

	return o.NonFungibleTokens
}

// GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field value
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesRI) GetNonFungibleTokensOk() ([]ListDepositAddressesRINonFungibleTokensInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonFungibleTokens, true
}

// SetNonFungibleTokens sets field value
func (o *ListDepositAddressesRI) SetNonFungibleTokens(v []ListDepositAddressesRINonFungibleTokensInner) {
	o.NonFungibleTokens = v
}

func (o ListDepositAddressesRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["confirmedBalance"] = o.ConfirmedBalance
	}
	if true {
		toSerialize["createdTimestamp"] = o.CreatedTimestamp
	}
	if true {
		toSerialize["fungibleTokens"] = o.FungibleTokens
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["nonFungibleTokens"] = o.NonFungibleTokens
	}
	return json.Marshal(toSerialize)
}

type NullableListDepositAddressesRI struct {
	value *ListDepositAddressesRI
	isSet bool
}

func (v NullableListDepositAddressesRI) Get() *ListDepositAddressesRI {
	return v.value
}

func (v *NullableListDepositAddressesRI) Set(val *ListDepositAddressesRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListDepositAddressesRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListDepositAddressesRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDepositAddressesRI(val *ListDepositAddressesRI) *NullableListDepositAddressesRI {
	return &NullableListDepositAddressesRI{value: val, isSet: true}
}

func (v NullableListDepositAddressesRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDepositAddressesRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


