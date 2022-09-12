# DeriveAndSyncNewReceivingAddressesRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 
**AddressFormat** | **string** | Represents the format of the address. | 
**AddressType** | **string** | Defines the address type. | 
**DerivationType** | **string** | Represents the derivation type | 
**Index** | **string** | Represents the output index. It refers to the UTXO sequence in the transaction outputs (vout). | 

## Methods

### NewDeriveAndSyncNewReceivingAddressesRI

`func NewDeriveAndSyncNewReceivingAddressesRI(address string, addressFormat string, addressType string, derivationType string, index string, ) *DeriveAndSyncNewReceivingAddressesRI`

NewDeriveAndSyncNewReceivingAddressesRI instantiates a new DeriveAndSyncNewReceivingAddressesRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeriveAndSyncNewReceivingAddressesRIWithDefaults

`func NewDeriveAndSyncNewReceivingAddressesRIWithDefaults() *DeriveAndSyncNewReceivingAddressesRI`

NewDeriveAndSyncNewReceivingAddressesRIWithDefaults instantiates a new DeriveAndSyncNewReceivingAddressesRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DeriveAndSyncNewReceivingAddressesRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressFormat

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressFormat() string`

GetAddressFormat returns the AddressFormat field if non-nil, zero value otherwise.

### GetAddressFormatOk

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressFormatOk() (*string, bool)`

GetAddressFormatOk returns a tuple with the AddressFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFormat

`func (o *DeriveAndSyncNewReceivingAddressesRI) SetAddressFormat(v string)`

SetAddressFormat sets AddressFormat field to given value.


### GetAddressType

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *DeriveAndSyncNewReceivingAddressesRI) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetDerivationType

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetDerivationType() string`

GetDerivationType returns the DerivationType field if non-nil, zero value otherwise.

### GetDerivationTypeOk

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetDerivationTypeOk() (*string, bool)`

GetDerivationTypeOk returns a tuple with the DerivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationType

`func (o *DeriveAndSyncNewReceivingAddressesRI) SetDerivationType(v string)`

SetDerivationType sets DerivationType field to given value.


### GetIndex

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DeriveAndSyncNewReceivingAddressesRI) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DeriveAndSyncNewReceivingAddressesRI) SetIndex(v string)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


