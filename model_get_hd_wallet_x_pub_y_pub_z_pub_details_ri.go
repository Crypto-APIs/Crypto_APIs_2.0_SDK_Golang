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

// GetHDWalletXPubYPubZPubDetailsRI struct for GetHDWalletXPubYPubZPubDetailsRI
type GetHDWalletXPubYPubZPubDetailsRI struct {
	// Specifies the confirmed coins balance of the Wallet.
	ConfirmedBalance string `json:"confirmedBalance"`
	// Defines the total currency received to the Wallet.
	TotalReceived *string `json:"totalReceived,omitempty"`
	// Defines the total currency spent from the Wallet.
	TotalSpent *string `json:"totalSpent,omitempty"`
}

// NewGetHDWalletXPubYPubZPubDetailsRI instantiates a new GetHDWalletXPubYPubZPubDetailsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHDWalletXPubYPubZPubDetailsRI(confirmedBalance string) *GetHDWalletXPubYPubZPubDetailsRI {
	this := GetHDWalletXPubYPubZPubDetailsRI{}
	this.ConfirmedBalance = confirmedBalance
	return &this
}

// NewGetHDWalletXPubYPubZPubDetailsRIWithDefaults instantiates a new GetHDWalletXPubYPubZPubDetailsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHDWalletXPubYPubZPubDetailsRIWithDefaults() *GetHDWalletXPubYPubZPubDetailsRI {
	this := GetHDWalletXPubYPubZPubDetailsRI{}
	return &this
}

// GetConfirmedBalance returns the ConfirmedBalance field value
func (o *GetHDWalletXPubYPubZPubDetailsRI) GetConfirmedBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmedBalance
}

// GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field value
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubDetailsRI) GetConfirmedBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmedBalance, true
}

// SetConfirmedBalance sets field value
func (o *GetHDWalletXPubYPubZPubDetailsRI) SetConfirmedBalance(v string) {
	o.ConfirmedBalance = v
}

// GetTotalReceived returns the TotalReceived field value if set, zero value otherwise.
func (o *GetHDWalletXPubYPubZPubDetailsRI) GetTotalReceived() string {
	if o == nil || o.TotalReceived == nil {
		var ret string
		return ret
	}
	return *o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubDetailsRI) GetTotalReceivedOk() (*string, bool) {
	if o == nil || o.TotalReceived == nil {
		return nil, false
	}
	return o.TotalReceived, true
}

// HasTotalReceived returns a boolean if a field has been set.
func (o *GetHDWalletXPubYPubZPubDetailsRI) HasTotalReceived() bool {
	if o != nil && o.TotalReceived != nil {
		return true
	}

	return false
}

// SetTotalReceived gets a reference to the given string and assigns it to the TotalReceived field.
func (o *GetHDWalletXPubYPubZPubDetailsRI) SetTotalReceived(v string) {
	o.TotalReceived = &v
}

// GetTotalSpent returns the TotalSpent field value if set, zero value otherwise.
func (o *GetHDWalletXPubYPubZPubDetailsRI) GetTotalSpent() string {
	if o == nil || o.TotalSpent == nil {
		var ret string
		return ret
	}
	return *o.TotalSpent
}

// GetTotalSpentOk returns a tuple with the TotalSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubDetailsRI) GetTotalSpentOk() (*string, bool) {
	if o == nil || o.TotalSpent == nil {
		return nil, false
	}
	return o.TotalSpent, true
}

// HasTotalSpent returns a boolean if a field has been set.
func (o *GetHDWalletXPubYPubZPubDetailsRI) HasTotalSpent() bool {
	if o != nil && o.TotalSpent != nil {
		return true
	}

	return false
}

// SetTotalSpent gets a reference to the given string and assigns it to the TotalSpent field.
func (o *GetHDWalletXPubYPubZPubDetailsRI) SetTotalSpent(v string) {
	o.TotalSpent = &v
}

func (o GetHDWalletXPubYPubZPubDetailsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["confirmedBalance"] = o.ConfirmedBalance
	}
	if o.TotalReceived != nil {
		toSerialize["totalReceived"] = o.TotalReceived
	}
	if o.TotalSpent != nil {
		toSerialize["totalSpent"] = o.TotalSpent
	}
	return json.Marshal(toSerialize)
}

type NullableGetHDWalletXPubYPubZPubDetailsRI struct {
	value *GetHDWalletXPubYPubZPubDetailsRI
	isSet bool
}

func (v NullableGetHDWalletXPubYPubZPubDetailsRI) Get() *GetHDWalletXPubYPubZPubDetailsRI {
	return v.value
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsRI) Set(val *GetHDWalletXPubYPubZPubDetailsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHDWalletXPubYPubZPubDetailsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHDWalletXPubYPubZPubDetailsRI(val *GetHDWalletXPubYPubZPubDetailsRI) *NullableGetHDWalletXPubYPubZPubDetailsRI {
	return &NullableGetHDWalletXPubYPubZPubDetailsRI{value: val, isSet: true}
}

func (v NullableGetHDWalletXPubYPubZPubDetailsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


