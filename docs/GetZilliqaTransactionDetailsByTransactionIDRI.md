# GetZilliqaTransactionDetailsByTransactionIDRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | [**GetZilliqaTransactionDetailsByTransactionIDRIFee**](GetZilliqaTransactionDetailsByTransactionIDRIFee.md) |  | 
**GasLimit** | **int32** | Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit. | 
**GasPrice** | **int32** | Defines the price of the gas. | 
**GasUsed** | **int32** | Defines how much of the gas for the block has been used. | 
**MinedInBlockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**Recipients** | [**[]GetZilliqaTransactionDetailsByTransactionIDRIRecipientsInner**](GetZilliqaTransactionDetailsByTransactionIDRIRecipientsInner.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]GetZilliqaTransactionDetailsByTransactionIDRISendersInner**](GetZilliqaTransactionDetailsByTransactionIDRISendersInner.md) | Represents an object of addresses that provide the funds. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionIndex** | **int32** | Defines the numeric representation of the transaction index. | 
**TransactionStatus** | **string** | Defines the status of the transaction, whether it is e.g. pending or complete. | 

## Methods

### NewGetZilliqaTransactionDetailsByTransactionIDRI

`func NewGetZilliqaTransactionDetailsByTransactionIDRI(fee GetZilliqaTransactionDetailsByTransactionIDRIFee, gasLimit int32, gasPrice int32, gasUsed int32, minedInBlockHash string, minedInBlockHeight int32, nonce int32, recipients []GetZilliqaTransactionDetailsByTransactionIDRIRecipientsInner, senders []GetZilliqaTransactionDetailsByTransactionIDRISendersInner, timestamp int32, transactionIndex int32, transactionStatus string, ) *GetZilliqaTransactionDetailsByTransactionIDRI`

NewGetZilliqaTransactionDetailsByTransactionIDRI instantiates a new GetZilliqaTransactionDetailsByTransactionIDRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZilliqaTransactionDetailsByTransactionIDRIWithDefaults

`func NewGetZilliqaTransactionDetailsByTransactionIDRIWithDefaults() *GetZilliqaTransactionDetailsByTransactionIDRI`

NewGetZilliqaTransactionDetailsByTransactionIDRIWithDefaults instantiates a new GetZilliqaTransactionDetailsByTransactionIDRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetFee() GetZilliqaTransactionDetailsByTransactionIDRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetFeeOk() (*GetZilliqaTransactionDetailsByTransactionIDRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetFee(v GetZilliqaTransactionDetailsByTransactionIDRIFee)`

SetFee sets Fee field to given value.


### GetGasLimit

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetGasPrice() int32`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetGasPriceOk() (*int32, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetGasPrice(v int32)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInBlockHash

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetNonce

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetRecipients

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetRecipients() []GetZilliqaTransactionDetailsByTransactionIDRIRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetRecipientsOk() (*[]GetZilliqaTransactionDetailsByTransactionIDRIRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetRecipients(v []GetZilliqaTransactionDetailsByTransactionIDRIRecipientsInner)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetSenders() []GetZilliqaTransactionDetailsByTransactionIDRISendersInner`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetSendersOk() (*[]GetZilliqaTransactionDetailsByTransactionIDRISendersInner, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetSenders(v []GetZilliqaTransactionDetailsByTransactionIDRISendersInner)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionIndex

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetTransactionIndex() int32`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetTransactionIndexOk() (*int32, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetTransactionIndex(v int32)`

SetTransactionIndex sets TransactionIndex field to given value.


### GetTransactionStatus

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetZilliqaTransactionDetailsByTransactionIDRI) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


