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

// GetWalletTransactionDetailsByTransactionIDRIBSZ Zcash
type GetWalletTransactionDetailsByTransactionIDRIBSZ struct {
	// It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions.
	BindingSig string `json:"bindingSig"`
	// Represents a block height after which the transaction will expire.
	ExpiryHeight int32 `json:"expiryHeight"`
	// Represents an encoding of a JoinSplitSig public validating key.
	JoinSplitPubKey string `json:"joinSplitPubKey"`
	// Is used to sign transactions that contain at least one JoinSplit description.
	JoinSplitSig string `json:"joinSplitSig"`
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int64 `json:"locktime"`
	// \"Overwinter\" is the network upgrade for the Zcash blockchain.
	Overwintered bool `json:"overwintered"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents a sequence of JoinSplit descriptions using BCTV14 proofs.
	VJoinSplit []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner `json:"vJoinSplit,omitempty"`
	// Object Array representation of transaction output descriptions
	VShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner `json:"vShieldedOutput,omitempty"`
	// Object Array representation of transaction spend descriptions
	VShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner `json:"vShieldedSpend,omitempty"`
	// String representation of the transaction value balance
	ValueBalance string `json:"valueBalance"`
	// Represents the transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction version group ID.
	VersionGroupId string `json:"versionGroupId"`
	// Object Array representation of transaction inputs
	Vin []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner `json:"vin"`
	// Object Array representation of transaction outputs
	Vout []ListTransactionsByBlockHeightRIBSZVoutInner `json:"vout"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSZ instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSZ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int64, overwintered bool, size int32, valueBalance string, version int32, versionGroupId string, vin []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner, vout []ListTransactionsByBlockHeightRIBSZVoutInner) *GetWalletTransactionDetailsByTransactionIDRIBSZ {
	this := GetWalletTransactionDetailsByTransactionIDRIBSZ{}
	this.BindingSig = bindingSig
	this.ExpiryHeight = expiryHeight
	this.JoinSplitPubKey = joinSplitPubKey
	this.JoinSplitSig = joinSplitSig
	this.Locktime = locktime
	this.Overwintered = overwintered
	this.Size = size
	this.ValueBalance = valueBalance
	this.Version = version
	this.VersionGroupId = versionGroupId
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSZWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSZ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSZWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSZ {
	this := GetWalletTransactionDetailsByTransactionIDRIBSZ{}
	return &this
}

// GetBindingSig returns the BindingSig field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetBindingSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindingSig
}

// GetBindingSigOk returns a tuple with the BindingSig field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetBindingSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindingSig, true
}

// SetBindingSig sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetBindingSig(v string) {
	o.BindingSig = v
}

// GetExpiryHeight returns the ExpiryHeight field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetExpiryHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryHeight
}

// GetExpiryHeightOk returns a tuple with the ExpiryHeight field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetExpiryHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryHeight, true
}

// SetExpiryHeight sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetExpiryHeight(v int32) {
	o.ExpiryHeight = v
}

// GetJoinSplitPubKey returns the JoinSplitPubKey field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinSplitPubKey
}

// GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinSplitPubKey, true
}

// SetJoinSplitPubKey sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetJoinSplitPubKey(v string) {
	o.JoinSplitPubKey = v
}

// GetJoinSplitSig returns the JoinSplitSig field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinSplitSig
}

// GetJoinSplitSigOk returns a tuple with the JoinSplitSig field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinSplitSig, true
}

// SetJoinSplitSig sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetJoinSplitSig(v string) {
	o.JoinSplitSig = v
}

// GetLocktime returns the Locktime field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetLocktime(v int64) {
	o.Locktime = v
}

// GetOverwintered returns the Overwintered field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetOverwintered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Overwintered
}

// GetOverwinteredOk returns a tuple with the Overwintered field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetOverwinteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overwintered, true
}

// SetOverwintered sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetOverwintered(v bool) {
	o.Overwintered = v
}

// GetSize returns the Size field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetSize(v int32) {
	o.Size = v
}

// GetVJoinSplit returns the VJoinSplit field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVJoinSplit() []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner {
	if o == nil || o.VJoinSplit == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner
		return ret
	}
	return o.VJoinSplit
}

