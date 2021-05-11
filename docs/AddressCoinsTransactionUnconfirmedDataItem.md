# AddressCoinsTransactionUnconfirmedDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**Address** | **string** | Defines the specific address to which the coin transaction has been sent and is pending confirmation. | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**Amount** | **string** | Defines the amount of coins sent with the transaction that is pending confirmation. | 
**Unit** | **string** | Defines the unit of the transaction, e.g. BTC. | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 
**FirstSeenInMempoolTimestamp** | **int32** | Defines the exact time the transaction has been first accepted into the mempool to await confirmation as timestamp. | 

## Methods

### NewAddressCoinsTransactionUnconfirmedDataItem

`func NewAddressCoinsTransactionUnconfirmedDataItem(blockchain string, network string, address string, transactionId string, amount string, unit string, direction string, firstSeenInMempoolTimestamp int32, ) *AddressCoinsTransactionUnconfirmedDataItem`

NewAddressCoinsTransactionUnconfirmedDataItem instantiates a new AddressCoinsTransactionUnconfirmedDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressCoinsTransactionUnconfirmedDataItemWithDefaults

`func NewAddressCoinsTransactionUnconfirmedDataItemWithDefaults() *AddressCoinsTransactionUnconfirmedDataItem`

NewAddressCoinsTransactionUnconfirmedDataItemWithDefaults instantiates a new AddressCoinsTransactionUnconfirmedDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTransactionId

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAmount

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDirection

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetFirstSeenInMempoolTimestamp

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetFirstSeenInMempoolTimestamp() int32`

GetFirstSeenInMempoolTimestamp returns the FirstSeenInMempoolTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenInMempoolTimestampOk

`func (o *AddressCoinsTransactionUnconfirmedDataItem) GetFirstSeenInMempoolTimestampOk() (*int32, bool)`

GetFirstSeenInMempoolTimestampOk returns a tuple with the FirstSeenInMempoolTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenInMempoolTimestamp

`func (o *AddressCoinsTransactionUnconfirmedDataItem) SetFirstSeenInMempoolTimestamp(v int32)`

SetFirstSeenInMempoolTimestamp sets FirstSeenInMempoolTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


