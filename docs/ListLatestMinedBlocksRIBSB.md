# ListLatestMinedBlocksRIBSB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | **string** | A sub-unit of BCH equal to 0.000001 BCH, or 100 Satoshi, and is the same as microbitcoincash (μBCH). Bits have two-decimal precision. | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**Difficulty** | Pointer to **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | [optional] 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Nonce** | Pointer to **int32** | Represents a random value that can be adjusted to satisfy the proof of work | [optional] 
**Size** | Pointer to **int32** | Represents the block size | [optional] 
**StrippedSize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**Version** | **int32** | Represents the version of the specific block on the blockchain. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 
**Weight** | **int32** | Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit. | 

## Methods

### NewListLatestMinedBlocksRIBSB

`func NewListLatestMinedBlocksRIBSB(bits string, chainwork string, merkleRoot string, strippedSize int32, version int32, versionHex string, weight int32, ) *ListLatestMinedBlocksRIBSB`

NewListLatestMinedBlocksRIBSB instantiates a new ListLatestMinedBlocksRIBSB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIBSBWithDefaults

`func NewListLatestMinedBlocksRIBSBWithDefaults() *ListLatestMinedBlocksRIBSB`

NewListLatestMinedBlocksRIBSBWithDefaults instantiates a new ListLatestMinedBlocksRIBSB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *ListLatestMinedBlocksRIBSB) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *ListLatestMinedBlocksRIBSB) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *ListLatestMinedBlocksRIBSB) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *ListLatestMinedBlocksRIBSB) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *ListLatestMinedBlocksRIBSB) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *ListLatestMinedBlocksRIBSB) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetDifficulty

`func (o *ListLatestMinedBlocksRIBSB) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *ListLatestMinedBlocksRIBSB) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *ListLatestMinedBlocksRIBSB) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *ListLatestMinedBlocksRIBSB) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetMerkleRoot

`func (o *ListLatestMinedBlocksRIBSB) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *ListLatestMinedBlocksRIBSB) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *ListLatestMinedBlocksRIBSB) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetNonce

`func (o *ListLatestMinedBlocksRIBSB) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListLatestMinedBlocksRIBSB) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListLatestMinedBlocksRIBSB) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ListLatestMinedBlocksRIBSB) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetSize

`func (o *ListLatestMinedBlocksRIBSB) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListLatestMinedBlocksRIBSB) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListLatestMinedBlocksRIBSB) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListLatestMinedBlocksRIBSB) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStrippedSize

`func (o *ListLatestMinedBlocksRIBSB) GetStrippedSize() int32`

GetStrippedSize returns the StrippedSize field if non-nil, zero value otherwise.

### GetStrippedSizeOk

`func (o *ListLatestMinedBlocksRIBSB) GetStrippedSizeOk() (*int32, bool)`

GetStrippedSizeOk returns a tuple with the StrippedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedSize

`func (o *ListLatestMinedBlocksRIBSB) SetStrippedSize(v int32)`

SetStrippedSize sets StrippedSize field to given value.


### GetVersion

`func (o *ListLatestMinedBlocksRIBSB) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListLatestMinedBlocksRIBSB) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListLatestMinedBlocksRIBSB) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *ListLatestMinedBlocksRIBSB) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *ListLatestMinedBlocksRIBSB) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *ListLatestMinedBlocksRIBSB) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.


### GetWeight

`func (o *ListLatestMinedBlocksRIBSB) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ListLatestMinedBlocksRIBSB) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ListLatestMinedBlocksRIBSB) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


