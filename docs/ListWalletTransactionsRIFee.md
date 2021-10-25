# ListWalletTransactionsRIFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the fee for the transaction. | 
**ConvertedAmount** | **string** | Defines the converted amount of the transaction&#39;s fee. | 
**ExchangeRateUnit** | **string** | Defines the exchange rate for the transaction&#39;s unit. | 
**Symbol** | **string** | Defines the unit of the transaction&#39;s fee. | 

## Methods

### NewListWalletTransactionsRIFee

`func NewListWalletTransactionsRIFee(amount string, convertedAmount string, exchangeRateUnit string, symbol string, ) *ListWalletTransactionsRIFee`

NewListWalletTransactionsRIFee instantiates a new ListWalletTransactionsRIFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIFeeWithDefaults

`func NewListWalletTransactionsRIFeeWithDefaults() *ListWalletTransactionsRIFee`

NewListWalletTransactionsRIFeeWithDefaults instantiates a new ListWalletTransactionsRIFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListWalletTransactionsRIFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRIFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRIFee) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetConvertedAmount

`func (o *ListWalletTransactionsRIFee) GetConvertedAmount() string`

GetConvertedAmount returns the ConvertedAmount field if non-nil, zero value otherwise.

### GetConvertedAmountOk

`func (o *ListWalletTransactionsRIFee) GetConvertedAmountOk() (*string, bool)`

GetConvertedAmountOk returns a tuple with the ConvertedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAmount

`func (o *ListWalletTransactionsRIFee) SetConvertedAmount(v string)`

SetConvertedAmount sets ConvertedAmount field to given value.


### GetExchangeRateUnit

`func (o *ListWalletTransactionsRIFee) GetExchangeRateUnit() string`

GetExchangeRateUnit returns the ExchangeRateUnit field if non-nil, zero value otherwise.

### GetExchangeRateUnitOk

`func (o *ListWalletTransactionsRIFee) GetExchangeRateUnitOk() (*string, bool)`

GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUnit

`func (o *ListWalletTransactionsRIFee) SetExchangeRateUnit(v string)`

SetExchangeRateUnit sets ExchangeRateUnit field to given value.


### GetSymbol

`func (o *ListWalletTransactionsRIFee) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListWalletTransactionsRIFee) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListWalletTransactionsRIFee) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


