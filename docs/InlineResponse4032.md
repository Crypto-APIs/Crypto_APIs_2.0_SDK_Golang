# InlineResponse4032

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**RequestId** | **string** | Defines the ID of the request. The &#x60;requestId&#x60; is generated by Crypto APIs and it&#39;s unique for every request. | 
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Error** | [**GetXRPRippleAddressDetailsE403**](GetXRPRippleAddressDetailsE403.md) |  | 

## Methods

### NewInlineResponse4032

`func NewInlineResponse4032(apiVersion string, requestId string, error_ GetXRPRippleAddressDetailsE403, ) *InlineResponse4032`

NewInlineResponse4032 instantiates a new InlineResponse4032 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse4032WithDefaults

`func NewInlineResponse4032WithDefaults() *InlineResponse4032`

NewInlineResponse4032WithDefaults instantiates a new InlineResponse4032 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *InlineResponse4032) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *InlineResponse4032) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *InlineResponse4032) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetRequestId

`func (o *InlineResponse4032) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InlineResponse4032) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InlineResponse4032) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetContext

`func (o *InlineResponse4032) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *InlineResponse4032) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *InlineResponse4032) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *InlineResponse4032) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse4032) GetError() GetXRPRippleAddressDetailsE403`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse4032) GetErrorOk() (*GetXRPRippleAddressDetailsE403, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse4032) SetError(v GetXRPRippleAddressDetailsE403)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

