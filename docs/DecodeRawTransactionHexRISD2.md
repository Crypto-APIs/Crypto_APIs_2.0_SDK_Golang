# DecodeRawTransactionHexRISD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**TransactionHash** | **string** | Represents the same as transactionId for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols hash is different from transactionId for SegWit transactions. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents transaction version number | 
**Vin** | [**[]DecodeRawTransactionHexRISD2Vin**](DecodeRawTransactionHexRISD2Vin.md) | Represents the transaction inputs. | 
**Vout** | [**[]DecodeRawTransactionHexRISD2Vout**](DecodeRawTransactionHexRISD2Vout.md) | Represents the transaction outputs. | 
**Weight** | Pointer to **int32** | Represents the size of a block, measured in weight units and including the segwit discount. | [optional] 

## Methods

### NewDecodeRawTransactionHexRISD2

`func NewDecodeRawTransactionHexRISD2(locktime int32, transactionHash string, vSize int32, version int32, vin []DecodeRawTransactionHexRISD2Vin, vout []DecodeRawTransactionHexRISD2Vout, ) *DecodeRawTransactionHexRISD2`

NewDecodeRawTransactionHexRISD2 instantiates a new DecodeRawTransactionHexRISD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISD2WithDefaults

`func NewDecodeRawTransactionHexRISD2WithDefaults() *DecodeRawTransactionHexRISD2`

NewDecodeRawTransactionHexRISD2WithDefaults instantiates a new DecodeRawTransactionHexRISD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *DecodeRawTransactionHexRISD2) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *DecodeRawTransactionHexRISD2) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *DecodeRawTransactionHexRISD2) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetTransactionHash

`func (o *DecodeRawTransactionHexRISD2) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *DecodeRawTransactionHexRISD2) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *DecodeRawTransactionHexRISD2) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetVSize

`func (o *DecodeRawTransactionHexRISD2) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *DecodeRawTransactionHexRISD2) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *DecodeRawTransactionHexRISD2) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *DecodeRawTransactionHexRISD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecodeRawTransactionHexRISD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecodeRawTransactionHexRISD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *DecodeRawTransactionHexRISD2) GetVin() []DecodeRawTransactionHexRISD2Vin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *DecodeRawTransactionHexRISD2) GetVinOk() (*[]DecodeRawTransactionHexRISD2Vin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *DecodeRawTransactionHexRISD2) SetVin(v []DecodeRawTransactionHexRISD2Vin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *DecodeRawTransactionHexRISD2) GetVout() []DecodeRawTransactionHexRISD2Vout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *DecodeRawTransactionHexRISD2) GetVoutOk() (*[]DecodeRawTransactionHexRISD2Vout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *DecodeRawTransactionHexRISD2) SetVout(v []DecodeRawTransactionHexRISD2Vout)`

SetVout sets Vout field to given value.


### GetWeight

`func (o *DecodeRawTransactionHexRISD2) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DecodeRawTransactionHexRISD2) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DecodeRawTransactionHexRISD2) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DecodeRawTransactionHexRISD2) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


