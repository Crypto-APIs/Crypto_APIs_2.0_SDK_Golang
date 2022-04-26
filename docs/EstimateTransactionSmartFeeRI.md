# EstimateTransactionSmartFeeRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmationTarget** | **int32** | Represents the confirmation target in blocks | 
**FeeRate** | **string** | Represents the Fee Rate value. | 
**Unit** | **string** | Defines the fee unit. | 

## Methods

### NewEstimateTransactionSmartFeeRI

`func NewEstimateTransactionSmartFeeRI(confirmationTarget int32, feeRate string, unit string, ) *EstimateTransactionSmartFeeRI`

NewEstimateTransactionSmartFeeRI instantiates a new EstimateTransactionSmartFeeRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateTransactionSmartFeeRIWithDefaults

`func NewEstimateTransactionSmartFeeRIWithDefaults() *EstimateTransactionSmartFeeRI`

NewEstimateTransactionSmartFeeRIWithDefaults instantiates a new EstimateTransactionSmartFeeRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmationTarget

`func (o *EstimateTransactionSmartFeeRI) GetConfirmationTarget() int32`

GetConfirmationTarget returns the ConfirmationTarget field if non-nil, zero value otherwise.

### GetConfirmationTargetOk

`func (o *EstimateTransactionSmartFeeRI) GetConfirmationTargetOk() (*int32, bool)`

GetConfirmationTargetOk returns a tuple with the ConfirmationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationTarget

`func (o *EstimateTransactionSmartFeeRI) SetConfirmationTarget(v int32)`

SetConfirmationTarget sets ConfirmationTarget field to given value.


### GetFeeRate

`func (o *EstimateTransactionSmartFeeRI) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *EstimateTransactionSmartFeeRI) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *EstimateTransactionSmartFeeRI) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetUnit

`func (o *EstimateTransactionSmartFeeRI) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *EstimateTransactionSmartFeeRI) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *EstimateTransactionSmartFeeRI) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


