# GetBlockDetailsByBlockHashResponseItemBlockchainSpecific

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the Proof of Work. | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**Bits** | **string** | Represents a specific sub-unit of Dash. Bits have two-decimal precision. | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**StrippedSize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**Version** | **int32** | Represents the version of the specific block on the blockchain. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 
**Weight** | **int32** | Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit. | 
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **string** | Defines the total gas limit of all transactions in the block. | 
**GasUsed** | **string** | Represents the total amount of gas used by all transactions in this block. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in seconds. | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block. | 
**Uncles** | **[]string** |  | 
**Merkleroot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 

## Methods

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecific

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecific(difficulty string, nonce string, size int32, bits string, chainwork string, merkleRoot string, strippedSize int32, version int32, versionHex string, weight int32, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, sha3Uncles string, totalDifficulty string, uncles []string, merkleroot string, ) *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecific instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecific object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificWithDefaults

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificWithDefaults() *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificWithDefaults instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecific object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetSize(v int32)`

SetSize sets Size field to given value.


### GetBits

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetMerkleRoot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetStrippedSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetStrippedSize() int32`

GetStrippedSize returns the StrippedSize field if non-nil, zero value otherwise.

### GetStrippedSizeOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetStrippedSizeOk() (*int32, bool)`

GetStrippedSizeOk returns a tuple with the StrippedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetStrippedSize(v int32)`

SetStrippedSize sets StrippedSize field to given value.


### GetVersion

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.


### GetWeight

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetExtraData

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetSha3Uncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetTotalDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetUncles(v []string)`

SetUncles sets Uncles field to given value.


### GetMerkleroot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetMerkleroot() string`

GetMerkleroot returns the Merkleroot field if non-nil, zero value otherwise.

### GetMerklerootOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) GetMerklerootOk() (*string, bool)`

GetMerklerootOk returns a tuple with the Merkleroot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleroot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecific) SetMerkleroot(v string)`

SetMerkleroot sets Merkleroot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


