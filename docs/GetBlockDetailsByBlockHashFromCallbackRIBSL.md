# GetBlockDetailsByBlockHashFromCallbackRIBSL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | **string** | Represents a specific sub-unit of Litecoin. Bits have two-decimal precision. | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Nonce** | **int64** | Represents a random value that can be adjusted to satisfy the proof of work | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**Strippedsize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**Version** | **int32** | Represents the version of the specific block on the blockchain. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 
**Weight** | **int32** | Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit. | 

## Methods

### NewGetBlockDetailsByBlockHashFromCallbackRIBSL

`func NewGetBlockDetailsByBlockHashFromCallbackRIBSL(bits string, chainwork string, difficulty string, merkleRoot string, nonce int64, size int32, strippedsize int32, version int32, versionHex string, weight int32, ) *GetBlockDetailsByBlockHashFromCallbackRIBSL`

NewGetBlockDetailsByBlockHashFromCallbackRIBSL instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHashFromCallbackRIBSLWithDefaults

`func NewGetBlockDetailsByBlockHashFromCallbackRIBSLWithDefaults() *GetBlockDetailsByBlockHashFromCallbackRIBSL`

NewGetBlockDetailsByBlockHashFromCallbackRIBSLWithDefaults instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetDifficulty

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetMerkleRoot

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetNonce(v int64)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStrippedsize

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetStrippedsize() int32`

GetStrippedsize returns the Strippedsize field if non-nil, zero value otherwise.

### GetStrippedsizeOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetStrippedsizeOk() (*int32, bool)`

GetStrippedsizeOk returns a tuple with the Strippedsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedsize

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetStrippedsize(v int32)`

SetStrippedsize sets Strippedsize field to given value.


### GetVersion

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.


### GetWeight

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


