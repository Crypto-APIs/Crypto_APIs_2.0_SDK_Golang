# TransactionRequestApprovalDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**RequiredApprovals** | **int32** | The required number of approvals needed to approve the transaction. | 
**RequiredRejections** | **int32** | The required number of rejections needed to reject the transaction. | 
**CurrentApprovals** | **int32** | The current number of approvals given for the transaction. | 
**CurrentRejections** | **int32** | The current number of rejections given for the transaction. | 

## Methods

### NewTransactionRequestApprovalDataItem

`func NewTransactionRequestApprovalDataItem(blockchain string, network string, requiredApprovals int32, requiredRejections int32, currentApprovals int32, currentRejections int32, ) *TransactionRequestApprovalDataItem`

NewTransactionRequestApprovalDataItem instantiates a new TransactionRequestApprovalDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestApprovalDataItemWithDefaults

`func NewTransactionRequestApprovalDataItemWithDefaults() *TransactionRequestApprovalDataItem`

NewTransactionRequestApprovalDataItemWithDefaults instantiates a new TransactionRequestApprovalDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *TransactionRequestApprovalDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *TransactionRequestApprovalDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *TransactionRequestApprovalDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *TransactionRequestApprovalDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionRequestApprovalDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionRequestApprovalDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRequiredApprovals

`func (o *TransactionRequestApprovalDataItem) GetRequiredApprovals() int32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *TransactionRequestApprovalDataItem) GetRequiredApprovalsOk() (*int32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *TransactionRequestApprovalDataItem) SetRequiredApprovals(v int32)`

SetRequiredApprovals sets RequiredApprovals field to given value.


### GetRequiredRejections

`func (o *TransactionRequestApprovalDataItem) GetRequiredRejections() int32`

GetRequiredRejections returns the RequiredRejections field if non-nil, zero value otherwise.

### GetRequiredRejectionsOk

`func (o *TransactionRequestApprovalDataItem) GetRequiredRejectionsOk() (*int32, bool)`

GetRequiredRejectionsOk returns a tuple with the RequiredRejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRejections

`func (o *TransactionRequestApprovalDataItem) SetRequiredRejections(v int32)`

SetRequiredRejections sets RequiredRejections field to given value.


### GetCurrentApprovals

`func (o *TransactionRequestApprovalDataItem) GetCurrentApprovals() int32`

GetCurrentApprovals returns the CurrentApprovals field if non-nil, zero value otherwise.

### GetCurrentApprovalsOk

`func (o *TransactionRequestApprovalDataItem) GetCurrentApprovalsOk() (*int32, bool)`

GetCurrentApprovalsOk returns a tuple with the CurrentApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentApprovals

`func (o *TransactionRequestApprovalDataItem) SetCurrentApprovals(v int32)`

SetCurrentApprovals sets CurrentApprovals field to given value.


### GetCurrentRejections

`func (o *TransactionRequestApprovalDataItem) GetCurrentRejections() int32`

GetCurrentRejections returns the CurrentRejections field if non-nil, zero value otherwise.

### GetCurrentRejectionsOk

`func (o *TransactionRequestApprovalDataItem) GetCurrentRejectionsOk() (*int32, bool)`

GetCurrentRejectionsOk returns a tuple with the CurrentRejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRejections

`func (o *TransactionRequestApprovalDataItem) SetCurrentRejections(v int32)`

SetCurrentRejections sets CurrentRejections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


