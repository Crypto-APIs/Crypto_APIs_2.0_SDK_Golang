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

// ListAllUnconfirmedTransactionsRIBSZ Zcash
type ListAllUnconfirmedTransactionsRIBSZ struct {
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
	VJoinSplit []ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner `json:"vJoinSplit"`
	// Object Array representation of transaction output descriptions
	VShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner `json:"vShieldedOutput"`
	// Object Array representation of transaction spend descriptions
	VShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner `json:"vShieldedSpend"`
	// Defines the transaction value balance.
	ValueBalance string `json:"valueBalance"`
	// Defines the version of the transaction.
	Version int32 `json:"version"`
	// Represents the transaction version group ID.
	VersionGroupId string `json:"versionGroupId"`
	// Object Array representation of transaction inputs
	Vin []GetTransactionDetailsByTransactionIDRIBSZVinInner `json:"vin"`
	// Object Array representation of transaction outputs
	Vout []GetTransactionDetailsByTransactionIDRIBSZVoutInner `json:"vout"`
}

// NewListAllUnconfirmedTransactionsRIBSZ instantiates a new ListAllUnconfirmedTransactionsRIBSZ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllUnconfirmedTransactionsRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int64, overwintered bool, size int32, vJoinSplit []ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner, vShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner, vShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner, valueBalance string, version int32, versionGroupId string, vin []GetTransactionDetailsByTransactionIDRIBSZVinInner, vout []GetTransactionDetailsByTransactionIDRIBSZVoutInner) *ListAllUnconfirmedTransactionsRIBSZ {
	this := ListAllUnconfirmedTransactionsRIBSZ{}
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

// NewListAllUnconfirmedTransactionsRIBSZWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSZ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllUnconfirmedTransactionsRIBSZWithDefaults() *ListAllUnconfirmedTransactionsRIBSZ {
	this := ListAllUnconfirmedTransactionsRIBSZ{}
	return &this
}

// GetBindingSig returns the BindingSig field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetBindingSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindingSig
}

// GetBindingSigOk returns a tuple with the BindingSig field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetBindingSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindingSig, true
}

// SetBindingSig sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetBindingSig(v string) {
	o.BindingSig = v
}

// GetExpiryHeight returns the ExpiryHeight field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetExpiryHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryHeight
}

// GetExpiryHeightOk returns a tuple with the ExpiryHeight field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetExpiryHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryHeight, true
}

// SetExpiryHeight sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetExpiryHeight(v int32) {
	o.ExpiryHeight = v
}

// GetJoinSplitPubKey returns the JoinSplitPubKey field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetJoinSplitPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinSplitPubKey
}

// GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetJoinSplitPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinSplitPubKey, true
}

// SetJoinSplitPubKey sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetJoinSplitPubKey(v string) {
	o.JoinSplitPubKey = v
}

// GetJoinSplitSig returns the JoinSplitSig field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetJoinSplitSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinSplitSig
}

// GetJoinSplitSigOk returns a tuple with the JoinSplitSig field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetJoinSplitSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinSplitSig, true
}

// SetJoinSplitSig sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetJoinSplitSig(v string) {
	o.JoinSplitSig = v
}

// GetLocktime returns the Locktime field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetLocktime(v int64) {
	o.Locktime = v
}

// GetOverwintered returns the Overwintered field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetOverwintered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Overwintered
}

// GetOverwinteredOk returns a tuple with the Overwintered field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetOverwinteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overwintered, true
}

// SetOverwintered sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetOverwintered(v bool) {
	o.Overwintered = v
}

// GetSize returns the Size field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetSize(v int32) {
	o.Size = v
}

// GetVJoinSplit returns the VJoinSplit field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVJoinSplit() []ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner
		return ret
	}

	return o.VJoinSplit
}

// GetVJoinSplitOk returns a tuple with the VJoinSplit field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVJoinSplitOk() ([]ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.VJoinSplit, true
}

// SetVJoinSplit sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetVJoinSplit(v []ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) {
	o.VJoinSplit = v
}

// GetVShieldedOutput returns the VShieldedOutput field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner
		return ret
	}

	return o.VShieldedOutput
}

// GetVShieldedOutputOk returns a tuple with the VShieldedOutput field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVShieldedOutputOk() ([]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.VShieldedOutput, true
}

// SetVShieldedOutput sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) {
	o.VShieldedOutput = v
}

// GetVShieldedSpend returns the VShieldedSpend field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner
		return ret
	}

	return o.VShieldedSpend
}

// GetVShieldedSpendOk returns a tuple with the VShieldedSpend field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVShieldedSpendOk() ([]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.VShieldedSpend, true
}

// SetVShieldedSpend sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner) {
	o.VShieldedSpend = v
}

// GetValueBalance returns the ValueBalance field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetValueBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueBalance
}

// GetValueBalanceOk returns a tuple with the ValueBalance field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetValueBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueBalance, true
}

// SetValueBalance sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetValueBalance(v string) {
	o.ValueBalance = v
}

// GetVersion returns the Version field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetVersion(v int32) {
	o.Version = v
}

// GetVersionGroupId returns the VersionGroupId field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVersionGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionGroupId
}

// GetVersionGroupIdOk returns a tuple with the VersionGroupId field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVersionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionGroupId, true
}

// SetVersionGroupId sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetVersionGroupId(v string) {
	o.VersionGroupId = v
}

// GetVin returns the Vin field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVin() []GetTransactionDetailsByTransactionIDRIBSZVinInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVinInner
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVinOk() ([]GetTransactionDetailsByTransactionIDRIBSZVinInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetVin(v []GetTransactionDetailsByTransactionIDRIBSZVinInner) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVout() []GetTransactionDetailsByTransactionIDRIBSZVoutInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSZVoutInner
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSZ) GetVoutOk() ([]GetTransactionDetailsByTransactionIDRIBSZVoutInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListAllUnconfirmedTransactionsRIBSZ) SetVout(v []GetTransactionDetailsByTransactionIDRIBSZVoutInner) {
	o.Vout = v
}

func (o ListAllUnconfirmedTransactionsRIBSZ) MarshalJSON() ([]byte, error) {
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

type NullableListAllUnconfirmedTransactionsRIBSZ struct {
	value *ListAllUnconfirmedTransactionsRIBSZ
	isSet bool
}

func (v NullableListAllUnconfirmedTransactionsRIBSZ) Get() *ListAllUnconfirmedTransactionsRIBSZ {
	return v.value
}

func (v *NullableListAllUnconfirmedTransactionsRIBSZ) Set(val *ListAllUnconfirmedTransactionsRIBSZ) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllUnconfirmedTransactionsRIBSZ) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllUnconfirmedTransactionsRIBSZ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllUnconfirmedTransactionsRIBSZ(val *ListAllUnconfirmedTransactionsRIBSZ) *NullableListAllUnconfirmedTransactionsRIBSZ {
	return &NullableListAllUnconfirmedTransactionsRIBSZ{value: val, isSet: true}
}

func (v NullableListAllUnconfirmedTransactionsRIBSZ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllUnconfirmedTransactionsRIBSZ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


