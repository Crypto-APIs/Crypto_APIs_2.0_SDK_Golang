# AddressCoinsTransactionConfirmedDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;mordor\&quot; are test networks. | 
**Address** | **string** | Defines the specific address to which the coin transaction has been sent and is confirmed. | 
**MinedInBlock** | [**AddressCoinsTransactionConfirmedDataItemMinedInBlock**](AddressCoinsTransactionConfirmedDataItemMinedInBlock.md) |  | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**Amount** | **string** | Defines the amount of coins sent with the confirmed transaction. | 
**Unit** | **string** | Defines the unit of the transaction, e.g. BTC. | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 

## Methods

### NewAddressCoinsTransactionConfirmedDataItem

`func NewAddressCoinsTransactionConfirmedDataItem(blockchain string, network string, address string, minedInBlock AddressCoinsTransactionConfirmedDataItemMinedInBlock, transactionId string, amount string, unit string, direction string, ) *AddressCoinsTransactionConfirmedDataItem`

NewAddressCoinsTransactionConfirmedDataItem instantiates a new AddressCoinsTransactionConfirmedDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressCoinsTransactionConfirmedDataItemWithDefaults

`func NewAddressCoinsTransactionConfirmedDataItemWithDefaults() *AddressCoinsTransactionConfirmedDataItem`

NewAddressCoinsTransactionConfirmedDataItemWithDefaults instantiates a new AddressCoinsTransactionConfirmedDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressCoinsTransactionConfirmedDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressCoinsTransactionConfirmedDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressCoinsTransactionConfirmedDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressCoinsTransactionConfirmedDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressCoinsTransactionConfirmedDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressCoinsTransactionConfirmedDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMinedInBlock

`func (o *AddressCoinsTransactionConfirmedDataItem) GetMinedInBlock() AddressCoinsTransactionConfirmedDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetMinedInBlockOk() (*AddressCoinsTransactionConfirmedDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *AddressCoinsTransactionConfirmedDataItem) SetMinedInBlock(v AddressCoinsTransactionConfirmedDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetTransactionId

`func (o *AddressCoinsTransactionConfirmedDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AddressCoinsTransactionConfirmedDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAmount

`func (o *AddressCoinsTransactionConfirmedDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressCoinsTransactionConfirmedDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *AddressCoinsTransactionConfirmedDataItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AddressCoinsTransactionConfirmedDataItem) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDirection

`func (o *AddressCoinsTransactionConfirmedDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressCoinsTransactionConfirmedDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressCoinsTransactionConfirmedDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


