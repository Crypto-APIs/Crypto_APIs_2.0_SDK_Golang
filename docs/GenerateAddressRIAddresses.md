# GenerateAddressRIAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the specific address that will be generated. | 
**Format** | **string** | Defines the form of transaction that is used. Keep in mind that we support more than one type of formats for example: p2pkh p2sh p2wpkh etc. and they are generated simultaneously in the response of the Generate Address endpoint, depending on the blockchain and network that has been chosen. For more information about supported formats - please check \&quot;What we support\&quot; section | 

## Methods

### NewGenerateAddressRIAddresses

`func NewGenerateAddressRIAddresses(address string, format string, ) *GenerateAddressRIAddresses`

NewGenerateAddressRIAddresses instantiates a new GenerateAddressRIAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateAddressRIAddressesWithDefaults

`func NewGenerateAddressRIAddressesWithDefaults() *GenerateAddressRIAddresses`

NewGenerateAddressRIAddressesWithDefaults instantiates a new GenerateAddressRIAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GenerateAddressRIAddresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GenerateAddressRIAddresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GenerateAddressRIAddresses) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFormat

`func (o *GenerateAddressRIAddresses) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GenerateAddressRIAddresses) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GenerateAddressRIAddresses) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


