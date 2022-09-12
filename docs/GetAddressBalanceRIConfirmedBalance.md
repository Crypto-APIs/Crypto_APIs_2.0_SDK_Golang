# GetAddressBalanceRIConfirmedBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the total balance of the address that is confirmed. It doesn&#39;t include unconfirmed transactions. | 
**Unit** | **string** | Represents the unit of the confirmed balance. | 

## Methods

### NewGetAddressBalanceRIConfirmedBalance

`func NewGetAddressBalanceRIConfirmedBalance(amount string, unit string, ) *GetAddressBalanceRIConfirmedBalance`

NewGetAddressBalanceRIConfirmedBalance instantiates a new GetAddressBalanceRIConfirmedBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddressBalanceRIConfirmedBalanceWithDefaults

`func NewGetAddressBalanceRIConfirmedBalanceWithDefaults() *GetAddressBalanceRIConfirmedBalance`

NewGetAddressBalanceRIConfirmedBalanceWithDefaults instantiates a new GetAddressBalanceRIConfirmedBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetAddressBalanceRIConfirmedBalance) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetAddressBalanceRIConfirmedBalance) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetAddressBalanceRIConfirmedBalance) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *GetAddressBalanceRIConfirmedBalance) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetAddressBalanceRIConfirmedBalance) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetAddressBalanceRIConfirmedBalance) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


