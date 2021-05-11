# CoinsForwardingSuccessDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**FromAddress** | **string** | Represents the hash of the address that provides the coins. | 
**ToAddress** | **string** | Represents the hash of the address to forward the coins to. | 
**ForwardedAmount** | **string** | Represents the amount of coins that have been forwarded. | 
**ForwardedUnit** | **string** | Represents the unit of coins that have been forwarded, e.g. BTC. | 
**SpentFeesAmount** | **string** | Represents the amount of the fee spent for the coins to be forwarded. | 
**SpentFeesUnit** | **string** | Represents the unit of the fee spent for the coins to be forwarded, e.g. BTC. | 
**TriggerTransactionId** | **string** | Defines the unique Transaction ID that triggered the coin forwarding. | 
**ForwardingTransactionId** | **string** | Defines the unique Transaction ID that forwarded the coins. | 

## Methods

### NewCoinsForwardingSuccessDataItem

`func NewCoinsForwardingSuccessDataItem(blockchain string, network string, fromAddress string, toAddress string, forwardedAmount string, forwardedUnit string, spentFeesAmount string, spentFeesUnit string, triggerTransactionId string, forwardingTransactionId string, ) *CoinsForwardingSuccessDataItem`

NewCoinsForwardingSuccessDataItem instantiates a new CoinsForwardingSuccessDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoinsForwardingSuccessDataItemWithDefaults

`func NewCoinsForwardingSuccessDataItemWithDefaults() *CoinsForwardingSuccessDataItem`

NewCoinsForwardingSuccessDataItemWithDefaults instantiates a new CoinsForwardingSuccessDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *CoinsForwardingSuccessDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *CoinsForwardingSuccessDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *CoinsForwardingSuccessDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *CoinsForwardingSuccessDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CoinsForwardingSuccessDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CoinsForwardingSuccessDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFromAddress

`func (o *CoinsForwardingSuccessDataItem) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CoinsForwardingSuccessDataItem) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CoinsForwardingSuccessDataItem) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *CoinsForwardingSuccessDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CoinsForwardingSuccessDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CoinsForwardingSuccessDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetForwardedAmount

`func (o *CoinsForwardingSuccessDataItem) GetForwardedAmount() string`

GetForwardedAmount returns the ForwardedAmount field if non-nil, zero value otherwise.

### GetForwardedAmountOk

`func (o *CoinsForwardingSuccessDataItem) GetForwardedAmountOk() (*string, bool)`

GetForwardedAmountOk returns a tuple with the ForwardedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedAmount

`func (o *CoinsForwardingSuccessDataItem) SetForwardedAmount(v string)`

SetForwardedAmount sets ForwardedAmount field to given value.


### GetForwardedUnit

`func (o *CoinsForwardingSuccessDataItem) GetForwardedUnit() string`

GetForwardedUnit returns the ForwardedUnit field if non-nil, zero value otherwise.

### GetForwardedUnitOk

`func (o *CoinsForwardingSuccessDataItem) GetForwardedUnitOk() (*string, bool)`

GetForwardedUnitOk returns a tuple with the ForwardedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedUnit

`func (o *CoinsForwardingSuccessDataItem) SetForwardedUnit(v string)`

SetForwardedUnit sets ForwardedUnit field to given value.


### GetSpentFeesAmount

`func (o *CoinsForwardingSuccessDataItem) GetSpentFeesAmount() string`

GetSpentFeesAmount returns the SpentFeesAmount field if non-nil, zero value otherwise.

### GetSpentFeesAmountOk

`func (o *CoinsForwardingSuccessDataItem) GetSpentFeesAmountOk() (*string, bool)`

GetSpentFeesAmountOk returns a tuple with the SpentFeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentFeesAmount

`func (o *CoinsForwardingSuccessDataItem) SetSpentFeesAmount(v string)`

SetSpentFeesAmount sets SpentFeesAmount field to given value.


### GetSpentFeesUnit

`func (o *CoinsForwardingSuccessDataItem) GetSpentFeesUnit() string`

GetSpentFeesUnit returns the SpentFeesUnit field if non-nil, zero value otherwise.

### GetSpentFeesUnitOk

`func (o *CoinsForwardingSuccessDataItem) GetSpentFeesUnitOk() (*string, bool)`

GetSpentFeesUnitOk returns a tuple with the SpentFeesUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentFeesUnit

`func (o *CoinsForwardingSuccessDataItem) SetSpentFeesUnit(v string)`

SetSpentFeesUnit sets SpentFeesUnit field to given value.


### GetTriggerTransactionId

`func (o *CoinsForwardingSuccessDataItem) GetTriggerTransactionId() string`

GetTriggerTransactionId returns the TriggerTransactionId field if non-nil, zero value otherwise.

### GetTriggerTransactionIdOk

`func (o *CoinsForwardingSuccessDataItem) GetTriggerTransactionIdOk() (*string, bool)`

GetTriggerTransactionIdOk returns a tuple with the TriggerTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTransactionId

`func (o *CoinsForwardingSuccessDataItem) SetTriggerTransactionId(v string)`

SetTriggerTransactionId sets TriggerTransactionId field to given value.


### GetForwardingTransactionId

`func (o *CoinsForwardingSuccessDataItem) GetForwardingTransactionId() string`

GetForwardingTransactionId returns the ForwardingTransactionId field if non-nil, zero value otherwise.

### GetForwardingTransactionIdOk

`func (o *CoinsForwardingSuccessDataItem) GetForwardingTransactionIdOk() (*string, bool)`

GetForwardingTransactionIdOk returns a tuple with the ForwardingTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingTransactionId

`func (o *CoinsForwardingSuccessDataItem) SetForwardingTransactionId(v string)`

SetForwardingTransactionId sets ForwardingTransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


