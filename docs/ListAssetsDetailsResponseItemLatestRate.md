# ListAssetsDetailsResponseItemLatestRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Specifies the amount of the latest price of the asset. | 
**CalculationTimestamp** | Pointer to **int32** | Defines when the price was calculated in UNIX timestamp. | [optional] 
**Unit** | **string** | Specifies the unit of the latest price of the asset. | 

## Methods

### NewListAssetsDetailsResponseItemLatestRate

`func NewListAssetsDetailsResponseItemLatestRate(amount string, unit string, ) *ListAssetsDetailsResponseItemLatestRate`

NewListAssetsDetailsResponseItemLatestRate instantiates a new ListAssetsDetailsResponseItemLatestRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetsDetailsResponseItemLatestRateWithDefaults

`func NewListAssetsDetailsResponseItemLatestRateWithDefaults() *ListAssetsDetailsResponseItemLatestRate`

NewListAssetsDetailsResponseItemLatestRateWithDefaults instantiates a new ListAssetsDetailsResponseItemLatestRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListAssetsDetailsResponseItemLatestRate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListAssetsDetailsResponseItemLatestRate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListAssetsDetailsResponseItemLatestRate) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCalculationTimestamp

`func (o *ListAssetsDetailsResponseItemLatestRate) GetCalculationTimestamp() int32`

GetCalculationTimestamp returns the CalculationTimestamp field if non-nil, zero value otherwise.

### GetCalculationTimestampOk

`func (o *ListAssetsDetailsResponseItemLatestRate) GetCalculationTimestampOk() (*int32, bool)`

GetCalculationTimestampOk returns a tuple with the CalculationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationTimestamp

`func (o *ListAssetsDetailsResponseItemLatestRate) SetCalculationTimestamp(v int32)`

SetCalculationTimestamp sets CalculationTimestamp field to given value.

### HasCalculationTimestamp

`func (o *ListAssetsDetailsResponseItemLatestRate) HasCalculationTimestamp() bool`

HasCalculationTimestamp returns a boolean if a field has been set.

### GetUnit

`func (o *ListAssetsDetailsResponseItemLatestRate) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListAssetsDetailsResponseItemLatestRate) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListAssetsDetailsResponseItemLatestRate) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


