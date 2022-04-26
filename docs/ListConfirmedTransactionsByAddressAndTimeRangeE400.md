# ListConfirmedTransactionsByAddressAndTimeRangeE400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetails**](BannedIpAddressDetails.md) |  | [optional] 

## Methods

### NewListConfirmedTransactionsByAddressAndTimeRangeE400

`func NewListConfirmedTransactionsByAddressAndTimeRangeE400(code string, message string, ) *ListConfirmedTransactionsByAddressAndTimeRangeE400`

NewListConfirmedTransactionsByAddressAndTimeRangeE400 instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeE400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressAndTimeRangeE400WithDefaults

`func NewListConfirmedTransactionsByAddressAndTimeRangeE400WithDefaults() *ListConfirmedTransactionsByAddressAndTimeRangeE400`

NewListConfirmedTransactionsByAddressAndTimeRangeE400WithDefaults instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeE400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) GetDetails() []BannedIpAddressDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) GetDetailsOk() (*[]BannedIpAddressDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) SetDetails(v []BannedIpAddressDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeE400) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


