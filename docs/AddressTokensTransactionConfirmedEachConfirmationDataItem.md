# AddressTokensTransactionConfirmedEachConfirmationDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;,  are test networks. | 
**Address** | **string** | Defines the specific address to which the transaction has been sent. | 
**MinedInBlock** | [**AddressTokensTransactionConfirmedDataItemMinedInBlock**](AddressTokensTransactionConfirmedDataItemMinedInBlock.md) |  | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**CurrentConfirmations** | **int32** | Defines the number of currently received confirmations for the transaction. | 
**TargetConfirmations** | **int32** | Defines the number of confirmation transactions requested as callbacks, i.e. the system can notify till the n-th confirmation. | 
**TokenType** | **string** | Defines the type of token sent with the transaction, e.g. ERC 20. | 
**Token** | [**AddressTokensTransactionConfirmedEachConfirmationToken**](AddressTokensTransactionConfirmedEachConfirmationToken.md) |  | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 

## Methods

### NewAddressTokensTransactionConfirmedEachConfirmationDataItem

`func NewAddressTokensTransactionConfirmedEachConfirmationDataItem(blockchain string, network string, address string, minedInBlock AddressTokensTransactionConfirmedDataItemMinedInBlock, transactionId string, currentConfirmations int32, targetConfirmations int32, tokenType string, token AddressTokensTransactionConfirmedEachConfirmationToken, direction string, ) *AddressTokensTransactionConfirmedEachConfirmationDataItem`

NewAddressTokensTransactionConfirmedEachConfirmationDataItem instantiates a new AddressTokensTransactionConfirmedEachConfirmationDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTokensTransactionConfirmedEachConfirmationDataItemWithDefaults

`func NewAddressTokensTransactionConfirmedEachConfirmationDataItemWithDefaults() *AddressTokensTransactionConfirmedEachConfirmationDataItem`

NewAddressTokensTransactionConfirmedEachConfirmationDataItemWithDefaults instantiates a new AddressTokensTransactionConfirmedEachConfirmationDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMinedInBlock

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetMinedInBlock() AddressTokensTransactionConfirmedDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetMinedInBlockOk() (*AddressTokensTransactionConfirmedDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetMinedInBlock(v AddressTokensTransactionConfirmedDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetTransactionId

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetCurrentConfirmations

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmations() int32`

GetCurrentConfirmations returns the CurrentConfirmations field if non-nil, zero value otherwise.

### GetCurrentConfirmationsOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetCurrentConfirmationsOk() (*int32, bool)`

GetCurrentConfirmationsOk returns a tuple with the CurrentConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfirmations

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetCurrentConfirmations(v int32)`

SetCurrentConfirmations sets CurrentConfirmations field to given value.


### GetTargetConfirmations

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmations() int32`

GetTargetConfirmations returns the TargetConfirmations field if non-nil, zero value otherwise.

### GetTargetConfirmationsOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTargetConfirmationsOk() (*int32, bool)`

GetTargetConfirmationsOk returns a tuple with the TargetConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfirmations

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetTargetConfirmations(v int32)`

SetTargetConfirmations sets TargetConfirmations field to given value.


### GetTokenType

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetToken

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetToken() AddressTokensTransactionConfirmedEachConfirmationToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetTokenOk() (*AddressTokensTransactionConfirmedEachConfirmationToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetToken(v AddressTokensTransactionConfirmedEachConfirmationToken)`

SetToken sets Token field to given value.


### GetDirection

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressTokensTransactionConfirmedEachConfirmationDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


