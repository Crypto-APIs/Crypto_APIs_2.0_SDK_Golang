# GetTransactionDetailsByTransactionIDFromCallbackRIFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | When isConfirmed is True - Defines the amount of the transaction fee  When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value. | 
**Unit** | **string** | Defines the unit of the fee. | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIFee

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIFee(amount string, unit string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIFee`

NewGetTransactionDetailsByTransactionIDFromCallbackRIFee instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIFeeWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIFeeWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIFee`

NewGetTransactionDetailsByTransactionIDFromCallbackRIFeeWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


