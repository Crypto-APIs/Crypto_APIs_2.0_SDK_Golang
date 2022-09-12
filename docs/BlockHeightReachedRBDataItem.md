# BlockHeightReachedRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDuplicates** | Pointer to **bool** | Specifies a flag that permits or denies the creation of duplicate addresses. | [optional] [default to false]
**BlockHeightReached** | **int64** | Represents the specified value of block height for which the callback will be received. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | 

## Methods

### NewBlockHeightReachedRBDataItem

`func NewBlockHeightReachedRBDataItem(blockHeightReached int64, callbackUrl string, ) *BlockHeightReachedRBDataItem`

NewBlockHeightReachedRBDataItem instantiates a new BlockHeightReachedRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockHeightReachedRBDataItemWithDefaults

`func NewBlockHeightReachedRBDataItemWithDefaults() *BlockHeightReachedRBDataItem`

NewBlockHeightReachedRBDataItemWithDefaults instantiates a new BlockHeightReachedRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDuplicates

`func (o *BlockHeightReachedRBDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *BlockHeightReachedRBDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *BlockHeightReachedRBDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *BlockHeightReachedRBDataItem) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetBlockHeightReached

`func (o *BlockHeightReachedRBDataItem) GetBlockHeightReached() int64`

GetBlockHeightReached returns the BlockHeightReached field if non-nil, zero value otherwise.

### GetBlockHeightReachedOk

`func (o *BlockHeightReachedRBDataItem) GetBlockHeightReachedOk() (*int64, bool)`

GetBlockHeightReachedOk returns a tuple with the BlockHeightReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeightReached

`func (o *BlockHeightReachedRBDataItem) SetBlockHeightReached(v int64)`

SetBlockHeightReached sets BlockHeightReached field to given value.


### GetCallbackSecretKey

`func (o *BlockHeightReachedRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *BlockHeightReachedRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *BlockHeightReachedRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *BlockHeightReachedRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *BlockHeightReachedRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *BlockHeightReachedRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *BlockHeightReachedRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


