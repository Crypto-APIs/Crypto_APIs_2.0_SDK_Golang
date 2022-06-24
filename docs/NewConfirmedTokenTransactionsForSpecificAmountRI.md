# NewConfirmedTokenTransactionsForSpecificAmountRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountHigherThan** | **int64** | Represents a specific amount of coins after which the system have to send a callback to customers&#39; callbackUrl. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | 
**ContractAddress** | **string** | Represents the contract address of the token, which controls its logic. It is not the address that holds the tokens. | 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the subscription was created in Unix Timestamp. | 
**EventType** | **string** | Defines the type of the specific event available for the customer to subscribe to for callback notification. | 
**IsActive** | **bool** | Defines whether the subscription is active or not. Set as boolean. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

## Methods

### NewNewConfirmedTokenTransactionsForSpecificAmountRI

`func NewNewConfirmedTokenTransactionsForSpecificAmountRI(amountHigherThan int64, callbackUrl string, contractAddress string, createdTimestamp int32, eventType string, isActive bool, referenceId string, ) *NewConfirmedTokenTransactionsForSpecificAmountRI`

NewNewConfirmedTokenTransactionsForSpecificAmountRI instantiates a new NewConfirmedTokenTransactionsForSpecificAmountRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedTokenTransactionsForSpecificAmountRIWithDefaults

`func NewNewConfirmedTokenTransactionsForSpecificAmountRIWithDefaults() *NewConfirmedTokenTransactionsForSpecificAmountRI`

NewNewConfirmedTokenTransactionsForSpecificAmountRIWithDefaults instantiates a new NewConfirmedTokenTransactionsForSpecificAmountRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountHigherThan

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetAmountHigherThan() int64`

GetAmountHigherThan returns the AmountHigherThan field if non-nil, zero value otherwise.

### GetAmountHigherThanOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetAmountHigherThanOk() (*int64, bool)`

GetAmountHigherThanOk returns a tuple with the AmountHigherThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountHigherThan

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetAmountHigherThan(v int64)`

SetAmountHigherThan sets AmountHigherThan field to given value.


### GetCallbackSecretKey

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetContractAddress

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetCreatedTimestamp

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetEventType

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetIsActive

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetReferenceId

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *NewConfirmedTokenTransactionsForSpecificAmountRI) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


