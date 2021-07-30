# GenerateDepositAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Specifies the specific address&#39;s unique string value. | 
**CreatedTimestamp** | **int32** | Defines the specific UNIX time when the deposit address was created. | 
**Label** | **string** | Represents a custom tag that customers can set up for their Wallets and addresses. E.g. custom label named \&quot;Special addresses\&quot;. | 

## Methods

### NewGenerateDepositAddressRI

`func NewGenerateDepositAddressRI(address string, createdTimestamp int32, label string, ) *GenerateDepositAddressRI`

NewGenerateDepositAddressRI instantiates a new GenerateDepositAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDepositAddressRIWithDefaults

`func NewGenerateDepositAddressRIWithDefaults() *GenerateDepositAddressRI`

NewGenerateDepositAddressRIWithDefaults instantiates a new GenerateDepositAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GenerateDepositAddressRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GenerateDepositAddressRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GenerateDepositAddressRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCreatedTimestamp

`func (o *GenerateDepositAddressRI) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *GenerateDepositAddressRI) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *GenerateDepositAddressRI) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetLabel

`func (o *GenerateDepositAddressRI) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GenerateDepositAddressRI) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GenerateDepositAddressRI) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


