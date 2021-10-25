# ListTransactionsByBlockHashRIBSZVJoinSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchor** | **string** | Defines a Merkle tree root of a note commitment tree which uniquely identifies a note commitment tree state given the assumed security properties of the Merkle tree’s  hash function. | 
**CipherTexts** | **[]string** |  | 
**Commitments** | **[]string** |  | 
**Macs** | **[]string** |  | 
**Nullifiers** | **[]string** |  | 
**OneTimePubKey** | **string** | Defines the one time public key. | 
**Proof** | **string** | Defines the proof. | 
**RandomSeed** | **string** | Represents a 256-bit seed that must be chosen independently at random for each JoinSplit description. | 
**VPubNew** | **string** | Defines the value that the joinSplit transfer will insert into the transparent transaction value pool. | 
**VPubOld** | **string** | Defines the value that the joinSplit transfer will remove from the transparent transaction value pool. | 

## Methods

### NewListTransactionsByBlockHashRIBSZVJoinSplit

`func NewListTransactionsByBlockHashRIBSZVJoinSplit(anchor string, cipherTexts []string, commitments []string, macs []string, nullifiers []string, oneTimePubKey string, proof string, randomSeed string, vPubNew string, vPubOld string, ) *ListTransactionsByBlockHashRIBSZVJoinSplit`

NewListTransactionsByBlockHashRIBSZVJoinSplit instantiates a new ListTransactionsByBlockHashRIBSZVJoinSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSZVJoinSplitWithDefaults

`func NewListTransactionsByBlockHashRIBSZVJoinSplitWithDefaults() *ListTransactionsByBlockHashRIBSZVJoinSplit`

NewListTransactionsByBlockHashRIBSZVJoinSplitWithDefaults instantiates a new ListTransactionsByBlockHashRIBSZVJoinSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchor

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetAnchor() string`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetAnchorOk() (*string, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetAnchor(v string)`

SetAnchor sets Anchor field to given value.


### GetCipherTexts

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetCipherTexts() []string`

GetCipherTexts returns the CipherTexts field if non-nil, zero value otherwise.

### GetCipherTextsOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetCipherTextsOk() (*[]string, bool)`

GetCipherTextsOk returns a tuple with the CipherTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherTexts

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetCipherTexts(v []string)`

SetCipherTexts sets CipherTexts field to given value.


### GetCommitments

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetCommitments() []string`

GetCommitments returns the Commitments field if non-nil, zero value otherwise.

### GetCommitmentsOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetCommitmentsOk() (*[]string, bool)`

GetCommitmentsOk returns a tuple with the Commitments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitments

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetCommitments(v []string)`

SetCommitments sets Commitments field to given value.


### GetMacs

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetMacs(v []string)`

SetMacs sets Macs field to given value.


### GetNullifiers

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetNullifiers() []string`

GetNullifiers returns the Nullifiers field if non-nil, zero value otherwise.

### GetNullifiersOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetNullifiersOk() (*[]string, bool)`

GetNullifiersOk returns a tuple with the Nullifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullifiers

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetNullifiers(v []string)`

SetNullifiers sets Nullifiers field to given value.


### GetOneTimePubKey

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetOneTimePubKey() string`

GetOneTimePubKey returns the OneTimePubKey field if non-nil, zero value otherwise.

### GetOneTimePubKeyOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetOneTimePubKeyOk() (*string, bool)`

GetOneTimePubKeyOk returns a tuple with the OneTimePubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePubKey

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetOneTimePubKey(v string)`

SetOneTimePubKey sets OneTimePubKey field to given value.


### GetProof

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetProof(v string)`

SetProof sets Proof field to given value.


### GetRandomSeed

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetRandomSeed() string`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetRandomSeedOk() (*string, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetRandomSeed(v string)`

SetRandomSeed sets RandomSeed field to given value.


### GetVPubNew

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetVPubNew() string`

GetVPubNew returns the VPubNew field if non-nil, zero value otherwise.

### GetVPubNewOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetVPubNewOk() (*string, bool)`

GetVPubNewOk returns a tuple with the VPubNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPubNew

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetVPubNew(v string)`

SetVPubNew sets VPubNew field to given value.


### GetVPubOld

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetVPubOld() string`

GetVPubOld returns the VPubOld field if non-nil, zero value otherwise.

### GetVPubOldOk

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) GetVPubOldOk() (*string, bool)`

GetVPubOldOk returns a tuple with the VPubOld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPubOld

`func (o *ListTransactionsByBlockHashRIBSZVJoinSplit) SetVPubOld(v string)`

SetVPubOld sets VPubOld field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


