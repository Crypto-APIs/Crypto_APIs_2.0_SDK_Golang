# ListLatestMinedBlocksRIBSBC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | **string** | A sub-unit of BCH equal to 0.000001 BCH, or 100 Satoshi, and is the same as microbitcoincash (μBCH). Bits have two-decimal precision. | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Version** | **int32** | Represents the version of the specific block on the blockchain. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 

## Methods

### NewListLatestMinedBlocksRIBSBC

`func NewListLatestMinedBlocksRIBSBC(bits string, chainwork string, merkleRoot string, version int32, versionHex string, ) *ListLatestMinedBlocksRIBSBC`

NewListLatestMinedBlocksRIBSBC instantiates a new ListLatestMinedBlocksRIBSBC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIBSBCWithDefaults

`func NewListLatestMinedBlocksRIBSBCWithDefaults() *ListLatestMinedBlocksRIBSBC`

NewListLatestMinedBlocksRIBSBCWithDefaults instantiates a new ListLatestMinedBlocksRIBSBC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *ListLatestMinedBlocksRIBSBC) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *ListLatestMinedBlocksRIBSBC) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *ListLatestMinedBlocksRIBSBC) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *ListLatestMinedBlocksRIBSBC) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *ListLatestMinedBlocksRIBSBC) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *ListLatestMinedBlocksRIBSBC) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetMerkleRoot

`func (o *ListLatestMinedBlocksRIBSBC) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *ListLatestMinedBlocksRIBSBC) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *ListLatestMinedBlocksRIBSBC) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetVersion

`func (o *ListLatestMinedBlocksRIBSBC) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListLatestMinedBlocksRIBSBC) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListLatestMinedBlocksRIBSBC) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *ListLatestMinedBlocksRIBSBC) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *ListLatestMinedBlocksRIBSBC) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *ListLatestMinedBlocksRIBSBC) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


