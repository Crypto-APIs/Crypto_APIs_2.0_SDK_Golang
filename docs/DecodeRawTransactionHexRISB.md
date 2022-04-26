# DecodeRawTransactionHexRISB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**TransactionHash** | **string** | Represents the same as transactionId for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols hash is different from transactionId for SegWit transactions. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]DecodeRawTransactionHexRISBVin**](DecodeRawTransactionHexRISBVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]DecodeRawTransactionHexRISBVout**](DecodeRawTransactionHexRISBVout.md) | Represents the transaction outputs. | 
**Weight** | Pointer to **int32** | Represents the size of Bitcoin block, measured in weight units and including the segwit discount. | [optional] 

## Methods

### NewDecodeRawTransactionHexRISB

`func NewDecodeRawTransactionHexRISB(locktime int32, transactionHash string, vSize int32, version int32, vin []DecodeRawTransactionHexRISBVin, vout []DecodeRawTransactionHexRISBVout, ) *DecodeRawTransactionHexRISB`

NewDecodeRawTransactionHexRISB instantiates a new DecodeRawTransactionHexRISB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISBWithDefaults

`func NewDecodeRawTransactionHexRISBWithDefaults() *DecodeRawTransactionHexRISB`

NewDecodeRawTransactionHexRISBWithDefaults instantiates a new DecodeRawTransactionHexRISB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *DecodeRawTransactionHexRISB) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *DecodeRawTransactionHexRISB) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *DecodeRawTransactionHexRISB) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetTransactionHash

`func (o *DecodeRawTransactionHexRISB) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *DecodeRawTransactionHexRISB) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *DecodeRawTransactionHexRISB) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetVSize

`func (o *DecodeRawTransactionHexRISB) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *DecodeRawTransactionHexRISB) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *DecodeRawTransactionHexRISB) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *DecodeRawTransactionHexRISB) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecodeRawTransactionHexRISB) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecodeRawTransactionHexRISB) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *DecodeRawTransactionHexRISB) GetVin() []DecodeRawTransactionHexRISBVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *DecodeRawTransactionHexRISB) GetVinOk() (*[]DecodeRawTransactionHexRISBVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *DecodeRawTransactionHexRISB) SetVin(v []DecodeRawTransactionHexRISBVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *DecodeRawTransactionHexRISB) GetVout() []DecodeRawTransactionHexRISBVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *DecodeRawTransactionHexRISB) GetVoutOk() (*[]DecodeRawTransactionHexRISBVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *DecodeRawTransactionHexRISB) SetVout(v []DecodeRawTransactionHexRISBVout)`

SetVout sets Vout field to given value.


### GetWeight

`func (o *DecodeRawTransactionHexRISB) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DecodeRawTransactionHexRISB) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DecodeRawTransactionHexRISB) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DecodeRawTransactionHexRISB) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


