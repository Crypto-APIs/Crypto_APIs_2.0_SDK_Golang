# CreateAutomaticTokensForwardingRequestBodyDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**ConfirmationsCount** | **string** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;SLOW\&quot;, \&quot;STANDARD\&quot; or \&quot;FAST\&quot;. | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the currency in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 
**ToAddress** | **string** | Represents the hash of the address the currency is forwarded to. | 
**TokenData** | [**CreateAutomaticTokensForwardingRequestBodyTokenData**](CreateAutomaticTokensForwardingRequestBodyTokenData.md) |  | 

## Methods

### NewCreateAutomaticTokensForwardingRequestBodyDataItem

`func NewCreateAutomaticTokensForwardingRequestBodyDataItem(callbackUrl string, confirmationsCount string, feePriority string, minimumTransferAmount string, toAddress string, tokenData CreateAutomaticTokensForwardingRequestBodyTokenData, ) *CreateAutomaticTokensForwardingRequestBodyDataItem`

NewCreateAutomaticTokensForwardingRequestBodyDataItem instantiates a new CreateAutomaticTokensForwardingRequestBodyDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomaticTokensForwardingRequestBodyDataItemWithDefaults

`func NewCreateAutomaticTokensForwardingRequestBodyDataItemWithDefaults() *CreateAutomaticTokensForwardingRequestBodyDataItem`

NewCreateAutomaticTokensForwardingRequestBodyDataItemWithDefaults instantiates a new CreateAutomaticTokensForwardingRequestBodyDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetConfirmationsCount() string`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetConfirmationsCountOk() (*string, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) SetConfirmationsCount(v string)`

SetConfirmationsCount sets ConfirmationsCount field to given value.


### GetFeePriority

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetMinimumTransferAmount

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.


### GetToAddress

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetTokenData

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetTokenData() CreateAutomaticTokensForwardingRequestBodyTokenData`

GetTokenData returns the TokenData field if non-nil, zero value otherwise.

### GetTokenDataOk

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) GetTokenDataOk() (*CreateAutomaticTokensForwardingRequestBodyTokenData, bool)`

GetTokenDataOk returns a tuple with the TokenData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenData

`func (o *CreateAutomaticTokensForwardingRequestBodyDataItem) SetTokenData(v CreateAutomaticTokensForwardingRequestBodyTokenData)`

SetTokenData sets TokenData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


