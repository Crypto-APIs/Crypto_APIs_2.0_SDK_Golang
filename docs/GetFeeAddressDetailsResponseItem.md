# GetFeeAddressDetailsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the specific fee address, which is always automatically generated. Users must fund it. | 
**Balance** | [**GetFeeAddressDetailsResponseItemBalance**](GetFeeAddressDetailsResponseItemBalance.md) |  | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the currency in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 

## Methods

### NewGetFeeAddressDetailsResponseItem

`func NewGetFeeAddressDetailsResponseItem(address string, balance GetFeeAddressDetailsResponseItemBalance, minimumTransferAmount string, ) *GetFeeAddressDetailsResponseItem`

NewGetFeeAddressDetailsResponseItem instantiates a new GetFeeAddressDetailsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeeAddressDetailsResponseItemWithDefaults

`func NewGetFeeAddressDetailsResponseItemWithDefaults() *GetFeeAddressDetailsResponseItem`

NewGetFeeAddressDetailsResponseItemWithDefaults instantiates a new GetFeeAddressDetailsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetFeeAddressDetailsResponseItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetFeeAddressDetailsResponseItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetFeeAddressDetailsResponseItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *GetFeeAddressDetailsResponseItem) GetBalance() GetFeeAddressDetailsResponseItemBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetFeeAddressDetailsResponseItem) GetBalanceOk() (*GetFeeAddressDetailsResponseItemBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetFeeAddressDetailsResponseItem) SetBalance(v GetFeeAddressDetailsResponseItemBalance)`

SetBalance sets Balance field to given value.


### GetMinimumTransferAmount

`func (o *GetFeeAddressDetailsResponseItem) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *GetFeeAddressDetailsResponseItem) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *GetFeeAddressDetailsResponseItem) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


