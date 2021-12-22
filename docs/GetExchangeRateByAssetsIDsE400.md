# GetExchangeRateByAssetsIDsE400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetails**](BannedIpAddressDetails.md) |  | [optional] 

## Methods

### NewGetExchangeRateByAssetsIDsE400

`func NewGetExchangeRateByAssetsIDsE400(code string, message string, ) *GetExchangeRateByAssetsIDsE400`

NewGetExchangeRateByAssetsIDsE400 instantiates a new GetExchangeRateByAssetsIDsE400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeRateByAssetsIDsE400WithDefaults

`func NewGetExchangeRateByAssetsIDsE400WithDefaults() *GetExchangeRateByAssetsIDsE400`

NewGetExchangeRateByAssetsIDsE400WithDefaults instantiates a new GetExchangeRateByAssetsIDsE400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetExchangeRateByAssetsIDsE400) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetExchangeRateByAssetsIDsE400) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetExchangeRateByAssetsIDsE400) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *GetExchangeRateByAssetsIDsE400) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetExchangeRateByAssetsIDsE400) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetExchangeRateByAssetsIDsE400) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *GetExchangeRateByAssetsIDsE400) GetDetails() []BannedIpAddressDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetExchangeRateByAssetsIDsE400) GetDetailsOk() (*[]BannedIpAddressDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetExchangeRateByAssetsIDsE400) SetDetails(v []BannedIpAddressDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetExchangeRateByAssetsIDsE400) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


