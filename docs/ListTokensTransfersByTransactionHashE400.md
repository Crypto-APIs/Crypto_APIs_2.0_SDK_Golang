# ListTokensTransfersByTransactionHashE400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetails**](BannedIpAddressDetails.md) |  | [optional] 

## Methods

### NewListTokensTransfersByTransactionHashE400

`func NewListTokensTransfersByTransactionHashE400(code string, message string, ) *ListTokensTransfersByTransactionHashE400`

NewListTokensTransfersByTransactionHashE400 instantiates a new ListTokensTransfersByTransactionHashE400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokensTransfersByTransactionHashE400WithDefaults

`func NewListTokensTransfersByTransactionHashE400WithDefaults() *ListTokensTransfersByTransactionHashE400`

NewListTokensTransfersByTransactionHashE400WithDefaults instantiates a new ListTokensTransfersByTransactionHashE400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ListTokensTransfersByTransactionHashE400) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTokensTransfersByTransactionHashE400) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTokensTransfersByTransactionHashE400) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ListTokensTransfersByTransactionHashE400) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListTokensTransfersByTransactionHashE400) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListTokensTransfersByTransactionHashE400) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *ListTokensTransfersByTransactionHashE400) GetDetails() []BannedIpAddressDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListTokensTransfersByTransactionHashE400) GetDetailsOk() (*[]BannedIpAddressDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListTokensTransfersByTransactionHashE400) SetDetails(v []BannedIpAddressDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListTokensTransfersByTransactionHashE400) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


