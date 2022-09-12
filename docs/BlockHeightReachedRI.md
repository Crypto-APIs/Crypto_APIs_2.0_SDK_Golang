# BlockHeightReachedRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHeightReached** | **int64** | Represents the specified value of block height for which the callback will be received. | 
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. We support ONLY httpS type of protocol. | 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the subscription was created in Unix Timestamp. | 
**IsActive** | **bool** | Defines whether the subscription is active or not. Set as boolean. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 

## Methods

### NewBlockHeightReachedRI

`func NewBlockHeightReachedRI(blockHeightReached int64, callbackSecretKey string, callbackUrl string, createdTimestamp int32, isActive bool, referenceId string, ) *BlockHeightReachedRI`

NewBlockHeightReachedRI instantiates a new BlockHeightReachedRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockHeightReachedRIWithDefaults

`func NewBlockHeightReachedRIWithDefaults() *BlockHeightReachedRI`

NewBlockHeightReachedRIWithDefaults instantiates a new BlockHeightReachedRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHeightReached

`func (o *BlockHeightReachedRI) GetBlockHeightReached() int64`

GetBlockHeightReached returns the BlockHeightReached field if non-nil, zero value otherwise.

### GetBlockHeightReachedOk

`func (o *BlockHeightReachedRI) GetBlockHeightReachedOk() (*int64, bool)`

GetBlockHeightReachedOk returns a tuple with the BlockHeightReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeightReached

`func (o *BlockHeightReachedRI) SetBlockHeightReached(v int64)`

SetBlockHeightReached sets BlockHeightReached field to given value.


### GetCallbackSecretKey

`func (o *BlockHeightReachedRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *BlockHeightReachedRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *BlockHeightReachedRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *BlockHeightReachedRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *BlockHeightReachedRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *BlockHeightReachedRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetCreatedTimestamp

`func (o *BlockHeightReachedRI) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *BlockHeightReachedRI) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *BlockHeightReachedRI) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetIsActive

`func (o *BlockHeightReachedRI) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BlockHeightReachedRI) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BlockHeightReachedRI) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetReferenceId

`func (o *BlockHeightReachedRI) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *BlockHeightReachedRI) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *BlockHeightReachedRI) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


