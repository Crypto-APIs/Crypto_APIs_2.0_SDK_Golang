# GetUnconfirmedOmniTransactionByTransactionIDTxidRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the sent tokens. | 
**Divisible** | **bool** | Defines whether the attribute can be divisible or not, as boolean. E.g., if it is \&quot;true\&quot;, the attribute is divisible. | 
**Mined** | **bool** | Defines whether the transaction has been mined or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is mined. | 
**PropertyId** | **int32** | Represents the identifier of the tokens to send. | 
**Recipients** | [**[]GetUnconfirmedOmniTransactionByTransactionIDTxidRIRecipients**](GetUnconfirmedOmniTransactionByTransactionIDTxidRIRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]GetUnconfirmedOmniTransactionByTransactionIDTxidRISenders**](GetUnconfirmedOmniTransactionByTransactionIDTxidRISenders.md) | Represents an object of addresses that provide the funds. | 
**Sent** | **bool** | Defines whether the transaction has been sent or not, as boolean. E.g. if set to \&quot;true\&quot;, it means the transaction is sent. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | String representation of the transaction identifier (txid) | 
**Type** | **string** | Defines the type of the transaction as a string. | 
**TypeInt** | **int32** | Defines the type of the transaction as a number. | 
**Version** | **int32** | Defines the specific version. | 
**Fee** | [**ListUnconfirmedOmniTransactionsByAddressRIFee**](ListUnconfirmedOmniTransactionsByAddressRIFee.md) |  | 

## Methods

### NewGetUnconfirmedOmniTransactionByTransactionIDTxidRI

`func NewGetUnconfirmedOmniTransactionByTransactionIDTxidRI(amount string, divisible bool, mined bool, propertyId int32, recipients []GetUnconfirmedOmniTransactionByTransactionIDTxidRIRecipients, senders []GetUnconfirmedOmniTransactionByTransactionIDTxidRISenders, sent bool, timestamp int32, transactionId string, type_ string, typeInt int32, version int32, fee ListUnconfirmedOmniTransactionsByAddressRIFee, ) *GetUnconfirmedOmniTransactionByTransactionIDTxidRI`

NewGetUnconfirmedOmniTransactionByTransactionIDTxidRI instantiates a new GetUnconfirmedOmniTransactionByTransactionIDTxidRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUnconfirmedOmniTransactionByTransactionIDTxidRIWithDefaults

`func NewGetUnconfirmedOmniTransactionByTransactionIDTxidRIWithDefaults() *GetUnconfirmedOmniTransactionByTransactionIDTxidRI`

NewGetUnconfirmedOmniTransactionByTransactionIDTxidRIWithDefaults instantiates a new GetUnconfirmedOmniTransactionByTransactionIDTxidRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDivisible

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetDivisible() bool`

GetDivisible returns the Divisible field if non-nil, zero value otherwise.

### GetDivisibleOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetDivisibleOk() (*bool, bool)`

GetDivisibleOk returns a tuple with the Divisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisible

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetDivisible(v bool)`

SetDivisible sets Divisible field to given value.


### GetMined

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetMined() bool`

GetMined returns the Mined field if non-nil, zero value otherwise.

### GetMinedOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetMinedOk() (*bool, bool)`

GetMinedOk returns a tuple with the Mined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMined

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetMined(v bool)`

SetMined sets Mined field to given value.


### GetPropertyId

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetRecipients

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetRecipients() []GetUnconfirmedOmniTransactionByTransactionIDTxidRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetRecipientsOk() (*[]GetUnconfirmedOmniTransactionByTransactionIDTxidRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetRecipients(v []GetUnconfirmedOmniTransactionByTransactionIDTxidRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetSenders() []GetUnconfirmedOmniTransactionByTransactionIDTxidRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetSendersOk() (*[]GetUnconfirmedOmniTransactionByTransactionIDTxidRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetSenders(v []GetUnconfirmedOmniTransactionByTransactionIDTxidRISenders)`

SetSenders sets Senders field to given value.


### GetSent

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetSent(v bool)`

SetSent sets Sent field to given value.


### GetTimestamp

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetType

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetType(v string)`

SetType sets Type field to given value.


### GetTypeInt

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetTypeInt() int32`

GetTypeInt returns the TypeInt field if non-nil, zero value otherwise.

### GetTypeIntOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetTypeIntOk() (*int32, bool)`

GetTypeIntOk returns a tuple with the TypeInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInt

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetTypeInt(v int32)`

SetTypeInt sets TypeInt field to given value.


### GetVersion

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFee

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetFee() ListUnconfirmedOmniTransactionsByAddressRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) GetFeeOk() (*ListUnconfirmedOmniTransactionsByAddressRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidRI) SetFee(v ListUnconfirmedOmniTransactionsByAddressRIFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


