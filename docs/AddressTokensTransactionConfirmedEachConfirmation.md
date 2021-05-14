# AddressTokensTransactionConfirmedEachConfirmation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**ReferenceId** | **string** | Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc. | 
**IdempotencyKey** | **string** | Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice. | 
**Data** | [**AddressTokensTransactionConfirmedEachConfirmationData**](AddressTokensTransactionConfirmedEachConfirmationData.md) |  | 

## Methods

### NewAddressTokensTransactionConfirmedEachConfirmation

`func NewAddressTokensTransactionConfirmedEachConfirmation(apiVersion string, referenceId string, idempotencyKey string, data AddressTokensTransactionConfirmedEachConfirmationData, ) *AddressTokensTransactionConfirmedEachConfirmation`

NewAddressTokensTransactionConfirmedEachConfirmation instantiates a new AddressTokensTransactionConfirmedEachConfirmation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTokensTransactionConfirmedEachConfirmationWithDefaults

`func NewAddressTokensTransactionConfirmedEachConfirmationWithDefaults() *AddressTokensTransactionConfirmedEachConfirmation`

NewAddressTokensTransactionConfirmedEachConfirmationWithDefaults instantiates a new AddressTokensTransactionConfirmedEachConfirmation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AddressTokensTransactionConfirmedEachConfirmation) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetReferenceId

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AddressTokensTransactionConfirmedEachConfirmation) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetIdempotencyKey

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *AddressTokensTransactionConfirmedEachConfirmation) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetData

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetData() AddressTokensTransactionConfirmedEachConfirmationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddressTokensTransactionConfirmedEachConfirmation) GetDataOk() (*AddressTokensTransactionConfirmedEachConfirmationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddressTokensTransactionConfirmedEachConfirmation) SetData(v AddressTokensTransactionConfirmedEachConfirmationData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

