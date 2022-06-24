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

// ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner struct for ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner
type ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner struct {
	// Defines a Merkle tree root of a note commitment tree which uniquely identifies a note commitment tree state given the assumed security properties of the Merkle tree’s  hash function.
	Anchor string `json:"anchor"`
	CipherTexts []string `json:"cipherTexts"`
	Commitments []string `json:"commitments"`
	Macs []string `json:"macs"`
	Nullifiers []string `json:"nullifiers"`
	// Defines the one time public key.
	OneTimePubKey string `json:"oneTimePubKey"`
	// Defines the proof.
	Proof string `json:"proof"`
	// Represents a 256-bit seed that must be chosen independently at random for each JoinSplit description.
	RandomSeed string `json:"randomSeed"`
	// Defines the value that the joinSplit transfer will insert into the transparent transaction value pool.
	VPubNew string `json:"vPubNew"`
	// Defines the value that the joinSplit transfer will remove from the transparent transaction value pool.
	VPubOld string `json:"vPubOld"`
}

// NewListConfirmedTransactionsByAddressRIBSZVJoinSplitInner instantiates a new ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSZVJoinSplitInner(anchor string, cipherTexts []string, commitments []string, macs []string, nullifiers []string, oneTimePubKey string, proof string, randomSeed string, vPubNew string, vPubOld string) *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner {
	this := ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner{}
	this.Anchor = anchor
	this.CipherTexts = cipherTexts
	this.Commitments = commitments
	this.Macs = macs
	this.Nullifiers = nullifiers
	this.OneTimePubKey = oneTimePubKey
	this.Proof = proof
	this.RandomSeed = randomSeed
	this.VPubNew = vPubNew
	this.VPubOld = vPubOld
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSZVJoinSplitInnerWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSZVJoinSplitInnerWithDefaults() *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner {
	this := ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner{}
	return &this
}

// GetAnchor returns the Anchor field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetAnchor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Anchor
}

// GetAnchorOk returns a tuple with the Anchor field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetAnchorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Anchor, true
}

// SetAnchor sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetAnchor(v string) {
	o.Anchor = v
}

// GetCipherTexts returns the CipherTexts field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetCipherTexts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CipherTexts
}

// GetCipherTextsOk returns a tuple with the CipherTexts field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetCipherTextsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CipherTexts, true
}

// SetCipherTexts sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetCipherTexts(v []string) {
	o.CipherTexts = v
}

// GetCommitments returns the Commitments field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetCommitments() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Commitments
}

// GetCommitmentsOk returns a tuple with the Commitments field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetCommitmentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commitments, true
}

// SetCommitments sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetCommitments(v []string) {
	o.Commitments = v
}

// GetMacs returns the Macs field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetMacs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetMacsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Macs, true
}

// SetMacs sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetMacs(v []string) {
	o.Macs = v
}

// GetNullifiers returns the Nullifiers field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetNullifiers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Nullifiers
}

// GetNullifiersOk returns a tuple with the Nullifiers field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetNullifiersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nullifiers, true
}

// SetNullifiers sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetNullifiers(v []string) {
	o.Nullifiers = v
}

// GetOneTimePubKey returns the OneTimePubKey field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetOneTimePubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OneTimePubKey
}

// GetOneTimePubKeyOk returns a tuple with the OneTimePubKey field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetOneTimePubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OneTimePubKey, true
}

// SetOneTimePubKey sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetOneTimePubKey(v string) {
	o.OneTimePubKey = v
}

// GetProof returns the Proof field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proof
}

// GetProofOk returns a tuple with the Proof field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proof, true
}

// SetProof sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetProof(v string) {
	o.Proof = v
}

// GetRandomSeed returns the RandomSeed field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetRandomSeed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RandomSeed
}

// GetRandomSeedOk returns a tuple with the RandomSeed field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetRandomSeedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RandomSeed, true
}

// SetRandomSeed sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetRandomSeed(v string) {
	o.RandomSeed = v
}

// GetVPubNew returns the VPubNew field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetVPubNew() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VPubNew
}

// GetVPubNewOk returns a tuple with the VPubNew field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetVPubNewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VPubNew, true
}

// SetVPubNew sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetVPubNew(v string) {
	o.VPubNew = v
}

// GetVPubOld returns the VPubOld field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetVPubOld() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VPubOld
}

// GetVPubOldOk returns a tuple with the VPubOld field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) GetVPubOldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VPubOld, true
}

// SetVPubOld sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) SetVPubOld(v string) {
	o.VPubOld = v
}

func (o ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["anchor"] = o.Anchor
	}
	if true {
		toSerialize["cipherTexts"] = o.CipherTexts
	}
	if true {
		toSerialize["commitments"] = o.Commitments
	}
	if true {
		toSerialize["macs"] = o.Macs
	}
	if true {
		toSerialize["nullifiers"] = o.Nullifiers
	}
	if true {
		toSerialize["oneTimePubKey"] = o.OneTimePubKey
	}
	if true {
		toSerialize["proof"] = o.Proof
	}
	if true {
		toSerialize["randomSeed"] = o.RandomSeed
	}
	if true {
		toSerialize["vPubNew"] = o.VPubNew
	}
	if true {
		toSerialize["vPubOld"] = o.VPubOld
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner struct {
	value *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) Get() *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) Set(val *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner(val *ListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) *NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner {
	return &NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSZVJoinSplitInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


