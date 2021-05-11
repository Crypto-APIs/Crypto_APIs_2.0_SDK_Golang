# AddressTokensTransactionUnconfirmedDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot;, \&quot;rinkeby\&quot; are test networks. | 
**Address** | **string** | Defines the specific address to which the token transaction has been sent and is pending confirmation. | 
**TransactionId** | **string** | Defines the unique ID of the specific transaction, i.e. its identification number. | 
**TokenType** | **string** | Defines the type of token sent with the transaction, e.g. ERC 20. | 
**Token** | [**AddressTokensTransactionUnconfirmedToken**](AddressTokensTransactionUnconfirmedToken.md) |  | 
**Direction** | **string** | Defines whether the transaction is \&quot;incoming\&quot; or \&quot;outgoing\&quot;. | 
**FirstSeenInMempoolTimestamp** | **int32** | Defines the exact time the transaction has been first accepted into the mempool to await confirmation as timestamp. | 

## Methods

### NewAddressTokensTransactionUnconfirmedDataItem

`func NewAddressTokensTransactionUnconfirmedDataItem(blockchain string, network string, address string, transactionId string, tokenType string, token AddressTokensTransactionUnconfirmedToken, direction string, firstSeenInMempoolTimestamp int32, ) *AddressTokensTransactionUnconfirmedDataItem`

NewAddressTokensTransactionUnconfirmedDataItem instantiates a new AddressTokensTransactionUnconfirmedDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTokensTransactionUnconfirmedDataItemWithDefaults

`func NewAddressTokensTransactionUnconfirmedDataItemWithDefaults() *AddressTokensTransactionUnconfirmedDataItem`

NewAddressTokensTransactionUnconfirmedDataItemWithDefaults instantiates a new AddressTokensTransactionUnconfirmedDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTransactionId

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTokenType

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetToken

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetToken() AddressTokensTransactionUnconfirmedToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetTokenOk() (*AddressTokensTransactionUnconfirmedToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetToken(v AddressTokensTransactionUnconfirmedToken)`

SetToken sets Token field to given value.


### GetDirection

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetFirstSeenInMempoolTimestamp

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetFirstSeenInMempoolTimestamp() int32`

GetFirstSeenInMempoolTimestamp returns the FirstSeenInMempoolTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenInMempoolTimestampOk

`func (o *AddressTokensTransactionUnconfirmedDataItem) GetFirstSeenInMempoolTimestampOk() (*int32, bool)`

GetFirstSeenInMempoolTimestampOk returns a tuple with the FirstSeenInMempoolTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenInMempoolTimestamp

`func (o *AddressTokensTransactionUnconfirmedDataItem) SetFirstSeenInMempoolTimestamp(v int32)`

SetFirstSeenInMempoolTimestamp sets FirstSeenInMempoolTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


