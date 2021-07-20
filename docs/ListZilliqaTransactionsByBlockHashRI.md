# ListZilliqaTransactionsByBlockHashRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | [**GetZilliqaTransactionDetailsByTransactionIDRIFee**](GetZilliqaTransactionDetailsByTransactionIDRIFee.md) |  | 
**GasLimit** | **int32** | Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit. | 
**GasPrice** | **int32** | Defines the price of the gas. | 
**GasUsed** | **int32** | Defines how much of the gas for the block has been used. | 
**MinedInBlockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**Nonce** | **int32** | Represents a random value that can be adjusted to satisfy the Proof of Work. | 
**Recipients** | [**[]ListZilliqaTransactionsByAddressRIRecipients**](ListZilliqaTransactionsByAddressRIRecipients.md) | Defines an object array of the transaction recipients. | 
**Senders** | [**[]ListZilliqaTransactionsByAddressRISenders**](ListZilliqaTransactionsByAddressRISenders.md) | Represents an object of addresses that provide the funds. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionHash** | **string** | Represents the hash of the transaction, which is its unique identifier. | 
**TransactionIndex** | **int32** | Defines the numeric representation of the transaction index. | 
**TransactionStatus** | **string** | Defines the status of the transaction, whether it is e.g. pending or complete. | 

## Methods

### NewListZilliqaTransactionsByBlockHashRI

`func NewListZilliqaTransactionsByBlockHashRI(fee GetZilliqaTransactionDetailsByTransactionIDRIFee, gasLimit int32, gasPrice int32, gasUsed int32, minedInBlockHeight int32, nonce int32, recipients []ListZilliqaTransactionsByAddressRIRecipients, senders []ListZilliqaTransactionsByAddressRISenders, timestamp int32, transactionHash string, transactionIndex int32, transactionStatus string, ) *ListZilliqaTransactionsByBlockHashRI`

NewListZilliqaTransactionsByBlockHashRI instantiates a new ListZilliqaTransactionsByBlockHashRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListZilliqaTransactionsByBlockHashRIWithDefaults

`func NewListZilliqaTransactionsByBlockHashRIWithDefaults() *ListZilliqaTransactionsByBlockHashRI`

NewListZilliqaTransactionsByBlockHashRIWithDefaults instantiates a new ListZilliqaTransactionsByBlockHashRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *ListZilliqaTransactionsByBlockHashRI) GetFee() GetZilliqaTransactionDetailsByTransactionIDRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetFeeOk() (*GetZilliqaTransactionDetailsByTransactionIDRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListZilliqaTransactionsByBlockHashRI) SetFee(v GetZilliqaTransactionDetailsByTransactionIDRIFee)`

SetFee sets Fee field to given value.


### GetGasLimit

`func (o *ListZilliqaTransactionsByBlockHashRI) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListZilliqaTransactionsByBlockHashRI) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListZilliqaTransactionsByBlockHashRI) GetGasPrice() int32`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetGasPriceOk() (*int32, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListZilliqaTransactionsByBlockHashRI) SetGasPrice(v int32)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListZilliqaTransactionsByBlockHashRI) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListZilliqaTransactionsByBlockHashRI) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInBlockHeight

`func (o *ListZilliqaTransactionsByBlockHashRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListZilliqaTransactionsByBlockHashRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetNonce

`func (o *ListZilliqaTransactionsByBlockHashRI) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListZilliqaTransactionsByBlockHashRI) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetRecipients

`func (o *ListZilliqaTransactionsByBlockHashRI) GetRecipients() []ListZilliqaTransactionsByAddressRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetRecipientsOk() (*[]ListZilliqaTransactionsByAddressRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListZilliqaTransactionsByBlockHashRI) SetRecipients(v []ListZilliqaTransactionsByAddressRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListZilliqaTransactionsByBlockHashRI) GetSenders() []ListZilliqaTransactionsByAddressRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetSendersOk() (*[]ListZilliqaTransactionsByAddressRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListZilliqaTransactionsByBlockHashRI) SetSenders(v []ListZilliqaTransactionsByAddressRISenders)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListZilliqaTransactionsByBlockHashRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListZilliqaTransactionsByBlockHashRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionIndex

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTransactionIndex() int32`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTransactionIndexOk() (*int32, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *ListZilliqaTransactionsByBlockHashRI) SetTransactionIndex(v int32)`

SetTransactionIndex sets TransactionIndex field to given value.


### GetTransactionStatus

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListZilliqaTransactionsByBlockHashRI) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListZilliqaTransactionsByBlockHashRI) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


