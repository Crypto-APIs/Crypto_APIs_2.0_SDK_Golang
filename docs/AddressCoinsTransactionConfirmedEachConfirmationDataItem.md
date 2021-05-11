# AddressCoinsTransactionConfirmedEachConfirmationDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**Address** | **string** | Defines the specific address to which the transaction has been sent. | 
**MinedInBlock** | [**AddressCoinsTransactionConfirmedEachConfirmationDataItemMinedInBlock**](AddressCoinsTransactionConfirmedEachConfirmationDataItemMinedInBlock.md) |  | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**CurrentConfirmations** | **int32** | Defines the number of currently received confirmations for the transaction. | 
**TargetConfirmations** | **int32** | Defines the number of confirmation transactions requested as callbacks, i.e. the system can notify till the n-th confirmation. | 
**Amount** | **string** | Defines the amount of coins sent with the confirmed transaction. | 
**Unit** | **string** | Defines the unit of the transaction, e.g. BTC. | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 

## Methods

### NewAddressCoinsTransactionConfirmedEachConfirmationDataItem

`func NewAddressCoinsTransactionConfirmedEachConfirmationDataItem(blockchain string, network string, address string, minedInBlock AddressCoinsTransactionConfirmedEachConfirmationDataItemMinedInBlock, transactionId string, currentConfirmations int32, targetConfirmations int32, amount string, unit string, direction string, ) *AddressCoinsTransactionConfirmedEachConfirmationDataItem`

NewAddressCoinsTransactionConfirmedEachConfirmationDataItem instantiates a new AddressCoinsTransactionConfirmedEachConfirmationDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressCoinsTransactionConfirmedEachConfirmationDataItemWithDefaults

`func NewAddressCoinsTransactionConfirmedEachConfirmationDataItemWithDefaults() *AddressCoinsTransactionConfirmedEachConfirmationDataItem`

NewAddressCoinsTransactionConfirmedEachConfirmationDataItemWithDefaults instantiates a new AddressCoinsTransactionConfirmedEachConfirmationDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMinedInBlock

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetMinedInBlock() AddressCoinsTransactionConfirmedEachConfirmationDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetMinedInBlockOk() (*AddressCoinsTransactionConfirmedEachConfirmationDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetMinedInBlock(v AddressCoinsTransactionConfirmedEachConfirmationDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetTransactionId

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetCurrentConfirmations

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmations() int32`

GetCurrentConfirmations returns the CurrentConfirmations field if non-nil, zero value otherwise.

### GetCurrentConfirmationsOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmationsOk() (*int32, bool)`

GetCurrentConfirmationsOk returns a tuple with the CurrentConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfirmations

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetCurrentConfirmations(v int32)`

SetCurrentConfirmations sets CurrentConfirmations field to given value.


### GetTargetConfirmations

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmations() int32`

GetTargetConfirmations returns the TargetConfirmations field if non-nil, zero value otherwise.

### GetTargetConfirmationsOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmationsOk() (*int32, bool)`

GetTargetConfirmationsOk returns a tuple with the TargetConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfirmations

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetTargetConfirmations(v int32)`

SetTargetConfirmations sets TargetConfirmations field to given value.


### GetAmount

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDirection

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressCoinsTransactionConfirmedEachConfirmationDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


