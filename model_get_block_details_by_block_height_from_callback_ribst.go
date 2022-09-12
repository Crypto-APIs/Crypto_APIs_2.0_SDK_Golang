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

// GetBlockDetailsByBlockHeightFromCallbackRIBST Tron
type GetBlockDetailsByBlockHeightFromCallbackRIBST struct {
	// Represents the bandwidth used for the transaction.
	BandwidthUsed string `json:"bandwidthUsed"`
	// Represents the block burned TRX.
	BurnedTrx string `json:"burnedTrx"`
	// Representats the used energy for the transaction.
	EnergyUsed string `json:"energyUsed"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
}

// NewGetBlockDetailsByBlockHeightFromCallbackRIBST instantiates a new GetBlockDetailsByBlockHeightFromCallbackRIBST object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHeightFromCallbackRIBST(bandwidthUsed string, burnedTrx string, energyUsed string, size int32) *GetBlockDetailsByBlockHeightFromCallbackRIBST {
	this := GetBlockDetailsByBlockHeightFromCallbackRIBST{}
	this.BandwidthUsed = bandwidthUsed
	this.BurnedTrx = burnedTrx
	this.EnergyUsed = energyUsed
	this.Size = size
	return &this
}

// NewGetBlockDetailsByBlockHeightFromCallbackRIBSTWithDefaults instantiates a new GetBlockDetailsByBlockHeightFromCallbackRIBST object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHeightFromCallbackRIBSTWithDefaults() *GetBlockDetailsByBlockHeightFromCallbackRIBST {
	this := GetBlockDetailsByBlockHeightFromCallbackRIBST{}
	return &this
}

// GetBandwidthUsed returns the BandwidthUsed field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetBandwidthUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandwidthUsed
}

// GetBandwidthUsedOk returns a tuple with the BandwidthUsed field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetBandwidthUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandwidthUsed, true
}

// SetBandwidthUsed sets field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) SetBandwidthUsed(v string) {
	o.BandwidthUsed = v
}

// GetBurnedTrx returns the BurnedTrx field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetBurnedTrx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BurnedTrx
}

// GetBurnedTrxOk returns a tuple with the BurnedTrx field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetBurnedTrxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BurnedTrx, true
}

// SetBurnedTrx sets field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) SetBurnedTrx(v string) {
	o.BurnedTrx = v
}

// GetEnergyUsed returns the EnergyUsed field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetEnergyUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnergyUsed
}

// GetEnergyUsedOk returns a tuple with the EnergyUsed field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetEnergyUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnergyUsed, true
}

// SetEnergyUsed sets field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) SetEnergyUsed(v string) {
	o.EnergyUsed = v
}

// GetSize returns the Size field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetBlockDetailsByBlockHeightFromCallbackRIBST) SetSize(v int32) {
	o.Size = v
}

func (o GetBlockDetailsByBlockHeightFromCallbackRIBST) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bandwidthUsed"] = o.BandwidthUsed
	}
	if true {
		toSerialize["burnedTrx"] = o.BurnedTrx
	}
	if true {
		toSerialize["energyUsed"] = o.EnergyUsed
	}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHeightFromCallbackRIBST struct {
	value *GetBlockDetailsByBlockHeightFromCallbackRIBST
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHeightFromCallbackRIBST) Get() *GetBlockDetailsByBlockHeightFromCallbackRIBST {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHeightFromCallbackRIBST) Set(val *GetBlockDetailsByBlockHeightFromCallbackRIBST) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHeightFromCallbackRIBST) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHeightFromCallbackRIBST) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHeightFromCallbackRIBST(val *GetBlockDetailsByBlockHeightFromCallbackRIBST) *NullableGetBlockDetailsByBlockHeightFromCallbackRIBST {
	return &NullableGetBlockDetailsByBlockHeightFromCallbackRIBST{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHeightFromCallbackRIBST) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHeightFromCallbackRIBST) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


