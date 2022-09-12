# ListLatestMinedBlocksRIBS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | **string** | Represents a specific sub-unit of Zcash. Bits have two-decimal precision | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the proof of work | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**StrippedSize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**Version** | **int32** | Represents the transaction version number. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 
**Weight** | **int32** | Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit. | 
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **int32** | Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit. | 
**GasUsed** | **int32** | Defines how much of the gas for the block has been used. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in second | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block | 
**Uncles** | **[]string** |  | 
**DsBlock** | **int32** | Represents the Directory Service block which contains metadata about the miners who participate in the consensus protocol. | 
**DsDifficulty** | **string** | Defines how difficult it is to mine the dsBlocks. | 
**DsLeader** | **string** | Represents a part of the DS Committee which leads the consensus protocol for the epoch. | 
**MicroBlocks** | **[]string** |  | 
**TotalCoins** | Pointer to [**ListLatestMinedBlocksRIBSXTotalCoins**](ListLatestMinedBlocksRIBSXTotalCoins.md) |  | [optional] 
**TotalFees** | [**ListLatestMinedBlocksRIBSXTotalFees**](ListLatestMinedBlocksRIBSXTotalFees.md) |  | 

## Methods

### NewListLatestMinedBlocksRIBS

`func NewListLatestMinedBlocksRIBS(bits string, chainwork string, difficulty string, merkleRoot string, nonce string, size int32, strippedSize int32, version int32, versionHex string, weight int32, extraData string, gasLimit int32, gasUsed int32, minedInSeconds int32, sha3Uncles string, totalDifficulty string, uncles []string, dsBlock int32, dsDifficulty string, dsLeader string, microBlocks []string, totalFees ListLatestMinedBlocksRIBSXTotalFees, ) *ListLatestMinedBlocksRIBS`

NewListLatestMinedBlocksRIBS instantiates a new ListLatestMinedBlocksRIBS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIBSWithDefaults

`func NewListLatestMinedBlocksRIBSWithDefaults() *ListLatestMinedBlocksRIBS`

NewListLatestMinedBlocksRIBSWithDefaults instantiates a new ListLatestMinedBlocksRIBS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *ListLatestMinedBlocksRIBS) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *ListLatestMinedBlocksRIBS) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *ListLatestMinedBlocksRIBS) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *ListLatestMinedBlocksRIBS) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *ListLatestMinedBlocksRIBS) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *ListLatestMinedBlocksRIBS) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetDifficulty

`func (o *ListLatestMinedBlocksRIBS) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *ListLatestMinedBlocksRIBS) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *ListLatestMinedBlocksRIBS) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetMerkleRoot

`func (o *ListLatestMinedBlocksRIBS) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *ListLatestMinedBlocksRIBS) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *ListLatestMinedBlocksRIBS) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetNonce

`func (o *ListLatestMinedBlocksRIBS) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListLatestMinedBlocksRIBS) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListLatestMinedBlocksRIBS) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *ListLatestMinedBlocksRIBS) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListLatestMinedBlocksRIBS) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListLatestMinedBlocksRIBS) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStrippedSize

`func (o *ListLatestMinedBlocksRIBS) GetStrippedSize() int32`

GetStrippedSize returns the StrippedSize field if non-nil, zero value otherwise.

### GetStrippedSizeOk

`func (o *ListLatestMinedBlocksRIBS) GetStrippedSizeOk() (*int32, bool)`

GetStrippedSizeOk returns a tuple with the StrippedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedSize

`func (o *ListLatestMinedBlocksRIBS) SetStrippedSize(v int32)`

SetStrippedSize sets StrippedSize field to given value.


### GetVersion

`func (o *ListLatestMinedBlocksRIBS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListLatestMinedBlocksRIBS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListLatestMinedBlocksRIBS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *ListLatestMinedBlocksRIBS) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *ListLatestMinedBlocksRIBS) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *ListLatestMinedBlocksRIBS) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.


### GetWeight

`func (o *ListLatestMinedBlocksRIBS) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ListLatestMinedBlocksRIBS) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ListLatestMinedBlocksRIBS) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetExtraData

`func (o *ListLatestMinedBlocksRIBS) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ListLatestMinedBlocksRIBS) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ListLatestMinedBlocksRIBS) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *ListLatestMinedBlocksRIBS) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListLatestMinedBlocksRIBS) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListLatestMinedBlocksRIBS) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *ListLatestMinedBlocksRIBS) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListLatestMinedBlocksRIBS) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListLatestMinedBlocksRIBS) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *ListLatestMinedBlocksRIBS) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *ListLatestMinedBlocksRIBS) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *ListLatestMinedBlocksRIBS) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetSha3Uncles

