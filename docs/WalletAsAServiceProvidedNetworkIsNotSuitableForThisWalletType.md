# WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetailsInner**](BannedIpAddressDetailsInner.md) |  | [optional] 

## Methods

### NewWalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType

`func NewWalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType(code string, message string, ) *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType`

NewWalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType instantiates a new WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletTypeWithDefaults

`func NewWalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletTypeWithDefaults() *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType`

NewWalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletTypeWithDefaults instantiates a new WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) GetDetails() []BannedIpAddressDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) GetDetailsOk() (*[]BannedIpAddressDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) SetDetails(v []BannedIpAddressDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *WalletAsAServiceProvidedNetworkIsNotSuitableForThisWalletType) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


