# ListZilliqaTransactionsByBlockHashR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**RequestId** | **string** | Defines the ID of the request. The &#x60;requestId&#x60; is generated by Crypto APIs and it&#39;s unique for every request. | 
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Data** | [**ListZilliqaTransactionsByBlockHashRData**](ListZilliqaTransactionsByBlockHashRData.md) |  | 

## Methods

### NewListZilliqaTransactionsByBlockHashR

`func NewListZilliqaTransactionsByBlockHashR(apiVersion string, requestId string, data ListZilliqaTransactionsByBlockHashRData, ) *ListZilliqaTransactionsByBlockHashR`

NewListZilliqaTransactionsByBlockHashR instantiates a new ListZilliqaTransactionsByBlockHashR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListZilliqaTransactionsByBlockHashRWithDefaults

`func NewListZilliqaTransactionsByBlockHashRWithDefaults() *ListZilliqaTransactionsByBlockHashR`

NewListZilliqaTransactionsByBlockHashRWithDefaults instantiates a new ListZilliqaTransactionsByBlockHashR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListZilliqaTransactionsByBlockHashR) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListZilliqaTransactionsByBlockHashR) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListZilliqaTransactionsByBlockHashR) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetRequestId

`func (o *ListZilliqaTransactionsByBlockHashR) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListZilliqaTransactionsByBlockHashR) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListZilliqaTransactionsByBlockHashR) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetContext

`func (o *ListZilliqaTransactionsByBlockHashR) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ListZilliqaTransactionsByBlockHashR) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ListZilliqaTransactionsByBlockHashR) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ListZilliqaTransactionsByBlockHashR) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *ListZilliqaTransactionsByBlockHashR) GetData() ListZilliqaTransactionsByBlockHashRData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListZilliqaTransactionsByBlockHashR) GetDataOk() (*ListZilliqaTransactionsByBlockHashRData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListZilliqaTransactionsByBlockHashR) SetData(v ListZilliqaTransactionsByBlockHashRData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


