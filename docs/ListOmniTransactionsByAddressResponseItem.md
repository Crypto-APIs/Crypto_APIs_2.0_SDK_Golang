# ListOmniTransactionsByAddressResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the sent tokens. | 
**Divisible** | **bool** | Defines whether the attribute can be divisible or not, as boolean. E.g., if it is \&quot;true\&quot;, the attribute is divisible. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**PropertyId** | **int32** | Represents the identifier of the tokens to send. | 
**Recipients** | [**[]ListOmniTransactionsByAddressResponseItemRecipients**](ListOmniTransactionsByAddressResponseItemRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]ListOmniTransactionsByAddressResponseItemSenders**](ListOmniTransactionsByAddressResponseItemSenders.md) | Represents an object of addresses that provide the funds. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Type** | **string** | Defines the type of the transaction as a string. | 
**TypeInt** | **int32** | Defines the type of the transaction as a number. | 
**Valid** | **bool** | Defines whether the transaction is valid or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is valid. | 
**Version** | **int32** | Defines the specific version. | 
**Fee** | [**ListOmniTransactionsByAddressResponseItemFee**](ListOmniTransactionsByAddressResponseItemFee.md) |  | 

## Methods

### NewListOmniTransactionsByAddressResponseItem

`func NewListOmniTransactionsByAddressResponseItem(amount string, divisible bool, minedInBlockHash string, minedInBlockHeight int32, propertyId int32, recipients []ListOmniTransactionsByAddressResponseItemRecipients, senders []ListOmniTransactionsByAddressResponseItemSenders, timestamp int32, transactionId string, type_ string, typeInt int32, valid bool, version int32, fee ListOmniTransactionsByAddressResponseItemFee, ) *ListOmniTransactionsByAddressResponseItem`

NewListOmniTransactionsByAddressResponseItem instantiates a new ListOmniTransactionsByAddressResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOmniTransactionsByAddressResponseItemWithDefaults

`func NewListOmniTransactionsByAddressResponseItemWithDefaults() *ListOmniTransactionsByAddressResponseItem`

NewListOmniTransactionsByAddressResponseItemWithDefaults instantiates a new ListOmniTransactionsByAddressResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListOmniTransactionsByAddressResponseItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListOmniTransactionsByAddressResponseItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDivisible

`func (o *ListOmniTransactionsByAddressResponseItem) GetDivisible() bool`

GetDivisible returns the Divisible field if non-nil, zero value otherwise.

### GetDivisibleOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetDivisibleOk() (*bool, bool)`

GetDivisibleOk returns a tuple with the Divisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisible

`func (o *ListOmniTransactionsByAddressResponseItem) SetDivisible(v bool)`

SetDivisible sets Divisible field to given value.


### GetMinedInBlockHash

`func (o *ListOmniTransactionsByAddressResponseItem) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListOmniTransactionsByAddressResponseItem) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *ListOmniTransactionsByAddressResponseItem) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListOmniTransactionsByAddressResponseItem) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetPropertyId

`func (o *ListOmniTransactionsByAddressResponseItem) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ListOmniTransactionsByAddressResponseItem) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetRecipients

`func (o *ListOmniTransactionsByAddressResponseItem) GetRecipients() []ListOmniTransactionsByAddressResponseItemRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetRecipientsOk() (*[]ListOmniTransactionsByAddressResponseItemRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListOmniTransactionsByAddressResponseItem) SetRecipients(v []ListOmniTransactionsByAddressResponseItemRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListOmniTransactionsByAddressResponseItem) GetSenders() []ListOmniTransactionsByAddressResponseItemSenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetSendersOk() (*[]ListOmniTransactionsByAddressResponseItemSenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListOmniTransactionsByAddressResponseItem) SetSenders(v []ListOmniTransactionsByAddressResponseItemSenders)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *ListOmniTransactionsByAddressResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListOmniTransactionsByAddressResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *ListOmniTransactionsByAddressResponseItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListOmniTransactionsByAddressResponseItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetType

`func (o *ListOmniTransactionsByAddressResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListOmniTransactionsByAddressResponseItem) SetType(v string)`

SetType sets Type field to given value.


### GetTypeInt

`func (o *ListOmniTransactionsByAddressResponseItem) GetTypeInt() int32`

GetTypeInt returns the TypeInt field if non-nil, zero value otherwise.

### GetTypeIntOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetTypeIntOk() (*int32, bool)`

GetTypeIntOk returns a tuple with the TypeInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInt

`func (o *ListOmniTransactionsByAddressResponseItem) SetTypeInt(v int32)`

SetTypeInt sets TypeInt field to given value.


### GetValid

`func (o *ListOmniTransactionsByAddressResponseItem) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ListOmniTransactionsByAddressResponseItem) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetVersion

`func (o *ListOmniTransactionsByAddressResponseItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListOmniTransactionsByAddressResponseItem) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFee

`func (o *ListOmniTransactionsByAddressResponseItem) GetFee() ListOmniTransactionsByAddressResponseItemFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListOmniTransactionsByAddressResponseItem) GetFeeOk() (*ListOmniTransactionsByAddressResponseItemFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListOmniTransactionsByAddressResponseItem) SetFee(v ListOmniTransactionsByAddressResponseItemFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


