# BannedIpAddressDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | Specifies an attribute of the error by name. | 
**Message** | **string** | Specifies the details of an attribute as part from the error. | 

## Methods

### NewBannedIpAddressDetails

`func NewBannedIpAddressDetails(attribute string, message string, ) *BannedIpAddressDetails`

NewBannedIpAddressDetails instantiates a new BannedIpAddressDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannedIpAddressDetailsWithDefaults

`func NewBannedIpAddressDetailsWithDefaults() *BannedIpAddressDetails`

NewBannedIpAddressDetailsWithDefaults instantiates a new BannedIpAddressDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *BannedIpAddressDetails) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *BannedIpAddressDetails) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *BannedIpAddressDetails) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetMessage

`func (o *BannedIpAddressDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BannedIpAddressDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BannedIpAddressDetails) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


