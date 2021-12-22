# AddressInternalTransactionConfirmedDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**Address** | **string** | Defines the specific address of the internal transaction. | 
**MinedInBlock** | [**AddressInternalTransactionConfirmedDataItemMinedInBlock**](AddressInternalTransactionConfirmedDataItemMinedInBlock.md) |  | 
**ParentTransactionId** | **string** | Defines the Parent Transaction&#39;s unique ID. | 
**OperationId** | **string** | Defines the specific operation&#39;s unique ID. | 
**Amount** | **string** | Defines the amount of coins sent with the confirmed transaction. | 
**Unit** | **string** | Defines the unit of the transaction, e.g. Gwei. | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 

## Methods

### NewAddressInternalTransactionConfirmedDataItem

`func NewAddressInternalTransactionConfirmedDataItem(blockchain string, network string, address string, minedInBlock AddressInternalTransactionConfirmedDataItemMinedInBlock, parentTransactionId string, operationId string, amount string, unit string, direction string, ) *AddressInternalTransactionConfirmedDataItem`

NewAddressInternalTransactionConfirmedDataItem instantiates a new AddressInternalTransactionConfirmedDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressInternalTransactionConfirmedDataItemWithDefaults

`func NewAddressInternalTransactionConfirmedDataItemWithDefaults() *AddressInternalTransactionConfirmedDataItem`

NewAddressInternalTransactionConfirmedDataItemWithDefaults instantiates a new AddressInternalTransactionConfirmedDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressInternalTransactionConfirmedDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressInternalTransactionConfirmedDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressInternalTransactionConfirmedDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressInternalTransactionConfirmedDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressInternalTransactionConfirmedDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressInternalTransactionConfirmedDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMinedInBlock

`func (o *AddressInternalTransactionConfirmedDataItem) GetMinedInBlock() AddressInternalTransactionConfirmedDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetMinedInBlockOk() (*AddressInternalTransactionConfirmedDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *AddressInternalTransactionConfirmedDataItem) SetMinedInBlock(v AddressInternalTransactionConfirmedDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetParentTransactionId

`func (o *AddressInternalTransactionConfirmedDataItem) GetParentTransactionId() string`

GetParentTransactionId returns the ParentTransactionId field if non-nil, zero value otherwise.

### GetParentTransactionIdOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetParentTransactionIdOk() (*string, bool)`

GetParentTransactionIdOk returns a tuple with the ParentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransactionId

`func (o *AddressInternalTransactionConfirmedDataItem) SetParentTransactionId(v string)`

SetParentTransactionId sets ParentTransactionId field to given value.


### GetOperationId

`func (o *AddressInternalTransactionConfirmedDataItem) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *AddressInternalTransactionConfirmedDataItem) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetAmount

`func (o *AddressInternalTransactionConfirmedDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressInternalTransactionConfirmedDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *AddressInternalTransactionConfirmedDataItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AddressInternalTransactionConfirmedDataItem) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetDirection

`func (o *AddressInternalTransactionConfirmedDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressInternalTransactionConfirmedDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressInternalTransactionConfirmedDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


