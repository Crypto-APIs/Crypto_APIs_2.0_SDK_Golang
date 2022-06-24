# ListWalletTransactionsRIInternalTransactionsInner

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

### NewListWalletTransactionsRIInternalTransactionsInner

`func NewListWalletTransactionsRIInternalTransactionsInner(amount string, convertedAmount string, exchangeRateUnit string, operationId string, recipient string, sender string, symbol string, ) *ListWalletTransactionsRIInternalTransactionsInner`

NewListWalletTransactionsRIInternalTransactionsInner instantiates a new ListWalletTransactionsRIInternalTransactionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIInternalTransactionsInnerWithDefaults

`func NewListWalletTransactionsRIInternalTransactionsInnerWithDefaults() *ListWalletTransactionsRIInternalTransactionsInner`

NewListWalletTransactionsRIInternalTransactionsInnerWithDefaults instantiates a new ListWalletTransactionsRIInternalTransactionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRIInternalTransactionsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetConvertedAmount

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetConvertedAmount() string`

GetConvertedAmount returns the ConvertedAmount field if non-nil, zero value otherwise.

### GetConvertedAmountOk

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetConvertedAmountOk() (*string, bool)`

GetConvertedAmountOk returns a tuple with the ConvertedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAmount

`func (o *ListWalletTransactionsRIInternalTransactionsInner) SetConvertedAmount(v string)`

SetConvertedAmount sets ConvertedAmount field to given value.


### GetExchangeRateUnit

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetExchangeRateUnit() string`

GetExchangeRateUnit returns the ExchangeRateUnit field if non-nil, zero value otherwise.

### GetExchangeRateUnitOk

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetExchangeRateUnitOk() (*string, bool)`

GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUnit

`func (o *ListWalletTransactionsRIInternalTransactionsInner) SetExchangeRateUnit(v string)`

SetExchangeRateUnit sets ExchangeRateUnit field to given value.


### GetOperationId

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *ListWalletTransactionsRIInternalTransactionsInner) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetRecipient

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ListWalletTransactionsRIInternalTransactionsInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListWalletTransactionsRIInternalTransactionsInner) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSymbol

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListWalletTransactionsRIInternalTransactionsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListWalletTransactionsRIInternalTransactionsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


