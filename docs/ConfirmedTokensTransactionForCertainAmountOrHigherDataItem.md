# ConfirmedTokensTransactionForCertainAmountOrHigherDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;mordor\&quot; are test networks. | 
**MinedInBlock** | [**AddressTokensTransactionConfirmedDataItemMinedInBlock**](AddressTokensTransactionConfirmedDataItemMinedInBlock.md) |  | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**TokenType** | **string** | Defines the type of token sent with the transaction, e.g. ERC 20. | 
**Token** | [**ConfirmedTokensTransactionForCertainAmountOrHigherToken**](ConfirmedTokensTransactionForCertainAmountOrHigherToken.md) |  | 

## Methods

### NewConfirmedTokensTransactionForCertainAmountOrHigherDataItem

`func NewConfirmedTokensTransactionForCertainAmountOrHigherDataItem(blockchain string, network string, minedInBlock AddressTokensTransactionConfirmedDataItemMinedInBlock, transactionId string, tokenType string, token ConfirmedTokensTransactionForCertainAmountOrHigherToken, ) *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem`

NewConfirmedTokensTransactionForCertainAmountOrHigherDataItem instantiates a new ConfirmedTokensTransactionForCertainAmountOrHigherDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmedTokensTransactionForCertainAmountOrHigherDataItemWithDefaults

`func NewConfirmedTokensTransactionForCertainAmountOrHigherDataItemWithDefaults() *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem`

NewConfirmedTokensTransactionForCertainAmountOrHigherDataItemWithDefaults instantiates a new ConfirmedTokensTransactionForCertainAmountOrHigherDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetMinedInBlock

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetMinedInBlock() AddressTokensTransactionConfirmedDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetMinedInBlockOk() (*AddressTokensTransactionConfirmedDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) SetMinedInBlock(v AddressTokensTransactionConfirmedDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetTransactionId

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTokenType

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetToken

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetToken() ConfirmedTokensTransactionForCertainAmountOrHigherToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) GetTokenOk() (*ConfirmedTokensTransactionForCertainAmountOrHigherToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfirmedTokensTransactionForCertainAmountOrHigherDataItem) SetToken(v ConfirmedTokensTransactionForCertainAmountOrHigherToken)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


