# BannedIpAddressErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | Specifies an attribute of the error by name. | 
**Message** | **string** | Specifies the details of an attribute as part from the error. | 

## Methods

### NewBannedIpAddressErrorDetails

`func NewBannedIpAddressErrorDetails(attribute string, message string, ) *BannedIpAddressErrorDetails`

NewBannedIpAddressErrorDetails instantiates a new BannedIpAddressErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannedIpAddressErrorDetailsWithDefaults

`func NewBannedIpAddressErrorDetailsWithDefaults() *BannedIpAddressErrorDetails`

NewBannedIpAddressErrorDetailsWithDefaults instantiates a new BannedIpAddressErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *BannedIpAddressErrorDetails) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *BannedIpAddressErrorDetails) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *BannedIpAddressErrorDetails) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetMessage

`func (o *BannedIpAddressErrorDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BannedIpAddressErrorDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BannedIpAddressErrorDetails) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


