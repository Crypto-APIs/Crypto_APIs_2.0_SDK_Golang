# GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the Proof of Work. | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**Bits** | **string** | Represents a specific sub-unit of Dash. Bits have two-decimal precision. | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Version** | **int32** | Represents the version of the specific block on the blockchain. | 
**VersionHex** | **string** | Is the hexadecimal string representation of the block&#39;s version. | 

## Methods

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash(difficulty string, nonce string, size int32, bits string, chainwork string, merkleRoot string, version int32, versionHex string, ) *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificDashWithDefaults

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificDashWithDefaults() *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificDashWithDefaults instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetSize(v int32)`

SetSize sets Size field to given value.


### GetBits

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetMerkleRoot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetVersion

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionHex

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetVersionHex() string`

GetVersionHex returns the VersionHex field if non-nil, zero value otherwise.

### GetVersionHexOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) GetVersionHexOk() (*string, bool)`

GetVersionHexOk returns a tuple with the VersionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHex

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificDash) SetVersionHex(v string)`

SetVersionHex sets VersionHex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