// GetVJoinSplitOk returns a tuple with the VJoinSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVJoinSplitOk() ([]GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner, bool) {
	if o == nil || o.VJoinSplit == nil {
		return nil, false
	}
	return o.VJoinSplit, true
}

// HasVJoinSplit returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) HasVJoinSplit() bool {
	if o != nil && o.VJoinSplit != nil {
		return true
	}

	return false
}

// SetVJoinSplit gets a reference to the given []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner and assigns it to the VJoinSplit field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVJoinSplit(v []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner) {
	o.VJoinSplit = v
}

// GetVShieldedOutput returns the VShieldedOutput field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner {
	if o == nil || o.VShieldedOutput == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner
		return ret
	}
	return o.VShieldedOutput
}

// GetVShieldedOutputOk returns a tuple with the VShieldedOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedOutputOk() ([]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner, bool) {
	if o == nil || o.VShieldedOutput == nil {
		return nil, false
	}
	return o.VShieldedOutput, true
}

// HasVShieldedOutput returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) HasVShieldedOutput() bool {
	if o != nil && o.VShieldedOutput != nil {
		return true
	}

	return false
}

// SetVShieldedOutput gets a reference to the given []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner and assigns it to the VShieldedOutput field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) {
	o.VShieldedOutput = v
}

// GetVShieldedSpend returns the VShieldedSpend field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner {
	if o == nil || o.VShieldedSpend == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner
		return ret
	}
	return o.VShieldedSpend
}

// GetVShieldedSpendOk returns a tuple with the VShieldedSpend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedSpendOk() ([]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner, bool) {
	if o == nil || o.VShieldedSpend == nil {
		return nil, false
	}
	return o.VShieldedSpend, true
}

// HasVShieldedSpend returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) HasVShieldedSpend() bool {
	if o != nil && o.VShieldedSpend != nil {
		return true
	}

	return false
}

// SetVShieldedSpend gets a reference to the given []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner and assigns it to the VShieldedSpend field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner) {
	o.VShieldedSpend = v
}

// GetValueBalance returns the ValueBalance field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetValueBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueBalance
}

// GetValueBalanceOk returns a tuple with the ValueBalance field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetValueBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueBalance, true
}

// SetValueBalance sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetValueBalance(v string) {
	o.ValueBalance = v
}

// GetVersion returns the Version field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVersion(v int32) {
	o.Version = v
}

// GetVersionGroupId returns the VersionGroupId field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersionGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionGroupId
}

// GetVersionGroupIdOk returns a tuple with the VersionGroupId field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionGroupId, true
}

// SetVersionGroupId sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVersionGroupId(v string) {
	o.VersionGroupId = v
}

// GetVin returns the Vin field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner {
	if o == nil {
		var ret []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVinOk() ([]GetWalletTransactionDetailsByTransactionIDRIBSZVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVout() []ListTransactionsByBlockHeightRIBSZVoutInner {
	if o == nil {
		var ret []ListTransactionsByBlockHeightRIBSZVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVoutOk() ([]ListTransactionsByBlockHeightRIBSZVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVout(v []ListTransactionsByBlockHeightRIBSZVoutInner) {
	o.Vout = v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBSZ) MarshalJSON() ([]byte, error) {
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
	if o.VJoinSplit != nil {
		toSerialize["vJoinSplit"] = o.VJoinSplit
	}
	if o.VShieldedOutput != nil {
		toSerialize["vShieldedOutput"] = o.VShieldedOutput
	}
	if o.VShieldedSpend != nil {
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

type NullableGetWalletTransactionDetailsByTransactionIDRIBSZ struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBSZ
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSZ) Get() *GetWalletTransactionDetailsByTransactionIDRIBSZ {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSZ) Set(val *GetWalletTransactionDetailsByTransactionIDRIBSZ) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSZ) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSZ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBSZ(val *GetWalletTransactionDetailsByTransactionIDRIBSZ) *NullableGetWalletTransactionDetailsByTransactionIDRIBSZ {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBSZ{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSZ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSZ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


