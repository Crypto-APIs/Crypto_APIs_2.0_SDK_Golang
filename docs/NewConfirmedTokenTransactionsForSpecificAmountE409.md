# NewConfirmedTokenTransactionsForSpecificAmountE409

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetailsInner**](BannedIpAddressDetailsInner.md) |  | [optional] 

## Methods

### NewNewConfirmedTokenTransactionsForSpecificAmountE409

`func NewNewConfirmedTokenTransactionsForSpecificAmountE409(code string, message string, ) *NewConfirmedTokenTransactionsForSpecificAmountE409`

NewNewConfirmedTokenTransactionsForSpecificAmountE409 instantiates a new NewConfirmedTokenTransactionsForSpecificAmountE409 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedTokenTransactionsForSpecificAmountE409WithDefaults

`func NewNewConfirmedTokenTransactionsForSpecificAmountE409WithDefaults() *NewConfirmedTokenTransactionsForSpecificAmountE409`

NewNewConfirmedTokenTransactionsForSpecificAmountE409WithDefaults instantiates a new NewConfirmedTokenTransactionsForSpecificAmountE409 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) GetDetails() []BannedIpAddressDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) GetDetailsOk() (*[]BannedIpAddressDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) SetDetails(v []BannedIpAddressDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NewConfirmedTokenTransactionsForSpecificAmountE409) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


