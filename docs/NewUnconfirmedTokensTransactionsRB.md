# NewUnconfirmedTokensTransactionsRB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Data** | [**NewUnconfirmedTokensTransactionsRBData**](NewUnconfirmedTokensTransactionsRBData.md) |  | 

## Methods

### NewNewUnconfirmedTokensTransactionsRB

`func NewNewUnconfirmedTokensTransactionsRB(data NewUnconfirmedTokensTransactionsRBData, ) *NewUnconfirmedTokensTransactionsRB`

NewNewUnconfirmedTokensTransactionsRB instantiates a new NewUnconfirmedTokensTransactionsRB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUnconfirmedTokensTransactionsRBWithDefaults

`func NewNewUnconfirmedTokensTransactionsRBWithDefaults() *NewUnconfirmedTokensTransactionsRB`

NewNewUnconfirmedTokensTransactionsRBWithDefaults instantiates a new NewUnconfirmedTokensTransactionsRB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *NewUnconfirmedTokensTransactionsRB) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *NewUnconfirmedTokensTransactionsRB) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *NewUnconfirmedTokensTransactionsRB) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *NewUnconfirmedTokensTransactionsRB) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *NewUnconfirmedTokensTransactionsRB) GetData() NewUnconfirmedTokensTransactionsRBData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NewUnconfirmedTokensTransactionsRB) GetDataOk() (*NewUnconfirmedTokensTransactionsRBData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NewUnconfirmedTokensTransactionsRB) SetData(v NewUnconfirmedTokensTransactionsRBData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


