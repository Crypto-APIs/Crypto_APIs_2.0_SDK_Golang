# WalletAsAServiceDepositAddressesLimitReachedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressErrorDetails**](BannedIpAddressErrorDetails.md) |  | [optional] 

## Methods

### NewWalletAsAServiceDepositAddressesLimitReachedError

`func NewWalletAsAServiceDepositAddressesLimitReachedError(code string, message string, ) *WalletAsAServiceDepositAddressesLimitReachedError`

NewWalletAsAServiceDepositAddressesLimitReachedError instantiates a new WalletAsAServiceDepositAddressesLimitReachedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletAsAServiceDepositAddressesLimitReachedErrorWithDefaults

`func NewWalletAsAServiceDepositAddressesLimitReachedErrorWithDefaults() *WalletAsAServiceDepositAddressesLimitReachedError`

NewWalletAsAServiceDepositAddressesLimitReachedErrorWithDefaults instantiates a new WalletAsAServiceDepositAddressesLimitReachedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) GetDetails() []BannedIpAddressErrorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) GetDetailsOk() (*[]BannedIpAddressErrorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) SetDetails(v []BannedIpAddressErrorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *WalletAsAServiceDepositAddressesLimitReachedError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


