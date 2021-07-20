# ListUnconfirmedOmniTransactionsByAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the sent tokens. | 
**Divisible** | **bool** | Defines whether the attribute can be divisible or not, as boolean. E.g., if it is \&quot;true\&quot;, the attribute is divisible. | 
**Mined** | **bool** | Defines whether the transaction has been mined or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is mined. | 
**PropertyId** | **int32** | Represents the identifier of the tokens to send. | 
**Recipients** | [**[]ListOmniTransactionsByAddressRIRecipients**](ListOmniTransactionsByAddressRIRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]ListUnconfirmedOmniTransactionsByAddressRISenders**](ListUnconfirmedOmniTransactionsByAddressRISenders.md) | Represents an object of addresses that provide the funds. | 
**Sent** | **bool** | Defines whether the transaction has been sent or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is sent. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Type** | **string** | Defines the type of the transaction as a string. | 
**TypeInt** | **int32** | Defines the type of the transaction as a number. | 
**Version** | **int32** | Defines the specific version. | 
**Fee** | [**ListUnconfirmedOmniTransactionsByAddressRIFee**](ListUnconfirmedOmniTransactionsByAddressRIFee.md) |  | 

## Methods

### NewListUnconfirmedOmniTransactionsByAddressRI

`func NewListUnconfirmedOmniTransactionsByAddressRI(amount string, divisible bool, mined bool, propertyId int32, recipients []ListOmniTransactionsByAddressRIRecipients, senders []ListUnconfirmedOmniTransactionsByAddressRISenders, sent bool, timestamp int32, transactionId string, type_ string, typeInt int32, version int32, fee ListUnconfirmedOmniTransactionsByAddressRIFee, ) *ListUnconfirmedOmniTransactionsByAddressRI`

NewListUnconfirmedOmniTransactionsByAddressRI instantiates a new ListUnconfirmedOmniTransactionsByAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedOmniTransactionsByAddressRIWithDefaults

`func NewListUnconfirmedOmniTransactionsByAddressRIWithDefaults() *ListUnconfirmedOmniTransactionsByAddressRI`

NewListUnconfirmedOmniTransactionsByAddressRIWithDefaults instantiates a new ListUnconfirmedOmniTransactionsByAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDivisible

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetDivisible() bool`

GetDivisible returns the Divisible field if non-nil, zero value otherwise.

### GetDivisibleOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetDivisibleOk() (*bool, bool)`

GetDivisibleOk returns a tuple with the Divisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisible

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetDivisible(v bool)`

SetDivisible sets Divisible field to given value.


### GetMined

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetMined() bool`

GetMined returns the Mined field if non-nil, zero value otherwise.

### GetMinedOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetMinedOk() (*bool, bool)`

GetMinedOk returns a tuple with the Mined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMined

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetMined(v bool)`

SetMined sets Mined field to given value.


### GetPropertyId

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetRecipients

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetRecipients() []ListOmniTransactionsByAddressRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetRecipientsOk() (*[]ListOmniTransactionsByAddressRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetRecipients(v []ListOmniTransactionsByAddressRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetSenders() []ListUnconfirmedOmniTransactionsByAddressRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetSendersOk() (*[]ListUnconfirmedOmniTransactionsByAddressRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetSenders(v []ListUnconfirmedOmniTransactionsByAddressRISenders)`

SetSenders sets Senders field to given value.


### GetSent

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetSent(v bool)`

SetSent sets Sent field to given value.


### GetTimestamp

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetType

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetType(v string)`

SetType sets Type field to given value.


### GetTypeInt

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetTypeInt() int32`

GetTypeInt returns the TypeInt field if non-nil, zero value otherwise.

### GetTypeIntOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetTypeIntOk() (*int32, bool)`

GetTypeIntOk returns a tuple with the TypeInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInt

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetTypeInt(v int32)`

SetTypeInt sets TypeInt field to given value.


### GetVersion

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFee

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetFee() ListUnconfirmedOmniTransactionsByAddressRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) GetFeeOk() (*ListUnconfirmedOmniTransactionsByAddressRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListUnconfirmedOmniTransactionsByAddressRI) SetFee(v ListUnconfirmedOmniTransactionsByAddressRIFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


