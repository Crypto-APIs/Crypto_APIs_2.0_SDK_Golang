# BlockMinedDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**Height** | **int32** | Defines the number of blocks in the blockchain preceding this specific block. | 
**Hash** | **string** | Represents the hash of the block&#39;s header, i.e. an output that has a fixed length. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in seconds since Unix Epoch time. | 

## Methods

### NewBlockMinedDataItem

`func NewBlockMinedDataItem(blockchain string, network string, height int32, hash string, timestamp int32, ) *BlockMinedDataItem`

NewBlockMinedDataItem instantiates a new BlockMinedDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockMinedDataItemWithDefaults

`func NewBlockMinedDataItemWithDefaults() *BlockMinedDataItem`

NewBlockMinedDataItemWithDefaults instantiates a new BlockMinedDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *BlockMinedDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *BlockMinedDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *BlockMinedDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *BlockMinedDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BlockMinedDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BlockMinedDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetHeight

`func (o *BlockMinedDataItem) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockMinedDataItem) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockMinedDataItem) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetHash

`func (o *BlockMinedDataItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlockMinedDataItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlockMinedDataItem) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTimestamp

`func (o *BlockMinedDataItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockMinedDataItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockMinedDataItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


