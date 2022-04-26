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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSZ Zcash
type GetTransactionDetailsByTransactionIDFromCallbackRIBSZ struct {
	// It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions.
	BindingSig string `json:"bindingSig"`
	// Represents a block height after which the transaction will expire.
	ExpiryHeight int32 `json:"expiryHeight"`
	// Represents an encoding of a JoinSplitSig public validating key.
	JoinSplitPubKey string `json:"joinSplitPubKey"`
	// Is used to sign transactions that contain at least one JoinSplit description.
	JoinSplitSig string `json:"joinSplitSig"`
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int64 `json:"locktime"`
	// \"Overwinter\" is the network upgrade for the Zcash blockchain.
	Overwintered bool `json:"overwintered"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents a sequence of JoinSplit descriptions using BCTV14 proofs.
	VJoinSplit []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit `json:"vJoinSplit"`
	// Object Array representation of transaction output descriptions
	VShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput `json:"vShieldedOutput"`
	// Object Array representation of transaction spend descriptions
	VShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend `json:"vShieldedSpend"`
	// String representation of the transaction value balance
	ValueBalance string `json:"valueBalance"`
	// Defines the version of the transaction.
	Version int32 `json:"version"`
	// Represents the transaction version group ID
	VersionGroupId string `json:"versionGroupId"`
	// Object Array representation of transaction inputs
	Vin []GetTransactionDetailsByTransactionIDRIBSZVin `json:"vin"`
	// Object Array representation of transaction outputs
	Vout []GetTransactionDetailsByTransactionIDRIBSZVout `json:"vout"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int64, overwintered bool, size int32, vJoinSplit []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit, vShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, vShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, valueBalance string, version int32, versionGroupId string, vin []GetTransactionDetailsByTransactionIDRIBSZVin, vout []GetTransactionDetailsByTransactionIDRIBSZVout) *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSZ{}
	this.BindingSig = bindingSig
	this.ExpiryHeight = expiryHeight
	this.JoinSplitPubKey = joinSplitPubKey
	this.JoinSplitSig = joinSplitSig
	this.Locktime = locktime
	this.Overwintered = overwintered
	this.Size = size
	this.VJoinSplit = vJoinSplit
	this.VShieldedOutput = vShieldedOutput
	this.VShieldedSpend = vShieldedSpend
	this.ValueBalance = valueBalance
	this.Version = version
	this.VersionGroupId = versionGroupId
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSZ{}
	return &this
}

// GetBindingSig returns the BindingSig field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetBindingSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindingSig
}

// GetBindingSigOk returns a tuple with the BindingSig field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetBindingSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindingSig, true
}

// SetBindingSig sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetBindingSig(v string) {
	o.BindingSig = v
}

// GetExpiryHeight returns the ExpiryHeight field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetExpiryHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryHeight
}

// GetExpiryHeightOk returns a tuple with the ExpiryHeight field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetExpiryHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryHeight, true
}

// SetExpiryHeight sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetExpiryHeight(v int32) {
	o.ExpiryHeight = v
}

// GetJoinSplitPubKey returns the JoinSplitPubKey field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinSplitPubKey
}

// GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinSplitPubKey, true
}

// SetJoinSplitPubKey sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetJoinSplitPubKey(v string) {
	o.JoinSplitPubKey = v
}

// GetJoinSplitSig returns the JoinSplitSig field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinSplitSig
}

// GetJoinSplitSigOk returns a tuple with the JoinSplitSig field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinSplitSig, true
}

// SetJoinSplitSig sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetJoinSplitSig(v string) {
	o.JoinSplitSig = v
}

// GetLocktime returns the Locktime field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetLocktime(v int64) {
	o.Locktime = v
}

// GetOverwintered returns the Overwintered field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetOverwintered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Overwintered
}

// GetOverwinteredOk returns a tuple with the Overwintered field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetOverwinteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overwintered, true
}

// SetOverwintered sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetOverwintered(v bool) {
	o.Overwintered = v
}

// GetSize returns the Size field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetSize(v int32) {
	o.Size = v
}

// GetVJoinSplit returns the VJoinSplit field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVJoinSplit() []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit
		return ret
	}

	return o.VJoinSplit
}

// GetVJoinSplitOk returns a tuple with the VJoinSplit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVJoinSplitOk() ([]GetTransactionDetailsByTransactionIDRIBSZVJoinSplit, bool) {
	if o == nil {
		return nil, false
	}
	return o.VJoinSplit, true
}

// SetVJoinSplit sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVJoinSplit(v []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit) {
	o.VJoinSplit = v
}

// GetVShieldedOutput returns the VShieldedOutput field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput
		return ret
	}

	return o.VShieldedOutput
}

// GetVShieldedOutputOk returns a tuple with the VShieldedOutput field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedOutputOk() ([]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.VShieldedOutput, true
}

// SetVShieldedOutput sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput) {
	o.VShieldedOutput = v
}

// GetVShieldedSpend returns the VShieldedSpend field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend
		return ret
	}

	return o.VShieldedSpend
}

// GetVShieldedSpendOk returns a tuple with the VShieldedSpend field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedSpendOk() ([]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, bool) {
	if o == nil {
		return nil, false
	}
	return o.VShieldedSpend, true
}

// SetVShieldedSpend sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) {
	o.VShieldedSpend = v
}

// GetValueBalance returns the ValueBalance field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetValueBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueBalance
}

// GetValueBalanceOk returns a tuple with the ValueBalance field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetValueBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueBalance, true
}

// SetValueBalance sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetValueBalance(v string) {
	o.ValueBalance = v
}

// GetVersion returns the Version field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVersion(v int32) {
	o.Version = v
}

// GetVersionGroupId returns the VersionGroupId field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersionGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionGroupId
}

// GetVersionGroupIdOk returns a tuple with the VersionGroupId field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionGroupId, true
}

// SetVersionGroupId sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVersionGroupId(v string) {
	o.VersionGroupId = v
}

// GetVin returns the Vin field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVin() []GetTransactionDetailsByTransactionIDRIBSZVin {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVinOk() ([]GetTransactionDetailsByTransactionIDRIBSZVin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVin(v []GetTransactionDetailsByTransactionIDRIBSZVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVout() []GetTransactionDetailsByTransactionIDRIBSZVout {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVoutOk() ([]GetTransactionDetailsByTransactionIDRIBSZVout, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVout(v []GetTransactionDetailsByTransactionIDRIBSZVout) {
	o.Vout = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bindingSig"] = o.BindingSig
	}
	if true {
		toSerialize["expiryHeight"] = o.ExpiryHeight
	}
	if true {
		toSerialize["joinSplitPubKey"] = o.JoinSplitPubKey
	}
	if true {
		toSerialize["joinSplitSig"] = o.JoinSplitSig
	}
	if true {
		toSerialize["locktime"] = o.Locktime
	}
	if true {
		toSerialize["overwintered"] = o.Overwintered
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["vJoinSplit"] = o.VJoinSplit
	}
	if true {
		toSerialize["vShieldedOutput"] = o.VShieldedOutput
	}
	if true {
		toSerialize["vShieldedSpend"] = o.VShieldedSpend
	}
	if true {
		toSerialize["valueBalance"] = o.ValueBalance
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["versionGroupId"] = o.VersionGroupId
	}
	if true {
		toSerialize["vin"] = o.Vin
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


