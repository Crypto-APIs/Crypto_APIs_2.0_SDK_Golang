# ListAllUnconfirmedTransactionsE403

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetailsInner**](BannedIpAddressDetailsInner.md) |  | [optional] 

## Methods

### NewListAllUnconfirmedTransactionsE403

`func NewListAllUnconfirmedTransactionsE403(code string, message string, ) *ListAllUnconfirmedTransactionsE403`

NewListAllUnconfirmedTransactionsE403 instantiates a new ListAllUnconfirmedTransactionsE403 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsE403WithDefaults

`func NewListAllUnconfirmedTransactionsE403WithDefaults() *ListAllUnconfirmedTransactionsE403`

NewListAllUnconfirmedTransactionsE403WithDefaults instantiates a new ListAllUnconfirmedTransactionsE403 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ListAllUnconfirmedTransactionsE403) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListAllUnconfirmedTransactionsE403) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListAllUnconfirmedTransactionsE403) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ListAllUnconfirmedTransactionsE403) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListAllUnconfirmedTransactionsE403) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListAllUnconfirmedTransactionsE403) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *ListAllUnconfirmedTransactionsE403) GetDetails() []BannedIpAddressDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListAllUnconfirmedTransactionsE403) GetDetailsOk() (*[]BannedIpAddressDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListAllUnconfirmedTransactionsE403) SetDetails(v []BannedIpAddressDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListAllUnconfirmedTransactionsE403) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


