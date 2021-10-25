# ListLatestMinedBlocksRIBSZ2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | **string** | Represents a specific sub-unit of Zcash. Bits have two-decimal precision | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**Merkleroot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the proof of work | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**Version** | **int32** | Represents the transaction version number. | 

## Methods

### NewListLatestMinedBlocksRIBSZ2

`func NewListLatestMinedBlocksRIBSZ2(bits string, chainwork string, merkleroot string, nonce string, size int32, version int32, ) *ListLatestMinedBlocksRIBSZ2`

NewListLatestMinedBlocksRIBSZ2 instantiates a new ListLatestMinedBlocksRIBSZ2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIBSZ2WithDefaults

`func NewListLatestMinedBlocksRIBSZ2WithDefaults() *ListLatestMinedBlocksRIBSZ2`

NewListLatestMinedBlocksRIBSZ2WithDefaults instantiates a new ListLatestMinedBlocksRIBSZ2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *ListLatestMinedBlocksRIBSZ2) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *ListLatestMinedBlocksRIBSZ2) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *ListLatestMinedBlocksRIBSZ2) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *ListLatestMinedBlocksRIBSZ2) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *ListLatestMinedBlocksRIBSZ2) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *ListLatestMinedBlocksRIBSZ2) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetMerkleroot

`func (o *ListLatestMinedBlocksRIBSZ2) GetMerkleroot() string`

GetMerkleroot returns the Merkleroot field if non-nil, zero value otherwise.

### GetMerklerootOk

`func (o *ListLatestMinedBlocksRIBSZ2) GetMerklerootOk() (*string, bool)`

GetMerklerootOk returns a tuple with the Merkleroot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleroot

`func (o *ListLatestMinedBlocksRIBSZ2) SetMerkleroot(v string)`

SetMerkleroot sets Merkleroot field to given value.


### GetNonce

`func (o *ListLatestMinedBlocksRIBSZ2) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListLatestMinedBlocksRIBSZ2) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListLatestMinedBlocksRIBSZ2) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *ListLatestMinedBlocksRIBSZ2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListLatestMinedBlocksRIBSZ2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListLatestMinedBlocksRIBSZ2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListLatestMinedBlocksRIBSZ2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListLatestMinedBlocksRIBSZ2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListLatestMinedBlocksRIBSZ2) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


