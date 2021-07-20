# GetOmniTransactionDetailsByTransactionIDTxidRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the sent tokens. | 
**Divisible** | **bool** | Defines whether the attribute can be divisible or not, as boolean. E.g., if it is \&quot;true\&quot;, the attribute is divisible. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**PropertyId** | **int32** | Represents the identifier of the tokens to send. | 
**Recipients** | [**[]ListOmniTransactionsByAddressRIRecipients**](ListOmniTransactionsByAddressRIRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]GetOmniTransactionDetailsByTransactionIDTxidRISenders**](GetOmniTransactionDetailsByTransactionIDTxidRISenders.md) | Represents an object of addresses that provide the funds. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Type** | **string** | Defines the type of the transaction as a string. | 
**TypeInt** | **int32** | Defines the type of the transaction as a number. | 
**Valid** | **bool** | Defines whether the transaction is valid or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is valid. | 
**Version** | **int32** | Defines the specific version. | 
**Fee** | [**ListUnconfirmedOmniTransactionsByAddressRIFee**](ListUnconfirmedOmniTransactionsByAddressRIFee.md) |  | 

## Methods

### NewGetOmniTransactionDetailsByTransactionIDTxidRI

`func NewGetOmniTransactionDetailsByTransactionIDTxidRI(amount string, divisible bool, minedInBlockHash string, minedInBlockHeight int32, propertyId int32, recipients []ListOmniTransactionsByAddressRIRecipients, senders []GetOmniTransactionDetailsByTransactionIDTxidRISenders, timestamp int32, transactionId string, type_ string, typeInt int32, valid bool, version int32, fee ListUnconfirmedOmniTransactionsByAddressRIFee, ) *GetOmniTransactionDetailsByTransactionIDTxidRI`

NewGetOmniTransactionDetailsByTransactionIDTxidRI instantiates a new GetOmniTransactionDetailsByTransactionIDTxidRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOmniTransactionDetailsByTransactionIDTxidRIWithDefaults

`func NewGetOmniTransactionDetailsByTransactionIDTxidRIWithDefaults() *GetOmniTransactionDetailsByTransactionIDTxidRI`

NewGetOmniTransactionDetailsByTransactionIDTxidRIWithDefaults instantiates a new GetOmniTransactionDetailsByTransactionIDTxidRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDivisible

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetDivisible() bool`

GetDivisible returns the Divisible field if non-nil, zero value otherwise.

### GetDivisibleOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetDivisibleOk() (*bool, bool)`

GetDivisibleOk returns a tuple with the Divisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisible

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetDivisible(v bool)`

SetDivisible sets Divisible field to given value.


### GetMinedInBlockHash

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetPropertyId

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetRecipients

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetRecipients() []ListOmniTransactionsByAddressRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetRecipientsOk() (*[]ListOmniTransactionsByAddressRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetRecipients(v []ListOmniTransactionsByAddressRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetSenders() []GetOmniTransactionDetailsByTransactionIDTxidRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetSendersOk() (*[]GetOmniTransactionDetailsByTransactionIDTxidRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetSenders(v []GetOmniTransactionDetailsByTransactionIDTxidRISenders)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetType

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetType(v string)`

SetType sets Type field to given value.


### GetTypeInt

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetTypeInt() int32`

GetTypeInt returns the TypeInt field if non-nil, zero value otherwise.

### GetTypeIntOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetTypeIntOk() (*int32, bool)`

GetTypeIntOk returns a tuple with the TypeInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInt

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetTypeInt(v int32)`

SetTypeInt sets TypeInt field to given value.


### GetValid

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetVersion

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFee

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetFee() ListUnconfirmedOmniTransactionsByAddressRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) GetFeeOk() (*ListUnconfirmedOmniTransactionsByAddressRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetOmniTransactionDetailsByTransactionIDTxidRI) SetFee(v ListUnconfirmedOmniTransactionsByAddressRIFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


