# GenerateAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]GenerateAddressRIAddresses**](GenerateAddressRIAddresses.md) |  | 
**PrivateKey** | **string** | Represents the privately known secret key used for authentication and encryption of the address. | 
**PublicKey** | **string** | Represents the publicly known key used for identification of the address. | 
**Wif** | **string** | Represents the Wallet Import Format which dictates the encoding that allows the copy of the private ECDSA key easily. | 

## Methods

### NewGenerateAddressRI

`func NewGenerateAddressRI(addresses []GenerateAddressRIAddresses, privateKey string, publicKey string, wif string, ) *GenerateAddressRI`

NewGenerateAddressRI instantiates a new GenerateAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateAddressRIWithDefaults

`func NewGenerateAddressRIWithDefaults() *GenerateAddressRI`

NewGenerateAddressRIWithDefaults instantiates a new GenerateAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GenerateAddressRI) GetAddresses() []GenerateAddressRIAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GenerateAddressRI) GetAddressesOk() (*[]GenerateAddressRIAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GenerateAddressRI) SetAddresses(v []GenerateAddressRIAddresses)`

SetAddresses sets Addresses field to given value.


### GetPrivateKey

`func (o *GenerateAddressRI) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GenerateAddressRI) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GenerateAddressRI) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPublicKey

`func (o *GenerateAddressRI) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GenerateAddressRI) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GenerateAddressRI) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetWif

`func (o *GenerateAddressRI) GetWif() string`

GetWif returns the Wif field if non-nil, zero value otherwise.

### GetWifOk

`func (o *GenerateAddressRI) GetWifOk() (*string, bool)`

GetWifOk returns a tuple with the Wif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWif

`func (o *GenerateAddressRI) SetWif(v string)`

SetWif sets Wif field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


