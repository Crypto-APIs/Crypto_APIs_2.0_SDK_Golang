# NewConfirmedTokensTransactionsAndEachConfirmationResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address of the transaction, per which the result is returned. | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the subscription was created in Unix Timestamp. | 
**EventType** | **string** | Defines the type of the specific event available for the customer to subscribe to for callback notification. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

## Methods

### NewNewConfirmedTokensTransactionsAndEachConfirmationResponseItem

`func NewNewConfirmedTokensTransactionsAndEachConfirmationResponseItem(address string, callbackUrl string, createdTimestamp int32, eventType string, referenceId string, ) *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem`

NewNewConfirmedTokensTransactionsAndEachConfirmationResponseItem instantiates a new NewConfirmedTokensTransactionsAndEachConfirmationResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedTokensTransactionsAndEachConfirmationResponseItemWithDefaults

`func NewNewConfirmedTokensTransactionsAndEachConfirmationResponseItemWithDefaults() *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem`

NewNewConfirmedTokensTransactionsAndEachConfirmationResponseItemWithDefaults instantiates a new NewConfirmedTokensTransactionsAndEachConfirmationResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCallbackUrl

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetCreatedTimestamp

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetEventType

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetReferenceId

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationResponseItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


