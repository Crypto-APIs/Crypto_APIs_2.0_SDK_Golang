# ListTransactionsByBlockHashRIBSBScriptPubKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Asm** | **string** | Represents the assembly of the script public key of the address. | 
**Hex** | **string** | Represents the hex of the script public key of the address. | 
**ReqSigs** | **int32** | Represents the required signatures. | 
**Type** | **string** | Represents the script type. | 

## Methods

### NewListTransactionsByBlockHashRIBSBScriptPubKey

`func NewListTransactionsByBlockHashRIBSBScriptPubKey(addresses []string, asm string, hex string, reqSigs int32, type_ string, ) *ListTransactionsByBlockHashRIBSBScriptPubKey`

NewListTransactionsByBlockHashRIBSBScriptPubKey instantiates a new ListTransactionsByBlockHashRIBSBScriptPubKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSBScriptPubKeyWithDefaults

`func NewListTransactionsByBlockHashRIBSBScriptPubKeyWithDefaults() *ListTransactionsByBlockHashRIBSBScriptPubKey`

NewListTransactionsByBlockHashRIBSBScriptPubKeyWithDefaults instantiates a new ListTransactionsByBlockHashRIBSBScriptPubKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetAsm

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetAsm() string`

GetAsm returns the Asm field if non-nil, zero value otherwise.

### GetAsmOk

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetAsmOk() (*string, bool)`

GetAsmOk returns a tuple with the Asm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsm

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) SetAsm(v string)`

SetAsm sets Asm field to given value.


### GetHex

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) SetHex(v string)`

SetHex sets Hex field to given value.


### GetReqSigs

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetReqSigs() int32`

GetReqSigs returns the ReqSigs field if non-nil, zero value otherwise.

### GetReqSigsOk

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetReqSigsOk() (*int32, bool)`

GetReqSigsOk returns a tuple with the ReqSigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSigs

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) SetReqSigs(v int32)`

SetReqSigs sets ReqSigs field to given value.


### GetType

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListTransactionsByBlockHashRIBSBScriptPubKey) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


