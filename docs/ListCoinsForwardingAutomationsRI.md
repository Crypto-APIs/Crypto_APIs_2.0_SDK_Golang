# ListCoinsForwardingAutomationsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | 
**ConfirmationsCountTrigger** | **int32** | Represents the total count of the transaction confirmations before triggering the event. | 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the automatic forwarding was created in Unix Timestamp. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;SLOW\&quot;, \&quot;STANDARD\&quot; OR \&quot;FAST\&quot;. | 
**FromAddress** | **string** | Represents the hash of the address that forwards the currency. | 
**MinimumTransferAmount** | **string** | Represents the minimum transfer amount of the currency in the &#x60;fromAddress&#x60; that can be allowed for an automatic forwarding. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific forwarding automation. | 
**ToAddress** | **string** | Represents the hash of the address the currency is forwarded to. | 

## Methods

### NewListCoinsForwardingAutomationsRI

`func NewListCoinsForwardingAutomationsRI(callbackUrl string, confirmationsCountTrigger int32, createdTimestamp int32, feePriority string, fromAddress string, minimumTransferAmount string, referenceId string, toAddress string, ) *ListCoinsForwardingAutomationsRI`

NewListCoinsForwardingAutomationsRI instantiates a new ListCoinsForwardingAutomationsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCoinsForwardingAutomationsRIWithDefaults

`func NewListCoinsForwardingAutomationsRIWithDefaults() *ListCoinsForwardingAutomationsRI`

NewListCoinsForwardingAutomationsRIWithDefaults instantiates a new ListCoinsForwardingAutomationsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *ListCoinsForwardingAutomationsRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ListCoinsForwardingAutomationsRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ListCoinsForwardingAutomationsRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCountTrigger

`func (o *ListCoinsForwardingAutomationsRI) GetConfirmationsCountTrigger() int32`

GetConfirmationsCountTrigger returns the ConfirmationsCountTrigger field if non-nil, zero value otherwise.

### GetConfirmationsCountTriggerOk

`func (o *ListCoinsForwardingAutomationsRI) GetConfirmationsCountTriggerOk() (*int32, bool)`

GetConfirmationsCountTriggerOk returns a tuple with the ConfirmationsCountTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCountTrigger

`func (o *ListCoinsForwardingAutomationsRI) SetConfirmationsCountTrigger(v int32)`

SetConfirmationsCountTrigger sets ConfirmationsCountTrigger field to given value.


### GetCreatedTimestamp

`func (o *ListCoinsForwardingAutomationsRI) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ListCoinsForwardingAutomationsRI) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ListCoinsForwardingAutomationsRI) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetFeePriority

`func (o *ListCoinsForwardingAutomationsRI) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *ListCoinsForwardingAutomationsRI) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *ListCoinsForwardingAutomationsRI) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetFromAddress

`func (o *ListCoinsForwardingAutomationsRI) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ListCoinsForwardingAutomationsRI) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ListCoinsForwardingAutomationsRI) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetMinimumTransferAmount

`func (o *ListCoinsForwardingAutomationsRI) GetMinimumTransferAmount() string`

GetMinimumTransferAmount returns the MinimumTransferAmount field if non-nil, zero value otherwise.

### GetMinimumTransferAmountOk

`func (o *ListCoinsForwardingAutomationsRI) GetMinimumTransferAmountOk() (*string, bool)`

GetMinimumTransferAmountOk returns a tuple with the MinimumTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTransferAmount

`func (o *ListCoinsForwardingAutomationsRI) SetMinimumTransferAmount(v string)`

SetMinimumTransferAmount sets MinimumTransferAmount field to given value.


### GetReferenceId

`func (o *ListCoinsForwardingAutomationsRI) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ListCoinsForwardingAutomationsRI) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ListCoinsForwardingAutomationsRI) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetToAddress

`func (o *ListCoinsForwardingAutomationsRI) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ListCoinsForwardingAutomationsRI) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ListCoinsForwardingAutomationsRI) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


