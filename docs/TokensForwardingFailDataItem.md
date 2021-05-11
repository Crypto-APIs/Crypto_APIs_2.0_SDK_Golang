# TokensForwardingFailDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**FromAddress** | **string** | Represents the hash of the address that provides the tokens. | 
**ToAddress** | **string** | Represents the hash of the address to forward the tokens to. | 
**TriggerTransactionId** | **string** | Defines the unique Transaction ID that triggered the token forwarding. | 
**ErrorCode** | **string** | Represents the error code received for the failed token forwarding. | 
**ErrorMessage** | **string** | Represents the error message received for the failed token forwarding. | 

## Methods

### NewTokensForwardingFailDataItem

`func NewTokensForwardingFailDataItem(blockchain string, network string, fromAddress string, toAddress string, triggerTransactionId string, errorCode string, errorMessage string, ) *TokensForwardingFailDataItem`

NewTokensForwardingFailDataItem instantiates a new TokensForwardingFailDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensForwardingFailDataItemWithDefaults

`func NewTokensForwardingFailDataItemWithDefaults() *TokensForwardingFailDataItem`

NewTokensForwardingFailDataItemWithDefaults instantiates a new TokensForwardingFailDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *TokensForwardingFailDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *TokensForwardingFailDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *TokensForwardingFailDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *TokensForwardingFailDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TokensForwardingFailDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TokensForwardingFailDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFromAddress

`func (o *TokensForwardingFailDataItem) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *TokensForwardingFailDataItem) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *TokensForwardingFailDataItem) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *TokensForwardingFailDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *TokensForwardingFailDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *TokensForwardingFailDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetTriggerTransactionId

`func (o *TokensForwardingFailDataItem) GetTriggerTransactionId() string`

GetTriggerTransactionId returns the TriggerTransactionId field if non-nil, zero value otherwise.

### GetTriggerTransactionIdOk

`func (o *TokensForwardingFailDataItem) GetTriggerTransactionIdOk() (*string, bool)`

GetTriggerTransactionIdOk returns a tuple with the TriggerTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTransactionId

`func (o *TokensForwardingFailDataItem) SetTriggerTransactionId(v string)`

SetTriggerTransactionId sets TriggerTransactionId field to given value.


### GetErrorCode

`func (o *TokensForwardingFailDataItem) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TokensForwardingFailDataItem) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TokensForwardingFailDataItem) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *TokensForwardingFailDataItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TokensForwardingFailDataItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TokensForwardingFailDataItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


