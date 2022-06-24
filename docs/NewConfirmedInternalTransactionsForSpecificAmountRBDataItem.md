# NewConfirmedInternalTransactionsForSpecificAmountRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDuplicates** | Pointer to **bool** | Specifies a flag that permits or denies the creation of duplicate addresses. | [optional] [default to false]
**AmountHigherThan** | **int64** | Represents a specific amount of coins after which the system have to send a callback to customers&#39; callbackUrl. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs 2.0. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | 

## Methods

### NewNewConfirmedInternalTransactionsForSpecificAmountRBDataItem

`func NewNewConfirmedInternalTransactionsForSpecificAmountRBDataItem(amountHigherThan int64, callbackUrl string, ) *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem`

NewNewConfirmedInternalTransactionsForSpecificAmountRBDataItem instantiates a new NewConfirmedInternalTransactionsForSpecificAmountRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedInternalTransactionsForSpecificAmountRBDataItemWithDefaults

`func NewNewConfirmedInternalTransactionsForSpecificAmountRBDataItemWithDefaults() *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem`

NewNewConfirmedInternalTransactionsForSpecificAmountRBDataItemWithDefaults instantiates a new NewConfirmedInternalTransactionsForSpecificAmountRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDuplicates

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetAmountHigherThan

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetAmountHigherThan() int64`

GetAmountHigherThan returns the AmountHigherThan field if non-nil, zero value otherwise.

### GetAmountHigherThanOk

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetAmountHigherThanOk() (*int64, bool)`

GetAmountHigherThanOk returns a tuple with the AmountHigherThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountHigherThan

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) SetAmountHigherThan(v int64)`

SetAmountHigherThan sets AmountHigherThan field to given value.


### GetCallbackSecretKey

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewConfirmedInternalTransactionsForSpecificAmountRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


