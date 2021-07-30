# BroadcastTransactionSuccessDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 

## Methods

### NewBroadcastTransactionSuccessDataItem

`func NewBroadcastTransactionSuccessDataItem(blockchain string, network string, transactionId string, ) *BroadcastTransactionSuccessDataItem`

NewBroadcastTransactionSuccessDataItem instantiates a new BroadcastTransactionSuccessDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTransactionSuccessDataItemWithDefaults

`func NewBroadcastTransactionSuccessDataItemWithDefaults() *BroadcastTransactionSuccessDataItem`

NewBroadcastTransactionSuccessDataItemWithDefaults instantiates a new BroadcastTransactionSuccessDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *BroadcastTransactionSuccessDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *BroadcastTransactionSuccessDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *BroadcastTransactionSuccessDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *BroadcastTransactionSuccessDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BroadcastTransactionSuccessDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BroadcastTransactionSuccessDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransactionId

`func (o *BroadcastTransactionSuccessDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BroadcastTransactionSuccessDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BroadcastTransactionSuccessDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


