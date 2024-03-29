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

// DeriveAndSyncNewReceivingAddressesRI struct for DeriveAndSyncNewReceivingAddressesRI
type DeriveAndSyncNewReceivingAddressesRI struct {
	// Represents the public address, which is a compressed and shortened form of a public key.
	Address string `json:"address"`
	// Represents the format of the address.
	AddressFormat string `json:"addressFormat"`
	// Defines the address type.
	AddressType string `json:"addressType"`
	// Represents the derivation type
	DerivationType string `json:"derivationType"`
	// Represents the output index. It refers to the UTXO sequence in the transaction outputs (vout).
	Index string `json:"index"`
}

// NewDeriveAndSyncNewReceivingAddressesRI instantiates a new DeriveAndSyncNewReceivingAddressesRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeriveAndSyncNewReceivingAddressesRI(address string, addressFormat string, addressType string, derivationType string, index string) *DeriveAndSyncNewReceivingAddressesRI {
	this := DeriveAndSyncNewReceivingAddressesRI{}
	this.Address = address
	this.AddressFormat = addressFormat
	this.AddressType = addressType
	this.DerivationType = derivationType
	this.Index = index
	return &this
}

// NewDeriveAndSyncNewReceivingAddressesRIWithDefaults instantiates a new DeriveAndSyncNewReceivingAddressesRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeriveAndSyncNewReceivingAddressesRIWithDefaults() *DeriveAndSyncNewReceivingAddressesRI {
	this := DeriveAndSyncNewReceivingAddressesRI{}
	return &this
}

// GetAddress returns the Address field value
func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DeriveAndSyncNewReceivingAddressesRI) SetAddress(v string) {
	o.Address = v
}

// GetAddressFormat returns the AddressFormat field value
func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressFormat
}

// GetAddressFormatOk returns a tuple with the AddressFormat field value
// and a boolean to check if the value has been set.
func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressFormat, true
}

// SetAddressFormat sets field value
func (o *DeriveAndSyncNewReceivingAddressesRI) SetAddressFormat(v string) {
	o.AddressFormat = v
}

// GetAddressType returns the AddressType field value
func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressType
}

// GetAddressTypeOk returns a tuple with the AddressType field value
// and a boolean to check if the value has been set.
func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressType, true
}

// SetAddressType sets field value
func (o *DeriveAndSyncNewReceivingAddressesRI) SetAddressType(v string) {
	o.AddressType = v
}

// GetDerivationType returns the DerivationType field value
func (o *DeriveAndSyncNewReceivingAddressesRI) GetDerivationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DerivationType
}

// GetDerivationTypeOk returns a tuple with the DerivationType field value
// and a boolean to check if the value has been set.
func (o *DeriveAndSyncNewReceivingAddressesRI) GetDerivationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DerivationType, true
}

// SetDerivationType sets field value
func (o *DeriveAndSyncNewReceivingAddressesRI) SetDerivationType(v string) {
	o.DerivationType = v
}

// GetIndex returns the Index field value
func (o *DeriveAndSyncNewReceivingAddressesRI) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *DeriveAndSyncNewReceivingAddressesRI) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *DeriveAndSyncNewReceivingAddressesRI) SetIndex(v string) {
	o.Index = v
}

func (o DeriveAndSyncNewReceivingAddressesRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["addressFormat"] = o.AddressFormat
	}
	if true {
		toSerialize["addressType"] = o.AddressType
	}
	if true {
		toSerialize["derivationType"] = o.DerivationType
	}
	if true {
		toSerialize["index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

type NullableDeriveAndSyncNewReceivingAddressesRI struct {
	value *DeriveAndSyncNewReceivingAddressesRI
	isSet bool
}

func (v NullableDeriveAndSyncNewReceivingAddressesRI) Get() *DeriveAndSyncNewReceivingAddressesRI {
	return v.value
}

func (v *NullableDeriveAndSyncNewReceivingAddressesRI) Set(val *DeriveAndSyncNewReceivingAddressesRI) {
	v.value = val
	v.isSet = true
}

func (v NullableDeriveAndSyncNewReceivingAddressesRI) IsSet() bool {
	return v.isSet
}

func (v *NullableDeriveAndSyncNewReceivingAddressesRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeriveAndSyncNewReceivingAddressesRI(val *DeriveAndSyncNewReceivingAddressesRI) *NullableDeriveAndSyncNewReceivingAddressesRI {
	return &NullableDeriveAndSyncNewReceivingAddressesRI{value: val, isSet: true}
}

func (v NullableDeriveAndSyncNewReceivingAddressesRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeriveAndSyncNewReceivingAddressesRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


