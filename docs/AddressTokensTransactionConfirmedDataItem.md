# AddressTokensTransactionConfirmedDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;,  are test networks. | 
**Address** | **string** | Defines the specific address to which the transaction has been sent. | 
**MinedInBlock** | [**AddressTokensTransactionConfirmedDataItemMinedInBlock**](AddressTokensTransactionConfirmedDataItemMinedInBlock.md) |  | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**TokenType** | **string** | Defines the type of token sent with the transaction, e.g. ERC 20. | 
**Token** | [**AddressTokensTransactionConfirmedToken**](AddressTokensTransactionConfirmedToken.md) |  | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 

## Methods

### NewAddressTokensTransactionConfirmedDataItem

`func NewAddressTokensTransactionConfirmedDataItem(blockchain string, network string, address string, minedInBlock AddressTokensTransactionConfirmedDataItemMinedInBlock, transactionId string, tokenType string, token AddressTokensTransactionConfirmedToken, direction string, ) *AddressTokensTransactionConfirmedDataItem`

NewAddressTokensTransactionConfirmedDataItem instantiates a new AddressTokensTransactionConfirmedDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTokensTransactionConfirmedDataItemWithDefaults

`func NewAddressTokensTransactionConfirmedDataItemWithDefaults() *AddressTokensTransactionConfirmedDataItem`

NewAddressTokensTransactionConfirmedDataItemWithDefaults instantiates a new AddressTokensTransactionConfirmedDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressTokensTransactionConfirmedDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressTokensTransactionConfirmedDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressTokensTransactionConfirmedDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressTokensTransactionConfirmedDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressTokensTransactionConfirmedDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressTokensTransactionConfirmedDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMinedInBlock

`func (o *AddressTokensTransactionConfirmedDataItem) GetMinedInBlock() AddressTokensTransactionConfirmedDataItemMinedInBlock`

GetMinedInBlock returns the MinedInBlock field if non-nil, zero value otherwise.

### GetMinedInBlockOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetMinedInBlockOk() (*AddressTokensTransactionConfirmedDataItemMinedInBlock, bool)`

GetMinedInBlockOk returns a tuple with the MinedInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlock

`func (o *AddressTokensTransactionConfirmedDataItem) SetMinedInBlock(v AddressTokensTransactionConfirmedDataItemMinedInBlock)`

SetMinedInBlock sets MinedInBlock field to given value.


### GetTransactionId

`func (o *AddressTokensTransactionConfirmedDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AddressTokensTransactionConfirmedDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTokenType

`func (o *AddressTokensTransactionConfirmedDataItem) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AddressTokensTransactionConfirmedDataItem) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetToken

`func (o *AddressTokensTransactionConfirmedDataItem) GetToken() AddressTokensTransactionConfirmedToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetTokenOk() (*AddressTokensTransactionConfirmedToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddressTokensTransactionConfirmedDataItem) SetToken(v AddressTokensTransactionConfirmedToken)`

SetToken sets Token field to given value.


### GetDirection

`func (o *AddressTokensTransactionConfirmedDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressTokensTransactionConfirmedDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressTokensTransactionConfirmedDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


