# ValidateAddress400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**RequestId** | **string** | Defines the ID of the request. The &#x60;requestId&#x60; is generated by Crypto APIs and it&#39;s unique for every request. | 
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Error** | [**ValidateAddressE400**](ValidateAddressE400.md) |  | 

## Methods

### NewValidateAddress400Response

`func NewValidateAddress400Response(apiVersion string, requestId string, error_ ValidateAddressE400, ) *ValidateAddress400Response`

NewValidateAddress400Response instantiates a new ValidateAddress400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateAddress400ResponseWithDefaults

`func NewValidateAddress400ResponseWithDefaults() *ValidateAddress400Response`

NewValidateAddress400ResponseWithDefaults instantiates a new ValidateAddress400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ValidateAddress400Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ValidateAddress400Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ValidateAddress400Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetRequestId

`func (o *ValidateAddress400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ValidateAddress400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ValidateAddress400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetContext

`func (o *ValidateAddress400Response) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ValidateAddress400Response) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ValidateAddress400Response) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ValidateAddress400Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetError

`func (o *ValidateAddress400Response) GetError() ValidateAddressE400`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidateAddress400Response) GetErrorOk() (*ValidateAddressE400, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidateAddress400Response) SetError(v ValidateAddressE400)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


