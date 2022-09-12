# GetExchangeRateByAssetSymbolsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculationTimestamp** | **int32** | Defines the time of the market data used to calculate the exchange rate in UNIX Timestamp. Oldest possible timestamp is 30 days. | 
**FromAssetId** | **string** | Defines the base asset Reference ID to get a rate for. | 
**FromAssetSymbol** | **string** | Defines the base asset symbol to get a rate for. | 
**Rate** | **string** | Defines the exchange rate between assets calculated by weighted average of the last trades in every exchange for the last 24 hours by giving more weight to exchanges with higher volume. | 
**ToAssetId** | **string** | Defines the relation asset Reference ID in which the base asset rate will be displayed. | 
**ToAssetSymbol** | **string** | Defines the relation asset symbol in which the base asset rate will be displayed. | 

## Methods

### NewGetExchangeRateByAssetSymbolsRI

`func NewGetExchangeRateByAssetSymbolsRI(calculationTimestamp int32, fromAssetId string, fromAssetSymbol string, rate string, toAssetId string, toAssetSymbol string, ) *GetExchangeRateByAssetSymbolsRI`

NewGetExchangeRateByAssetSymbolsRI instantiates a new GetExchangeRateByAssetSymbolsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeRateByAssetSymbolsRIWithDefaults

`func NewGetExchangeRateByAssetSymbolsRIWithDefaults() *GetExchangeRateByAssetSymbolsRI`

NewGetExchangeRateByAssetSymbolsRIWithDefaults instantiates a new GetExchangeRateByAssetSymbolsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculationTimestamp

`func (o *GetExchangeRateByAssetSymbolsRI) GetCalculationTimestamp() int32`

GetCalculationTimestamp returns the CalculationTimestamp field if non-nil, zero value otherwise.

### GetCalculationTimestampOk

`func (o *GetExchangeRateByAssetSymbolsRI) GetCalculationTimestampOk() (*int32, bool)`

GetCalculationTimestampOk returns a tuple with the CalculationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationTimestamp

`func (o *GetExchangeRateByAssetSymbolsRI) SetCalculationTimestamp(v int32)`

SetCalculationTimestamp sets CalculationTimestamp field to given value.


### GetFromAssetId

`func (o *GetExchangeRateByAssetSymbolsRI) GetFromAssetId() string`

GetFromAssetId returns the FromAssetId field if non-nil, zero value otherwise.

### GetFromAssetIdOk

`func (o *GetExchangeRateByAssetSymbolsRI) GetFromAssetIdOk() (*string, bool)`

GetFromAssetIdOk returns a tuple with the FromAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAssetId

`func (o *GetExchangeRateByAssetSymbolsRI) SetFromAssetId(v string)`

SetFromAssetId sets FromAssetId field to given value.


### GetFromAssetSymbol

`func (o *GetExchangeRateByAssetSymbolsRI) GetFromAssetSymbol() string`

GetFromAssetSymbol returns the FromAssetSymbol field if non-nil, zero value otherwise.

### GetFromAssetSymbolOk

`func (o *GetExchangeRateByAssetSymbolsRI) GetFromAssetSymbolOk() (*string, bool)`

GetFromAssetSymbolOk returns a tuple with the FromAssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAssetSymbol

`func (o *GetExchangeRateByAssetSymbolsRI) SetFromAssetSymbol(v string)`

SetFromAssetSymbol sets FromAssetSymbol field to given value.


### GetRate

`func (o *GetExchangeRateByAssetSymbolsRI) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetExchangeRateByAssetSymbolsRI) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetExchangeRateByAssetSymbolsRI) SetRate(v string)`

SetRate sets Rate field to given value.


### GetToAssetId

`func (o *GetExchangeRateByAssetSymbolsRI) GetToAssetId() string`

GetToAssetId returns the ToAssetId field if non-nil, zero value otherwise.

### GetToAssetIdOk

`func (o *GetExchangeRateByAssetSymbolsRI) GetToAssetIdOk() (*string, bool)`

GetToAssetIdOk returns a tuple with the ToAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAssetId

`func (o *GetExchangeRateByAssetSymbolsRI) SetToAssetId(v string)`

SetToAssetId sets ToAssetId field to given value.


### GetToAssetSymbol

`func (o *GetExchangeRateByAssetSymbolsRI) GetToAssetSymbol() string`

GetToAssetSymbol returns the ToAssetSymbol field if non-nil, zero value otherwise.

### GetToAssetSymbolOk

`func (o *GetExchangeRateByAssetSymbolsRI) GetToAssetSymbolOk() (*string, bool)`

GetToAssetSymbolOk returns a tuple with the ToAssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAssetSymbol

`func (o *GetExchangeRateByAssetSymbolsRI) SetToAssetSymbol(v string)`

SetToAssetSymbol sets ToAssetSymbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


