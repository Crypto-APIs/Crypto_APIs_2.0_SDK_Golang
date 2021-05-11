# TransactionMinedDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**MinedInBlock** | [**TransactionMinedDataItemMinedInBlock**](TransactionMinedDataItemMinedInBlock.md) |  | 

## Methods

### NewTransactionMinedDataItem

`func NewTransactionMinedDataItem(blockchain string, network string, transactionId string, minedInBlock TransactionMinedDataItemMinedInBlock, ) *TransactionMinedDataItem`

NewTransactionMinedDataItem instantiates a new TransactionMinedDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMinedDataItemWithDefaults

`func NewTransactionMinedDataItemWithDefaults() *TransactionMinedDataItem`

NewTransactionMinedDataItemWithDefaults instantiates a new TransactionMinedDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *TransactionMinedDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *TransactionMinedDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *TransactionMinedDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *TransactionMinedDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionMinedDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionMinedDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransactionId

`func (o *TransactionMinedDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionMinedDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionMinedDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetMinedInBlock

`func (o *TransactionMinedDataItem) GetMinedInBlock() TransactionMinedDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *TransactionMinedDataItem) GetMinedInBlockOk() (*TransactionMinedDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *TransactionMinedDataItem) SetMinedInBlock(v TransactionMinedDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


