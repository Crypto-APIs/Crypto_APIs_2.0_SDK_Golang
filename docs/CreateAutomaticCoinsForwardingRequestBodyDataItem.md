# CreateAutomaticCoinsForwardingRequestBodyDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**ConfirmationsCount** | **int32** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the currency in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 
**ToAddress** | **string** | Represents the hash of the address the currency is forwarded to. | 

## Methods

### NewCreateAutomaticCoinsForwardingRequestBodyDataItem

`func NewCreateAutomaticCoinsForwardingRequestBodyDataItem(callbackSecretKey string, callbackUrl string, confirmationsCount int32, feePriority string, minimumTransferAmount string, toAddress string, ) *CreateAutomaticCoinsForwardingRequestBodyDataItem`

NewCreateAutomaticCoinsForwardingRequestBodyDataItem instantiates a new CreateAutomaticCoinsForwardingRequestBodyDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomaticCoinsForwardingRequestBodyDataItemWithDefaults

`func NewCreateAutomaticCoinsForwardingRequestBodyDataItemWithDefaults() *CreateAutomaticCoinsForwardingRequestBodyDataItem`

NewCreateAutomaticCoinsForwardingRequestBodyDataItemWithDefaults instantiates a new CreateAutomaticCoinsForwardingRequestBodyDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetConfirmationsCount() int32`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetConfirmationsCountOk() (*int32, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) SetConfirmationsCount(v int32)`

SetConfirmationsCount sets ConfirmationsCount field to given value.


### GetFeePriority

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetMinimumTransferAmount

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.


### GetToAddress

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateAutomaticCoinsForwardingRequestBodyDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


