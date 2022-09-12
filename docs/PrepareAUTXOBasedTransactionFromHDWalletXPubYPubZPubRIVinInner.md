# PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Representation of the address | 
**Change** | Pointer to **int64** | Representation of the change value | [optional] 
**DerivationIndex** | Pointer to **int64** | Representation of the derivation index of the xpub address. | [optional] 
**OutputIndex** | **int32** | Representation of the output index | 
**Satoshis** | **int64** | Representation of the satoshis value | 
**Script** | **string** | Representation of the script string | 
**Sighash** | **string** | Representation of the hash that should be signed. | 
**TransactionId** | **string** | Represents the reference transaction identifier. | 

## Methods

### NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner

`func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner(address string, outputIndex int32, satoshis int64, script string, sighash string, transactionId string, ) *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner`

NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInnerWithDefaults

`func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInnerWithDefaults() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner`

NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInnerWithDefaults instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChange

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetChange() int64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetChangeOk() (*int64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetChange(v int64)`

SetChange sets Change field to given value.

### HasChange

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetDerivationIndex

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetDerivationIndex() int64`

GetDerivationIndex returns the DerivationIndex field if non-nil, zero value otherwise.

### GetDerivationIndexOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetDerivationIndexOk() (*int64, bool)`

GetDerivationIndexOk returns a tuple with the DerivationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationIndex

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetDerivationIndex(v int64)`

SetDerivationIndex sets DerivationIndex field to given value.

### HasDerivationIndex

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) HasDerivationIndex() bool`

HasDerivationIndex returns a boolean if a field has been set.

### GetOutputIndex

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.


### GetSatoshis

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetSatoshis() int64`

GetSatoshis returns the Satoshis field if non-nil, zero value otherwise.

### GetSatoshisOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetSatoshisOk() (*int64, bool)`

GetSatoshisOk returns a tuple with the Satoshis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshis

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetSatoshis(v int64)`

SetSatoshis sets Satoshis field to given value.


### GetScript

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetScript(v string)`

SetScript sets Script field to given value.


### GetSighash

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetSighash() string`

GetSighash returns the Sighash field if non-nil, zero value otherwise.

### GetSighashOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetSighashOk() (*string, bool)`

GetSighashOk returns a tuple with the Sighash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSighash

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetSighash(v string)`

SetSighash sets Sighash field to given value.


### GetTransactionId

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRIVinInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


