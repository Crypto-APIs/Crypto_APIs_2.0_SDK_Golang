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

// SyncHDWalletXPubYPubZPubRI struct for SyncHDWalletXPubYPubZPubRI
type SyncHDWalletXPubYPubZPubRI struct {
	// Defines the account extended publicly known key which is used to derive all child public keys.
	ExtendedPublicKey string `json:"extendedPublicKey"`
}

// NewSyncHDWalletXPubYPubZPubRI instantiates a new SyncHDWalletXPubYPubZPubRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncHDWalletXPubYPubZPubRI(extendedPublicKey string) *SyncHDWalletXPubYPubZPubRI {
	this := SyncHDWalletXPubYPubZPubRI{}
	this.ExtendedPublicKey = extendedPublicKey
	return &this
}

// NewSyncHDWalletXPubYPubZPubRIWithDefaults instantiates a new SyncHDWalletXPubYPubZPubRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncHDWalletXPubYPubZPubRIWithDefaults() *SyncHDWalletXPubYPubZPubRI {
	this := SyncHDWalletXPubYPubZPubRI{}
	return &this
}

// GetExtendedPublicKey returns the ExtendedPublicKey field value
func (o *SyncHDWalletXPubYPubZPubRI) GetExtendedPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtendedPublicKey
}

// GetExtendedPublicKeyOk returns a tuple with the ExtendedPublicKey field value
// and a boolean to check if the value has been set.
func (o *SyncHDWalletXPubYPubZPubRI) GetExtendedPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtendedPublicKey, true
}

// SetExtendedPublicKey sets field value
func (o *SyncHDWalletXPubYPubZPubRI) SetExtendedPublicKey(v string) {
	o.ExtendedPublicKey = v
}

func (o SyncHDWalletXPubYPubZPubRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extendedPublicKey"] = o.ExtendedPublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableSyncHDWalletXPubYPubZPubRI struct {
	value *SyncHDWalletXPubYPubZPubRI
	isSet bool
}

func (v NullableSyncHDWalletXPubYPubZPubRI) Get() *SyncHDWalletXPubYPubZPubRI {
	return v.value
}

func (v *NullableSyncHDWalletXPubYPubZPubRI) Set(val *SyncHDWalletXPubYPubZPubRI) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncHDWalletXPubYPubZPubRI) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncHDWalletXPubYPubZPubRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncHDWalletXPubYPubZPubRI(val *SyncHDWalletXPubYPubZPubRI) *NullableSyncHDWalletXPubYPubZPubRI {
	return &NullableSyncHDWalletXPubYPubZPubRI{value: val, isSet: true}
}

func (v NullableSyncHDWalletXPubYPubZPubRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncHDWalletXPubYPubZPubRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


