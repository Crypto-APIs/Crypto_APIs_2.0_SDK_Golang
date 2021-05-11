# NewUnconfirmedCoinsTransactionsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Data** | [**NewUnconfirmedCoinsTransactionsRequestBodyData**](NewUnconfirmedCoinsTransactionsRequestBodyData.md) |  | 

## Methods

### NewNewUnconfirmedCoinsTransactionsRequestBody

`func NewNewUnconfirmedCoinsTransactionsRequestBody(data NewUnconfirmedCoinsTransactionsRequestBodyData, ) *NewUnconfirmedCoinsTransactionsRequestBody`

NewNewUnconfirmedCoinsTransactionsRequestBody instantiates a new NewUnconfirmedCoinsTransactionsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUnconfirmedCoinsTransactionsRequestBodyWithDefaults

`func NewNewUnconfirmedCoinsTransactionsRequestBodyWithDefaults() *NewUnconfirmedCoinsTransactionsRequestBody`

NewNewUnconfirmedCoinsTransactionsRequestBodyWithDefaults instantiates a new NewUnconfirmedCoinsTransactionsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *NewUnconfirmedCoinsTransactionsRequestBody) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *NewUnconfirmedCoinsTransactionsRequestBody) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *NewUnconfirmedCoinsTransactionsRequestBody) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *NewUnconfirmedCoinsTransactionsRequestBody) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *NewUnconfirmedCoinsTransactionsRequestBody) GetData() NewUnconfirmedCoinsTransactionsRequestBodyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NewUnconfirmedCoinsTransactionsRequestBody) GetDataOk() (*NewUnconfirmedCoinsTransactionsRequestBodyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NewUnconfirmedCoinsTransactionsRequestBody) SetData(v NewUnconfirmedCoinsTransactionsRequestBodyData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


