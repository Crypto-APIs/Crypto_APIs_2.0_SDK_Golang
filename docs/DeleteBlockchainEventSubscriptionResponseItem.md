# DeleteBlockchainEventSubscriptionResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the subscription was created in Unix Timestamp. | 
**EventType** | **string** | Defines the type of the specific event available for the customer to subscribe to for callback notification. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

## Methods

### NewDeleteBlockchainEventSubscriptionResponseItem

`func NewDeleteBlockchainEventSubscriptionResponseItem(callbackSecretKey string, callbackUrl string, createdTimestamp int32, eventType string, referenceId string, ) *DeleteBlockchainEventSubscriptionResponseItem`

NewDeleteBlockchainEventSubscriptionResponseItem instantiates a new DeleteBlockchainEventSubscriptionResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBlockchainEventSubscriptionResponseItemWithDefaults

`func NewDeleteBlockchainEventSubscriptionResponseItemWithDefaults() *DeleteBlockchainEventSubscriptionResponseItem`

NewDeleteBlockchainEventSubscriptionResponseItemWithDefaults instantiates a new DeleteBlockchainEventSubscriptionResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *DeleteBlockchainEventSubscriptionResponseItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *DeleteBlockchainEventSubscriptionResponseItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetCreatedTimestamp

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DeleteBlockchainEventSubscriptionResponseItem) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetEventType

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *DeleteBlockchainEventSubscriptionResponseItem) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetReferenceId

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *DeleteBlockchainEventSubscriptionResponseItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *DeleteBlockchainEventSubscriptionResponseItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


