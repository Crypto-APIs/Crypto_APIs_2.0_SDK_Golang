# GetZilliqaAddressDetailsE403

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetailsInner**](BannedIpAddressDetailsInner.md) |  | [optional] 

## Methods

### NewGetZilliqaAddressDetailsE403

`func NewGetZilliqaAddressDetailsE403(code string, message string, ) *GetZilliqaAddressDetailsE403`

NewGetZilliqaAddressDetailsE403 instantiates a new GetZilliqaAddressDetailsE403 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZilliqaAddressDetailsE403WithDefaults

`func NewGetZilliqaAddressDetailsE403WithDefaults() *GetZilliqaAddressDetailsE403`

NewGetZilliqaAddressDetailsE403WithDefaults instantiates a new GetZilliqaAddressDetailsE403 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetZilliqaAddressDetailsE403) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetZilliqaAddressDetailsE403) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetZilliqaAddressDetailsE403) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *GetZilliqaAddressDetailsE403) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetZilliqaAddressDetailsE403) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetZilliqaAddressDetailsE403) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *GetZilliqaAddressDetailsE403) GetDetails() []BannedIpAddressDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetZilliqaAddressDetailsE403) GetDetailsOk() (*[]BannedIpAddressDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetZilliqaAddressDetailsE403) SetDetails(v []BannedIpAddressDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetZilliqaAddressDetailsE403) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


