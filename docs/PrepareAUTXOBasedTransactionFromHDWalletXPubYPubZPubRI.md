# PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** | Representation of the additional data | [optional] 
**Fee** | **string** | When isConfirmed is True - Defines the amount of the transaction fee When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value. | 
**FeePerByte** | Pointer to **string** | Defines the fee per byte value | [optional] 
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Replaceable** | **bool** | Representation of whether the transaction is replaceable | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Vin** | [**[]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner**](PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner.md) | Represents the transaction inputs. | 
**Vout** | [**[]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner**](PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI

`func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI(fee string, locktime int64, replaceable bool, size int32, vin []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner, vout []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner, ) *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI`

NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults

`func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI`

NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetFee

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetFee(v string)`

SetFee sets Fee field to given value.


### GetFeePerByte

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFeePerByte() string`

GetFeePerByte returns the FeePerByte field if non-nil, zero value otherwise.

### GetFeePerByteOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetFeePerByteOk() (*string, bool)`

GetFeePerByteOk returns a tuple with the FeePerByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePerByte

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetFeePerByte(v string)`

SetFeePerByte sets FeePerByte field to given value.

### HasFeePerByte

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) HasFeePerByte() bool`

HasFeePerByte returns a boolean if a field has been set.

### GetLocktime

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetReplaceable

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetReplaceable() bool`

GetReplaceable returns the Replaceable field if non-nil, zero value otherwise.

### GetReplaceableOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetReplaceableOk() (*bool, bool)`

GetReplaceableOk returns a tuple with the Replaceable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceable

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetReplaceable(v bool)`

SetReplaceable sets Replaceable field to given value.


### GetSize

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVin

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVin() []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVinOk() (*[]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetVin(v []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVout() []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) GetVoutOk() (*[]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRI) SetVout(v []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


