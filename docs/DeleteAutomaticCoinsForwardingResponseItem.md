# DeleteAutomaticCoinsForwardingResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**ConfirmationsCount** | **string** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | 
**CreatedTimestamp** | **string** | Defines the specific time/date when the automatic forwarding was created in Unix Timestamp. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;SLOW\&quot;, \&quot;STANDARD\&quot; or \&quot;FAST\&quot;. | 
**FromAddress** | **string** | Represents the hash of the address that forwards the currency. | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the currency in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 
**ToAddress** | **string** | Represents the hash of the address the currency is forwarded to. | 

## Methods

### NewDeleteAutomaticCoinsForwardingResponseItem

`func NewDeleteAutomaticCoinsForwardingResponseItem(callbackUrl string, confirmationsCount string, createdTimestamp string, feePriority string, fromAddress string, minimumTransferAmount string, referenceId string, toAddress string, ) *DeleteAutomaticCoinsForwardingResponseItem`

NewDeleteAutomaticCoinsForwardingResponseItem instantiates a new DeleteAutomaticCoinsForwardingResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAutomaticCoinsForwardingResponseItemWithDefaults

`func NewDeleteAutomaticCoinsForwardingResponseItemWithDefaults() *DeleteAutomaticCoinsForwardingResponseItem`

NewDeleteAutomaticCoinsForwardingResponseItemWithDefaults instantiates a new DeleteAutomaticCoinsForwardingResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetConfirmationsCount() string`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetConfirmationsCountOk() (*string, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetConfirmationsCount(v string)`

SetConfirmationsCount sets ConfirmationsCount field to given value.


### GetCreatedTimestamp

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetFeePriority

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetFromAddress

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetMinimumTransferAmount

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.


### GetReferenceId

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetToAddress

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *DeleteAutomaticCoinsForwardingResponseItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *DeleteAutomaticCoinsForwardingResponseItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


