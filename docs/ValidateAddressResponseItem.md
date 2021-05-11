# ValidateAddressResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the specific address that will be checked if it&#39;s valid or not. | 
**IsValid** | **bool** | Defines whether the address is valid or not. Set as boolean. | 

## Methods

### NewValidateAddressResponseItem

`func NewValidateAddressResponseItem(address string, isValid bool, ) *ValidateAddressResponseItem`

NewValidateAddressResponseItem instantiates a new ValidateAddressResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateAddressResponseItemWithDefaults

`func NewValidateAddressResponseItemWithDefaults() *ValidateAddressResponseItem`

NewValidateAddressResponseItemWithDefaults instantiates a new ValidateAddressResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ValidateAddressResponseItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ValidateAddressResponseItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ValidateAddressResponseItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetIsValid

`func (o *ValidateAddressResponseItem) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ValidateAddressResponseItem) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ValidateAddressResponseItem) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


