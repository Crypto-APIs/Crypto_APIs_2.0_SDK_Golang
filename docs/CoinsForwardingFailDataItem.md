# CoinsForwardingFailDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**FromAddress** | **string** | Represents the hash of the address that provides the coins. | 
**ToAddress** | **string** | Represents the hash of the address to forward the coins to. | 
**TriggerTransactionId** | **string** | Defines the unique Transaction ID that triggered the coin forwarding. | 
**ErrorCode** | **string** | Represents the error code received for the failed coin forwarding. | 
**ErrorMessage** | **string** | Represents the error message received for the failed coin forwarding. | 

## Methods

### NewCoinsForwardingFailDataItem

`func NewCoinsForwardingFailDataItem(blockchain string, network string, fromAddress string, toAddress string, triggerTransactionId string, errorCode string, errorMessage string, ) *CoinsForwardingFailDataItem`

NewCoinsForwardingFailDataItem instantiates a new CoinsForwardingFailDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoinsForwardingFailDataItemWithDefaults

`func NewCoinsForwardingFailDataItemWithDefaults() *CoinsForwardingFailDataItem`

NewCoinsForwardingFailDataItemWithDefaults instantiates a new CoinsForwardingFailDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *CoinsForwardingFailDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *CoinsForwardingFailDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *CoinsForwardingFailDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *CoinsForwardingFailDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CoinsForwardingFailDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CoinsForwardingFailDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFromAddress

`func (o *CoinsForwardingFailDataItem) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CoinsForwardingFailDataItem) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CoinsForwardingFailDataItem) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *CoinsForwardingFailDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CoinsForwardingFailDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CoinsForwardingFailDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetTriggerTransactionId

`func (o *CoinsForwardingFailDataItem) GetTriggerTransactionId() string`

GetTriggerTransactionId returns the TriggerTransactionId field if non-nil, zero value otherwise.

### GetTriggerTransactionIdOk

`func (o *CoinsForwardingFailDataItem) GetTriggerTransactionIdOk() (*string, bool)`

GetTriggerTransactionIdOk returns a tuple with the TriggerTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTransactionId

`func (o *CoinsForwardingFailDataItem) SetTriggerTransactionId(v string)`

SetTriggerTransactionId sets TriggerTransactionId field to given value.


### GetErrorCode

`func (o *CoinsForwardingFailDataItem) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CoinsForwardingFailDataItem) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CoinsForwardingFailDataItem) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *CoinsForwardingFailDataItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CoinsForwardingFailDataItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CoinsForwardingFailDataItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