`func (o *ListLatestMinedBlocksRIBS) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *ListLatestMinedBlocksRIBS) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *ListLatestMinedBlocksRIBS) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetTotalDifficulty

`func (o *ListLatestMinedBlocksRIBS) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *ListLatestMinedBlocksRIBS) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *ListLatestMinedBlocksRIBS) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *ListLatestMinedBlocksRIBS) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *ListLatestMinedBlocksRIBS) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *ListLatestMinedBlocksRIBS) SetUncles(v []string)`

SetUncles sets Uncles field to given value.


### GetDsBlock

`func (o *ListLatestMinedBlocksRIBS) GetDsBlock() int32`

GetDsBlock returns the DsBlock field if non-nil, zero value otherwise.

### GetDsBlockOk

`func (o *ListLatestMinedBlocksRIBS) GetDsBlockOk() (*int32, bool)`

GetDsBlockOk returns a tuple with the DsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsBlock

`func (o *ListLatestMinedBlocksRIBS) SetDsBlock(v int32)`

SetDsBlock sets DsBlock field to given value.


### GetDsDifficulty

`func (o *ListLatestMinedBlocksRIBS) GetDsDifficulty() string`

GetDsDifficulty returns the DsDifficulty field if non-nil, zero value otherwise.

### GetDsDifficultyOk

`func (o *ListLatestMinedBlocksRIBS) GetDsDifficultyOk() (*string, bool)`

GetDsDifficultyOk returns a tuple with the DsDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsDifficulty

`func (o *ListLatestMinedBlocksRIBS) SetDsDifficulty(v string)`

SetDsDifficulty sets DsDifficulty field to given value.


### GetDsLeader

`func (o *ListLatestMinedBlocksRIBS) GetDsLeader() string`

GetDsLeader returns the DsLeader field if non-nil, zero value otherwise.

### GetDsLeaderOk

`func (o *ListLatestMinedBlocksRIBS) GetDsLeaderOk() (*string, bool)`

GetDsLeaderOk returns a tuple with the DsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsLeader

`func (o *ListLatestMinedBlocksRIBS) SetDsLeader(v string)`

SetDsLeader sets DsLeader field to given value.


### GetMicroBlocks

`func (o *ListLatestMinedBlocksRIBS) GetMicroBlocks() []string`

GetMicroBlocks returns the MicroBlocks field if non-nil, zero value otherwise.

### GetMicroBlocksOk

`func (o *ListLatestMinedBlocksRIBS) GetMicroBlocksOk() (*[]string, bool)`

GetMicroBlocksOk returns a tuple with the MicroBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroBlocks

`func (o *ListLatestMinedBlocksRIBS) SetMicroBlocks(v []string)`

SetMicroBlocks sets MicroBlocks field to given value.


### GetTotalCoins

`func (o *ListLatestMinedBlocksRIBS) GetTotalCoins() ListLatestMinedBlocksRIBSXTotalCoins`

GetTotalCoins returns the TotalCoins field if non-nil, zero value otherwise.

### GetTotalCoinsOk

`func (o *ListLatestMinedBlocksRIBS) GetTotalCoinsOk() (*ListLatestMinedBlocksRIBSXTotalCoins, bool)`

GetTotalCoinsOk returns a tuple with the TotalCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCoins

`func (o *ListLatestMinedBlocksRIBS) SetTotalCoins(v ListLatestMinedBlocksRIBSXTotalCoins)`

SetTotalCoins sets TotalCoins field to given value.

### HasTotalCoins

`func (o *ListLatestMinedBlocksRIBS) HasTotalCoins() bool`

HasTotalCoins returns a boolean if a field has been set.

### GetTotalFees

`func (o *ListLatestMinedBlocksRIBS) GetTotalFees() ListLatestMinedBlocksRIBSXTotalFees`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *ListLatestMinedBlocksRIBS) GetTotalFeesOk() (*ListLatestMinedBlocksRIBSXTotalFees, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *ListLatestMinedBlocksRIBS) SetTotalFees(v ListLatestMinedBlocksRIBSXTotalFees)`

SetTotalFees sets TotalFees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


