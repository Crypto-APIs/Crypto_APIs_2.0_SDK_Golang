# GetWalletDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositAddressesCount** | **int32** | Specifies the count of deposit addresses in the Wallet. | 
**Name** | **string** | Defines the name of the Wallet given to it by the user. | 

## Methods

### NewGetWalletDetailsRI

`func NewGetWalletDetailsRI(depositAddressesCount int32, name string, ) *GetWalletDetailsRI`

NewGetWalletDetailsRI instantiates a new GetWalletDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletDetailsRIWithDefaults

`func NewGetWalletDetailsRIWithDefaults() *GetWalletDetailsRI`

NewGetWalletDetailsRIWithDefaults instantiates a new GetWalletDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositAddressesCount

`func (o *GetWalletDetailsRI) GetDepositAddressesCount() int32`

GetDepositAddressesCount returns the DepositAddressesCount field if non-nil, zero value otherwise.

### GetDepositAddressesCountOk

`func (o *GetWalletDetailsRI) GetDepositAddressesCountOk() (*int32, bool)`

GetDepositAddressesCountOk returns a tuple with the DepositAddressesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAddressesCount

`func (o *GetWalletDetailsRI) SetDepositAddressesCount(v int32)`

SetDepositAddressesCount sets DepositAddressesCount field to given value.


### GetName

`func (o *GetWalletDetailsRI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetWalletDetailsRI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetWalletDetailsRI) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


