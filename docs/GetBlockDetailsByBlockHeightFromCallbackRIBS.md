# GetBlockDetailsByBlockHeightFromCallbackRIBS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | **string** | Represents a specific sub-unit of Zcash. Bits have two-decimal precision | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Nonce** | **string** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**StrippedSize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**Version** | **int32** | Represents the transaction version number. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 
**Weight** | **int32** | Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit. | 
**Strippedsize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in seconds. | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block. | 
**Uncles** | **[]string** |  | 
**DsBlock** | **int32** | Represents the Directory Service block which contains metadata about the miners who participate in the consensus protocol. | 
**DsDifficulty** | **string** | Defines how difficult it is to mine the dsBlocks. | 
**DsLeader** | **string** | Represents a part of the DS Committee which leads the consensus protocol for the epoch. | 
**MicroBlocks** | **[]string** |  | 
**TotalCoins** | [**GetLatestMinedXRPRippleBlockRITotalCoins**](GetLatestMinedXRPRippleBlockRITotalCoins.md) |  | 
**TotalFees** | [**GetLatestMinedXRPRippleBlockRITotalFees**](GetLatestMinedXRPRippleBlockRITotalFees.md) |  | 
**BandwidthUsed** | **string** | Represents the bandwidth used for the transaction. | 
**BurnedTrx** | **string** | Represents the block burned TRX. | 
**EnergyUsed** | **string** | Representats the used energy for the transaction. | 

## Methods

### NewGetBlockDetailsByBlockHeightFromCallbackRIBS

`func NewGetBlockDetailsByBlockHeightFromCallbackRIBS(bits string, chainwork string, difficulty string, merkleRoot string, nonce string, size int32, strippedSize int32, version int32, versionHex string, weight int32, strippedsize int32, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, sha3Uncles string, totalDifficulty string, uncles []string, dsBlock int32, dsDifficulty string, dsLeader string, microBlocks []string, totalCoins GetLatestMinedXRPRippleBlockRITotalCoins, totalFees GetLatestMinedXRPRippleBlockRITotalFees, bandwidthUsed string, burnedTrx string, energyUsed string, ) *GetBlockDetailsByBlockHeightFromCallbackRIBS`

NewGetBlockDetailsByBlockHeightFromCallbackRIBS instantiates a new GetBlockDetailsByBlockHeightFromCallbackRIBS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHeightFromCallbackRIBSWithDefaults

`func NewGetBlockDetailsByBlockHeightFromCallbackRIBSWithDefaults() *GetBlockDetailsByBlockHeightFromCallbackRIBS`

NewGetBlockDetailsByBlockHeightFromCallbackRIBSWithDefaults instantiates a new GetBlockDetailsByBlockHeightFromCallbackRIBS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetMerkleRoot

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStrippedSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetStrippedSize() int32`

GetStrippedSize returns the StrippedSize field if non-nil, zero value otherwise.

### GetStrippedSizeOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetStrippedSizeOk() (*int32, bool)`

GetStrippedSizeOk returns a tuple with the StrippedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetStrippedSize(v int32)`

SetStrippedSize sets StrippedSize field to given value.


### GetVersion

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.


### GetWeight

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetStrippedsize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetStrippedsize() int32`

GetStrippedsize returns the Strippedsize field if non-nil, zero value otherwise.

### GetStrippedsizeOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetStrippedsizeOk() (*int32, bool)`

GetStrippedsizeOk returns a tuple with the Strippedsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedsize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetStrippedsize(v int32)`

SetStrippedsize sets Strippedsize field to given value.


### GetExtraData

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetSha3Uncles

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetTotalDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetUncles(v []string)`

SetUncles sets Uncles field to given value.


### GetDsBlock

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDsBlock() int32`

GetDsBlock returns the DsBlock field if non-nil, zero value otherwise.

### GetDsBlockOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDsBlockOk() (*int32, bool)`

GetDsBlockOk returns a tuple with the DsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsBlock

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetDsBlock(v int32)`

SetDsBlock sets DsBlock field to given value.


### GetDsDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDsDifficulty() string`

GetDsDifficulty returns the DsDifficulty field if non-nil, zero value otherwise.

### GetDsDifficultyOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDsDifficultyOk() (*string, bool)`

GetDsDifficultyOk returns a tuple with the DsDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetDsDifficulty(v string)`

SetDsDifficulty sets DsDifficulty field to given value.


### GetDsLeader

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDsLeader() string`

GetDsLeader returns the DsLeader field if non-nil, zero value otherwise.

### GetDsLeaderOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetDsLeaderOk() (*string, bool)`

GetDsLeaderOk returns a tuple with the DsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsLeader

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetDsLeader(v string)`

SetDsLeader sets DsLeader field to given value.


### GetMicroBlocks

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetMicroBlocks() []string`

GetMicroBlocks returns the MicroBlocks field if non-nil, zero value otherwise.

### GetMicroBlocksOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetMicroBlocksOk() (*[]string, bool)`

GetMicroBlocksOk returns a tuple with the MicroBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroBlocks

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetMicroBlocks(v []string)`

SetMicroBlocks sets MicroBlocks field to given value.


### GetTotalCoins

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetTotalCoins() GetLatestMinedXRPRippleBlockRITotalCoins`

GetTotalCoins returns the TotalCoins field if non-nil, zero value otherwise.

### GetTotalCoinsOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetTotalCoinsOk() (*GetLatestMinedXRPRippleBlockRITotalCoins, bool)`

GetTotalCoinsOk returns a tuple with the TotalCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCoins

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetTotalCoins(v GetLatestMinedXRPRippleBlockRITotalCoins)`

SetTotalCoins sets TotalCoins field to given value.


### GetTotalFees

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetTotalFees() GetLatestMinedXRPRippleBlockRITotalFees`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetTotalFeesOk() (*GetLatestMinedXRPRippleBlockRITotalFees, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetTotalFees(v GetLatestMinedXRPRippleBlockRITotalFees)`

SetTotalFees sets TotalFees field to given value.


### GetBandwidthUsed

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetBandwidthUsed() string`

GetBandwidthUsed returns the BandwidthUsed field if non-nil, zero value otherwise.

### GetBandwidthUsedOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetBandwidthUsedOk() (*string, bool)`

GetBandwidthUsedOk returns a tuple with the BandwidthUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthUsed

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetBandwidthUsed(v string)`

SetBandwidthUsed sets BandwidthUsed field to given value.


### GetBurnedTrx

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetBurnedTrx() string`

GetBurnedTrx returns the BurnedTrx field if non-nil, zero value otherwise.

### GetBurnedTrxOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetBurnedTrxOk() (*string, bool)`

GetBurnedTrxOk returns a tuple with the BurnedTrx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnedTrx

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetBurnedTrx(v string)`

SetBurnedTrx sets BurnedTrx field to given value.


### GetEnergyUsed

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetEnergyUsed() string`

GetEnergyUsed returns the EnergyUsed field if non-nil, zero value otherwise.

### GetEnergyUsedOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) GetEnergyUsedOk() (*string, bool)`

GetEnergyUsedOk returns a tuple with the EnergyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyUsed

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBS) SetEnergyUsed(v string)`

SetEnergyUsed sets EnergyUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


