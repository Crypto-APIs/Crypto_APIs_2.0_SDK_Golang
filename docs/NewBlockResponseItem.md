# NewBlockResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the subscription was created in Unix Timestamp. | 
**IsActive** | **bool** | Defines whether the subscription is active or not. Set as boolean. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

## Methods

### NewNewBlockResponseItem

`func NewNewBlockResponseItem(callbackSecretKey string, callbackUrl string, createdTimestamp int32, isActive bool, referenceId string, ) *NewBlockResponseItem`

NewNewBlockResponseItem instantiates a new NewBlockResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBlockResponseItemWithDefaults

`func NewNewBlockResponseItemWithDefaults() *NewBlockResponseItem`

NewNewBlockResponseItemWithDefaults instantiates a new NewBlockResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *NewBlockResponseItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewBlockResponseItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewBlockResponseItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *NewBlockResponseItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewBlockResponseItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewBlockResponseItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetCreatedTimestamp

`func (o *NewBlockResponseItem) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NewBlockResponseItem) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NewBlockResponseItem) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetIsActive

`func (o *NewBlockResponseItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NewBlockResponseItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NewBlockResponseItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetReferenceId

`func (o *NewBlockResponseItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *NewBlockResponseItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *NewBlockResponseItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


