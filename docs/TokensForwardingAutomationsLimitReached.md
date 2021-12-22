# TokensForwardingAutomationsLimitReached

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetails**](BannedIpAddressDetails.md) |  | [optional] 

## Methods

### NewTokensForwardingAutomationsLimitReached

`func NewTokensForwardingAutomationsLimitReached(code string, message string, ) *TokensForwardingAutomationsLimitReached`

NewTokensForwardingAutomationsLimitReached instantiates a new TokensForwardingAutomationsLimitReached object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensForwardingAutomationsLimitReachedWithDefaults

`func NewTokensForwardingAutomationsLimitReachedWithDefaults() *TokensForwardingAutomationsLimitReached`

NewTokensForwardingAutomationsLimitReachedWithDefaults instantiates a new TokensForwardingAutomationsLimitReached object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TokensForwardingAutomationsLimitReached) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TokensForwardingAutomationsLimitReached) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TokensForwardingAutomationsLimitReached) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *TokensForwardingAutomationsLimitReached) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TokensForwardingAutomationsLimitReached) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TokensForwardingAutomationsLimitReached) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *TokensForwardingAutomationsLimitReached) GetDetails() []BannedIpAddressDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TokensForwardingAutomationsLimitReached) GetDetailsOk() (*[]BannedIpAddressDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TokensForwardingAutomationsLimitReached) SetDetails(v []BannedIpAddressDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TokensForwardingAutomationsLimitReached) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


