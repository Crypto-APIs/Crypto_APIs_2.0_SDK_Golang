# NewConfirmedTokensTransactionsAndEachConfirmationRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address of the transaction, per which the result is returned. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | 
**ConfirmationsCount** | Pointer to **int32** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | [optional] 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the subscription was created in Unix Timestamp. | 
**EventType** | **string** | Defines the type of the specific event available for the customer to subscribe to for callback notification. | 
**IsActive** | **bool** | Defines whether the subscription is active or not. Set as boolean. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

## Methods

### NewNewConfirmedTokensTransactionsAndEachConfirmationRI

`func NewNewConfirmedTokensTransactionsAndEachConfirmationRI(address string, callbackUrl string, createdTimestamp int32, eventType string, isActive bool, referenceId string, ) *NewConfirmedTokensTransactionsAndEachConfirmationRI`

NewNewConfirmedTokensTransactionsAndEachConfirmationRI instantiates a new NewConfirmedTokensTransactionsAndEachConfirmationRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedTokensTransactionsAndEachConfirmationRIWithDefaults

`func NewNewConfirmedTokensTransactionsAndEachConfirmationRIWithDefaults() *NewConfirmedTokensTransactionsAndEachConfirmationRI`

NewNewConfirmedTokensTransactionsAndEachConfirmationRIWithDefaults instantiates a new NewConfirmedTokensTransactionsAndEachConfirmationRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCallbackSecretKey

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetConfirmationsCount() int32`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetConfirmationsCountOk() (*int32, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetConfirmationsCount(v int32)`

SetConfirmationsCount sets ConfirmationsCount field to given value.

### HasConfirmationsCount

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) HasConfirmationsCount() bool`

HasConfirmationsCount returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetEventType

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetIsActive

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetReferenceId

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *NewConfirmedTokensTransactionsAndEachConfirmationRI) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


