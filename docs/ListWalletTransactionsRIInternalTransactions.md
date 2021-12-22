# ListWalletTransactionsRIInternalTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the specific amount of the transaction. | 
**ConvertedAmount** | **string** | Represents the converted amount. | 
**ExchangeRateUnit** | **string** | Defines the base asset symbol to get a rate for. | 
**OperationId** | **string** | Represents the unique internal transaction ID in regards to the parent transaction (type trace address). | 
**Recipient** | **string** | Represents the recipient address with the respective amount. | 
**Sender** | **string** | Represents the sender address with the respective amount. | 
**Symbol** | **string** | Represents the unique unit symbol. | 

## Methods

### NewListWalletTransactionsRIInternalTransactions

`func NewListWalletTransactionsRIInternalTransactions(amount string, convertedAmount string, exchangeRateUnit string, operationId string, recipient string, sender string, symbol string, ) *ListWalletTransactionsRIInternalTransactions`

NewListWalletTransactionsRIInternalTransactions instantiates a new ListWalletTransactionsRIInternalTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIInternalTransactionsWithDefaults

`func NewListWalletTransactionsRIInternalTransactionsWithDefaults() *ListWalletTransactionsRIInternalTransactions`

NewListWalletTransactionsRIInternalTransactionsWithDefaults instantiates a new ListWalletTransactionsRIInternalTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListWalletTransactionsRIInternalTransactions) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRIInternalTransactions) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRIInternalTransactions) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetConvertedAmount

`func (o *ListWalletTransactionsRIInternalTransactions) GetConvertedAmount() string`

GetConvertedAmount returns the ConvertedAmount field if non-nil, zero value otherwise.

### GetConvertedAmountOk

`func (o *ListWalletTransactionsRIInternalTransactions) GetConvertedAmountOk() (*string, bool)`

GetConvertedAmountOk returns a tuple with the ConvertedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAmount

`func (o *ListWalletTransactionsRIInternalTransactions) SetConvertedAmount(v string)`

SetConvertedAmount sets ConvertedAmount field to given value.


### GetExchangeRateUnit

`func (o *ListWalletTransactionsRIInternalTransactions) GetExchangeRateUnit() string`

GetExchangeRateUnit returns the ExchangeRateUnit field if non-nil, zero value otherwise.

### GetExchangeRateUnitOk

`func (o *ListWalletTransactionsRIInternalTransactions) GetExchangeRateUnitOk() (*string, bool)`

GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUnit

`func (o *ListWalletTransactionsRIInternalTransactions) SetExchangeRateUnit(v string)`

SetExchangeRateUnit sets ExchangeRateUnit field to given value.


### GetOperationId

`func (o *ListWalletTransactionsRIInternalTransactions) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *ListWalletTransactionsRIInternalTransactions) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *ListWalletTransactionsRIInternalTransactions) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetRecipient

`func (o *ListWalletTransactionsRIInternalTransactions) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ListWalletTransactionsRIInternalTransactions) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ListWalletTransactionsRIInternalTransactions) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *ListWalletTransactionsRIInternalTransactions) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListWalletTransactionsRIInternalTransactions) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListWalletTransactionsRIInternalTransactions) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSymbol

`func (o *ListWalletTransactionsRIInternalTransactions) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListWalletTransactionsRIInternalTransactions) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListWalletTransactionsRIInternalTransactions) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


