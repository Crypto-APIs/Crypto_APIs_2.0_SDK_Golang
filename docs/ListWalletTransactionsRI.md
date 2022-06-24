# ListWalletTransactionsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | Defines the direction of the transaction, e.g. incoming. | 
**Fee** | [**ListWalletTransactionsRIFee**](ListWalletTransactionsRIFee.md) |  | 
**FungibleTokens** | Pointer to [**[]ListWalletTransactionsRIFungibleTokensInner**](ListWalletTransactionsRIFungibleTokensInner.md) | Represents fungible tokens&#39;es detailed information | [optional] 
**InternalTransactions** | Pointer to [**[]ListWalletTransactionsRIInternalTransactionsInner**](ListWalletTransactionsRIInternalTransactionsInner.md) |  | [optional] 
**NonFungibleTokens** | Pointer to [**[]ListWalletTransactionsRINonFungibleTokensInner**](ListWalletTransactionsRINonFungibleTokensInner.md) | Represents non-fungible tokens&#39;es detailed information. | [optional] 
**Recipients** | [**[]ListWalletTransactionsRIRecipientsInner**](ListWalletTransactionsRIRecipientsInner.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | [**[]ListWalletTransactionsRISendersInner**](ListWalletTransactionsRISendersInner.md) | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Status** | **string** | Defines the status of the transaction, if it is confirmed or unconfirmed. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | Represents the unique TD of the transaction. | 
**Value** | [**ListWalletTransactionsRIValue**](ListWalletTransactionsRIValue.md) |  | 

## Methods

### NewListWalletTransactionsRI

`func NewListWalletTransactionsRI(direction string, fee ListWalletTransactionsRIFee, recipients []ListWalletTransactionsRIRecipientsInner, senders []ListWalletTransactionsRISendersInner, status string, timestamp int32, transactionId string, value ListWalletTransactionsRIValue, ) *ListWalletTransactionsRI`

NewListWalletTransactionsRI instantiates a new ListWalletTransactionsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIWithDefaults

`func NewListWalletTransactionsRIWithDefaults() *ListWalletTransactionsRI`

NewListWalletTransactionsRIWithDefaults instantiates a new ListWalletTransactionsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *ListWalletTransactionsRI) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ListWalletTransactionsRI) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ListWalletTransactionsRI) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetFee

`func (o *ListWalletTransactionsRI) GetFee() ListWalletTransactionsRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListWalletTransactionsRI) GetFeeOk() (*ListWalletTransactionsRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListWalletTransactionsRI) SetFee(v ListWalletTransactionsRIFee)`

SetFee sets Fee field to given value.


### GetFungibleTokens

`func (o *ListWalletTransactionsRI) GetFungibleTokens() []ListWalletTransactionsRIFungibleTokensInner`

GetFungibleTokens returns the FungibleTokens field if non-nil, zero value otherwise.

### GetFungibleTokensOk

`func (o *ListWalletTransactionsRI) GetFungibleTokensOk() (*[]ListWalletTransactionsRIFungibleTokensInner, bool)`

GetFungibleTokensOk returns a tuple with the FungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFungibleTokens

`func (o *ListWalletTransactionsRI) SetFungibleTokens(v []ListWalletTransactionsRIFungibleTokensInner)`

SetFungibleTokens sets FungibleTokens field to given value.

### HasFungibleTokens

`func (o *ListWalletTransactionsRI) HasFungibleTokens() bool`

HasFungibleTokens returns a boolean if a field has been set.

### GetInternalTransactions

`func (o *ListWalletTransactionsRI) GetInternalTransactions() []ListWalletTransactionsRIInternalTransactionsInner`

GetInternalTransactions returns the InternalTransactions field if non-nil, zero value otherwise.

### GetInternalTransactionsOk

`func (o *ListWalletTransactionsRI) GetInternalTransactionsOk() (*[]ListWalletTransactionsRIInternalTransactionsInner, bool)`

GetInternalTransactionsOk returns a tuple with the InternalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransactions

`func (o *ListWalletTransactionsRI) SetInternalTransactions(v []ListWalletTransactionsRIInternalTransactionsInner)`

SetInternalTransactions sets InternalTransactions field to given value.

### HasInternalTransactions

`func (o *ListWalletTransactionsRI) HasInternalTransactions() bool`

HasInternalTransactions returns a boolean if a field has been set.

### GetNonFungibleTokens

`func (o *ListWalletTransactionsRI) GetNonFungibleTokens() []ListWalletTransactionsRINonFungibleTokensInner`

GetNonFungibleTokens returns the NonFungibleTokens field if non-nil, zero value otherwise.

### GetNonFungibleTokensOk

`func (o *ListWalletTransactionsRI) GetNonFungibleTokensOk() (*[]ListWalletTransactionsRINonFungibleTokensInner, bool)`

GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFungibleTokens

`func (o *ListWalletTransactionsRI) SetNonFungibleTokens(v []ListWalletTransactionsRINonFungibleTokensInner)`

SetNonFungibleTokens sets NonFungibleTokens field to given value.

### HasNonFungibleTokens

`func (o *ListWalletTransactionsRI) HasNonFungibleTokens() bool`

HasNonFungibleTokens returns a boolean if a field has been set.

### GetRecipients

`func (o *ListWalletTransactionsRI) GetRecipients() []ListWalletTransactionsRIRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListWalletTransactionsRI) GetRecipientsOk() (*[]ListWalletTransactionsRIRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListWalletTransactionsRI) SetRecipients(v []ListWalletTransactionsRIRecipientsInner)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListWalletTransactionsRI) GetSenders() []ListWalletTransactionsRISendersInner`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListWalletTransactionsRI) GetSendersOk() (*[]ListWalletTransactionsRISendersInner, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListWalletTransactionsRI) SetSenders(v []ListWalletTransactionsRISendersInner)`

SetSenders sets Senders field to given value.


### GetStatus

`func (o *ListWalletTransactionsRI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListWalletTransactionsRI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListWalletTransactionsRI) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ListWalletTransactionsRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListWalletTransactionsRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListWalletTransactionsRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *ListWalletTransactionsRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListWalletTransactionsRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListWalletTransactionsRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetValue

`func (o *ListWalletTransactionsRI) GetValue() ListWalletTransactionsRIValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListWalletTransactionsRI) GetValueOk() (*ListWalletTransactionsRIValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListWalletTransactionsRI) SetValue(v ListWalletTransactionsRIValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


