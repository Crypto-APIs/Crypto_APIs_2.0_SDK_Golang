# BroadcastLocallySignedTransactionRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**SignedTransactionHex** | **string** | Represents the signed transaction&#39;s specific hex. | 

## Methods

### NewBroadcastLocallySignedTransactionRBDataItem

`func NewBroadcastLocallySignedTransactionRBDataItem(callbackUrl string, signedTransactionHex string, ) *BroadcastLocallySignedTransactionRBDataItem`

NewBroadcastLocallySignedTransactionRBDataItem instantiates a new BroadcastLocallySignedTransactionRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastLocallySignedTransactionRBDataItemWithDefaults

`func NewBroadcastLocallySignedTransactionRBDataItemWithDefaults() *BroadcastLocallySignedTransactionRBDataItem`

NewBroadcastLocallySignedTransactionRBDataItemWithDefaults instantiates a new BroadcastLocallySignedTransactionRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *BroadcastLocallySignedTransactionRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *BroadcastLocallySignedTransactionRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *BroadcastLocallySignedTransactionRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *BroadcastLocallySignedTransactionRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *BroadcastLocallySignedTransactionRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *BroadcastLocallySignedTransactionRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *BroadcastLocallySignedTransactionRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetSignedTransactionHex

`func (o *BroadcastLocallySignedTransactionRBDataItem) GetSignedTransactionHex() string`

GetSignedTransactionHex returns the SignedTransactionHex field if non-nil, zero value otherwise.

### GetSignedTransactionHexOk

`func (o *BroadcastLocallySignedTransactionRBDataItem) GetSignedTransactionHexOk() (*string, bool)`

GetSignedTransactionHexOk returns a tuple with the SignedTransactionHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTransactionHex

`func (o *BroadcastLocallySignedTransactionRBDataItem) SetSignedTransactionHex(v string)`

SetSignedTransactionHex sets SignedTransactionHex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


