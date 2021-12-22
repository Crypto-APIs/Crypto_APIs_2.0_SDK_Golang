# AddressInternalTransactionConfirmedEachConfirmationDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**Address** | **string** | Defines the specific address of the internal transaction. | 
**MinedInBlock** | [**AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock**](AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock.md) |  | 
**ParentTransactionId** | **string** | Defines the Parent Transaction&#39;s unique ID. | 
**OperationId** | **string** | Defines the specific operation&#39;s unique ID. | 
**CurrentConfirmations** | **int32** | Defines the number of currently received confirmations for the transaction. | 
**TargetConfirmations** | **int32** | Defines the number of confirmation transactions requested as callbacks, i.e. the system can notify till the n-th confirmation. | 
**Amount** | **string** | Defines the amount of coins sent with the confirmed transaction. | 
**Unit** | **string** | Defines the unit of the transaction, e.g. Gwei. | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 

## Methods

### NewAddressInternalTransactionConfirmedEachConfirmationDataItem

`func NewAddressInternalTransactionConfirmedEachConfirmationDataItem(blockchain string, network string, address string, minedInBlock AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock, parentTransactionId string, operationId string, currentConfirmations int32, targetConfirmations int32, amount string, unit string, direction string, ) *AddressInternalTransactionConfirmedEachConfirmationDataItem`

NewAddressInternalTransactionConfirmedEachConfirmationDataItem instantiates a new AddressInternalTransactionConfirmedEachConfirmationDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressInternalTransactionConfirmedEachConfirmationDataItemWithDefaults

`func NewAddressInternalTransactionConfirmedEachConfirmationDataItemWithDefaults() *AddressInternalTransactionConfirmedEachConfirmationDataItem`

NewAddressInternalTransactionConfirmedEachConfirmationDataItemWithDefaults instantiates a new AddressInternalTransactionConfirmedEachConfirmationDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMinedInBlock

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetMinedInBlock() AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetMinedInBlockOk() (*AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetMinedInBlock(v AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetParentTransactionId

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetParentTransactionId() string`

GetParentTransactionId returns the ParentTransactionId field if non-nil, zero value otherwise.

### GetParentTransactionIdOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetParentTransactionIdOk() (*string, bool)`

GetParentTransactionIdOk returns a tuple with the ParentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransactionId

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetParentTransactionId(v string)`

SetParentTransactionId sets ParentTransactionId field to given value.


### GetOperationId

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetCurrentConfirmations

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmations() int32`

GetCurrentConfirmations returns the CurrentConfirmations field if non-nil, zero value otherwise.

### GetCurrentConfirmationsOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmationsOk() (*int32, bool)`

GetCurrentConfirmationsOk returns a tuple with the CurrentConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfirmations

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetCurrentConfirmations(v int32)`

SetCurrentConfirmations sets CurrentConfirmations field to given value.


### GetTargetConfirmations

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmations() int32`

GetTargetConfirmations returns the TargetConfirmations field if non-nil, zero value otherwise.

### GetTargetConfirmationsOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmationsOk() (*int32, bool)`

GetTargetConfirmationsOk returns a tuple with the TargetConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfirmations

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetTargetConfirmations(v int32)`

SetTargetConfirmations sets TargetConfirmations field to given value.


### GetAmount

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDirection

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressInternalTransactionConfirmedEachConfirmationDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


