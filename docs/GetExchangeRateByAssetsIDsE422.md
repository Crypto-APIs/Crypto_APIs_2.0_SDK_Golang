# GetExchangeRateByAssetsIDsE422

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetails**](BannedIpAddressDetails.md) |  | [optional] 

## Methods

### NewGetExchangeRateByAssetsIDsE422

`func NewGetExchangeRateByAssetsIDsE422(code string, message string, ) *GetExchangeRateByAssetsIDsE422`

NewGetExchangeRateByAssetsIDsE422 instantiates a new GetExchangeRateByAssetsIDsE422 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeRateByAssetsIDsE422WithDefaults

`func NewGetExchangeRateByAssetsIDsE422WithDefaults() *GetExchangeRateByAssetsIDsE422`

NewGetExchangeRateByAssetsIDsE422WithDefaults instantiates a new GetExchangeRateByAssetsIDsE422 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetExchangeRateByAssetsIDsE422) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetExchangeRateByAssetsIDsE422) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetExchangeRateByAssetsIDsE422) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *GetExchangeRateByAssetsIDsE422) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetExchangeRateByAssetsIDsE422) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetExchangeRateByAssetsIDsE422) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *GetExchangeRateByAssetsIDsE422) GetDetails() []BannedIpAddressDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetExchangeRateByAssetsIDsE422) GetDetailsOk() (*[]BannedIpAddressDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetExchangeRateByAssetsIDsE422) SetDetails(v []BannedIpAddressDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetExchangeRateByAssetsIDsE422) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


