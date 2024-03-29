# GetExchangeRateByAssetsIDsRI

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

### NewGetExchangeRateByAssetsIDsRI

`func NewGetExchangeRateByAssetsIDsRI(calculationTimestamp int32, fromAssetId string, fromAssetSymbol string, rate string, toAssetId string, toAssetSymbol string, ) *GetExchangeRateByAssetsIDsRI`

NewGetExchangeRateByAssetsIDsRI instantiates a new GetExchangeRateByAssetsIDsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeRateByAssetsIDsRIWithDefaults

`func NewGetExchangeRateByAssetsIDsRIWithDefaults() *GetExchangeRateByAssetsIDsRI`

NewGetExchangeRateByAssetsIDsRIWithDefaults instantiates a new GetExchangeRateByAssetsIDsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculationTimestamp

`func (o *GetExchangeRateByAssetsIDsRI) GetCalculationTimestamp() int32`

GetCalculationTimestamp returns the CalculationTimestamp field if non-nil, zero value otherwise.

### GetCalculationTimestampOk

`func (o *GetExchangeRateByAssetsIDsRI) GetCalculationTimestampOk() (*int32, bool)`

GetCalculationTimestampOk returns a tuple with the CalculationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationTimestamp

`func (o *GetExchangeRateByAssetsIDsRI) SetCalculationTimestamp(v int32)`

SetCalculationTimestamp sets CalculationTimestamp field to given value.


### GetFromAssetId

`func (o *GetExchangeRateByAssetsIDsRI) GetFromAssetId() string`

GetFromAssetId returns the FromAssetId field if non-nil, zero value otherwise.

### GetFromAssetIdOk

`func (o *GetExchangeRateByAssetsIDsRI) GetFromAssetIdOk() (*string, bool)`

GetFromAssetIdOk returns a tuple with the FromAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAssetId

`func (o *GetExchangeRateByAssetsIDsRI) SetFromAssetId(v string)`

SetFromAssetId sets FromAssetId field to given value.


### GetFromAssetSymbol

`func (o *GetExchangeRateByAssetsIDsRI) GetFromAssetSymbol() string`

GetFromAssetSymbol returns the FromAssetSymbol field if non-nil, zero value otherwise.

### GetFromAssetSymbolOk

`func (o *GetExchangeRateByAssetsIDsRI) GetFromAssetSymbolOk() (*string, bool)`

GetFromAssetSymbolOk returns a tuple with the FromAssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAssetSymbol

`func (o *GetExchangeRateByAssetsIDsRI) SetFromAssetSymbol(v string)`

SetFromAssetSymbol sets FromAssetSymbol field to given value.


### GetRate

`func (o *GetExchangeRateByAssetsIDsRI) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetExchangeRateByAssetsIDsRI) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetExchangeRateByAssetsIDsRI) SetRate(v string)`

SetRate sets Rate field to given value.


### GetToAssetId

`func (o *GetExchangeRateByAssetsIDsRI) GetToAssetId() string`

GetToAssetId returns the ToAssetId field if non-nil, zero value otherwise.

### GetToAssetIdOk

`func (o *GetExchangeRateByAssetsIDsRI) GetToAssetIdOk() (*string, bool)`

GetToAssetIdOk returns a tuple with the ToAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAssetId

`func (o *GetExchangeRateByAssetsIDsRI) SetToAssetId(v string)`

SetToAssetId sets ToAssetId field to given value.


### GetToAssetSymbol

`func (o *GetExchangeRateByAssetsIDsRI) GetToAssetSymbol() string`

GetToAssetSymbol returns the ToAssetSymbol field if non-nil, zero value otherwise.

### GetToAssetSymbolOk

`func (o *GetExchangeRateByAssetsIDsRI) GetToAssetSymbolOk() (*string, bool)`

GetToAssetSymbolOk returns a tuple with the ToAssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAssetSymbol

`func (o *GetExchangeRateByAssetsIDsRI) SetToAssetSymbol(v string)`

SetToAssetSymbol sets ToAssetSymbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


