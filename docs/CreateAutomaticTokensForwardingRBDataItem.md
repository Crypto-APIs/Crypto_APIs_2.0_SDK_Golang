# CreateAutomaticTokensForwardingRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | 
**ConfirmationsCount** | **string** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;SLOW\&quot;, \&quot;STANDARD\&quot; or \&quot;FAST\&quot;. | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the currency in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 
**ToAddress** | **string** | Represents the hash of the address the currency is forwarded to. | 
**TokenData** | [**CreateAutomaticTokensForwardingRBTokenData**](CreateAutomaticTokensForwardingRBTokenData.md) |  | 

## Methods

### NewCreateAutomaticTokensForwardingRBDataItem

`func NewCreateAutomaticTokensForwardingRBDataItem(callbackUrl string, confirmationsCount string, feePriority string, minimumTransferAmount string, toAddress string, tokenData CreateAutomaticTokensForwardingRBTokenData, ) *CreateAutomaticTokensForwardingRBDataItem`

NewCreateAutomaticTokensForwardingRBDataItem instantiates a new CreateAutomaticTokensForwardingRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomaticTokensForwardingRBDataItemWithDefaults

`func NewCreateAutomaticTokensForwardingRBDataItemWithDefaults() *CreateAutomaticTokensForwardingRBDataItem`

NewCreateAutomaticTokensForwardingRBDataItemWithDefaults instantiates a new CreateAutomaticTokensForwardingRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateAutomaticTokensForwardingRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateAutomaticTokensForwardingRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateAutomaticTokensForwardingRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetConfirmationsCount() string`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetConfirmationsCountOk() (*string, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *CreateAutomaticTokensForwardingRBDataItem) SetConfirmationsCount(v string)`

SetConfirmationsCount sets ConfirmationsCount field to given value.


### GetFeePriority

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateAutomaticTokensForwardingRBDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetMinimumTransferAmount

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *CreateAutomaticTokensForwardingRBDataItem) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.


### GetToAddress

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateAutomaticTokensForwardingRBDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetTokenData

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetTokenData() CreateAutomaticTokensForwardingRBTokenData`

GetTokenData returns the TokenData field if non-nil, zero value otherwise.

### GetTokenDataOk

`func (o *CreateAutomaticTokensForwardingRBDataItem) GetTokenDataOk() (*CreateAutomaticTokensForwardingRBTokenData, bool)`

GetTokenDataOk returns a tuple with the TokenData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenData

`func (o *CreateAutomaticTokensForwardingRBDataItem) SetTokenData(v CreateAutomaticTokensForwardingRBTokenData)`

SetTokenData sets TokenData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


