# ListTransactionsByBlockHashRIBSD2ScriptPubKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Asm** | **string** | Represents the assembly of the script public key of the address. | 
**Hex** | **string** | Represents the hex of the script public key of the address. | 
**ReqSigs** | **int32** | Represents the required signatures. | 
**Type** | **string** | Represents the script type. | 

## Methods

### NewListTransactionsByBlockHashRIBSD2ScriptPubKey

`func NewListTransactionsByBlockHashRIBSD2ScriptPubKey(addresses []string, asm string, hex string, reqSigs int32, type_ string, ) *ListTransactionsByBlockHashRIBSD2ScriptPubKey`

NewListTransactionsByBlockHashRIBSD2ScriptPubKey instantiates a new ListTransactionsByBlockHashRIBSD2ScriptPubKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSD2ScriptPubKeyWithDefaults

`func NewListTransactionsByBlockHashRIBSD2ScriptPubKeyWithDefaults() *ListTransactionsByBlockHashRIBSD2ScriptPubKey`

NewListTransactionsByBlockHashRIBSD2ScriptPubKeyWithDefaults instantiates a new ListTransactionsByBlockHashRIBSD2ScriptPubKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetAsm

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetAsm() string`

GetAsm returns the Asm field if non-nil, zero value otherwise.

### GetAsmOk

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetAsmOk() (*string, bool)`

GetAsmOk returns a tuple with the Asm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsm

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) SetAsm(v string)`

SetAsm sets Asm field to given value.


### GetHex

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) SetHex(v string)`

SetHex sets Hex field to given value.


### GetReqSigs

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetReqSigs() int32`

GetReqSigs returns the ReqSigs field if non-nil, zero value otherwise.

### GetReqSigsOk

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetReqSigsOk() (*int32, bool)`

GetReqSigsOk returns a tuple with the ReqSigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSigs

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) SetReqSigs(v int32)`

SetReqSigs sets ReqSigs field to given value.


### GetType

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListTransactionsByBlockHashRIBSD2ScriptPubKey) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


