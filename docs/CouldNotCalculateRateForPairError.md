# CouldNotCalculateRateForPairError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressErrorDetails**](BannedIpAddressErrorDetails.md) |  | [optional] 

## Methods

### NewCouldNotCalculateRateForPairError

`func NewCouldNotCalculateRateForPairError(code string, message string, ) *CouldNotCalculateRateForPairError`

NewCouldNotCalculateRateForPairError instantiates a new CouldNotCalculateRateForPairError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouldNotCalculateRateForPairErrorWithDefaults

`func NewCouldNotCalculateRateForPairErrorWithDefaults() *CouldNotCalculateRateForPairError`

NewCouldNotCalculateRateForPairErrorWithDefaults instantiates a new CouldNotCalculateRateForPairError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CouldNotCalculateRateForPairError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CouldNotCalculateRateForPairError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CouldNotCalculateRateForPairError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *CouldNotCalculateRateForPairError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CouldNotCalculateRateForPairError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CouldNotCalculateRateForPairError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *CouldNotCalculateRateForPairError) GetDetails() []BannedIpAddressErrorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CouldNotCalculateRateForPairError) GetDetailsOk() (*[]BannedIpAddressErrorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CouldNotCalculateRateForPairError) SetDetails(v []BannedIpAddressErrorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CouldNotCalculateRateForPairError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


