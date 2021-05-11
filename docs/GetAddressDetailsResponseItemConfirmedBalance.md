# GetAddressDetailsResponseItemConfirmedBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the total balance of the address that is confirmed. It doesn&#39;t include unconfirmed transactions. | 
**Unit** | **string** | Defines the unit of the confirmed balance amount, e.g. BTC, ETH, XRP. | 

## Methods

### NewGetAddressDetailsResponseItemConfirmedBalance

`func NewGetAddressDetailsResponseItemConfirmedBalance(amount string, unit string, ) *GetAddressDetailsResponseItemConfirmedBalance`

NewGetAddressDetailsResponseItemConfirmedBalance instantiates a new GetAddressDetailsResponseItemConfirmedBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddressDetailsResponseItemConfirmedBalanceWithDefaults

`func NewGetAddressDetailsResponseItemConfirmedBalanceWithDefaults() *GetAddressDetailsResponseItemConfirmedBalance`

NewGetAddressDetailsResponseItemConfirmedBalanceWithDefaults instantiates a new GetAddressDetailsResponseItemConfirmedBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetAddressDetailsResponseItemConfirmedBalance) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetAddressDetailsResponseItemConfirmedBalance) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetAddressDetailsResponseItemConfirmedBalance) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *GetAddressDetailsResponseItemConfirmedBalance) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetAddressDetailsResponseItemConfirmedBalance) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetAddressDetailsResponseItemConfirmedBalance) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


