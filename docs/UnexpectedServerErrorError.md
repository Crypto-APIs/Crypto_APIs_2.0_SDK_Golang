# UnexpectedServerErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressErrorDetails**](BannedIpAddressErrorDetails.md) |  | [optional] 

## Methods

### NewUnexpectedServerErrorError

`func NewUnexpectedServerErrorError(code string, message string, ) *UnexpectedServerErrorError`

NewUnexpectedServerErrorError instantiates a new UnexpectedServerErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnexpectedServerErrorErrorWithDefaults

`func NewUnexpectedServerErrorErrorWithDefaults() *UnexpectedServerErrorError`

NewUnexpectedServerErrorErrorWithDefaults instantiates a new UnexpectedServerErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnexpectedServerErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnexpectedServerErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnexpectedServerErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *UnexpectedServerErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnexpectedServerErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnexpectedServerErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *UnexpectedServerErrorError) GetDetails() []BannedIpAddressErrorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UnexpectedServerErrorError) GetDetailsOk() (*[]BannedIpAddressErrorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UnexpectedServerErrorError) SetDetails(v []BannedIpAddressErrorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *UnexpectedServerErrorError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


