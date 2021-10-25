# ListWalletTransactionsRIValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the transaction amount. | 
**ConvertedAmount** | **string** | Defines the converted amount of the transaction as a string. | 
**ExchangeRateUnit** | **string** | Defines the exchange rate&#39;s unit. | 
**Symbol** | **string** | Defines the unit of the transaction&#39;s amount. | 
**TokenIdentifier** | Pointer to **string** | Defines the token&#39;s identifier of the transaction&#39;s amount. | [optional] 

## Methods

### NewListWalletTransactionsRIValue

`func NewListWalletTransactionsRIValue(amount string, convertedAmount string, exchangeRateUnit string, symbol string, ) *ListWalletTransactionsRIValue`

NewListWalletTransactionsRIValue instantiates a new ListWalletTransactionsRIValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIValueWithDefaults

`func NewListWalletTransactionsRIValueWithDefaults() *ListWalletTransactionsRIValue`

NewListWalletTransactionsRIValueWithDefaults instantiates a new ListWalletTransactionsRIValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListWalletTransactionsRIValue) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRIValue) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRIValue) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetConvertedAmount

`func (o *ListWalletTransactionsRIValue) GetConvertedAmount() string`

GetConvertedAmount returns the ConvertedAmount field if non-nil, zero value otherwise.

### GetConvertedAmountOk

`func (o *ListWalletTransactionsRIValue) GetConvertedAmountOk() (*string, bool)`

GetConvertedAmountOk returns a tuple with the ConvertedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAmount

`func (o *ListWalletTransactionsRIValue) SetConvertedAmount(v string)`

SetConvertedAmount sets ConvertedAmount field to given value.


### GetExchangeRateUnit

`func (o *ListWalletTransactionsRIValue) GetExchangeRateUnit() string`

GetExchangeRateUnit returns the ExchangeRateUnit field if non-nil, zero value otherwise.

### GetExchangeRateUnitOk

`func (o *ListWalletTransactionsRIValue) GetExchangeRateUnitOk() (*string, bool)`

GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUnit

`func (o *ListWalletTransactionsRIValue) SetExchangeRateUnit(v string)`

SetExchangeRateUnit sets ExchangeRateUnit field to given value.


### GetSymbol

`func (o *ListWalletTransactionsRIValue) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListWalletTransactionsRIValue) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListWalletTransactionsRIValue) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenIdentifier

`func (o *ListWalletTransactionsRIValue) GetTokenIdentifier() string`

GetTokenIdentifier returns the TokenIdentifier field if non-nil, zero value otherwise.

### GetTokenIdentifierOk

`func (o *ListWalletTransactionsRIValue) GetTokenIdentifierOk() (*string, bool)`

GetTokenIdentifierOk returns a tuple with the TokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdentifier

`func (o *ListWalletTransactionsRIValue) SetTokenIdentifier(v string)`

SetTokenIdentifier sets TokenIdentifier field to given value.

### HasTokenIdentifier

`func (o *ListWalletTransactionsRIValue) HasTokenIdentifier() bool`

HasTokenIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


