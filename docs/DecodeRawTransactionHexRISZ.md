# DecodeRawTransactionHexRISZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**Saplinged** | **bool** | Defines if the transaction includes sapling or not. | 
**TransactionHash** | **string** | Represents the same as transactionId for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols hash is different from transactionId for SegWit transactions. | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**Version** | **int32** | Represents the transaction version number. | 
**VersionGroupId** | **string** | Represents the transaction version group ID | 
**Vin** | [**[]DecodeRawTransactionHexRISZVin**](DecodeRawTransactionHexRISZVin.md) | Represents the Inputs of the transaction | 
**Vout** | [**[]DecodeRawTransactionHexRISZVout**](DecodeRawTransactionHexRISZVout.md) | Represents the Inputs of the transaction | 

## Methods

### NewDecodeRawTransactionHexRISZ

`func NewDecodeRawTransactionHexRISZ(expiryHeight int32, locktime int32, overwintered bool, saplinged bool, transactionHash string, valueBalance string, version int32, versionGroupId string, vin []DecodeRawTransactionHexRISZVin, vout []DecodeRawTransactionHexRISZVout, ) *DecodeRawTransactionHexRISZ`

NewDecodeRawTransactionHexRISZ instantiates a new DecodeRawTransactionHexRISZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISZWithDefaults

`func NewDecodeRawTransactionHexRISZWithDefaults() *DecodeRawTransactionHexRISZ`

NewDecodeRawTransactionHexRISZWithDefaults instantiates a new DecodeRawTransactionHexRISZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryHeight

`func (o *DecodeRawTransactionHexRISZ) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *DecodeRawTransactionHexRISZ) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *DecodeRawTransactionHexRISZ) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetLocktime

`func (o *DecodeRawTransactionHexRISZ) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *DecodeRawTransactionHexRISZ) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *DecodeRawTransactionHexRISZ) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetOverwintered

`func (o *DecodeRawTransactionHexRISZ) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *DecodeRawTransactionHexRISZ) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *DecodeRawTransactionHexRISZ) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSaplinged

`func (o *DecodeRawTransactionHexRISZ) GetSaplinged() bool`

GetSaplinged returns the Saplinged field if non-nil, zero value otherwise.

### GetSaplingedOk

`func (o *DecodeRawTransactionHexRISZ) GetSaplingedOk() (*bool, bool)`

GetSaplingedOk returns a tuple with the Saplinged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaplinged

`func (o *DecodeRawTransactionHexRISZ) SetSaplinged(v bool)`

SetSaplinged sets Saplinged field to given value.


### GetTransactionHash

`func (o *DecodeRawTransactionHexRISZ) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *DecodeRawTransactionHexRISZ) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *DecodeRawTransactionHexRISZ) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetValueBalance

`func (o *DecodeRawTransactionHexRISZ) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *DecodeRawTransactionHexRISZ) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *DecodeRawTransactionHexRISZ) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersion

`func (o *DecodeRawTransactionHexRISZ) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecodeRawTransactionHexRISZ) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecodeRawTransactionHexRISZ) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionGroupId

`func (o *DecodeRawTransactionHexRISZ) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *DecodeRawTransactionHexRISZ) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *DecodeRawTransactionHexRISZ) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetVin

`func (o *DecodeRawTransactionHexRISZ) GetVin() []DecodeRawTransactionHexRISZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *DecodeRawTransactionHexRISZ) GetVinOk() (*[]DecodeRawTransactionHexRISZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *DecodeRawTransactionHexRISZ) SetVin(v []DecodeRawTransactionHexRISZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *DecodeRawTransactionHexRISZ) GetVout() []DecodeRawTransactionHexRISZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *DecodeRawTransactionHexRISZ) GetVoutOk() (*[]DecodeRawTransactionHexRISZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *DecodeRawTransactionHexRISZ) SetVout(v []DecodeRawTransactionHexRISZVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


