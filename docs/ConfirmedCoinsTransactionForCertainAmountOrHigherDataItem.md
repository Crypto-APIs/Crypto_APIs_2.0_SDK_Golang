# ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**MinedInBlock** | [**AddressCoinsTransactionConfirmedDataItemMinedInBlock**](AddressCoinsTransactionConfirmedDataItemMinedInBlock.md) |  | 
**Amount** | **string** | Defines the amount of coins sent with the confirmed transaction. | 
**Unit** | **string** | Defines the unit of the transaction, e.g. BTC. | 

## Methods

### NewConfirmedCoinsTransactionForCertainAmountOrHigherDataItem

`func NewConfirmedCoinsTransactionForCertainAmountOrHigherDataItem(blockchain string, network string, transactionId string, minedInBlock AddressCoinsTransactionConfirmedDataItemMinedInBlock, amount string, unit string, ) *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem`

NewConfirmedCoinsTransactionForCertainAmountOrHigherDataItem instantiates a new ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmedCoinsTransactionForCertainAmountOrHigherDataItemWithDefaults

`func NewConfirmedCoinsTransactionForCertainAmountOrHigherDataItemWithDefaults() *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem`

NewConfirmedCoinsTransactionForCertainAmountOrHigherDataItemWithDefaults instantiates a new ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransactionId

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetMinedInBlock

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetMinedInBlock() AddressCoinsTransactionConfirmedDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetMinedInBlockOk() (*AddressCoinsTransactionConfirmedDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) SetMinedInBlock(v AddressCoinsTransactionConfirmedDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetAmount

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ConfirmedCoinsTransactionForCertainAmountOrHigherDataItem) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


