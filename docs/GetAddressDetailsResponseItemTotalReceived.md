# GetAddressDetailsResponseItemTotalReceived

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the total amount of all coins received to the address, based on confirmed transactions. | 
**Unit** | **string** | Defines the unit of the received amount, e.g. BTC, ETH, XRP. | 

## Methods

### NewGetAddressDetailsResponseItemTotalReceived

`func NewGetAddressDetailsResponseItemTotalReceived(amount string, unit string, ) *GetAddressDetailsResponseItemTotalReceived`

NewGetAddressDetailsResponseItemTotalReceived instantiates a new GetAddressDetailsResponseItemTotalReceived object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddressDetailsResponseItemTotalReceivedWithDefaults

`func NewGetAddressDetailsResponseItemTotalReceivedWithDefaults() *GetAddressDetailsResponseItemTotalReceived`

NewGetAddressDetailsResponseItemTotalReceivedWithDefaults instantiates a new GetAddressDetailsResponseItemTotalReceived object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetAddressDetailsResponseItemTotalReceived) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetAddressDetailsResponseItemTotalReceived) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetAddressDetailsResponseItemTotalReceived) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *GetAddressDetailsResponseItemTotalReceived) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetAddressDetailsResponseItemTotalReceived) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetAddressDetailsResponseItemTotalReceived) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

