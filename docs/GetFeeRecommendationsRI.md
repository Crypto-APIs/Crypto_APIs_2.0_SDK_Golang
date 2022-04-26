# GetFeeRecommendationsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Defines the unit of the fee, e.g. BTC. | 
**Fast** | **string** | Fast fee per byte calculated from unconfirmed transactions | 
**Slow** | **string** | Slow fee per byte calculated from unconfirmed transactions | 
**Standard** | **string** | Standard fee per byte calculated from unconfirmed transactions | 
**FeeCushionMultiplier** | **string** | Represents the fee cushion multiplier used to multiply the base fee. | 

## Methods

### NewGetFeeRecommendationsRI

`func NewGetFeeRecommendationsRI(unit string, fast string, slow string, standard string, feeCushionMultiplier string, ) *GetFeeRecommendationsRI`

NewGetFeeRecommendationsRI instantiates a new GetFeeRecommendationsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeeRecommendationsRIWithDefaults

`func NewGetFeeRecommendationsRIWithDefaults() *GetFeeRecommendationsRI`

NewGetFeeRecommendationsRIWithDefaults instantiates a new GetFeeRecommendationsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *GetFeeRecommendationsRI) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetFeeRecommendationsRI) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetFeeRecommendationsRI) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetFast

`func (o *GetFeeRecommendationsRI) GetFast() string`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *GetFeeRecommendationsRI) GetFastOk() (*string, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *GetFeeRecommendationsRI) SetFast(v string)`

SetFast sets Fast field to given value.


### GetSlow

`func (o *GetFeeRecommendationsRI) GetSlow() string`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *GetFeeRecommendationsRI) GetSlowOk() (*string, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *GetFeeRecommendationsRI) SetSlow(v string)`

SetSlow sets Slow field to given value.


### GetStandard

`func (o *GetFeeRecommendationsRI) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *GetFeeRecommendationsRI) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *GetFeeRecommendationsRI) SetStandard(v string)`

SetStandard sets Standard field to given value.


### GetFeeCushionMultiplier

`func (o *GetFeeRecommendationsRI) GetFeeCushionMultiplier() string`

GetFeeCushionMultiplier returns the FeeCushionMultiplier field if non-nil, zero value otherwise.

### GetFeeCushionMultiplierOk

`func (o *GetFeeRecommendationsRI) GetFeeCushionMultiplierOk() (*string, bool)`

GetFeeCushionMultiplierOk returns a tuple with the FeeCushionMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCushionMultiplier

`func (o *GetFeeRecommendationsRI) SetFeeCushionMultiplier(v string)`

SetFeeCushionMultiplier sets FeeCushionMultiplier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


