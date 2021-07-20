# GetFeeAddressDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the specific fee address, which is always automatically generated. Users must fund it. | 
**Balance** | [**GetFeeAddressDetailsRIBalance**](GetFeeAddressDetailsRIBalance.md) |  | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the currency in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 

## Methods

### NewGetFeeAddressDetailsRI

`func NewGetFeeAddressDetailsRI(address string, balance GetFeeAddressDetailsRIBalance, minimumTransferAmount string, ) *GetFeeAddressDetailsRI`

NewGetFeeAddressDetailsRI instantiates a new GetFeeAddressDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeeAddressDetailsRIWithDefaults

`func NewGetFeeAddressDetailsRIWithDefaults() *GetFeeAddressDetailsRI`

NewGetFeeAddressDetailsRIWithDefaults instantiates a new GetFeeAddressDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetFeeAddressDetailsRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetFeeAddressDetailsRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetFeeAddressDetailsRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *GetFeeAddressDetailsRI) GetBalance() GetFeeAddressDetailsRIBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetFeeAddressDetailsRI) GetBalanceOk() (*GetFeeAddressDetailsRIBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetFeeAddressDetailsRI) SetBalance(v GetFeeAddressDetailsRIBalance)`

SetBalance sets Balance field to given value.


### GetMinimumTransferAmount

`func (o *GetFeeAddressDetailsRI) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *GetFeeAddressDetailsRI) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *GetFeeAddressDetailsRI) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


