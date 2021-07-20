# ValidateAddressR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**RequestId** | **string** | Defines the ID of the request. The &#x60;requestId&#x60; is generated by Crypto APIs and it&#39;s unique for every request. | 
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Data** | [**ValidateAddressRData**](ValidateAddressRData.md) |  | 

## Methods

### NewValidateAddressR

`func NewValidateAddressR(apiVersion string, requestId string, data ValidateAddressRData, ) *ValidateAddressR`

NewValidateAddressR instantiates a new ValidateAddressR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateAddressRWithDefaults

`func NewValidateAddressRWithDefaults() *ValidateAddressR`

NewValidateAddressRWithDefaults instantiates a new ValidateAddressR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ValidateAddressR) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ValidateAddressR) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ValidateAddressR) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetRequestId

`func (o *ValidateAddressR) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ValidateAddressR) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ValidateAddressR) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetContext

`func (o *ValidateAddressR) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ValidateAddressR) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ValidateAddressR) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ValidateAddressR) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *ValidateAddressR) GetData() ValidateAddressRData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValidateAddressR) GetDataOk() (*ValidateAddressRData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValidateAddressR) SetData(v ValidateAddressRData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

