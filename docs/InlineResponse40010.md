# InlineResponse40010

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**RequestId** | **string** | Defines the ID of the request. The &#x60;requestId&#x60; is generated by Crypto APIs and it&#39;s unique for every request. | 
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Error** | [**ListConfirmedTransactionsByAddressE400**](ListConfirmedTransactionsByAddressE400.md) |  | 

## Methods

### NewInlineResponse40010

`func NewInlineResponse40010(apiVersion string, requestId string, error_ ListConfirmedTransactionsByAddressE400, ) *InlineResponse40010`

NewInlineResponse40010 instantiates a new InlineResponse40010 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse40010WithDefaults

`func NewInlineResponse40010WithDefaults() *InlineResponse40010`

NewInlineResponse40010WithDefaults instantiates a new InlineResponse40010 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *InlineResponse40010) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *InlineResponse40010) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *InlineResponse40010) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetRequestId

`func (o *InlineResponse40010) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InlineResponse40010) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InlineResponse40010) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetContext

`func (o *InlineResponse40010) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *InlineResponse40010) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *InlineResponse40010) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *InlineResponse40010) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse40010) GetError() ListConfirmedTransactionsByAddressE400`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse40010) GetErrorOk() (*ListConfirmedTransactionsByAddressE400, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse40010) SetError(v ListConfirmedTransactionsByAddressE400)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

