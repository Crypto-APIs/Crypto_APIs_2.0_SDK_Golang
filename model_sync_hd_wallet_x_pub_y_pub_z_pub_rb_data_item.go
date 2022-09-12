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

// SyncHDWalletXPubYPubZPubRBDataItem struct for SyncHDWalletXPubYPubZPubRBDataItem
type SyncHDWalletXPubYPubZPubRBDataItem struct {
	// Defines the account extended publicly known key which is used to derive all child public keys.
	ExtendedPublicKey string `json:"extendedPublicKey"`
}

// NewSyncHDWalletXPubYPubZPubRBDataItem instantiates a new SyncHDWalletXPubYPubZPubRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncHDWalletXPubYPubZPubRBDataItem(extendedPublicKey string) *SyncHDWalletXPubYPubZPubRBDataItem {
	this := SyncHDWalletXPubYPubZPubRBDataItem{}
	this.ExtendedPublicKey = extendedPublicKey
	return &this
}

// NewSyncHDWalletXPubYPubZPubRBDataItemWithDefaults instantiates a new SyncHDWalletXPubYPubZPubRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncHDWalletXPubYPubZPubRBDataItemWithDefaults() *SyncHDWalletXPubYPubZPubRBDataItem {
	this := SyncHDWalletXPubYPubZPubRBDataItem{}
	return &this
}

// GetExtendedPublicKey returns the ExtendedPublicKey field value
func (o *SyncHDWalletXPubYPubZPubRBDataItem) GetExtendedPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtendedPublicKey
}

// GetExtendedPublicKeyOk returns a tuple with the ExtendedPublicKey field value
// and a boolean to check if the value has been set.
func (o *SyncHDWalletXPubYPubZPubRBDataItem) GetExtendedPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtendedPublicKey, true
}

// SetExtendedPublicKey sets field value
func (o *SyncHDWalletXPubYPubZPubRBDataItem) SetExtendedPublicKey(v string) {
	o.ExtendedPublicKey = v
}

func (o SyncHDWalletXPubYPubZPubRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extendedPublicKey"] = o.ExtendedPublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableSyncHDWalletXPubYPubZPubRBDataItem struct {
	value *SyncHDWalletXPubYPubZPubRBDataItem
	isSet bool
}

func (v NullableSyncHDWalletXPubYPubZPubRBDataItem) Get() *SyncHDWalletXPubYPubZPubRBDataItem {
	return v.value
}

func (v *NullableSyncHDWalletXPubYPubZPubRBDataItem) Set(val *SyncHDWalletXPubYPubZPubRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncHDWalletXPubYPubZPubRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncHDWalletXPubYPubZPubRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncHDWalletXPubYPubZPubRBDataItem(val *SyncHDWalletXPubYPubZPubRBDataItem) *NullableSyncHDWalletXPubYPubZPubRBDataItem {
	return &NullableSyncHDWalletXPubYPubZPubRBDataItem{value: val, isSet: true}
}

func (v NullableSyncHDWalletXPubYPubZPubRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncHDWalletXPubYPubZPubRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


