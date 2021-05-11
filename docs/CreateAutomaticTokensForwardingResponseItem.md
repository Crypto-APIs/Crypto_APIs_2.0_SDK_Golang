# CreateAutomaticTokensForwardingResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**ConfirmationsCount** | **int32** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the automatic forwarding was created in Unix Timestamp. | 
**FeeAddress** | **string** | Represents the specific fee address, which is always automatically generated. Users must fund it. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;SLOW\&quot;, \&quot;STANDARD\&quot; or \&quot;FAST\&quot;. | 
**FromAddress** | **string** | Represents the hash of the address that forwards the tokens. | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the tokens in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 
**ToAddress** | **string** | Represents the hash of the address the tokens are forwarded to. | 
**TokenData** | [**CreateAutomaticTokensForwardingResponseItemTokenData**](CreateAutomaticTokensForwardingResponseItemTokenData.md) |  | 

## Methods

### NewCreateAutomaticTokensForwardingResponseItem

`func NewCreateAutomaticTokensForwardingResponseItem(callbackUrl string, confirmationsCount int32, createdTimestamp int32, feeAddress string, feePriority string, fromAddress string, minimumTransferAmount string, referenceId string, toAddress string, tokenData CreateAutomaticTokensForwardingResponseItemTokenData, ) *CreateAutomaticTokensForwardingResponseItem`

NewCreateAutomaticTokensForwardingResponseItem instantiates a new CreateAutomaticTokensForwardingResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomaticTokensForwardingResponseItemWithDefaults

`func NewCreateAutomaticTokensForwardingResponseItemWithDefaults() *CreateAutomaticTokensForwardingResponseItem`

NewCreateAutomaticTokensForwardingResponseItemWithDefaults instantiates a new CreateAutomaticTokensForwardingResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *CreateAutomaticTokensForwardingResponseItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateAutomaticTokensForwardingResponseItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *CreateAutomaticTokensForwardingResponseItem) GetConfirmationsCount() int32`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetConfirmationsCountOk() (*int32, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *CreateAutomaticTokensForwardingResponseItem) SetConfirmationsCount(v int32)`

SetConfirmationsCount sets ConfirmationsCount field to given value.


### GetCreatedTimestamp

`func (o *CreateAutomaticTokensForwardingResponseItem) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *CreateAutomaticTokensForwardingResponseItem) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetFeeAddress

`func (o *CreateAutomaticTokensForwardingResponseItem) GetFeeAddress() string`

GetFeeAddress returns the FeeAddress field if non-nil, zero value otherwise.

### GetFeeAddressOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetFeeAddressOk() (*string, bool)`

GetFeeAddressOk returns a tuple with the FeeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAddress

`func (o *CreateAutomaticTokensForwardingResponseItem) SetFeeAddress(v string)`

SetFeeAddress sets FeeAddress field to given value.


### GetFeePriority

`func (o *CreateAutomaticTokensForwardingResponseItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateAutomaticTokensForwardingResponseItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetFromAddress

`func (o *CreateAutomaticTokensForwardingResponseItem) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CreateAutomaticTokensForwardingResponseItem) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetMinimumTransferAmount

`func (o *CreateAutomaticTokensForwardingResponseItem) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *CreateAutomaticTokensForwardingResponseItem) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.


### GetReferenceId

`func (o *CreateAutomaticTokensForwardingResponseItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreateAutomaticTokensForwardingResponseItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetToAddress

`func (o *CreateAutomaticTokensForwardingResponseItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateAutomaticTokensForwardingResponseItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetTokenData

`func (o *CreateAutomaticTokensForwardingResponseItem) GetTokenData() CreateAutomaticTokensForwardingResponseItemTokenData`

GetTokenData returns the TokenData field if non-nil, zero value otherwise.

### GetTokenDataOk

`func (o *CreateAutomaticTokensForwardingResponseItem) GetTokenDataOk() (*CreateAutomaticTokensForwardingResponseItemTokenData, bool)`

GetTokenDataOk returns a tuple with the TokenData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenData

`func (o *CreateAutomaticTokensForwardingResponseItem) SetTokenData(v CreateAutomaticTokensForwardingResponseItemTokenData)`

SetTokenData sets TokenData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


