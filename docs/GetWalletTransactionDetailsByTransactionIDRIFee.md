# GetWalletTransactionDetailsByTransactionIDRIFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | When isConfirmed is True - Defines the amount of the transaction fee  When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value. | 
**Unit** | **string** |  | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIFee

`func NewGetWalletTransactionDetailsByTransactionIDRIFee(amount string, unit string, ) *GetWalletTransactionDetailsByTransactionIDRIFee`

NewGetWalletTransactionDetailsByTransactionIDRIFee instantiates a new GetWalletTransactionDetailsByTransactionIDRIFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIFeeWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIFeeWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIFee`

NewGetWalletTransactionDetailsByTransactionIDRIFeeWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetWalletTransactionDetailsByTransactionIDRIFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetWalletTransactionDetailsByTransactionIDRIFee) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *GetWalletTransactionDetailsByTransactionIDRIFee) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIFee) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetWalletTransactionDetailsByTransactionIDRIFee) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


