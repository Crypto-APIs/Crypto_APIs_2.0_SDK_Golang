# TokensForwardingSuccessData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** | Represents the Crypto APIs 2.0 product which sends the callback. | 
**Event** | **string** | Defines the specific event, for which a callback subscription is set. | 
**Item** | [**TokensForwardingSuccessDataItem**](TokensForwardingSuccessDataItem.md) |  | 

## Methods

### NewTokensForwardingSuccessData

`func NewTokensForwardingSuccessData(product string, event string, item TokensForwardingSuccessDataItem, ) *TokensForwardingSuccessData`

NewTokensForwardingSuccessData instantiates a new TokensForwardingSuccessData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensForwardingSuccessDataWithDefaults

`func NewTokensForwardingSuccessDataWithDefaults() *TokensForwardingSuccessData`

NewTokensForwardingSuccessDataWithDefaults instantiates a new TokensForwardingSuccessData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *TokensForwardingSuccessData) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *TokensForwardingSuccessData) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *TokensForwardingSuccessData) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetEvent

`func (o *TokensForwardingSuccessData) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TokensForwardingSuccessData) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TokensForwardingSuccessData) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetItem

`func (o *TokensForwardingSuccessData) GetItem() TokensForwardingSuccessDataItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TokensForwardingSuccessData) GetItemOk() (*TokensForwardingSuccessDataItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TokensForwardingSuccessData) SetItem(v TokensForwardingSuccessDataItem)`

SetItem sets Item field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


