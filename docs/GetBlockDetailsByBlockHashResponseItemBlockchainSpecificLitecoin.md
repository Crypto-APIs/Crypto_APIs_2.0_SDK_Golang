# GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the Proof of Work. | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**Bits** | **string** | Represents a specific sub-unit of Litecoin. Bits have two-decimal precision. | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**StrippedSize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**Version** | **int32** | Represents the version of the specific block on the blockchain. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 
**Weight** | **int32** | Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit. | 

## Methods

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin(difficulty string, nonce string, size int32, bits string, chainwork string, merkleRoot string, strippedSize int32, version int32, versionHex string, weight int32, ) *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoinWithDefaults

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoinWithDefaults() *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoinWithDefaults instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetSize(v int32)`

SetSize sets Size field to given value.


### GetBits

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetMerkleRoot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetStrippedSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetStrippedSize() int32`

GetStrippedSize returns the StrippedSize field if non-nil, zero value otherwise.

### GetStrippedSizeOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetStrippedSizeOk() (*int32, bool)`

GetStrippedSizeOk returns a tuple with the StrippedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetStrippedSize(v int32)`

SetStrippedSize sets StrippedSize field to given value.


### GetVersion

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.


### GetWeight

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificLitecoin) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


