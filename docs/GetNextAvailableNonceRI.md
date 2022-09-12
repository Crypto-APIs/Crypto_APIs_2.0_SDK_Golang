# GetNextAvailableNonceRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextAvailableNonce** | **int64** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 

## Methods

### NewGetNextAvailableNonceRI

`func NewGetNextAvailableNonceRI(nextAvailableNonce int64, ) *GetNextAvailableNonceRI`

NewGetNextAvailableNonceRI instantiates a new GetNextAvailableNonceRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNextAvailableNonceRIWithDefaults

`func NewGetNextAvailableNonceRIWithDefaults() *GetNextAvailableNonceRI`

NewGetNextAvailableNonceRIWithDefaults instantiates a new GetNextAvailableNonceRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextAvailableNonce

`func (o *GetNextAvailableNonceRI) GetNextAvailableNonce() int64`

GetNextAvailableNonce returns the NextAvailableNonce field if non-nil, zero value otherwise.

### GetNextAvailableNonceOk

`func (o *GetNextAvailableNonceRI) GetNextAvailableNonceOk() (*int64, bool)`

GetNextAvailableNonceOk returns a tuple with the NextAvailableNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableNonce

`func (o *GetNextAvailableNonceRI) SetNextAvailableNonce(v int64)`

SetNextAvailableNonce sets NextAvailableNonce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


