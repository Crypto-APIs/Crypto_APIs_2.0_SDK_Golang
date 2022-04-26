# DecodeXAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressTag** | **int32** | Defines a specific Tag that is an additional XRP address feature. It helps identifying a transaction recipient beyond a wallet address. | 
**ClassicAddress** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 

## Methods

### NewDecodeXAddressRI

`func NewDecodeXAddressRI(addressTag int32, classicAddress string, ) *DecodeXAddressRI`

NewDecodeXAddressRI instantiates a new DecodeXAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeXAddressRIWithDefaults

`func NewDecodeXAddressRIWithDefaults() *DecodeXAddressRI`

NewDecodeXAddressRIWithDefaults instantiates a new DecodeXAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressTag

`func (o *DecodeXAddressRI) GetAddressTag() int32`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *DecodeXAddressRI) GetAddressTagOk() (*int32, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *DecodeXAddressRI) SetAddressTag(v int32)`

SetAddressTag sets AddressTag field to given value.


### GetClassicAddress

`func (o *DecodeXAddressRI) GetClassicAddress() string`

GetClassicAddress returns the ClassicAddress field if non-nil, zero value otherwise.

### GetClassicAddressOk

`func (o *DecodeXAddressRI) GetClassicAddressOk() (*string, bool)`

GetClassicAddressOk returns a tuple with the ClassicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicAddress

`func (o *DecodeXAddressRI) SetClassicAddress(v string)`

SetClassicAddress sets ClassicAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


