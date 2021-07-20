# ValidateAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the specific address that will be checked if it&#39;s valid or not. | 
**IsValid** | **bool** | Defines whether the address is valid or not. Set as boolean. | 

## Methods

### NewValidateAddressRI

`func NewValidateAddressRI(address string, isValid bool, ) *ValidateAddressRI`

NewValidateAddressRI instantiates a new ValidateAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateAddressRIWithDefaults

`func NewValidateAddressRIWithDefaults() *ValidateAddressRI`

NewValidateAddressRIWithDefaults instantiates a new ValidateAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ValidateAddressRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ValidateAddressRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ValidateAddressRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetIsValid

`func (o *ValidateAddressRI) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ValidateAddressRI) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ValidateAddressRI) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


