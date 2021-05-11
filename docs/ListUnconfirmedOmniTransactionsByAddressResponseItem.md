# ListUnconfirmedOmniTransactionsByAddressResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the sent tokens. | 
**Divisible** | **bool** | Defines whether the attribute can be divisible or not, as boolean. E.g., if it is \&quot;true\&quot;, the attribute is divisible. | 
**Mined** | **bool** | Defines whether the transaction has been mined or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is mined. | 
**PropertyId** | **int32** | Represents the identifier of the tokens to send. | 
**Recipients** | [**[]ListOmniTransactionsByAddressResponseItemRecipients**](ListOmniTransactionsByAddressResponseItemRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]ListUnconfirmedOmniTransactionsByAddressResponseItemSenders**](ListUnconfirmedOmniTransactionsByAddressResponseItemSenders.md) | Represents an object of addresses that provide the funds. | 
**Sent** | **bool** | Defines whether the transaction has been sent or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is sent. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Type** | **string** | Defines the type of the transaction as a string. | 
**TypeInt** | **int32** | Defines the type of the transaction as a number. | 
**Version** | **int32** | Defines the specific version. | 
**Fee** | [**ListUnconfirmedOmniTransactionsByAddressResponseItemFee**](ListUnconfirmedOmniTransactionsByAddressResponseItemFee.md) |  | 

## Methods

### NewListUnconfirmedOmniTransactionsByAddressResponseItem

`func NewListUnconfirmedOmniTransactionsByAddressResponseItem(amount string, divisible bool, mined bool, propertyId int32, recipients []ListOmniTransactionsByAddressResponseItemRecipients, senders []ListUnconfirmedOmniTransactionsByAddressResponseItemSenders, sent bool, timestamp int32, transactionId string, type_ string, typeInt int32, version int32, fee ListUnconfirmedOmniTransactionsByAddressResponseItemFee, ) *ListUnconfirmedOmniTransactionsByAddressResponseItem`

NewListUnconfirmedOmniTransactionsByAddressResponseItem instantiates a new ListUnconfirmedOmniTransactionsByAddressResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedOmniTransactionsByAddressResponseItemWithDefaults

`func NewListUnconfirmedOmniTransactionsByAddressResponseItemWithDefaults() *ListUnconfirmedOmniTransactionsByAddressResponseItem`

NewListUnconfirmedOmniTransactionsByAddressResponseItemWithDefaults instantiates a new ListUnconfirmedOmniTransactionsByAddressResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDivisible

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetDivisible() bool`

GetDivisible returns the Divisible field if non-nil, zero value otherwise.

### GetDivisibleOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetDivisibleOk() (*bool, bool)`

GetDivisibleOk returns a tuple with the Divisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisible

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetDivisible(v bool)`

SetDivisible sets Divisible field to given value.


### GetMined

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetMined() bool`

GetMined returns the Mined field if non-nil, zero value otherwise.

### GetMinedOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetMinedOk() (*bool, bool)`

GetMinedOk returns a tuple with the Mined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMined

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetMined(v bool)`

SetMined sets Mined field to given value.


### GetPropertyId

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetRecipients

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetRecipients() []ListOmniTransactionsByAddressResponseItemRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetRecipientsOk() (*[]ListOmniTransactionsByAddressResponseItemRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetRecipients(v []ListOmniTransactionsByAddressResponseItemRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetSenders() []ListUnconfirmedOmniTransactionsByAddressResponseItemSenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetSendersOk() (*[]ListUnconfirmedOmniTransactionsByAddressResponseItemSenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetSenders(v []ListUnconfirmedOmniTransactionsByAddressResponseItemSenders)`

SetSenders sets Senders field to given value.


### GetSent

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetSent(v bool)`

SetSent sets Sent field to given value.


### GetTimestamp

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetType

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetType(v string)`

SetType sets Type field to given value.


### GetTypeInt

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetTypeInt() int32`

GetTypeInt returns the TypeInt field if non-nil, zero value otherwise.

### GetTypeIntOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetTypeIntOk() (*int32, bool)`

GetTypeIntOk returns a tuple with the TypeInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInt

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetTypeInt(v int32)`

SetTypeInt sets TypeInt field to given value.


### GetVersion

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFee

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetFee() ListUnconfirmedOmniTransactionsByAddressResponseItemFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) GetFeeOk() (*ListUnconfirmedOmniTransactionsByAddressResponseItemFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListUnconfirmedOmniTransactionsByAddressResponseItem) SetFee(v ListUnconfirmedOmniTransactionsByAddressResponseItemFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


