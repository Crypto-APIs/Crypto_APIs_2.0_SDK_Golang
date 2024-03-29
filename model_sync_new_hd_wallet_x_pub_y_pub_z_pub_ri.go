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

// SyncNewHDWalletXPubYPubZPubRI struct for SyncNewHDWalletXPubYPubZPubRI
type SyncNewHDWalletXPubYPubZPubRI struct {
	// Defines the account extended publicly known key which is used to derive all child public keys.
	ExtendedPublicKey string `json:"extendedPublicKey"`
}

// NewSyncNewHDWalletXPubYPubZPubRI instantiates a new SyncNewHDWalletXPubYPubZPubRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncNewHDWalletXPubYPubZPubRI(extendedPublicKey string) *SyncNewHDWalletXPubYPubZPubRI {
	this := SyncNewHDWalletXPubYPubZPubRI{}
	this.ExtendedPublicKey = extendedPublicKey
	return &this
}

// NewSyncNewHDWalletXPubYPubZPubRIWithDefaults instantiates a new SyncNewHDWalletXPubYPubZPubRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncNewHDWalletXPubYPubZPubRIWithDefaults() *SyncNewHDWalletXPubYPubZPubRI {
	this := SyncNewHDWalletXPubYPubZPubRI{}
	return &this
}

// GetExtendedPublicKey returns the ExtendedPublicKey field value
func (o *SyncNewHDWalletXPubYPubZPubRI) GetExtendedPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtendedPublicKey
}

// GetExtendedPublicKeyOk returns a tuple with the ExtendedPublicKey field value
// and a boolean to check if the value has been set.
func (o *SyncNewHDWalletXPubYPubZPubRI) GetExtendedPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtendedPublicKey, true
}

// SetExtendedPublicKey sets field value
func (o *SyncNewHDWalletXPubYPubZPubRI) SetExtendedPublicKey(v string) {
	o.ExtendedPublicKey = v
}

func (o SyncNewHDWalletXPubYPubZPubRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extendedPublicKey"] = o.ExtendedPublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableSyncNewHDWalletXPubYPubZPubRI struct {
	value *SyncNewHDWalletXPubYPubZPubRI
	isSet bool
}

func (v NullableSyncNewHDWalletXPubYPubZPubRI) Get() *SyncNewHDWalletXPubYPubZPubRI {
	return v.value
}

func (v *NullableSyncNewHDWalletXPubYPubZPubRI) Set(val *SyncNewHDWalletXPubYPubZPubRI) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncNewHDWalletXPubYPubZPubRI) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncNewHDWalletXPubYPubZPubRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncNewHDWalletXPubYPubZPubRI(val *SyncNewHDWalletXPubYPubZPubRI) *NullableSyncNewHDWalletXPubYPubZPubRI {
	return &NullableSyncNewHDWalletXPubYPubZPubRI{value: val, isSet: true}
}

func (v NullableSyncNewHDWalletXPubYPubZPubRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncNewHDWalletXPubYPubZPubRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


