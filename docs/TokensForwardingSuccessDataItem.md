# TokensForwardingSuccessDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**FromAddress** | **string** | Represents the hash of the address that provides the tokens. | 
**ToAddress** | **string** | Represents the hash of the address to forward the tokens to. | 
**SpentFeesAmount** | **string** | Represents the amount of the fee spent for the tokens to be forwarded. | 
**SpentFeesUnit** | **string** | Represents the unit of the fee spent for the tokens to be forwarded, e.g. BTC. | 
**TriggerTransactionId** | **string** | Defines the unique Transaction ID that triggered the token forwarding. | 
**ForwardingTransactionId** | **string** | Defines the unique Transaction ID that forwarded the tokens. | 
**TokenType** | **string** | Defines the type of token sent with the transaction, e.g. ERC 20. | 
**Token** | [**TokensForwardingSuccessToken**](TokensForwardingSuccessToken.md) |  | 

## Methods

### NewTokensForwardingSuccessDataItem

`func NewTokensForwardingSuccessDataItem(blockchain string, network string, fromAddress string, toAddress string, spentFeesAmount string, spentFeesUnit string, triggerTransactionId string, forwardingTransactionId string, tokenType string, token TokensForwardingSuccessToken, ) *TokensForwardingSuccessDataItem`

NewTokensForwardingSuccessDataItem instantiates a new TokensForwardingSuccessDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensForwardingSuccessDataItemWithDefaults

`func NewTokensForwardingSuccessDataItemWithDefaults() *TokensForwardingSuccessDataItem`

NewTokensForwardingSuccessDataItemWithDefaults instantiates a new TokensForwardingSuccessDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *TokensForwardingSuccessDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *TokensForwardingSuccessDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *TokensForwardingSuccessDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *TokensForwardingSuccessDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TokensForwardingSuccessDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TokensForwardingSuccessDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFromAddress

`func (o *TokensForwardingSuccessDataItem) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *TokensForwardingSuccessDataItem) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *TokensForwardingSuccessDataItem) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *TokensForwardingSuccessDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *TokensForwardingSuccessDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *TokensForwardingSuccessDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetSpentFeesAmount

`func (o *TokensForwardingSuccessDataItem) GetSpentFeesAmount() string`

GetSpentFeesAmount returns the SpentFeesAmount field if non-nil, zero value otherwise.

### GetSpentFeesAmountOk

`func (o *TokensForwardingSuccessDataItem) GetSpentFeesAmountOk() (*string, bool)`

GetSpentFeesAmountOk returns a tuple with the SpentFeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentFeesAmount

`func (o *TokensForwardingSuccessDataItem) SetSpentFeesAmount(v string)`

SetSpentFeesAmount sets SpentFeesAmount field to given value.


### GetSpentFeesUnit

`func (o *TokensForwardingSuccessDataItem) GetSpentFeesUnit() string`

GetSpentFeesUnit returns the SpentFeesUnit field if non-nil, zero value otherwise.

### GetSpentFeesUnitOk

`func (o *TokensForwardingSuccessDataItem) GetSpentFeesUnitOk() (*string, bool)`

GetSpentFeesUnitOk returns a tuple with the SpentFeesUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentFeesUnit

`func (o *TokensForwardingSuccessDataItem) SetSpentFeesUnit(v string)`

SetSpentFeesUnit sets SpentFeesUnit field to given value.


### GetTriggerTransactionId

`func (o *TokensForwardingSuccessDataItem) GetTriggerTransactionId() string`

GetTriggerTransactionId returns the TriggerTransactionId field if non-nil, zero value otherwise.

### GetTriggerTransactionIdOk

`func (o *TokensForwardingSuccessDataItem) GetTriggerTransactionIdOk() (*string, bool)`

GetTriggerTransactionIdOk returns a tuple with the TriggerTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTransactionId

`func (o *TokensForwardingSuccessDataItem) SetTriggerTransactionId(v string)`

SetTriggerTransactionId sets TriggerTransactionId field to given value.


### GetForwardingTransactionId

`func (o *TokensForwardingSuccessDataItem) GetForwardingTransactionId() string`

GetForwardingTransactionId returns the ForwardingTransactionId field if non-nil, zero value otherwise.

### GetForwardingTransactionIdOk

`func (o *TokensForwardingSuccessDataItem) GetForwardingTransactionIdOk() (*string, bool)`

GetForwardingTransactionIdOk returns a tuple with the ForwardingTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingTransactionId

`func (o *TokensForwardingSuccessDataItem) SetForwardingTransactionId(v string)`

SetForwardingTransactionId sets ForwardingTransactionId field to given value.


### GetTokenType

`func (o *TokensForwardingSuccessDataItem) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokensForwardingSuccessDataItem) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokensForwardingSuccessDataItem) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetToken

`func (o *TokensForwardingSuccessDataItem) GetToken() TokensForwardingSuccessToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokensForwardingSuccessDataItem) GetTokenOk() (*TokensForwardingSuccessToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokensForwardingSuccessDataItem) SetToken(v TokensForwardingSuccessToken)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


