# ListOmniTokensByAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** | Defines the balance of the Omni tokens to send in the address. | 
**Frozen** | **string** | Defines the amount frozen by the issuer (applies to managed properties only). | 
**Name** | **string** | Defines the name of the Omni tokens to send. | 
**PropertyId** | **int32** | Represents the identifier of the tokens to send. | 
**Reserved** | **string** | Represents the amount reserved by sell offers and accepts. | 

## Methods

### NewListOmniTokensByAddressRI

`func NewListOmniTokensByAddressRI(balance string, frozen string, name string, propertyId int32, reserved string, ) *ListOmniTokensByAddressRI`

NewListOmniTokensByAddressRI instantiates a new ListOmniTokensByAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOmniTokensByAddressRIWithDefaults

`func NewListOmniTokensByAddressRIWithDefaults() *ListOmniTokensByAddressRI`

NewListOmniTokensByAddressRIWithDefaults instantiates a new ListOmniTokensByAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *ListOmniTokensByAddressRI) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ListOmniTokensByAddressRI) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ListOmniTokensByAddressRI) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetFrozen

`func (o *ListOmniTokensByAddressRI) GetFrozen() string`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *ListOmniTokensByAddressRI) GetFrozenOk() (*string, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *ListOmniTokensByAddressRI) SetFrozen(v string)`

SetFrozen sets Frozen field to given value.


### GetName

`func (o *ListOmniTokensByAddressRI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOmniTokensByAddressRI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOmniTokensByAddressRI) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyId

`func (o *ListOmniTokensByAddressRI) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ListOmniTokensByAddressRI) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ListOmniTokensByAddressRI) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetReserved

`func (o *ListOmniTokensByAddressRI) GetReserved() string`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *ListOmniTokensByAddressRI) GetReservedOk() (*string, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *ListOmniTokensByAddressRI) SetReserved(v string)`

SetReserved sets Reserved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


